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
import { Group, Navbar, rem, Skeleton } from "@mantine/core";
import { useQuery, useQueryClient } from "react-query";
import { HeaderUserMenu } from "@/components/layout/HeaderUserMenu";
import { getUserId } from "@/helpers/constants";
import { getUserDetails } from "@/services/UserRequestService";
import { UserDownloads } from "@/components/layout/Header/_Partials/UserDownloads";
import { UserNotifications } from "@/components/layout/Header/_Partials/UserNotifications";
import { ThemeSwitcher } from "@/components/layout/Header/_Partials/ThemeSwitcher";
import { navbarScroll } from "@/components/layout/Navbar/_partials/NavbarScroll";
import { BillingLinks } from "@/components/layout/Navbar/_partials/BillingLinks";
import { OrganizationLogo } from "@/components/layout/Navbar/_partials/OrganizationLogo";
import { AdminLinks } from "@/components/layout/Navbar/_partials/SystemHealthLinks";
import { SearchSpotlight } from "@/components/layout/Header/Search";
import { DispatchLinks } from "@/components/layout/Navbar/_partials/DispatchLinks";
import { MainLinks } from "@/components/layout/Navbar/_partials/MainLinks";
import { MCode } from "@/components/common/Code";
import { useNavbarStyles } from "@/assets/styles/AsideStyles";

export function AsideMenu(): React.ReactElement {
  const { classes } = useNavbarStyles();
  const queryClient = useQueryClient();
  const userId = getUserId() || "";

  const { data: userData, isLoading: isUserDataLoading } = useQuery({
    queryKey: ["user", userId],
    queryFn: () => {
      if (!userId) {
        return Promise.resolve(null);
      }
      return getUserDetails(userId);
    },
    initialData: () => queryClient.getQueryData(["user", userId]),
    staleTime: Infinity, // never refetch
  });

  return (
    <Navbar
      hiddenBreakpoint="sm"
      height="100%"
      width={{ sm: 300 }}
      p="md"
      zIndex={10}
      className={classes.navbar}
    >
      <Group className={classes.header} position="apart">
        <OrganizationLogo />
        <MCode
          backgroundColor="rgba(112, 72, 232, .5)"
          color="white"
          sx={{ fontWeight: 700 }}
        >
          SANDBOX
        </MCode>
      </Group>

      <Navbar.Section className={classes.section}>
        {isUserDataLoading ? (
          <Group my={15}>
            <Skeleton ml={rem(15)} width={rem(250)} height={rem(40)} circle />
            <div>
              <Skeleton width={rem(120)} height={rem(15)} />
              <Skeleton mt={rem(5)} width={rem(150)} height={rem(15)} />
            </div>
          </Group>
        ) : (
          userData && <HeaderUserMenu user={userData} />
        )}
      </Navbar.Section>

      <SearchSpotlight />

      <Navbar.Section className={classes.section}>
        <UserDownloads />
        <UserNotifications />
        <ThemeSwitcher />
      </Navbar.Section>
      <Navbar.Section grow className={classes.links} component={navbarScroll}>
        <div className={classes.linksInner}>
          {/* Main Application Links */}
          <MainLinks />

          {/* Dispatch Links */}
          <DispatchLinks />

          {/* Billing Links */}
          <BillingLinks />

          {/* Admin Links */}
          <AdminLinks />
        </div>
      </Navbar.Section>
    </Navbar>
  );
}
