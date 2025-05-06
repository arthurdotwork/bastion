import AuthForm from "@/components/auth.form.tsx";
import { createFileRoute } from "@tanstack/react-router";
import { z } from "zod";

const Page = () => {
	return <AuthForm />;
};

const authSearchParams = z.object({ email: z.string().nullish() });

export const Route = createFileRoute("/(auth)/_layout/auth")({
	component: Page,
	validateSearch: (search: z.infer<typeof authSearchParams>) => authSearchParams.parse(search),
});
