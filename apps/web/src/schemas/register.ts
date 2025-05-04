import { z } from "zod";

export const authRegisterFormSchema = z
	.object({
		email: z.string().email(),
		password: z.string(),
		passwordConfirmation: z.string(),
	})
	.refine((data) => data.password === data.passwordConfirmation, {
		message: "Passwords do not match",
		path: ["passwordConfirmation"],
	});
