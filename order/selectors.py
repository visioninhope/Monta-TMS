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

from django.db.models import QuerySet

from movements.models import Movement
from order import models
from stops.models import Stop
from utils.types import MODEL_UUID


def get_order_by_id(*, order_id: MODEL_UUID) -> models.Order:
    """Get an Order model instance by its ID.

    Args:
        order_id (str): The ID of the order.

    Returns:
        models.Order: The Order model instance.
    """
    return models.Order.objects.get(id=order_id)

def get_order_movements(*, order: models.Order) -> QuerySet[Movement]:
    """Get the movements of an order.

    Args:
        order (models.Order): The order.

    Returns:
        QuerySet[Movement]: QuerySet of the movements of the order.
    """
    return Movement.objects.filter(order=order)

def get_order_stops(*, order: models.Order) -> QuerySet[Stop]:
    """Get the stops of an order.

    Args:
        order (models.Order): The order.

    Returns:
        QuerySett[Stop]: QuerySet of the stops of the order.
    """
    movements = get_order_movements(order=order)
    return Stop.objects.filter(movement__in=movements).select_related('movement')