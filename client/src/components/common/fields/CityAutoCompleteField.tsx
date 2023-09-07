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
import React, { useRef, useState } from "react";
import { Autocomplete, Loader } from "@mantine/core";
import { IconAlertTriangle } from "@tabler/icons-react";
import { AutocompleteProps } from "@mantine/core/lib/Autocomplete/Autocomplete";
import { UseFormReturnType } from "@mantine/form";
import { stateData } from "./StateSelect";
import { useFormStyles } from "@/assets/styles/FormStyles";

interface CityAutoCompleteFieldProps<TFormValues extends object>
  extends Omit<AutocompleteProps, "data" | "form" | "name"> {
  form: UseFormReturnType<TFormValues, (values: TFormValues) => TFormValues>;
  stateSelection: string;
  name: keyof TFormValues;
}

export function CityAutoCompleteField<TFormValues extends object>({
  form,
  stateSelection,
  name,
  ...rest
}: CityAutoCompleteFieldProps<TFormValues>) {
  const { classes } = useFormStyles();
  const error = form.errors[name as string];
  const timeoutRef = useRef<number>(-1);
  const [loading, setLoading] = useState<boolean>(false);
  const [data, setData] = useState<string[]>([]);

  const getCity = async () => {
    const selectedState = stateData.find(
      (state) => state.value === stateSelection,
    );
    const stateLabel = selectedState ? selectedState.label : "";
    const response = await fetch(
      "https://countriesnow.space/api/v0.1/countries/state/cities",
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          country: "United States",
          state: stateLabel,
        }),
      },
    );
    return response.json();
  };

  const handleChange = (val: any) => {
    window.clearTimeout(timeoutRef.current);
    form.setFieldValue(name as string, val);
    setData([]);

    if (val.trim().length === 0) {
      setLoading(false);
    } else {
      setLoading(true);
      timeoutRef.current = window.setTimeout(() => {
        setLoading(false);
        getCity().then((res) => {
          setData(res.data);
        });
      }, 500);
    }
  };

  return (
    <Autocomplete
      {...rest}
      {...form.getInputProps(name as string)}
      data={data ?? []}
      error={error}
      className={classes.fields}
      onChange={handleChange}
      rightSection={
        loading ? (
          <Loader size={24} />
        ) : (
          error && (
            <IconAlertTriangle
              stroke={1.5}
              size="1.1rem"
              className={classes.invalidIcon}
            />
          )
        )
      }
    />
  );
}