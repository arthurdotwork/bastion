import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import "./main.css";

import { ThemeProvider } from "@/providers/theme.tsx";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { RouterProvider, createRouter } from "@tanstack/react-router";
import { routeTree } from "./routeTree.gen.ts";

const queryClient = new QueryClient();

const router = createRouter({
	routeTree,
	context: {
		queryClient: queryClient,
	},
});

declare module "@tanstack/react-router" {
	interface Register {
		router: typeof router;
	}
}

const root = document.getElementById("root");
if (root) {
	const rootElement = createRoot(root);
	rootElement.render(
		<StrictMode>
			<QueryClientProvider client={queryClient}>
				<ThemeProvider defaultTheme={"system"} storageKey={"bastion-ui"}>
					<RouterProvider router={router} />
				</ThemeProvider>
			</QueryClientProvider>
		</StrictMode>,
	);
}
