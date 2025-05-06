import { SwitchTheme } from "@/components/theme.tsx";
import { Breadcrumb, BreadcrumbItem, BreadcrumbLink, BreadcrumbList, BreadcrumbPage, BreadcrumbSeparator } from "@/components/ui/breadcrumb";
import { CommandDialog, CommandEmpty, CommandGroup, CommandInput, CommandItem, CommandList } from "@/components/ui/command";
import { Separator } from "@/components/ui/separator.tsx";
import { Sidebar, SidebarContent, SidebarFooter, SidebarGroup, SidebarHeader, SidebarInset, SidebarMenuButton, SidebarProvider, SidebarTrigger } from "@/components/ui/sidebar.tsx";
import { apiClient } from "@/lib/client.ts";
import useHotkeys from "@reecelucas/react-use-hotkeys";
import { Outlet, createFileRoute, redirect, useNavigate } from "@tanstack/react-router";
import { DoorClosed, LogOut, ShieldCheck, SlidersVertical } from "lucide-react";
import { useState } from "react";

const LayoutComponent = () => {
	const [open, setOpen] = useState(false);

	useHotkeys("Meta+k", () => {
		setOpen((open) => !open);
	});

	const navigate = useNavigate();

	const logout = async () => {
		localStorage.removeItem("bastion_at");
		await navigate({ to: "/auth", search: { email: "" }, viewTransition: true });
	};

	return (
		<>
			<SidebarProvider>
				<Sidebar variant={"inset"}>
					<SidebarHeader className={"p-4"}>
						<div className={"flex items-center"}>
							<div className="flex h-6 w-6 items-center justify-center rounded-md bg-primary text-primary-foreground mr-2">
								<ShieldCheck className="size-4" />
							</div>
							<p className={"text-sm"}>Bastion Inc.</p>
						</div>
					</SidebarHeader>
					<SidebarContent>
						<SidebarGroup />
						<SidebarGroup />
					</SidebarContent>
					<SidebarFooter>
						<SidebarMenuButton onClick={() => navigate({ to: "/settings" })}>
							<SlidersVertical className="size-4 mr-2" />
							Settings
						</SidebarMenuButton>
						<SidebarMenuButton onClick={logout}>
							<LogOut className="size-4 mr-2" />
							Logout
						</SidebarMenuButton>
					</SidebarFooter>
				</Sidebar>
				<SidebarInset>
					<main>
						<CommandDialog open={open} onOpenChange={setOpen}>
							<CommandInput placeholder="Type a command or search..." />
							<CommandList>
								<CommandEmpty>No results found.</CommandEmpty>
								<CommandGroup heading={"Actions"}>
									<CommandItem onSelect={logout}>
										<DoorClosed />
										Logout
									</CommandItem>
									<CommandItem onSelect={() => setOpen((open) => !open)}>
										<DoorClosed />
										Close
									</CommandItem>
								</CommandGroup>
							</CommandList>
						</CommandDialog>
						<header className="flex h-16 shrink-0 items-center justify-between gap-2 px-4">
							<div className="flex items-center gap-2">
								<SidebarTrigger />
								<Separator orientation="vertical" className="mr-2 h-4" />
								<Breadcrumb>
									<BreadcrumbList>
										<BreadcrumbItem className="hidden md:block">
											<BreadcrumbLink href="#">Building Your Application</BreadcrumbLink>
										</BreadcrumbItem>
										<BreadcrumbSeparator className="hidden md:block" />
										<BreadcrumbItem>
											<BreadcrumbPage>Data Fetching</BreadcrumbPage>
										</BreadcrumbItem>
									</BreadcrumbList>
								</Breadcrumb>
							</div>
							<SwitchTheme />
						</header>
						<Outlet />
					</main>
				</SidebarInset>
			</SidebarProvider>
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
			await apiClient.get("http://localhost:8080/v1/authenticate");
		} catch (e: unknown) {
			localStorage.removeItem("bastion_at");
			return redirect({ to: "/auth", search: { email: "" } });
		}
	},
});
