import { CommandDialog, CommandEmpty, CommandGroup, CommandInput, CommandItem, CommandList } from "@/components/ui/command.tsx";
import useHotkeys from "@reecelucas/react-use-hotkeys";
import { DoorClosed } from "lucide-react";
import { useState } from "react";

type CommandProps = {
	logout: () => void;
};

export const Command = ({ logout }: CommandProps) => {
	const [open, setOpen] = useState(false);

	useHotkeys("Meta+k", () => {
		setOpen((open) => !open);
	});

	return (
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
	);
};
