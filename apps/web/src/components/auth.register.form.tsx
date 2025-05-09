import { cn } from "@/lib/utils.ts";
import { ShieldCheck } from "lucide-react";

import { Button } from "@/components/ui/button.tsx";
import { Input } from "@/components/ui/input.tsx";
import { Link, useNavigate } from "@tanstack/react-router";

import { zodResolver } from "@hookform/resolvers/zod";
import { useForm, useWatch } from "react-hook-form";

import { Form, FormControl, FormDescription, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form.tsx";
import type { z } from "zod";

import { useRegisterMutation } from "@/hooks/register.ts";
import { authRegisterFormSchema } from "@/schemas/register.ts";

const AuthRegisterForm = () => {
	const form = useForm<z.infer<typeof authRegisterFormSchema>>({
		resolver: zodResolver(authRegisterFormSchema),
		defaultValues: {
			email: "",
			password: "",
			passwordConfirmation: "",
		},
	});

	const emailValue = useWatch({ control: form.control, name: "email" });
	const navigate = useNavigate();

	const registerUserMutation = useRegisterMutation({
		onSuccess: async () => {
			await navigate({ to: "/auth", search: { email: emailValue } });
		},
		onError: () => {},
	});

	const onSubmit = (formData: z.infer<typeof authRegisterFormSchema>) => {
		registerUserMutation.mutate(formData);
	};

	return (
		<div className="grid min-h-svh lg:grid-cols-2">
			<div className="flex flex-col gap-4 p-6 md:p-10">
				<div className="flex justify-center gap-2 md:justify-start">
					<Link to={"/"} className="flex items-center gap-2 font-medium">
						<div className="flex h-6 w-6 items-center justify-center rounded-md bg-primary text-primary-foreground">
							<ShieldCheck className="size-4" />
						</div>
						Bastion Inc.
					</Link>
				</div>
				<div className="flex flex-1 items-center justify-center">
					<div className="w-full max-w-md">
						<Form {...form}>
							<form onSubmit={form.handleSubmit(onSubmit)} className={cn("flex flex-col gap-6")}>
								<div className="flex flex-col items-center gap-2 text-center">
									<h1 className="text-2xl font-bold">Create your account</h1>
									<p className="text-balance text-sm text-muted-foreground">Enter your email below to create your account</p>
								</div>
								<div className="grid gap-6">
									<div className="grid gap-2">
										<FormField
											control={form.control}
											name={"email"}
											render={({ field }) => (
												<FormItem>
													<FormLabel>Email</FormLabel>
													<FormControl>
														<Input placeholder={"email@example.com"} {...field} className={"text-sm"} />
													</FormControl>
													<FormDescription>This is the email you will use to login.</FormDescription>
													<FormMessage />
												</FormItem>
											)}
										/>
									</div>
									<div className="grid gap-2">
										<FormField
											control={form.control}
											name={"password"}
											render={({ field }) => (
												<FormItem>
													<FormLabel>Password</FormLabel>
													<FormControl>
														<Input type="password" placeholder={"********"} {...field} className={"text-sm"} />
													</FormControl>
													<FormDescription>This is the password you will use to login.</FormDescription>
													<FormMessage />
												</FormItem>
											)}
										/>
									</div>
									<div className="grid gap-2">
										<FormField
											control={form.control}
											name={"passwordConfirmation"}
											render={({ field }) => (
												<FormItem>
													<FormLabel>Confirm your password</FormLabel>
													<FormControl>
														<Input type="password" placeholder={"********"} {...field} className={"text-sm"} />
													</FormControl>
													<FormDescription>Please confirm your password.</FormDescription>
													<FormMessage />
												</FormItem>
											)}
										/>
									</div>
									<Button type="submit" className="w-full">
										Create my account
									</Button>
									{/*<div className="relative text-center text-sm after:absolute after:inset-0 after:top-1/2 after:z-0 after:flex after:items-center after:border-t after:border-border">
									<span className="relative z-10 bg-background px-2 text-muted-foreground">
										Or continue with
									</span>
								</div>
								<Button variant="outline" className="w-full">
									<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
										<path
											d="M12 .297c-6.63 0-12 5.373-12 12 0 5.303 3.438 9.8 8.205 11.385.6.113.82-.258.82-.577 0-.285-.01-1.04-.015-2.04-3.338.724-4.042-1.61-4.042-1.61C4.422 18.07 3.633 17.7 3.633 17.7c-1.087-.744.084-.729.084-.729 1.205.084 1.838 1.236 1.838 1.236 1.07 1.835 2.809 1.305 3.495.998.108-.776.417-1.305.76-1.605-2.665-.3-5.466-1.332-5.466-5.93 0-1.31.465-2.38 1.235-3.22-.135-.303-.54-1.523.105-3.176 0 0 1.005-.322 3.3 1.23.96-.267 1.98-.399 3-.405 1.02.006 2.04.138 3 .405 2.28-1.552 3.285-1.23 3.285-1.23.645 1.653.24 2.873.12 3.176.765.84 1.23 1.91 1.23 3.22 0 4.61-2.805 5.625-5.475 5.92.42.36.81 1.096.81 2.22 0 1.606-.015 2.896-.015 3.286 0 .315.21.69.825.57C20.565 22.092 24 17.592 24 12.297c0-6.627-5.373-12-12-12"
											fill="currentColor"
										/>
									</svg>
									Login with GitHub
								</Button>*/}
								</div>
								<div className="text-center text-sm">
									Already have an account?{" "}
									<Link to={"/auth"} search={{ email: emailValue }} className={cn("underline underline-offset-4")}>
										Sign in
									</Link>
								</div>
							</form>
						</Form>
					</div>
				</div>
			</div>
			<div className="relative hidden bg-neutral-900 lg:flex items-col justify-center items-center">
				<ShieldCheck size={"128"} color={"white"} />
			</div>
		</div>
	);
};

export default AuthRegisterForm;
