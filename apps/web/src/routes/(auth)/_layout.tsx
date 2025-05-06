import { Outlet, createFileRoute, redirect } from "@tanstack/react-router";

const LayoutComponent = () => {
	return (
		<>
			<Outlet />
		</>
	);
};

export const Route = createFileRoute("/(auth)/_layout")({
	component: LayoutComponent,
	beforeLoad: () => {
		if (localStorage.getItem("bastion_at")) {
			return redirect({ to: "/dashboard" });
		}
	},
});
