import { Outlet, createRootRouteWithContext } from '@tanstack/react-router';
import { TanStackRouterDevtools } from '@tanstack/react-router-devtools';

import TanstackQueryLayout from '../integrations/tanstack-query/devTools.tsx';

import type { QueryClient } from '@tanstack/react-query';

interface RouterContext {
	queryClient: QueryClient;
}

export const Route = createRootRouteWithContext<RouterContext>()({
	component: () => (
		<>
			<Outlet />

			<TanstackQueryLayout />
			<TanStackRouterDevtools />
		</>
	),
});
