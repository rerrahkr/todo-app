import type { Todo } from "@/types/todo";
import axios from "axios";
import { makeApiError } from "./errors";
import { z } from "zod";
import { todoSchema } from "@/types/todo";

const getTodosResponseSchema = z.object({
  todos: z.array(todoSchema),
});

type GetTodosResponse = z.infer<typeof getTodosResponseSchema>;

/**
 * Get all todos from the backend.
 *
 * @param abortSignal An optional AbortSignal to cancel the request.
 * @returns A list of todos.
 * @throws `ApiError` if the request fails or the response is invalid.
 */
export async function getTodos(abortSignal?: AbortSignal): Promise<Todo[]> {
  try {
    const response = await axios.get<GetTodosResponse>(
      `${import.meta.env.VITE_BACKEND_URI}/todos`,
      {
        signal: abortSignal,
      }
    );

    const todos = getTodosResponseSchema.parse(response.data);
    return todos.todos;
  } catch (err: unknown) {
    throw makeApiError(err);
  }
}
