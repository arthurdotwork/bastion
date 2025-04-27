import { createFileRoute } from "@tanstack/react-router";

const Page = () => {
	return <h1>Dashboard</h1>;
};

export const Route = createFileRoute("/(app)/_layout/dashboard")({
	component: Page,
});
