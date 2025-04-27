import { createFileRoute } from "@tanstack/react-router";

const Page = () => {
	return <h1>index</h1>;
};

export const Route = createFileRoute("/")({ component: Page });
