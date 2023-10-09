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

import { getUserOrganizationId } from "@/lib/auth";
import { getOrganizationDetails } from "@/services/OrganizationRequestService";
import { useQuery, useQueryClient } from "react-query";
import { Link } from "react-router-dom";

export function OrganizationLogo() {
  const queryClient = useQueryClient();

  // Get User organization data
  const organizationId = getUserOrganizationId() || "";
  const { data: organizationData, isLoading: IsOrgDataLoading } = useQuery({
    queryKey: ["organization", organizationId],
    queryFn: () => {
      if (!organizationId) {
        return Promise.resolve(null);
      }
      return getOrganizationDetails(organizationId);
    },
    initialData: () =>
      queryClient.getQueryData(["organization", organizationId]),
    staleTime: Infinity, // never refetch
  });

  if (IsOrgDataLoading) {
    return (
      <div className="flex flex-1 items-center justify-between space-x-2 md:justify-end">
        <div className="animate-pulse flex space-x-4">
          <div className="rounded-lg bg-black dark:bg-white opacity-10 h-10 w-36"></div>
        </div>
      </div>
    );
  }

  if (organizationData && organizationData.logo) {
    return (
      <Link to="/" style={{ textDecoration: "none" }}>
        <img src={organizationData?.logo} alt="Organization Logo" />
      </Link>
    );
  }

  return (
    <Link
      className="mr-6 flex items-center space-x-2 font-semibold text-accent-foreground"
      to="/"
    >
      {organizationData?.name}
    </Link>
  );
}
