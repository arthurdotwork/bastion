import { queryOptions, useSuspenseQuery } from "@tanstack/react-query";
import { createFileRoute } from "@tanstack/react-router";

const getDataQueryOptions = () =>
	queryOptions({
		queryKey: ["about"],
		queryFn: (): Promise<string> => {
			return new Promise((resolve) => {
				setTimeout(() => {
					resolve("ok");
				}, 1000);
			});
		},
		staleTime: 1000,
	});

const Page = () => {
	const q = useSuspenseQuery(getDataQueryOptions());
	console.log(q.data);

	return <h1>About {q.data}</h1>;
};

export const Route = createFileRoute("/about")({
	component: Page,
	loader: async ({ context }) =>
		context.queryClient.ensureQueryData(getDataQueryOptions()),
});
