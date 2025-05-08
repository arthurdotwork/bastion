import { requestToSnakeCase, responseToCamelCase } from "@alice-health/ky-hooks-change-case";
import ky, { type BeforeRequestHook, type KyRequest, type KyResponse } from "ky";

const tokenizerHook = ((request: KyRequest) => {
	const token = localStorage.getItem("bastion_at");
	if (token) {
		request.headers.set("Authorization", `Bearer ${token}`);
	}
}) as BeforeRequestHook;

export const apiClient = ky.extend({
	prefixUrl: "http://localhost:8080",
	hooks: {
		beforeRequest: [tokenizerHook, requestToSnakeCase],
		afterResponse: [
			responseToCamelCase,
			(request: KyRequest, _, response: KyResponse) => {
				const parsedRequestURL = new URL(request.url);
				if (response.status === 401 && parsedRequestURL.pathname !== "/v1/authenticate") {
					// TODO: Handle 401 Unauthorized properly.
					alert(`Unauthorized access from ${request.url}. Please log in again.`);
				}
			},
		],
	},
});
