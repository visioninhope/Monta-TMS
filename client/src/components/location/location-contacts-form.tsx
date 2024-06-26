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

import { type LocationFormValues as FormValues } from "@/types/location";
import { faOctagonExclamation } from "@fortawesome/pro-duotone-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { useFieldArray, type Control } from "react-hook-form";
import { InputField } from "../common/fields/input";
import { Button } from "../ui/button";
import { Form, FormControl, FormGroup } from "../ui/form";
import { ScrollArea } from "../ui/scroll-area";

export function LocationContactForm({
  control,
}: {
  control: Control<FormValues>;
}) {
  const { fields, append, remove } = useFieldArray({
    control,
    name: "contacts",
    keyName: "id",
  });

  const handleAddContact = () => {
    append({ name: "", emailAddress: "", phoneNumber: "" });
  };

  return (
    <Form className="flex size-full flex-col">
      {fields.length > 0 ? (
        <>
          <ScrollArea className="h-[70vh] p-4">
            {fields.map((field, index) => (
              <FormGroup key={field.id} className="border-b pb-2">
                <FormControl>
                  <InputField
                    control={control}
                    name={`contacts.${index}.name`}
                    label="Name"
                    placeholder="Name"
                    description="Enter the full name of the primary contact for this location."
                    rules={{ required: true }}
                  />
                </FormControl>
                <FormControl>
                  <InputField
                    control={control}
                    name={`contacts.${index}.emailAddress`}
                    label="Email Address"
                    placeholder="Email Address"
                    description="Provide the email address for direct communication with the location's contact."
                  />
                </FormControl>
                <FormControl>
                  <InputField
                    control={control}
                    name={`contacts.${index}.phoneNumber`}
                    label="Phone"
                    placeholder="Phone"
                    description="Input the telephone number for reaching the location's contact."
                  />
                </FormControl>
                <div className="flex max-w-sm flex-col justify-between gap-1">
                  <div className="min-h-[2em]">
                    <Button
                      size="sm"
                      className="bg-background hover:bg-background text-red-600 hover:text-red-700"
                      type="button"
                      onClick={() => remove(index)}
                    >
                      Remove
                    </Button>
                  </div>
                </div>
              </FormGroup>
            ))}
          </ScrollArea>
          <Button
            type="button"
            size="sm"
            className="my-4 w-[200px]"
            onClick={handleAddContact}
          >
            Add Another Contact
          </Button>
        </>
      ) : (
        <div className="mt-48 flex grow flex-col items-center justify-center">
          <span className="mb-4 text-6xl">
            <FontAwesomeIcon icon={faOctagonExclamation} />
          </span>
          <p className="mb-4">No contacts yet. Please add a new contact.</p>
          <Button type="button" size="sm" onClick={handleAddContact}>
            Add Contact
          </Button>
        </div>
      )}
    </Form>
  );
}
