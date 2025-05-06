import type { onError, onSuccess } from "@/hooks/types.ts";
import { apiClient } from "@/lib/client.ts";
import type { authRegisterFormSchema } from "@/schemas/register";
import { useMutation } from "@tanstack/react-query";
import type { HTTPError } from "ky";
import type { z } from "zod";

export const useRegisterMutation = ({
	onSuccess,
	onError,
}: {
	onSuccess: onSuccess<unknown>;
	onError: onError<HTTPError>;
}) => {
	return useMutation({
		mutationFn: async (data: z.infer<typeof authRegisterFormSchema>) => {
			return apiClient.post("/v1/register", { json: data }).json();
		},
		onSuccess,
		onError,
	});
};
