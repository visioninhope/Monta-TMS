"""
COPYRIGHT 2022 MONTA

This file is part of Monta.

Monta is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

Monta is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with Monta.  If not, see <https://www.gnu.org/licenses/>.
"""
import uuid

import pytest
from django.core import mail
from django.core.exceptions import ValidationError

from billing import selectors
from billing.models import (
    BillingControl,
    BillingException,
    BillingHistory,
    BillingQueue,
)
from billing.services import mass_order_billing
from customer.factories import CustomerBillingProfileFactory, CustomerFactory
from order.models import Order
from order.tests.factories import OrderFactory
from organization.models import Organization
from utils.models import StatusChoices

pytestmark = pytest.mark.django_db


def test_bill_orders(
    organization,
    user,
    worker,
) -> None:
    order = OrderFactory(status="C")

    customer = CustomerFactory(organization=organization)

    BillingQueue.objects.create(
        organization=user.organization,
        order_type=order.order_type,
        order=order,
        revenue_code=order.revenue_code,
        customer=customer,
        worker=worker,
        commodity=order.commodity,
        bol_number=order.bol_number,
        user=user,
    )

    mass_order_billing.mass_order_billing_service(
        task_id=str(uuid.uuid4()), user_id=str(user.id)
    )

    billing_queue = BillingQueue.objects.all()
    billing_history = BillingHistory.objects.get(order=order)

    billing_history.refresh_from_db()

    assert billing_queue.count() == 0
    assert billing_history.order == order
    assert billing_history.organization == order.organization
    assert billing_history.order_type == order.order_type
    assert billing_history.revenue_code == order.revenue_code
    assert billing_history.customer == order.customer
    assert billing_history.commodity == order.commodity
    assert billing_history.bol_number == order.bol_number
    assert (
        billing_history.invoice_number
        == f"{user.organization.invoice_control.invoice_number_prefix}00001"
    )

    order.refresh_from_db()
    assert order.billed is True
    assert mail.outbox[0].subject == f"New invoice from {user.organization.name}"
    assert (
        mail.outbox[0].body
        == f"Please see attached invoice for invoice: {order.pro_number}"
    )


def test_invoice_number_generation(organization, customer, user, worker) -> None:
    """
    Test that invoice number is generated for each new invoice
    """
    order = OrderFactory(status="C")
    invoice = BillingQueue.objects.create(
        organization=user.organization,
        order_type=order.order_type,
        order=order,
        revenue_code=order.revenue_code,
        customer=customer,
        worker=worker,
        commodity=order.commodity,
        bol_number=order.bol_number,
        user=user,
    )
    assert invoice.invoice_number is not None
    assert (
        invoice.invoice_number
        == f"{user.organization.invoice_control.invoice_number_prefix}00001"
    )


def test_invoice_number_increments(organization, customer, user, worker) -> None:
    """
    Test that invoice number increments by 1 for each new invoice
    """
    order = OrderFactory(status="C")
    order_2 = OrderFactory(status="C")
    invoice = BillingQueue.objects.create(
        organization=user.organization,
        order_type=order.order_type,
        order=order,
        revenue_code=order.revenue_code,
        customer=customer,
        worker=worker,
        commodity=order.commodity,
        bol_number=order.bol_number,
        user=user,
    )
    second_invoice = BillingQueue.objects.create(
        organization=user.organization,
        order_type=order_2.order_type,
        order=order_2,
        revenue_code=order_2.revenue_code,
        customer=customer,
        worker=worker,
        commodity=order_2.commodity,
        bol_number=order_2.bol_number,
        user=user,
    )

    assert invoice.invoice_number is not None
    assert (
        invoice.invoice_number
        == f"{user.organization.invoice_control.invoice_number_prefix}00001"
    )
    assert second_invoice.invoice_number is not None
    assert (
        second_invoice.invoice_number
        == f"{user.organization.invoice_control.invoice_number_prefix}00002"
    )


def test_unbilled_order_in_billing_history(order) -> None:
    """
    Test ValidationError is thrown when adding an order in billing history
    that hasn't billed.
    """

    with pytest.raises(ValidationError) as excinfo:
        BillingHistory.objects.create(
            organization=order.organization,
            order=order,
        )

    assert excinfo.value.message_dict["order"] == [
        "Order has not been billed. Please try again with a different order."
    ]


def test_billing_control_hook(organization) -> None:
    """
    Test that the billing control hook is created when a new organization is
    created.
    """
    assert organization.billing_control is not None


def test_auto_bill_criteria_required_when_auto_bill_true(organization) -> None:
    """
    Test if `auto_bill_orders` is true & `auto_bill_criteria` is blank that a
    `ValidationError` is thrown.
    """
    billing_control = organization.billing_control

    with pytest.raises(ValidationError) as excinfo:
        billing_control.auto_bill_orders = True
        billing_control.auto_bill_criteria = None
        billing_control.full_clean()

    assert excinfo.value.message_dict["auto_bill_criteria"] == [
        "Auto Billing criteria is required when `Auto Bill Orders` is on. Please try again."
    ]


def test_auto_bill_criteria_choices_is_invalid(organization) -> None:
    """
    Test when passing invalid choice to `auto_bill_criteria` that a
    `ValidationError` is thrown.
    """

    billing_control = organization.billing_control

    with pytest.raises(ValidationError) as excinfo:
        billing_control.auto_bill_criteria = "invalid"
        billing_control.full_clean()

    assert excinfo.value.message_dict["auto_bill_criteria"] == [
        "Value 'invalid' is not a valid choice."
    ]


@pytest.mark.parametrize(
    "org, order_transfer_criteria, expected_orders",
    [
        (
            # organization with READY_AND_COMPLETED order_transfer_criteria
            Organization(
                billing_control=BillingControl(
                    order_transfer_criteria=BillingControl.OrderTransferCriteriaChoices.READY_AND_COMPLETED
                )
            ),
            BillingControl.OrderTransferCriteriaChoices.READY_AND_COMPLETED,
            [
                Order(
                    billed=False,
                    status=StatusChoices.COMPLETED,
                    ready_to_bill=True,
                    transferred_to_billing=False,
                    billing_transfer_date=None,
                )
            ],
        ),
        (
            # organization with COMPLETED order_transfer_criteria
            Organization(
                billing_control=BillingControl(
                    order_transfer_criteria=BillingControl.OrderTransferCriteriaChoices.COMPLETED
                )
            ),
            BillingControl.OrderTransferCriteriaChoices.COMPLETED,
            [
                Order(
                    billed=False,
                    status=StatusChoices.COMPLETED,
                    ready_to_bill=False,
                    transferred_to_billing=False,
                    billing_transfer_date=None,
                )
            ],
        ),
        (
            # organization with READY_TO_BILL order_transfer_criteria
            Organization(
                billing_control=BillingControl(
                    order_transfer_criteria=BillingControl.OrderTransferCriteriaChoices.READY_TO_BILL
                )
            ),
            BillingControl.OrderTransferCriteriaChoices.READY_TO_BILL,
            [
                Order(
                    billed=False,
                    status=StatusChoices.IN_PROGRESS,
                    ready_to_bill=True,
                    transferred_to_billing=False,
                    billing_transfer_date=None,
                )
            ],
        ),
        (
            # organization with NO order_transfer_criteria set
            Organization(billing_control=BillingControl(order_transfer_criteria=None)),
            None,
            None,
        ),
    ],
)
def test_get_billable_orders(org, order_transfer_criteria, expected_orders):
    """
    Test that the correct orders are returned when using the `get_billable_orders`
    selector.
    """
    # set the order_transfer_criteria on the organization's billing_control
    org.billing_control.order_transfer_criteria = order_transfer_criteria
    # set the orders on the organization
    org.orders_set = expected_orders

    # call the get_billable_orders function and check the result
    result = selectors.get_billable_orders(organization=org)
    if expected_orders is None:
        assert result is None
    else:
        for order, expected_order in zip(result, expected_orders):
            assert order == expected_order


def test_get_billing_queue_information(order):
    """
    Test that the correct billing queue is returned when using the
    `get_billing_queue_information` selector.
    """
    order.ready_to_bill = True
    order.status = StatusChoices.COMPLETED
    order.save()

    billing_queue = BillingQueue.objects.create(
        organization=order.organization,
        order=order,
    )

    result = selectors.get_billing_queue_information(order=order)
    assert result == billing_queue


def test_cannot_delete_billing_history(organization, order) -> None:
    """
    Test that if the organization has remove_billing_history as false that
    the billing history cannot be deleted.
    """
    organization.billing_control.remove_billing_history = False
    organization.billing_control.save()

    order.billed = True
    order.save()

    billing_history = BillingHistory.objects.create(
        organization=organization, order=order
    )

    with pytest.raises(ValueError) as excinfo:
        billing_history.delete()

    assert (
        excinfo.value.__str__()
        == "Records are not allowed to be removed from billing history."
    )


def test_can_delete_billing_history(organization, order) -> None:
    """
    Test that if the organization has remove_billing_history as true that the billing
    history can be deleted.
    """
    organization.billing_control.remove_billing_history = True
    organization.billing_control.save()

    order.billed = True
    order.save()

    billing_history = BillingHistory.objects.create(
        organization=organization, order=order
    )

    billing_history.delete()

    assert BillingHistory.objects.count() == 0


def test_generate_invoice_number_before_save(order) -> None:
    """
    Test that the invoice number is generated before the save method is called.
    """

    order.status = StatusChoices.COMPLETED
    order.ready_to_bill = True

    billing_queue = BillingQueue.objects.create(
        organization=order.organization, order=order
    )

    assert (
        billing_queue.invoice_number
        == f"{order.organization.invoice_control.invoice_number_prefix}00001"
    )


def test_save_order_details_to_billing_history_before_save(order) -> None:
    """
    Test that the order details are saved to the billing history before the
    save method is called.
    """
    order.billed = True
    order.save()

    billing_history = BillingHistory.objects.create(
        organization=order.organization, order=order
    )

    assert billing_history.pieces == order.pieces
    assert billing_history.order_type == order.order_type
    assert billing_history.weight == order.weight
    assert billing_history.mileage == order.mileage
    assert billing_history.revenue_code == order.revenue_code
    assert billing_history.commodity == order.commodity
    assert billing_history.bol_number == order.bol_number
    assert billing_history.customer == order.customer
    assert billing_history.other_charge_total == order.other_charge_amount
    assert billing_history.freight_charge_amount == order.freight_charge_amount
    assert billing_history.total_amount == order.sub_total