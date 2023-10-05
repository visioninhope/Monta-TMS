/*
 * COPYRIGHT(c) 2023 MONTA
 *
 * This file is part of Monta.
 *
 * The Monta software is licensed under the Business Source License 1.1. You are granted the right
 * to copy, modify, and redistribute the software, but only for non-production use or with a total
 * of less than three server instances. Starting from the Change Date (November 16, 2026), the
 * software will be made available under version 2 or later of the GNU General Public License.
 * If you use the software in violation of this license, your rights under the license will be
 * terminated automatically. The software is provided "as is," and the Licensor disclaims all
 * warranties and conditions. If you use this license's text or the "Business Source License" name
 * and trademark, you must comply with the Licensor's covenants, which include specifying the
 * Change License as the GPL Version 2.0 or a compatible license, specifying an Additional Use
 * Grant, and not modifying the license in any other way.
 */

import React, { useMemo } from "react";
import { MRT_ColumnDef } from "mantine-react-table";
import { Badge } from "@mantine/core";
import { MontaTable } from "@/components/common/table/MontaTable";
import { EquipmentManufacturer } from "@/types/equipment";
import { useEquipManufacturerTableStore as store } from "@/stores/EquipmentStore";
import { CreateEquipManufacturerModal } from "@/components/equipment-manufacturer/CreateEquipManfacturerModal";
import { EquipManufacturerDrawer } from "@/components/equipment-manufacturer/EquipManufacturerDrawer";
import { TChoiceProps } from "@/types";

export function EquipManufacturerTable() {
  const columns = useMemo<MRT_ColumnDef<EquipmentManufacturer>[]>(
    () => [
      {
        id: "status",
        accessorKey: "status",
        header: "Status",
        filterFn: "equals",
        Cell: ({ cell }) => (
          <Badge
            color={cell.getValue() === "A" ? "green" : "red"}
            variant="filled"
            radius="xs"
          >
            {cell.getValue() === "A" ? "Active" : "Inactive"}
          </Badge>
        ),
        mantineFilterSelectProps: {
          data: [
            { value: "", label: "All" },
            { value: "A", label: "Active" },
            { value: "I", label: "Inactive" },
          ] as TChoiceProps[],
        },
        filterVariant: "select",
      },
      {
        accessorKey: "name", // access nested data with dot notation
        header: "Name",
      },
      {
        accessorKey: "description",
        header: "Description",
      },
    ],
    [],
  );

  return (
    <MontaTable
      store={store}
      link="/equipment_manufacturers"
      columns={columns}
      displayDeleteModal
      TableCreateDrawer={CreateEquipManufacturerModal}
      TableDrawer={EquipManufacturerDrawer}
      tableQueryKey="equipment-manufacturer-table-data"
      exportModelName="EquipmentManufacturer"
      name="Equipment Manufacturer"
    />
  );
}
