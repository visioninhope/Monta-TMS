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

import factory


class OrganizationFactory(factory.django.DjangoModelFactory):
    """
    Organization factory class
    """

    class Meta:
        """
        Metaclass for OrganizationFactory
        """

        model = "organization.Organization"

    name = factory.Faker("company", locale="en_US")
    scac_code = "RNDM"


class DepotFactory(factory.django.DjangoModelFactory):
    """
    Depot factory class
    """

    class Meta:
        """
        Metaclass for DepotFactory
        """

        model = "organization.Depot"
        django_get_or_create = ("organization",)

    name = factory.Faker("company", locale="en_US")
    organization = factory.SubFactory(OrganizationFactory)
