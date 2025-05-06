import { Button } from "@/components/ui/button.tsx";
import { ThemeEnum, useTheme } from "@/providers/theme.tsx";
import { Moon, Sun } from "lucide-react";

export const SwitchTheme = () => {
	const { theme, setTheme } = useTheme();

	const switchTheme = () => {
		const newTheme = theme === "dark" ? ThemeEnum.light : ThemeEnum.dark;
		setTheme(newTheme);
	};

	return (
		<>
			<Button onClick={switchTheme}>{theme === "dark" ? <Moon /> : <Sun />}</Button>
		</>
	);
};
