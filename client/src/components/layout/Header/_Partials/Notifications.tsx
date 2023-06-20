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
import { QueryKey, useQuery, useQueryClient } from "react-query";
import { getUserNotifications } from "@/requests/UserRequestFactory";
import { getUserId } from "@/lib/utils";
import { headerStore } from "@/stores/HeaderStore";
import { Badge, Group, Skeleton, Text } from "@mantine/core";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faEnvelope, faEnvelopeOpen } from "@fortawesome/pro-duotone-svg-icons";
import { UserNotification, Notification } from "@/types/apps/accounts";
import { formatTimestamp } from "@/utils/date";

export const Notifications: React.FC = () => {
  const [notificationsMenuOpen] = headerStore.use("notificationsMenuOpen");
  const userId = getUserId() || "";
  const queryClient = useQueryClient();

  const { data: notificationsData, isLoading: notificationsLoading } = useQuery<
    UserNotification,
    QueryKey
  >({
    queryKey: ["userNotifications", userId],
    queryFn: getUserNotifications,
    initialData: () => {
      return queryClient.getQueryData(["userNotifications", userId]);
    },
    enabled: notificationsMenuOpen,
  });

  if (notificationsLoading) {
    return <Skeleton width={200} height={250} />;
  }

  if (!notificationsData || notificationsData?.unread_list.length === 0) {
    return (
      <div
        style={{
          display: "flex",
          flexDirection: "column",
          justifyContent: "center",
          alignItems: "center",
          height: "100%",
          width: "100%",
          marginTop: "30%",
        }}
      >
        <FontAwesomeIcon icon={faEnvelopeOpen} size="3x" />
        <Text>You have no notifications</Text>
      </div>
    );
  }

  const notificationItems = notificationsData?.unread_list.map(
    (notification: Notification) => {
      const humanReadableTimestamp = formatTimestamp(notification.timestamp);
      return (
        <>
          <Group mt={5} mb={10} key={notification.id}>
            <FontAwesomeIcon icon={faEnvelope} />
            <div style={{ flex: 1 }}>
              <Text fw={700} size="xs">
                {notification.verb}
              </Text>
              <Text size="xs">{notification.description}</Text>
            </div>
            <Badge size="xs" radius="xs" variant="outline" color="violet">
              {humanReadableTimestamp}
            </Badge>
          </Group>
        </>
      );
    }
  );

  return <>{notificationItems}</>;
};