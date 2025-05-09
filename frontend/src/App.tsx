import { EditTodoDialog } from "@/components/EditTodoDialog";
import { TodoList } from "@/components/TodoList";
import { type Todo, type TodoEditableFields } from "@/types/todo";
import AddIcon from "@mui/icons-material/Add";
import Container from "@mui/material/Container";
import CssBaseline from "@mui/material/CssBaseline";
import Fab from "@mui/material/Fab";
import type React from "react";
import { useState } from "react";
import { useTodos } from "./hooks/load";
import Masonry from "@mui/lab/Masonry";
import { seq } from "@/utils";
import Skeleton from "@mui/material/Skeleton";
import { addTodo } from "./api/add-todo";
import { getTodos } from "./api/get-todos";
import { ApiError } from "./api/errors";
import { deleteTodo } from "./api/delete-todo";
import { updateTodo } from "./api/update-todos";

type DialogOpenStatus = {
  state: "closed" | "openNew" | "openEdit";
  title: string;
};

function App(): React.JSX.Element {
  const [todo, setTodo] = useState<Todo>();
  const { todos, setTodos, isPending } = useTodos();

  const [dialogOpenStatus, setDialogOpenStatus] = useState<DialogOpenStatus>({
    state: "closed",
    title: "",
  });

  function handleItemClicked(cardId: number) {
    setTodo(todos.find((todo) => todo.id === cardId));
    setDialogOpenStatus({
      state: "openEdit",
      title: "Edit Todo Item",
    });
  }

  function handleAddClicked() {
    setDialogOpenStatus({
      state: "openNew",
      title: "Create Todo Item",
    });
  }

  function handleClosedDialog() {
    setDialogOpenStatus((prev) => ({
      state: "closed",
      title: prev.title,
    }));
  }

  async function handleSubmitCreateTodo(fields: TodoEditableFields) {
    try {
      await addTodo(fields);
      setTodos(await getTodos());
      handleClosedDialog();
    } catch (err: unknown) {
      window.alert(err instanceof ApiError ? err.message : "Unknown error!");
    }
  }

  async function handleSubmitEditTodo(fields: TodoEditableFields) {
    const cardId = todo?.id;
    if (cardId == undefined) {
      window.alert("Todo ID is undefined!");
      return;
    }

    try {
      await updateTodo(cardId, fields);
      setTodos(await getTodos());
      handleClosedDialog();
    } catch (err: unknown) {
      window.alert(err instanceof ApiError ? err.message : "Unknown error!");
    }
  }

  async function handleItemChecked(cardId: number) {
    try {
      await deleteTodo(cardId);
      setTodos(await getTodos());
    } catch (err: unknown) {
      window.alert(err instanceof ApiError ? err.message : "Unknown error!");
    }
  }

  return (
    <>
      <CssBaseline />
      <Container sx={{ my: 4 }}>
        {isPending ? (
          <Masonry
            columns={{
              xs: 1,
              sm: 2,
              md: 3,
              lg: 4,
            }}
            spacing={2}
          >
            {[...seq(6)].map((i) => (
              <Skeleton key={`skeleton${i}`} variant="rounded" height={60} />
            ))}
          </Masonry>
        ) : (
          <TodoList
            todos={todos}
            onCheckItem={handleItemChecked}
            onClickItem={handleItemClicked}
            columns={{
              xs: 1,
              sm: 2,
              md: 3,
              lg: 4,
            }}
            spacing={2}
          />
        )}
        <Fab
          color="primary"
          aria-label="Add new todo"
          sx={{
            position: "fixed",
            right: 16,
            bottom: 16,
          }}
          onClick={handleAddClicked}
        >
          <AddIcon />
        </Fab>
      </Container>
      <EditTodoDialog
        onClose={handleClosedDialog}
        open={dialogOpenStatus.state !== "closed"}
        onSubmit={
          dialogOpenStatus.state === "openEdit"
            ? handleSubmitEditTodo
            : handleSubmitCreateTodo
        }
        title={dialogOpenStatus.title}
        defaults={
          todo && dialogOpenStatus.state === "openEdit"
            ? {
                content: todo.content,
              }
            : {
                content: "",
              }
        }
      />
    </>
  );
}

export default App;
