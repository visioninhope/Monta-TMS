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

import textwrap
import uuid
from typing import final, Any

from django.db import models
from django.urls import reverse
from django.utils.translation import gettext_lazy as _

from utils.models import ChoiceField, GenericModel


class Route(GenericModel):
    """
    Stores route information related to `orders.Orders` model
    """

    id = models.UUIDField(
        primary_key=True,
        default=uuid.uuid4,
        editable=False,
        unique=True,
    )
    origin_location = models.ForeignKey(
        "location.Location",
        on_delete=models.CASCADE,
        related_name="origin_location",
        help_text=_("Origin of the route"),
        verbose_name=_("Origin Location"),
    )
    destination_location = models.ForeignKey(
        "location.Location",
        on_delete=models.CASCADE,
        related_name="destination_location",
        help_text=_("Destination of the route"),
        verbose_name=_("Destination Location"),
    )
    total_mileage = models.FloatField(
        _("Total Mileage"),
        blank=True,
        null=True,
        help_text=_("Total Mile from origin to destination"),
    )
    duration = models.PositiveIntegerField(
        _("Duration"),
        null=True,
        blank=True,
        help_text=_("Duration of route from origin to destination"),
    )
    distance_method = models.CharField(
        _("Distance Model"),
        max_length=50,
        blank=True,
        help_text=_("Distance model used to calculate the route"),
    )

    class Meta:
        """
        Metaclass for Route
        """

        verbose_name = _("Route")
        verbose_name_plural = _("Routes")
        ordering = ("origin_location", "destination_location")
        indexes = [
            models.Index(fields=["total_mileage", "duration"]),
        ]
        db_table = "route"

    def __str__(self) -> str:
        """Route string representation

        Returns:
            str: Route string representation
        """
        return textwrap.wrap(
            f"{self.origin_location} - {self.destination_location}", 50
        )[0]

    def get_absolute_url(self) -> str:
        """Route absolute URL

        Returns:
            str: Route absolute URL
        """
        return reverse("route:detail", kwargs={"pk": self.pk})


class RouteControl(GenericModel):
    """
    Store Route Control information related to an `organization.Organization`
    """

    @final
    class RouteAvoidanceChoices(models.TextChoices):
        """
        Google Route Avoidance Choices
        """

        TOLLS = "tolls", "Tolls"
        HIGHWAYS = "highways", "Highways"
        FERRIES = "ferries", "Ferries"

    @final
    class RouteModelChoices(models.TextChoices):
        """
        Google Route Model Choices
        """

        BEST_GUESS = "best_guess", "Best Guess"
        OPTIMISTIC = "optimistic", "Optimistic"
        PESSIMISTIC = "pessimistic", "Pessimistic"

    @final
    class RouteDistanceUnitChoices(models.TextChoices):
        """
        Google Route Distance Unit Choices
        """

        METRIC = "metric", "Metric"
        IMPERIAL = "imperial", "Imperial"

    @final
    class DistanceMethodChoices(models.TextChoices):
        """
        Distance method choices for Order model
        """

        GOOGLE = "Google", _("Google")
        MONTA = "Monta", _("Monta")

    id = models.UUIDField(
        primary_key=True,
        default=uuid.uuid4,
        editable=False,
        unique=True,
    )
    organization = models.OneToOneField(
        "organization.Organization",
        on_delete=models.CASCADE,
        related_name="route_control",
        verbose_name=_("Organization"),
        help_text=_("Organization related to this route control"),
    )
    distance_method = ChoiceField(
        _("Distance Method"),
        choices=DistanceMethodChoices.choices,
        default=DistanceMethodChoices.MONTA,
        help_text=_("Distance method for the company."),
    )
    mileage_unit = ChoiceField(
        _("Mileage Unit"),
        choices=RouteDistanceUnitChoices.choices,
        default=RouteDistanceUnitChoices.IMPERIAL,
        help_text=_("The mileage unit that the organization uses"),
    )
    generate_routes = models.BooleanField(
        _("Generate Routes"),
        default=False,
        help_text=_("Automatically generate routes for orders"),
    )

    class Meta:
        """
        Metaclass for RouteControl
        """

        verbose_name = _("Route Control")
        verbose_name_plural = _("Route Controls")
        ordering = ("organization",)
        db_table = "route_control"

    def __str__(self) -> str:
        """Route Control string representation

        Returns:
            str: Route Control string representation
        """
        return str(self.organization)

    def get_absolute_url(self) -> str:
        """Route Control absolute URL

        Returns:
            str: Route Control absolute URL
        """
        return reverse("route:control-detail", kwargs={"pk": self.pk})
