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
import { Avatar, Badge, Box, Button, Menu, Text, Tooltip } from "@mantine/core";
import { formatDate, formatDateToHumanReadable } from "@/utils/date";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {
  faUser,
  faUserGear,
  faUserMinus,
} from "@fortawesome/pro-duotone-svg-icons";
import { MRT_ColumnDef } from "mantine-react-table";
import { User } from "@/types/apps/accounts";
import { faChevronDown } from "@fortawesome/pro-solid-svg-icons";
import { userTableStore } from "@/stores/UserTableStore";

export const UserTableColumns = (): MRT_ColumnDef<User>[] => {
  return [
    {
      id: "status",
      accessorFn: (originalRow) => (originalRow.is_active ? "true" : "false"),
      header: "Status",
      filterFn: "equals",
      Cell: ({ cell }) => (
        <Badge
          color={cell.getValue() === "true" ? "green" : "red"}
          variant="outline"
          radius="xs"
        >
          {cell.getValue() === "true" ? "Active" : "Inactive"}
        </Badge>
      ),
      mantineFilterSelectProps: {
        data: [
          { value: "", label: "All" },
          { value: "true", label: "Active" },
          { value: "false", label: "Inactive" },
        ] as any,
      },
      filterVariant: "select",
    },
    {
      accessorFn: (row) =>
        `${row.profile?.first_name} ${row.profile?.last_name}`,
      id: "name",
      header: "Name",
      size: 250,
      Cell: ({ renderedCellValue, row }) => (
        <Box
          sx={{
            display: "flex",
            alignItems: "center",
            gap: "16px",
          }}
        >
          {row.original.profile?.profile_picture ? (
            <Avatar
              src={row.original.profile?.profile_picture}
              alt={"Test"}
              radius="xl"
              size={30}
            />
          ) : (
            <Avatar color="blue" radius="xl" size={30}>
              {row.original.profile?.first_name.charAt(0)}
              {row.original.profile?.last_name.charAt(0)}
            </Avatar>
          )}
          <span>{renderedCellValue}</span>
        </Box>
      ),
    },
    {
      accessorKey: "email",
      header: "Email",
    },
    {
      id: "date_joined",
      header: "Date Joined",
      accessorFn: (row) => {
        if (row.date_joined) {
          return formatDateToHumanReadable(row.date_joined);
        } else {
          return null;
        }
      },
      Cell: ({ row }) => {
        if (!row.original.date_joined) {
          return <Text>Never</Text>;
        }

        return formatDate(row.original.date_joined);
      },
    },
    {
      id: "last_login",
      header: "Last Login",
      accessorFn: (row) => {
        if (row.last_login) {
          return formatDateToHumanReadable(row.last_login);
        } else {
          return null;
        }
      },
      Cell: ({ renderedCellValue, row }) => {
        if (!row.original.last_login) {
          return <Text>Never</Text>;
        }
        const tooltipDate = formatDate(row.original.last_login);

        return (
          <Tooltip withArrow position="left" label={tooltipDate}>
            <Text>{renderedCellValue}</Text>
          </Tooltip>
        );
      },
    },
    {
      id: "actions",
      header: "Actions",
      Cell: ({ row }) => (
        <>
          <Menu width={200} shadow="md" withArrow offset={5} position="bottom">
            <Menu.Target>
              <Button
                variant="light"
                color="gray"
                size="xs"
                rightIcon={<FontAwesomeIcon icon={faChevronDown} size="sm" />}
              >
                Actions
              </Button>
            </Menu.Target>
            <Menu.Dropdown>
              <Menu.Label>User Actions</Menu.Label>
              <Menu.Item
                icon={<FontAwesomeIcon icon={faUser} />}
                onClick={() => {
                  userTableStore.set("selectedUser", row.original);
                  userTableStore.set("viewUserModalOpen", true);
                }}
              >
                View User Profile
              </Menu.Item>
              <Menu.Item icon={<FontAwesomeIcon icon={faUserGear} />}>
                Edit User Profile
              </Menu.Item>
              <Menu.Item
                color="red"
                icon={<FontAwesomeIcon icon={faUserMinus} />}
                onClick={() => {
                  userTableStore.set("selectedUser", row.original);
                  userTableStore.set("deleteUserModalOpen", true);
                }}
              >
                Delete User Profile
              </Menu.Item>
            </Menu.Dropdown>
          </Menu>
        </>
      ),
    },
  ];
};