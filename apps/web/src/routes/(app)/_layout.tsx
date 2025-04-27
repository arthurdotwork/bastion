import { Outlet, createFileRoute, redirect } from "@tanstack/react-router";

export const Route = createFileRoute("/(app)/_layout")({
	component: RouteComponent,
	beforeLoad: () => {
		return redirect({ to: "/auth" });
	},
});

function RouteComponent() {
	return (
		<>
			<Outlet />
		</>
	);
}
