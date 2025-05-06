import AuthRegisterForm from "@/components/auth.register.form.tsx";
import { createFileRoute } from "@tanstack/react-router";

const Page = () => {
	return <AuthRegisterForm />;
};

export const Route = createFileRoute("/(auth)/_layout/register")({ component: Page });
