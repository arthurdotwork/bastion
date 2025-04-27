import type { QueryClient } from "@tanstack/react-query";
import { ReactQueryDevtools } from "@tanstack/react-query-devtools";
import { Outlet, createRootRouteWithContext } from "@tanstack/react-router";
import { TanStackRouterDevtools } from "@tanstack/react-router-devtools";

interface rootContext {
	queryClient: QueryClient;
}

export const Route = createRootRouteWithContext<rootContext>()({
	component: () => (
		<>
			<Outlet />

			<TanStackRouterDevtools />
			<ReactQueryDevtools />
		</>
	),
});
