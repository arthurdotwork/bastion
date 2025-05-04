import type { onError, onSuccess } from "@/hooks/types.ts";
import type { authRegisterFormSchema } from "@/schemas/register";
import { useMutation } from "@tanstack/react-query";
import ky from "ky";
import type { z } from "zod";

export const useRegisterMutation = ({
	onSuccess,
	onError,
}: {
	onSuccess: onSuccess<unknown>;
	onError: onError;
}) => {
	return useMutation({
		mutationFn: async (data: z.infer<typeof authRegisterFormSchema>) => {
			return ky
				.post("http://localhost:8080/v1/register", { json: data })
				.json();
		},
		onSuccess,
		onError,
	});
};
