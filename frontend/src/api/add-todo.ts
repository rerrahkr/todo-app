import axios from "axios";
import { makeApiError } from "./errors";
import { type TodoEditableFields } from "@/types/todo";

/**
 * Get todo item from the backend.
 *
 * @param fields The fields to create.
 * @throws `ApiError` if the request fails or the response is invalid.
 */
export async function addTodo(fields: TodoEditableFields): Promise<void> {
  try {
    await axios.post<TodoEditableFields>(
      `${import.meta.env.VITE_BACKEND_URI}/todos`,
      fields
    );
  } catch (err: unknown) {
    throw makeApiError(err);
  }
}
