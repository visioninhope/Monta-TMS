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


class MovementFactory(factory.django.DjangoModelFactory):
    """
    Factory for Movement Model.
    """

    class Meta:
        """
        Metaclass for MovementFactory
        """

        model = "movements.Movement"
        django_get_or_create = (
            "organization",
            "order",
            "equipment",
            "primary_worker",
        )

    id = factory.Faker("uuid4")
    organization = factory.SubFactory("organization.factories.OrganizationFactory")
    order = factory.SubFactory("order.tests.factories.OrderFactory")
    equipment = factory.SubFactory("equipment.tests.factories.EquipmentFactory")
    primary_worker = factory.SubFactory("worker.factories.WorkerFactory")
    secondary_worker = factory.SubFactory("worker.factories.WorkerFactory")
    status = "N"
