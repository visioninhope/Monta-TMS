# --------------------------------------------------------------------------------------------------
#  COPYRIGHT(c) 2023 MONTA                                                                         -
#                                                                                                  -
#  This file is part of Monta.                                                                     -
#                                                                                                  -
#  The Monta software is licensed under the Business Source License 1.1. You are granted the right -
#  to copy, modify, and redistribute the software, but only for non-production use or with a total -
#  of less than three server instances. Starting from the Change Date (November 16, 2026), the     -
#  software will be made available under version 2 or later of the GNU General Public License.     -
#  If you use the software in violation of this license, your rights under the license will be     -
#  terminated automatically. The software is provided "as is," and the Licensor disclaims all      -
#  warranties and conditions. If you use this license's text or the "Business Source License" name -
#  and trademark, you must comply with the Licensor's covenants, which include specifying the      -
#  Change License as the GPL Version 2.0 or a compatible license, specifying an Additional Use     -
#  Grant, and not modifying the license in any other way.                                          -
# --------------------------------------------------------------------------------------------------
from io import BytesIO

from django.shortcuts import render
from django.template.loader import get_template
from django.utils import timezone
from weasyprint import HTML

from billing.models import BillingQueue
from invoicing import models
from order.selectors import get_order_stops


def render_invoice_to_pdf(
    *, template_src: str, invoice: BillingQueue, context_dict: dict | None = None
):
    """Render invoice to pdf file.

    Args:
        invoice (BillingQueue): The invoice to render.

    Returns:
        str: The file name of the generated pdf file.
    """
    if not isinstance(template_src, str):
        raise ValueError("Invalid template name")

    if context_dict is None:
        context_dict = {}

    template = get_template(template_src)
    html = template.render(context_dict)

    pdf_file = BytesIO()
    HTML(string=html).write_pdf(pdf_file)

    file_name = f"invoice_{invoice.invoice_number}.pdf"

    return file_name, pdf_file


def generate_invoice(*, invoice_control: models.InvoiceControl, invoice: BillingQueue):
    """Generate invoice from invoice control and invoice.

    Args:
        invoice_control (InvoiceControl): The invoice control to use.
        invoice (BillingQueue): The invoice to generate.
    """
    bill_date = timezone.now().date()

    file_name, pdf_file = render_invoice_to_pdf(
        template_src="pdf/invoice-1.html",
        invoice=invoice,
        context_dict={
            "invoice": invoice,
            "bill_date": bill_date,
            "invoice_control": invoice_control,
            "user": invoice.user,
        },
    )

    return invoice


def invoice_view(request, invoice_id):
    invoice = BillingQueue.objects.get(pk=invoice_id)
    invoice_control = invoice.organization.invoice_control
    bill_date = timezone.now().date()

    # Get stop count for invoice order
    stop_count = get_order_stops(order=invoice.order).count()

    return render(
        request,
        "pdf/invoice-1.html",
        {
            "invoice": invoice,
            "bill_date": bill_date,
            "stop_count": stop_count,
            "invoice_control": invoice_control,
            "user": invoice.user,
        },
    )
