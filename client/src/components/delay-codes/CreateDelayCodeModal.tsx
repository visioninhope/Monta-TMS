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

import React from "react";
import { Button, Group, Modal, SimpleGrid } from "@mantine/core";
import { notifications } from "@mantine/notifications";
import { useForm, UseFormReturnType, yupResolver } from "@mantine/form";
import { useDelayCodeStore as store } from "@/stores/DispatchStore";
import { useFormStyles } from "@/assets/styles/FormStyles";
import { ValidatedTextInput } from "@/components/common/fields/TextInput";
import { ValidatedTextArea } from "@/components/common/fields/TextArea";
import { DelayCode, DelayCodeFormValues as FormValues } from "@/types/dispatch";
import { delayCodeSchema } from "@/lib/schemas/DispatchSchema";
import { SwitchInput } from "@/components/common/fields/SwitchInput";
import { useCustomMutation } from "@/hooks/useCustomMutation";
import { TableStoreProps } from "@/types/tables";
import { statusChoices } from "@/lib/constants";
import { SelectInput } from "@/components/common/fields/SelectInput";

export function DelayCodeForm({
  form,
}: {
  form: UseFormReturnType<FormValues>;
}) {
  const { classes } = useFormStyles();

  return (
    <div className={classes.div}>
      <SimpleGrid cols={2} breakpoints={[{ maxWidth: "sm", cols: 1 }]}>
        <SelectInput<FormValues>
          data={statusChoices}
          name="status"
          placeholder="Status"
          label="Status"
          description="Status of the delay code"
          form={form}
          variant="filled"
          withAsterisk
        />
        <ValidatedTextInput<FormValues>
          form={form}
          name="code"
          label="Code"
          maxLength={4}
          placeholder="Code"
          description="Unique Code of the delay code"
          withAsterisk
        />
      </SimpleGrid>
      <ValidatedTextArea<FormValues>
        form={form}
        name="description"
        label="Description"
        description="Description of the delay code."
        placeholder="Description"
        withAsterisk
      />
      <SwitchInput<FormValues>
        form={form}
        name="fCarrierOrDriver"
        label="F. Carrier/Driver"
        description="Fault of carrier or driver?"
      />
    </div>
  );
}

export function CreateDelayCodeModalForm() {
  const { classes } = useFormStyles();
  const [loading, setLoading] = React.useState<boolean>(false);

  const form = useForm<FormValues>({
    validate: yupResolver(delayCodeSchema),
    initialValues: {
      status: "A",
      code: "",
      description: "",
      fCarrierOrDriver: false,
    },
  });

  const mutation = useCustomMutation<FormValues, TableStoreProps<DelayCode>>(
    form,
    store,
    notifications,
    {
      method: "POST",
      path: "/delay_codes/",
      successMessage: "Delay Code created successfully.",
      queryKeysToInvalidate: ["delay-code-table-data"],
      additionalInvalidateQueries: ["delayCodes"],
      closeModal: true,
      errorMessage: "Failed to create delay code.",
    },
    () => setLoading(false),
  );

  const submitForm = (values: FormValues) => {
    setLoading(true);
    mutation.mutate(values);
  };

  return (
    <form onSubmit={form.onSubmit((values) => submitForm(values))}>
      <DelayCodeForm form={form} />
      <Group position="right" mt="md">
        <Button type="submit" className={classes.control} loading={loading}>
          Submit
        </Button>
      </Group>
    </form>
  );
}

export function CreateDelayCodeModal() {
  const [showCreateModal, setShowCreateModal] = store.use("createModalOpen");

  return (
    <Modal.Root
      opened={showCreateModal}
      onClose={() => setShowCreateModal(false)}
    >
      <Modal.Overlay />
      <Modal.Content>
        <Modal.Header>
          <Modal.Title>Create Delay Code</Modal.Title>
          <Modal.CloseButton />
        </Modal.Header>
        <Modal.Body>
          <CreateDelayCodeModalForm />
        </Modal.Body>
      </Modal.Content>
    </Modal.Root>
  );
}
