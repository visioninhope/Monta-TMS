# -*- coding: utf-8 -*-
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

import textwrap
from typing import final

from django.db import models
from django.urls import reverse
from django.utils.translation import gettext_lazy as _
from localflavor.us.models import USStateField  # type: ignore

from core.models import GenericModel
from organization.models import Organization


class OrderControl(GenericModel):
    """
    Stores the order control information for a related :model:`organization.Organization`.
    """

    organization = models.OneToOneField(
        Organization,
        on_delete=models.CASCADE,
        verbose_name=_("Organization"),
        related_name="order_control",
        related_query_name="order_controls",
    )
    auto_rate_orders = models.BooleanField(
        _("Auto Rate"),
        default=True,
        help_text=_("Auto rate orders"),
    )
    calculate_distance = models.BooleanField(
        _("Calculate Distance"),
        default=True,
        help_text=_("Calculate distance for the order"),
    )
    enforce_bill_to = models.BooleanField(
        _("Enforce Bill To"),
        default=False,
        help_text=_("Enforce bill to to being enter when entering an order."),
    )
    enforce_rev_code = models.BooleanField(
        _("Enforce Rev Code"),
        default=False,
        help_text=_("Enforce rev code code being entered when entering an order."),
    )
    enforce_shipper = models.BooleanField(
        _("Enforce Shipper"),
        default=False,
        help_text=_("Enforce shipper when putting in an order."),
    )
    enforce_cancel_comm = models.BooleanField(
        _("Enforce Voided Comm"),
        default=False,
        help_text=_("Enforce comment when cancelling an order."),
    )

    generate_routes = models.BooleanField(
        _("Generate Routes"),
        default=False,
        help_text=_("Generate routes for the organization"),
    )

    class Meta:
        """
        Metaclass for OrderControl
        """

        verbose_name: str = _("Order Control")
        verbose_name_plural: str = _("Order Controls")
        ordering: list[str] = ["organization"]

    def __str__(self) -> str:
        """Order control string representation

        Returns:
            str: Order control string representation
        """
        return textwrap.wrap(self.organization.name, 50)[0]

    def get_absolute_url(self) -> str:
        """Order control absolute url

        Returns:
            str: Order control absolute url
        """
        return reverse("order_control:detail", kwargs={"pk": self.pk})


class GoogleAPI(GenericModel):
    """
    Stores the Google API information for a related :model:`organization.Organization`.
    """

    @final
    class GoogleRouteAvoidanceChoices(models.TextChoices):
        """Google Route Avoidance Choices"""

        TOLLS = "tolls", "Tolls"
        HIGHWAYS = "highways", "Highways"
        FERRIES = "ferries", "Ferries"

    @final
    class GoogleRouteModelChoices(models.TextChoices):
        """Google Route Model Choices"""

        BEST_GUESS = "best_guess", "Best Guess"
        OPTIMISTIC = "optimistic", "Optimistic"
        PESSIMISTIC = "pessimistic", "Pessimistic"

    @final
    class GoogleRouteDistanceUnitChoices(models.TextChoices):
        """Google Route Distance Unit Choices"""

        METRIC = "metric", "Metric"
        IMPERIAL = "imperial", "Imperial"

    organization = models.OneToOneField(
        Organization,
        on_delete=models.CASCADE,
        verbose_name=_("Organization"),
        related_name="google_api",
        related_query_name="google_apis",
    )
    api_key = models.CharField(
        _("API Key"),
        max_length=255,
        help_text=_("Google API Key for the organization."),
    )
    mileage_unit = models.CharField(
        _("Mileage Unit"),
        max_length=255,
        choices=GoogleRouteDistanceUnitChoices.choices,
        default=GoogleRouteDistanceUnitChoices.IMPERIAL,
        help_text=_("The mileage unit that the organization uses"),
    )
    traffic_model = models.CharField(
        _("Traffic Model"),
        max_length=255,
        choices=GoogleRouteModelChoices.choices,
        default=GoogleRouteModelChoices.BEST_GUESS,
        help_text=_("The traffic model that the organization uses"),
    )
    add_customer_location = models.BooleanField(
        _("Add Customer Location"),
        default=False,
        help_text=_("Add customer location through google places"),
    )
    add_location = models.BooleanField(
        _("Add Location"),
        default=False,
        help_text=_("Add location through google places"),
    )

    class Meta:
        """
        Metaclass for GoogleAPI
        """

        verbose_name = _("Google API")
        verbose_name_plural = _("Google APIs")
        ordering: list[str] = ["organization"]

    def __str__(self) -> str:
        """Google API string representation

        Returns:
            str: Google API string representation
        """
        return textwrap.wrap(self.organization.name, 50)[0]

    def get_absolute_url(self) -> str:
        """Google API absolute url

        Returns:
            str: Google API absolute url
        """
        return reverse("google_api:detail", kwargs={"pk": self.pk})


class CommentType(GenericModel):
    """
    Stores the comment type information for a related :model:`organization.Organization`.
    """

    name = models.CharField(
        _("Name"),
        max_length=255,
        help_text=_("Comment type name"),
    )
    description = models.TextField(
        _("Description"),
        max_length=255,
        help_text=_("Comment type description"),
    )

    class Meta:
        """
        Metaclass for CommentType
        """

        verbose_name: str = _("Comment Type")
        verbose_name_plural: str = _("Comment Types")
        ordering: list[str] = ["organization"]

    def __str__(self) -> str:
        """Comment type string representation

        Returns:
            str: Comment type string representation
        """
        return textwrap.wrap(self.name, 50)[0]

    def get_absolute_url(self) -> str:
        """Comment type absolute url

        Returns:
            str: Comment type absolute url
        """
        return reverse("comment_type:detail", kwargs={"pk": self.pk})
