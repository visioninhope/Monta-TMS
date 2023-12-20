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
import { Notification, UserNotification } from "@/types/accounts";
import { Skeleton } from "@/components/ui/skeleton";
import nothingFound from "@/assets/images/there-is-nothing-here.png";
import { formatTimestamp } from "@/lib/date";
import { truncateText } from "@/lib/utils";
import { Badge } from "@/components/ui/badge";
import React from "react";

export function Notifications({
  notification,
  notificationLoading,
}: {
  notification: UserNotification;
  notificationLoading: boolean;
}) {
  if (notificationLoading) {
    return <Skeleton className="h-80" />;
  }

  if (!notification || notification.unreadList.length === 0) {
    return (
      <div className="flex flex-col justify-content-center items-center h-full w-full mt-10">
        <img src={nothingFound} alt="Nothing Found" className="h-40 w-40" />
        <h3 className="text-2xl font-medium">All Caught up!</h3>
        <p className="text-sm text-muted-foreground">
          You have no unread notifications
        </p>
      </div>
    );
  }

  const notificationItems = notification?.unreadList.map(
    (notification: Notification) => {
      const humanReadableTime = formatTimestamp(notification.timestamp);

      return (
        <div
          key={notification.id}
          className="group flex flex-col space-y-2 px-4 py-2 border-b border-accent cursor-pointer hover:bg-accent"
        >
          <div className="flex items-center justify-between">
            <p className="text-sm font-semibold leading-none">
              {truncateText(notification.verb, 25)}
            </p>
            <Badge
              withDot={false}
              className="p-0.5 text-xs select-none bg-accent group-hover:bg-accent-foreground group-hover:text-accent text-accent-forground"
            >
              {humanReadableTime}
            </Badge>
          </div>
          <p className="text-xs text-muted-foreground">
            {truncateText(notification.description, 40)}
          </p>
        </div>
      );
    },
  );

  return <>{notificationItems}</>;
}