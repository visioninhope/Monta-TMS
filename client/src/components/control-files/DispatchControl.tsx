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

import { usePageStyles } from "@/styles/PageStyles";
import { useQuery, useQueryClient } from "react-query";
import { getDispatchControl } from "@/requests/OrganizationRequestFactory";
import { Card, Divider, Skeleton, Text } from "@mantine/core";
import React from "react";
import { DispatchControlForm } from "@/components/control-files/_partials/DispatchControlForm";

const DispatchControlPage = () => {
  const { classes } = usePageStyles();
  const queryClient = useQueryClient();

  const { data: dispatchControlData, isLoading: isDispatchControlDataLoading } =
    useQuery({
      queryKey: ["dispatchControl"],
      queryFn: () => getDispatchControl(),
      initialData: () => {
        return queryClient.getQueryData(["dispatchControl"]);
      },
      staleTime: Infinity,
    });

  // Store first element of dispatchControlData in variable
  const dispatchControlDataArray = dispatchControlData?.[0];

  return (
    <>
      {isDispatchControlDataLoading ? (
        <Skeleton height={400}></Skeleton>
      ) : (
        <Card className={classes.card} withBorder>
          <Text fz="xl" fw={700} className={classes.text}>
            Dispatch Controls
          </Text>

          <Divider my={10} />
          {dispatchControlDataArray && (
            <DispatchControlForm dispatchControl={dispatchControlDataArray} />
          )}
        </Card>
      )}
    </>
  );
};

export default DispatchControlPage;