import { CustomSidebar } from "@/components/sidebar.tsx";
import { apiClient } from "@/lib/client";
import { Outlet, createFileRoute, redirect, useNavigate } from "@tanstack/react-router";

const LayoutComponent = () => {
	const navigate = useNavigate();

	const logout = async () => {
		localStorage.removeItem("bastion_at");
		await navigate({ to: "/auth", search: { email: "" }, viewTransition: true });
	};

	return (
		<>
			<CustomSidebar logout={logout}>
				<Outlet />
			</CustomSidebar>
		</>
	);
};

export const Route = createFileRoute("/(app)/_layout")({
	component: LayoutComponent,
	beforeLoad: async () => {
		if (!localStorage.getItem("bastion_at")) {
			return redirect({ to: "/auth", search: { email: "" } });
		}

		try {
			await apiClient.get("v1/authenticate").json();
		} catch (error) {
			localStorage.removeItem("bastion_at");
			return redirect({ to: "/auth", search: { email: "" } });
		}
	},
});
