import type { DefaultError } from "@tanstack/react-query";

export type onSuccess<T> = (data: T) => void;
export type onError = (data: DefaultError) => void;
