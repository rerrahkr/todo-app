import { ZodError } from "zod";
import axios from "axios";

/**
 * Error class for API.
 */
export class ApiError extends Error {
  public inner?: Error;

  constructor(message: string, inner?: Error) {
    super(message);
    this.name = "ApiError";
    this.inner = inner;
    this.stack = inner?.stack;
  }
}

function makeMessage(err: unknown): string {
  if (axios.isAxiosError(err)) {
    if (err.response) {
      return `${err.response.status}: ${err.response.data}`;
    } else {
      return `Invalid request!`;
    }
  } else if (err instanceof ZodError) {
    return `Unexpected response!`;
  } else if (err instanceof Error) {
    return `${err.name}: ${err.message}`;
  } else {
    return "Unknown error!";
  }
}

/**
 * Generate an ApiError from an unknown error.
 *
 * @param err An error that is wrapped in ApiError.
 * @returns A new ApiError instance.
 */
export function makeApiError(err: unknown): ApiError {
  return new ApiError(makeMessage(err), err as Error);
}
