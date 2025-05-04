import axios from "axios";
import { makeApiError } from "./errors";
import { type TodoEditableFields } from "@/types/todo";

/**
 * Update specific todo item in the backend.
 *
 * @param id The ID of the todo item to update.
 * @param fields The fields of todo.
 * @throws `ApiError` if the request fails or the response is invalid.
 */
export async function updateTodo(
  id: number,
  fields: TodoEditableFields
): Promise<void> {
  try {
    await axios.patch<TodoEditableFields>(
      `${import.meta.env.VITE_BACKEND_URI}/todos/${id}`,
      fields
    );
  } catch (err: unknown) {
    throw makeApiError(err);
  }
}
