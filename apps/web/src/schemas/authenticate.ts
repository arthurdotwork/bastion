import { z } from "zod";

export const authAuthenticateWithPasswordFormSchema = z.object({
	email: z.string().email(),
	password: z.string(),
});

export const authAuthenticateWithPasswordResponseSchema = z.object({
	accessToken: z.string(),
	refreshToken: z.string(),
});
