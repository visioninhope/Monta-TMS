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
import { Card, Flex, Text } from "@mantine/core";

export function LoadingScreen(): React.ReactElement {
  return (
    <Flex
      direction={{ base: "column", sm: "row" }}
      justify={{ sm: "center" }}
      align={{ sm: "center" }}
      style={{ height: "90vh" }}
    >
      <Card padding="xl">
        <Text weight={500} size="lg">
          Monta is loading. Please wait.
        </Text>
        <Text mt="xs" color="dimmed" size="sm">
          If the operation exceeds a duration of 10 seconds, kindly verify the
          status of your internet connectivity. <br />
          In case of persistent difficulty, please get in touch with your
          designated system administrator.
        </Text>
      </Card>
    </Flex>
  );
}