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

import { ORGANIZATION_ID_KEY } from "@/helpers/constants";

/**
 * Retrieves the CSRF token from the user's cookies.
 * @returns The CSRF token, or undefined if it was not found.
 */
export const getUserCSRFToken = (): string | undefined => {
  const csrfCookie = document.cookie
    .split("; ")
    .find((row) => row.startsWith("csrftoken="));

  if (csrfCookie) {
    return csrfCookie.split("=")[1];
  }
  console.error("No CSRF token found");
};

/**
 * Retrieves the specified item from the user's session storage.
 * @param key - The key of the item to retrieve.
 * @returns The item from session storage, parsed as JSON, or null if it was not found.
 */
export const getSessionItem = (key: string): any => {
  const item = sessionStorage.getItem(key);
  return item ? JSON.parse(item) : null;
};

/**
 * Retrieves the current user's organization ID from session storage.
 * @returns The organization's ID, or null if it was not found.
 */
export const getUserOrganizationId = (): string | null => {
  const userOrganization = sessionStorage.getItem(ORGANIZATION_ID_KEY);
  if (userOrganization) {
    return userOrganization;
  }
  return null;
};

/**
 * Clears all cookies from the browser.
 * @returns void
 * @see https://stackoverflow.com/a/179514
 */
export function clearAllCookies() {
  const cookies = document.cookie.split(";");
  cookies.forEach((cookie) => {
    const name = cookie.split("=")[0].trim();
    document.cookie = `${name}=; expires=Thu, 01 Jan 1970 00:00:01 GMT;`;
  });
}