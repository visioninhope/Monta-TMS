/*
 * COPYRIGHT(c) 2024 Trenova
 *
 * This file is part of Trenova.
 *
 * The Trenova software is licensed under the Business Source License 1.1. You are granted the right
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

import { TableChangeAlertForm } from "@/components/table-change-alerts/table-change-sheet";
import { Button } from "@/components/ui/button";
import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetFooter,
  SheetHeader,
  SheetTitle,
} from "@/components/ui/sheet";
import { useCustomMutation } from "@/hooks/useCustomMutation";
import { formatDate } from "@/lib/date";
import { cn } from "@/lib/utils";
import { tableChangeAlertSchema } from "@/lib/validations/OrganizationSchema";
import { useTableStore } from "@/stores/TableStore";
import {
  TableChangeAlert,
  TableChangeAlertFormValues,
} from "@/types/organization";
import { TableSheetProps } from "@/types/tables";
import { yupResolver } from "@hookform/resolvers/yup";
import React from "react";
import { useForm } from "react-hook-form";
function TableChangeEditForm({
  tableChangeAlert,
  open,
  onOpenChange,
}: {
  tableChangeAlert: TableChangeAlert;
  open: boolean;
  onOpenChange: (open: boolean) => void;
}) {
  const [isSubmitting, setIsSubmitting] = React.useState<boolean>(false);

  const { handleSubmit, control, reset, watch } =
    useForm<TableChangeAlertFormValues>({
      resolver: yupResolver(tableChangeAlertSchema),
      defaultValues: {
        status: tableChangeAlert.status,
        name: tableChangeAlert.name,
        source: tableChangeAlert.source,
        databaseAction: tableChangeAlert.databaseAction,
        description: tableChangeAlert.description,
        table: tableChangeAlert?.table || "",
        topic: tableChangeAlert.topic || "",
        emailProfile: tableChangeAlert.emailProfile,
        emailRecipients: tableChangeAlert.emailRecipients,
        conditionalLogic: tableChangeAlert.conditionalLogic,
        customSubject: tableChangeAlert.customSubject,
        effectiveDate: tableChangeAlert.effectiveDate,
        expirationDate: tableChangeAlert.expirationDate,
      },
    });

  const mutation = useCustomMutation<TableChangeAlertFormValues>(
    control,
    {
      method: "PUT",
      path: `/table_change_alerts/${tableChangeAlert.id}/`,
      successMessage: "Table Change Alert updated successfully.",
      queryKeysToInvalidate: ["table-change-alert-data"],
      closeModal: true,
      errorMessage: "Failed to update table change alert.",
    },
    () => setIsSubmitting(false),
    reset,
  );

  const onSubmit = (values: TableChangeAlertFormValues) => {
    setIsSubmitting(true);
    mutation.mutate(values);
  };

  return (
    <form
      onSubmit={handleSubmit(onSubmit)}
      className="flex h-full flex-col overflow-y-auto"
    >
      <TableChangeAlertForm control={control} open={open} watch={watch} />
      <SheetFooter className="mb-12">
        <Button
          type="reset"
          variant="secondary"
          onClick={() => onOpenChange(false)}
          className="w-full"
        >
          Cancel
        </Button>
        <Button type="submit" isLoading={isSubmitting} className="w-full">
          Save Changes
        </Button>
      </SheetFooter>
    </form>
  );
}

export function TableChangeAlertEditSheet({
  onOpenChange,
  open,
}: TableSheetProps) {
  const [tableChangeAlert] = useTableStore.use(
    "currentRecord",
  ) as TableChangeAlert[];

  if (!tableChangeAlert) return null;

  return (
    <Sheet open={open} onOpenChange={onOpenChange}>
      <SheetContent className={cn("w-full xl:w-1/2")}>
        <SheetHeader>
          <SheetTitle>{tableChangeAlert && tableChangeAlert.name}</SheetTitle>
          <SheetDescription>
            Last updated on{" "}
            {tableChangeAlert && formatDate(tableChangeAlert.modified)}
          </SheetDescription>
        </SheetHeader>
        {tableChangeAlert && (
          <TableChangeEditForm
            tableChangeAlert={tableChangeAlert}
            open={open}
            onOpenChange={onOpenChange}
          />
        )}
      </SheetContent>
    </Sheet>
  );
}
