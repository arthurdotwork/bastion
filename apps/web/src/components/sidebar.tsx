import { Command } from "@/components/command.tsx";
import { SwitchTheme } from "@/components/theme.tsx";
import { Breadcrumb, BreadcrumbItem, BreadcrumbLink, BreadcrumbList, BreadcrumbPage, BreadcrumbSeparator } from "@/components/ui/breadcrumb.tsx";
import { Separator } from "@/components/ui/separator.tsx";
import { Sidebar, SidebarContent, SidebarFooter, SidebarGroup, SidebarHeader, SidebarInset, SidebarMenuButton, SidebarProvider, SidebarTrigger } from "@/components/ui/sidebar.tsx";
import { useNavigate } from "@tanstack/react-router";
import { LogOut, ShieldCheck, SlidersVertical } from "lucide-react";
import type { ReactNode } from "react";

type SidebarProps = {
	logout: () => void;
	children?: ReactNode;
};

export const CustomSidebar = ({ logout, children }: SidebarProps) => {
	const navigate = useNavigate();

	return (
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
					<Command logout={logout} />
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
					{children}
				</main>
			</SidebarInset>
		</SidebarProvider>
	);
};
