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

import React, { Suspense } from "react";
import {
  Avatar,
  Badge,
  Box,
  Button,
  Card,
  Divider,
  Flex,
  Skeleton,
  Text,
} from "@mantine/core";
import { usePageStyles } from "@/styles/PageStyles";
import { Customer } from "@/types/apps/customer";
import { upperFirst } from "@/lib/utils";
import { customerStore as store } from "@/stores/CustomerStore";
import { EditCustomerModal } from "@/components/customer/view/_partials/EditCustomerModal";

ViewCustomerNavbar.defaultProps = {
  customer: null,
};

type ViewCustomerNavbarProps = {
  customer?: Customer | null;
  isLoading: boolean;
};

export function ViewCustomerNavbar({
  customer,
  isLoading,
}: ViewCustomerNavbarProps) {
  const { classes } = usePageStyles();

  const mapStatusToBadge = (status: string) => {
    switch (status) {
      case "A":
        return (
          <Badge color="green" variant="filled" radius="xs" my={10}>
            Active
          </Badge>
        );
      case "I":
        return (
          <Badge radius="xs" variant="filled" my={10} color="red">
            Inactive
          </Badge>
        );
      default:
        return (
          <Badge radius="xs" variant="filled" my={10} color="yellow">
            Unknown
          </Badge>
        );
    }
  };

  const getFirstAndLastChar = (name: string) => {
    const firstChar = name.charAt(0).toUpperCase();
    const lastChar = name.charAt(name.length - 1).toUpperCase();
    return `${firstChar}${lastChar}`;
  };

  return (
    <Flex>
      {isLoading ? (
        <Skeleton height={800} width={400} />
      ) : (
        <>
          <Card className={classes.card}>
            <Box
              style={{
                textAlign: "center",
                display: "flex",
                flexDirection: "column",
                alignItems: "center",
              }}
            >
              <Avatar
                src={null}
                alt={customer?.name}
                size="xl"
                radius={50}
                color="blue"
                my={20}
              >
                {getFirstAndLastChar(customer?.name as string)}
              </Avatar>
              <Text className={classes.text} fw={600}>
                {upperFirst(customer?.code as string)}
              </Text>
              <Text color="dimmed" size="sm">
                {upperFirst(customer?.name as string)}
              </Text>
            </Box>
            <Box
              style={{
                display: "flex",
                justifyContent: "space-between",
                alignItems: "center",
              }}
            >
              <Text className={classes.text} fw={600} size="lg" my={20}>
                Details
              </Text>
              <Button
                onClick={() => store.set("editModalOpen", true)}
                size="xs"
              >
                Edit
              </Button>
            </Box>
            <Divider mb={20} />

            {mapStatusToBadge(customer?.status as string)}

            <Box>
              <Text className={classes.text} fw={500} size="sm">
                Code
              </Text>
              <Text color="dimmed">{customer?.code}</Text>
            </Box>
            <Box my={10}>
              <Text className={classes.text} fw={500} size="sm">
                Name
              </Text>
              <Text color="dimmed">{customer?.name}</Text>
            </Box>
            <Box my={10}>
              <Text className={classes.text} fw={500} size="sm">
                Address Line 1
              </Text>
              <Text color="dimmed">{customer?.address_line_1}</Text>
            </Box>
            <Box my={10}>
              <Text className={classes.text} fw={500} size="sm">
                Address Line 2
              </Text>
              <Text color="dimmed">{customer?.address_line_2}</Text>
            </Box>
            <Box my={10}>
              <Text className={classes.text} fw={500} size="sm">
                City
              </Text>
              <Text color="dimmed">{customer?.city}</Text>
            </Box>
            <Box my={10}>
              <Text className={classes.text} fw={500} size="sm">
                Zip Code
              </Text>
              <Text color="dimmed">{customer?.zip_code}</Text>
            </Box>
            <Box my={10}>
              <Text className={classes.text} fw={500} size="sm">
                Customer Advocate
              </Text>
              <Text color="dimmed">{customer?.advocate_full_name || ""}</Text>
            </Box>
            <Box my={10}>
              <Text className={classes.text} fw={500} size="sm">
                Last Ship Date
              </Text>
              <Text color="dimmed">
                {customer?.customer_shipment_metrics.last_shipment_date || ""}
              </Text>
            </Box>
            <Box my={10}>
              <Text className={classes.text} fw={500} size="sm">
                Last Bill Date
              </Text>
              <Text color="dimmed">
                {customer?.customer_shipment_metrics.last_bill_date || ""}
              </Text>
            </Box>
          </Card>
          <Suspense fallback={<Skeleton height={500} />}>
            <EditCustomerModal customer={customer} />
          </Suspense>
        </>
      )}
    </Flex>
  );
}
