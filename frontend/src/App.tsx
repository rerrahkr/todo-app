import { EditTodoDialog } from "@/components/EditTodoDialog";
import { TodoList } from "@/components/TodoList";
import type { Todo, TodoEditableFields } from "@/types/todo";
import AddIcon from "@mui/icons-material/Add";
import Container from "@mui/material/Container";
import CssBaseline from "@mui/material/CssBaseline";
import Fab from "@mui/material/Fab";
import type React from "react";
import { useState } from "react";

function* range(begin: number, end: number, step = 1): Generator<number> {
  for (let i = begin; i < end; i += step) {
    yield i;
  }
}

type DialogOpenStatus = {
  state: "closed" | "openNew" | "openEdit";
  title: string;
};

function App(): React.JSX.Element {
  const [todo, setTodo] = useState<Todo>();

  const [todos, setTodos] = useState<Todo[]>(
    [...range(1, 20)].map(
      (i): Todo => ({
        id: i,
        content: `Todo ${i}`,
        createdAt: new Date(),
        updatedAt: new Date(),
      })
    )
  );

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
    setTodos((prev) => {
      const newTodo: Todo = {
        id: prev.length === 0 ? 1 : prev[prev.length - 1].id + 1,
        createdAt: new Date(),
        updatedAt: new Date(),
        ...fields,
      };
      return [...prev, newTodo];
    });

    console.log("create");
  }

  async function handleSubmitEditTodo(fields: TodoEditableFields) {
    const cardId = todo?.id;
    if (cardId == undefined) {
      return;
    }

    setTodos((prev) => {
      const i = prev.findIndex((todo) => todo.id === cardId);
      if (i !== -1) {
        prev[i] = {
          ...prev[i],
          ...fields,
        };
      }
      return prev;
    });

    console.log("update");
  }

  async function handleItemChecked(cardId: number) {
    setTodos((prev) => prev.filter((todo) => todo.id !== cardId));

    console.log("delete");
  }

  return (
    <>
      <CssBaseline />
      <Container sx={{ my: 4 }}>
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
