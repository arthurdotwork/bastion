import { createFileRoute } from "@tanstack/react-router";

const Page = () => {
	return <></>;
};

export const Route = createFileRoute("/(app)/_layout/dashboard")({
	component: Page,
});
