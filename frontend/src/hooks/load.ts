import { useState, useEffect, useTransition } from "react";
import type { Todo } from "@/types/todo";
import { ApiError } from "@/api/errors";
import { getTodos } from "@/api/get-todos";

export function useTodos() {
  const [todos, setTodos] = useState<Todo[]>([]);

  const [isPending, startTransition] = useTransition();

  useEffect(() => {
    startTransition(async () => {
      try {
        const newTodos = await getTodos();
        setTodos(newTodos);
      } catch (err: unknown) {
        window.alert(err instanceof ApiError ? err.message : "Unknown error!");
      }
    });
  }, []);

  return {
    todos,
    setTodos,
    isPending,
  };
}
