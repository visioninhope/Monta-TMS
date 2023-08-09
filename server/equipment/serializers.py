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

from typing import Any

from equipment import models
from utils.serializers import GenericSerializer


class EquipmentTypeDetailSerializer(GenericSerializer):
    """A serializer for the EquipmentTypeDetail model

    The serializer provides default operations for creating, update and deleting
    Equipment Type Detail, as well as listing and retrieving them.
    """

    class Meta:
        """
        A class representing the metadata for the `EquipmentTypeDetailSerializer`
        class.
        """

        model = models.EquipmentTypeDetail
        extra_read_only_fields = ("equipment_type",)


class EquipmentTypeSerializer(GenericSerializer):
    """A serializer for the EquipmentType model.

    The serializer provides default operations for creating, updating, and deleting
    Equipment Types, as well as listing and retrieving them.
    """

    equipment_type_details = EquipmentTypeDetailSerializer(required=False)

    class Meta:
        """
        A class representing the metadata for the `EquipmentTypeSerializer` class.
        """

        model = models.EquipmentType
        extra_fields = ("equipment_type_details",)

    def create(self, validated_data: Any) -> models.EquipmentType:
        """Create new Equipment Type

        Args:
            validated_data (Any): Validated data

        Returns:
            models.EquipmentType: Created EquipmentType
        """
        detail_data = validated_data.pop("equipment_type_details", {})
        organization = super().get_organization
        business_unit = super().get_business_unit

        equipment_type = models.EquipmentType.objects.create(
            organization=organization, business_unit=business_unit, **validated_data
        )

        if detail_data:
            if details := models.EquipmentTypeDetail.objects.get(  # type: ignore
                organization=organization,
                business_unit=business_unit,
                equipment_type=equipment_type,
            ):
                details.delete()

            models.EquipmentTypeDetail.objects.create(
                organization=organization,
                business_unit=business_unit,
                equipment_type=equipment_type,
                **detail_data,
            )

        return equipment_type

    def update(  # type: ignore
        self, instance: models.EquipmentType, validated_data: Any
    ) -> models.EquipmentType:
        """Update Equipment Type

        Args:
            instance (models.EquipmentType): EquipmentType instance
            validated_data (Any): Validated data

        Returns:
            models.EquipmentType: Updated EquipmentType
        """

        detail_data = validated_data.pop("equipment_type_details", {})

        instance.name = validated_data.get("name", instance.name)
        instance.description = validated_data.get("description", instance.description)
        instance.save()

        if detail_data:
            instance.equipment_type_details.update_details(**detail_data)

        return instance


class EquipmentManufacturerSerializer(GenericSerializer):
    """A serializer for the EquipmentManufacturer Model

    The serializer provides default operations for creating, update and deleting
    Equipment Manufacturer, as well as listing and retrieving them.
    """

    class Meta:
        """
        A class representing the metadata for the `EquipmentManufacturerSerializer`
        class.
        """

        model = models.EquipmentManufacturer


class TractorSerializer(GenericSerializer):
    """A serializer for the Tractor model

    The serializer provides default operations for creating, update and deleting
    Tractors, as well as listing and retrieving them.
    """

    class Meta:
        """
        A class representing the metadata for the `TractorSerializer` class.
        """

        model = models.Tractor


class TrailerSerializer(GenericSerializer):
    """A serializer for the Trailer model

    The serializer provides default operations for creating, update and deleting
    Trailers, as well as listing and retrieving them.
    """

    class Meta:
        """
        A class representing the metadata for the `TrailerSerializer` class.
        """

        model = models.Trailer


class EquipmentMaintenancePlanSerializer(GenericSerializer):
    """A serializer for the EquipmentMaintenancePlan model

    The serializer provides default operations for creating, update and deleting
    Equipment Maintenance Plan, as well as listing and retrieving them.
    """

    class Meta:
        """
        A class representing the metadata for the `EquipmentMaintenancePlanSerializer`
        class.
        """

        model = models.EquipmentMaintenancePlan