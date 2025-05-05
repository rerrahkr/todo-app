import { useState, useEffect, useTransition } from "react";
import type { Todo } from "@/types/todo";
import { ApiError } from "@/api/errors";
import { getTodos } from "@/api/get-todos";

export function useTodos() {
  const [todos, setTodos] = useState<Todo[]>([]);

  const [isPending, startTransition] = useTransition();

  useEffect(() => {
    const abortController = new AbortController();

    startTransition(async () => {
      try {
        const newTodos = await getTodos(abortController.signal);
        setTodos(newTodos);
      } catch (err: unknown) {
        window.alert(err instanceof ApiError ? err.message : "Unknown error!");
      }
    });

    return () => {
      abortController.abort();
    };
  }, []);

  return {
    todos,
    setTodos,
    isPending,
  };
}
