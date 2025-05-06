import { Button } from "@/components/ui/button.tsx";
import { Dialog, DialogContent, DialogDescription, DialogFooter, DialogHeader, DialogTitle } from "@/components/ui/dialog.tsx";
import { createFileRoute, useRouter } from "@tanstack/react-router";
import { useState } from "react";

export const Route = createFileRoute("/(app)/_layout/settings")({
	component: RouteComponent,
});

function RouteComponent() {
	const [open, setOpen] = useState(true);

	const router = useRouter();
	const onClose = () => {
		setOpen(false);
	};

	return (
		<Dialog open={open} onOpenChange={onClose}>
			<DialogContent
				className="sm:max-w-[425px]"
				onAnimationEnd={() => {
					if (!open) {
						console.log("canGoBack?", router.history.canGoBack());
						if (router.history.canGoBack()) {
							router.history.back();
							return;
						}

						router.navigate({ to: "/dashboard" });
					}
				}}
			>
				<DialogHeader>
					<DialogTitle>Edit profile</DialogTitle>
					<DialogDescription>Make changes to your profile here. Click save when you're done.</DialogDescription>
				</DialogHeader>
				<DialogFooter>
					<Button>Save changes</Button>
				</DialogFooter>
			</DialogContent>
		</Dialog>
	);
}
