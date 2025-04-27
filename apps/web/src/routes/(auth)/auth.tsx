import AuthForm from "@/components/auth.form.tsx";
import { createFileRoute } from "@tanstack/react-router";

const Page = () => {
	return <AuthForm />;
};

export const Route = createFileRoute("/(auth)/auth")({ component: Page });
