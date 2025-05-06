import type { onError, onSuccess } from "@/hooks/types.ts";
import { apiClient } from "@/lib/client.ts";
import type { authAuthenticateWithPasswordFormSchema, authAuthenticateWithPasswordResponseSchema } from "@/schemas/authenticate.ts";
import { useMutation } from "@tanstack/react-query";
import type { HTTPError } from "ky";
import type { z } from "zod";

export const useAuthenticateWithPasswordMutation = ({
	onSuccess,
	onError,
}: {
	onSuccess: onSuccess<z.infer<typeof authAuthenticateWithPasswordResponseSchema>>;
	onError: onError<HTTPError>;
}) => {
	return useMutation({
		mutationFn: async (data: z.infer<typeof authAuthenticateWithPasswordFormSchema>) => {
			return apiClient.post("http://localhost:8080/v1/authenticate", { json: data }).json<z.infer<typeof authAuthenticateWithPasswordResponseSchema>>();
		},
		onSuccess,
		onError,
	});
};
