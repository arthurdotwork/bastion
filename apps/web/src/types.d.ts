declare module "@alice-health/ky-hooks-change-case" {
	import type { Hook } from "ky";

	// Define the type for the main export
	// This is a guess based on the module name, adjust as needed
	export const requestToSnakeCase: Hook;
	export const responseToCamelCase: Hook;

	// Add any other exports the module has
	// For example:
	// export function someOtherFunction(): void;

	// Export default (if the module has a default export)
	// export default changeCaseHook;
}
