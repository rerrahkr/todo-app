import axios from "axios";
import { makeApiError } from "./errors";

/**
 * Delete todo item from the backend.
 *
 * @param id The ID of the todo item to delete.
 * @throws `ApiError` if the request fails or the response is invalid.
 */
export async function deleteTodo(id: number): Promise<void> {
  try {
    await axios.delete(`${import.meta.env.VITE_BACKEND_URI}/todos/${id}`);
  } catch (err: unknown) {
    throw makeApiError(err);
  }
}
