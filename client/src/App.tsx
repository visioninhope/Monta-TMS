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

import { BrowserRouter } from "react-router-dom";
import React, { Suspense, useEffect } from "react";
import LoadingScreen from "./components/LoadingScreen";
import ProtectedRoutes from "./routing/ProtectedRoutes";
import useVerifyToken from "./hooks/withTokenVerification";
import { useAuthStore } from "@/stores/authStore";
import "./styles/App.css";
import {
  ColorScheme,
  ColorSchemeProvider,
  MantineProvider,
} from "@mantine/core";
import { QueryClient, QueryClientProvider } from "react-query";
import { useHotkeys, useLocalStorage } from "@mantine/hooks";
import { Notifications } from "@mantine/notifications";
import { ModalsProvider } from "@mantine/modals";

function App() {
  useVerifyToken();
  const initialLoading = useAuthStore((state) => state.initialLoading);

  if (initialLoading) {
    return <LoadingScreen />;
  }

  const [colorScheme, setColorScheme] = useLocalStorage<ColorScheme>({
    key: "monta-color-scheme",
    defaultValue: "light",
    getInitialValueInEffect: true,
  });
  const toggleColorScheme = (value?: ColorScheme) =>
    setColorScheme(value || (colorScheme === "dark" ? "light" : "dark"));

  // TODO(wolfred): remove in production leave for dev because it's useful for me to switch.
  useHotkeys([["mod+J", () => toggleColorScheme()]]); // Acorn said this wasn't standard, and it would confuse users.

  useEffect(() => {
    document.body.className =
      colorScheme === "dark" ? "dark-theme" : "light-theme";
  }, [colorScheme]);

  const queryClient = new QueryClient({
    defaultOptions: {
      queries: {
        retry: false,
      },
    },
  });

  return (
    <>
      <ColorSchemeProvider
        colorScheme={colorScheme}
        toggleColorScheme={toggleColorScheme}
      >
        <MantineProvider
          theme={{
            colorScheme,
          }}
          withGlobalStyles
          withNormalizeCSS
          withCSSVariables
        >
          <Notifications limit={5} position="top-right" zIndex={2077} />
          <QueryClientProvider client={queryClient}>
            <BrowserRouter>
              <Suspense fallback={<LoadingScreen />}>
                <ProtectedRoutes />
              </Suspense>
            </BrowserRouter>
          </QueryClientProvider>
        </MantineProvider>
      </ColorSchemeProvider>
    </>
  );
}

export default App;
