import type { QueryClient } from "@tanstack/react-query";
import { Outlet, createRootRouteWithContext } from "@tanstack/react-router";

interface rootContext {
	queryClient: QueryClient;
}

export const Route = createRootRouteWithContext<rootContext>()({
	component: () => (
		<>
			<Outlet />

			{/*<TanStackRouterDevtools />*/}
			{/*<ReactQueryDevtools />*/}
		</>
	),
});
