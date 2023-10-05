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

import { Button, Group, Modal, SimpleGrid } from "@mantine/core";
import React from "react";
import { notifications } from "@mantine/notifications";
import { useForm, UseFormReturnType, yupResolver } from "@mantine/form";
import { accessorialChargeTableStore as store } from "@/stores/BillingStores";
import { useFormStyles } from "@/assets/styles/FormStyles";
import {
  AccessorialCharge,
  AccessorialChargeFormValues as FormValues,
} from "@/types/billing";
import { accessorialChargeSchema as Schema } from "@/lib/schemas/BillingSchema";
import { ValidatedTextInput } from "@/components/common/fields/TextInput";
import { ValidatedTextArea } from "@/components/common/fields/TextArea";
import { SelectInput } from "@/components/common/fields/SelectInput";
import { fuelMethodChoices } from "@/utils/apps/billing";
import { SwitchInput } from "@/components/common/fields/SwitchInput";
import { useCustomMutation } from "@/hooks/useCustomMutation";
import { TableStoreProps } from "@/types/tables";
import { statusChoices } from "@/lib/constants";

export function ACForm({ form }: { form: UseFormReturnType<FormValues> }) {
  const { classes } = useFormStyles();

  return (
    <div className={classes.div}>
      <SimpleGrid cols={2} breakpoints={[{ maxWidth: "sm", cols: 1 }]}>
        <SelectInput<FormValues>
          form={form}
          data={statusChoices}
          name="status"
          placeholder="Status"
          label="Status"
          description="Status of the accessorial charge"
          variant="filled"
          withAsterisk
        />
        <ValidatedTextInput<FormValues>
          form={form}
          className={classes.fields}
          name="code"
          label="Code"
          description="Code for the accessorial charge"
          placeholder="Code"
          variant="filled"
          withAsterisk
        />
      </SimpleGrid>
      <ValidatedTextArea<FormValues>
        form={form}
        className={classes.fields}
        name="description"
        label="Description"
        description="Description of the accessorial charge"
        placeholder="Description"
        variant="filled"
      />
      <SimpleGrid cols={2} breakpoints={[{ maxWidth: "sm", cols: 1 }]}>
        <ValidatedTextInput<FormValues>
          form={form}
          className={classes.fields}
          name="chargeAmount"
          label="Charge Amount"
          placeholder="Charge Amount"
          description="Charge amount for the accessorial charge"
          variant="filled"
          withAsterisk
        />
        <SelectInput<FormValues>
          form={form}
          data={fuelMethodChoices}
          className={classes.fields}
          name="method"
          label="Fuel Method"
          description="Method for calculating the accessorial charge"
          placeholder="Fuel Method"
          variant="filled"
          withAsterisk
        />
        <SwitchInput<FormValues>
          form={form}
          className={classes.fields}
          name="isDetention"
          label="Detention"
          description="Is this a detention charge?"
          placeholder="Detention"
          variant="filled"
        />
      </SimpleGrid>
    </div>
  );
}

export function CreateACModalForm() {
  const { classes } = useFormStyles();
  const [loading, setLoading] = React.useState<boolean>(false);

  const form = useForm<FormValues>({
    validate: yupResolver(Schema),
    initialValues: {
      status: "A",
      code: "",
      description: "",
      isDetention: false,
      chargeAmount: 1,
      method: "D",
    },
  });

  const mutation = useCustomMutation<
    FormValues,
    TableStoreProps<AccessorialCharge>
  >(
    form,
    store,
    notifications,
    {
      method: "POST",
      path: "/accessorial_charges/",
      successMessage: "Accessorial Charge created successfully.",
      queryKeysToInvalidate: ["accessorial-charges-table-data"],
      additionalInvalidateQueries: ["accessorialCharges"],
      closeModal: true,
      errorMessage: "Failed to create accessorial charge.",
    },
    () => setLoading(false),
  );

  const submitForm = (values: FormValues) => {
    setLoading(true);
    mutation.mutate(values);
  };

  return (
    <form onSubmit={form.onSubmit((values) => submitForm(values))}>
      <ACForm form={form} />
      <Group position="right" mt="md">
        <Button
          color="white"
          type="submit"
          className={classes.control}
          loading={loading}
        >
          Submit
        </Button>
      </Group>
    </form>
  );
}

export function CreateACModal() {
  const [showCreateModal, setShowCreateModal] = store.use("createModalOpen");

  return (
    <Modal.Root
      opened={showCreateModal}
      onClose={() => setShowCreateModal(false)}
      size="30%"
    >
      <Modal.Overlay />
      <Modal.Content>
        <Modal.Header>
          <Modal.Title>Create Accessorial Charge</Modal.Title>
          <Modal.CloseButton />
        </Modal.Header>
        <Modal.Body>
          <CreateACModalForm />
        </Modal.Body>
      </Modal.Content>
    </Modal.Root>
  );
}
