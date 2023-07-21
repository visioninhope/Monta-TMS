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

import React, { useState } from "react";
import { getUserId } from "@/lib/utils";
import { animated, useSpring } from "@react-spring/web";
import { useMutation, useQueryClient } from "react-query";
import { useDrag } from "react-use-gesture";
import axios from "@/lib/AxiosConfig";
import { notifications } from "@mantine/notifications";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faXmark } from "@fortawesome/pro-solid-svg-icons";
import { Button, Menu } from "@mantine/core";
import { IconProp } from "@fortawesome/fontawesome-svg-core";

type SwippableMenuItemProps<T extends Record<string, any>> = {
  item: T;
  icon: IconProp;
};

const useDeleteMutation = (userId: string) => {
  const queryClient = useQueryClient();
  return useMutation((id: string) => axios.delete(`/user_reports/${id}/`), {
    onSuccess: () => {
      queryClient.invalidateQueries(["userReport", userId]);
    },
    onError: () => {
      notifications.show({
        title: "Error",
        message: "An error occurred while deleting the report",
        color: "red",
        withCloseButton: true,
        icon: <FontAwesomeIcon icon={faXmark} />,
        autoClose: 10_000, // 10 seconds
      });
    },
  });
};

export function SwippableMenuItem<T extends Record<string, any>>({
  item,
  icon,
}: SwippableMenuItemProps<T>) {
  const [status, setStatus] = useState("normal"); // can be "normal", "swiped", or "deleted"
  const userId = getUserId() || "";
  const [{ x }, set] = useSpring(() => ({ x: 0 }));

  const bind = useDrag(({ down, movement: [mx], cancel, last }) => {
    set({
      x: down ? mx : 0,
      immediate: down,
      config: { tension: 300, friction: 30 },
    });

    if (!down && mx < -100) {
      setStatus("swiped");
      cancel!();
    } else if (last && status !== "swiped") {
      setStatus("normal");
    }
  });

  const handleDelete = (id: string) => {
    setStatus("deleted");
    useDeleteMutation(userId).mutate(id);
  };

  const handleCancel = () => {
    setStatus("normal");
    set({ x: 0 });
  };

  const handleContextMenu = (event: React.MouseEvent) => {
    event.preventDefault();
  };

  return status === "deleted" ? null : (
    <animated.div
      {...bind()}
      style={{ transform: x.interpolate((x) => `translateX(${x}px)`) }}
      onContextMenu={handleContextMenu}
    >
      <div style={{ display: status === "swiped" ? "none" : "block" }}>
        <Menu.Item
          key={item.id}
          icon={<FontAwesomeIcon icon={icon} />}
          component="a"
          onContextMenu={handleContextMenu}
          href={item.report}
        >
          {item.file_name}
        </Menu.Item>
      </div>
      {status === "swiped" && (
        <div>
          <Button
            w="50%"
            radius="xs"
            color="red"
            // variant="light"
            onClick={() => handleDelete(item.id)}
          >
            Delete
          </Button>
          <Button
            w="50%"
            radius="xs"
            variant="subtle"
            color="gray"
            onClick={handleCancel}
          >
            Cancel
          </Button>
        </div>
      )}
    </animated.div>
  );
}