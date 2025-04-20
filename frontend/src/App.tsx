import { EditTodoDialog } from "@/components/EditTodoDialog";
import { TodoList } from "@/components/TodoList";
import type { Todo } from "@/types/todo";
import Container from "@mui/material/Container";
import CssBaseline from "@mui/material/CssBaseline";
import type React from "react";
import { useState } from "react";

function* range(begin: number, end: number, step = 1): Generator<number> {
  for (let i = begin; i < end; i += step) {
    yield i;
  }
}

const TODOS: Todo[] = [...range(1, 11)].map(
  (i): Todo => ({
    id: i,
    content: `Todo ${i}`,
    createdAt: new Date(),
    updatedAt: new Date(),
  })
);

function App(): React.JSX.Element {
  const [todo, setTodo] = useState<Todo>();

  const [dialogIsOpened, setDialogIsOpened] = useState<boolean>(false);

  function handleChecked(cardId: number) {}

  function handleClicked(cardId: number) {
    setTodo(TODOS.find((todo) => todo.id === cardId));
    setDialogIsOpened(true);
  }

  function handleCloseDialog() {
    setDialogIsOpened(false);
  }

  return (
    <>
      <CssBaseline />
      <Container sx={{ my: 4 }}>
        <TodoList
          todos={TODOS}
          onCheckItem={handleChecked}
          onClickItem={handleClicked}
          columns={4}
          spacing={2}
        />
      </Container>
      {todo && (
        <EditTodoDialog
          todo={todo}
          open={dialogIsOpened}
          onClose={handleCloseDialog}
        />
      )}
    </>
  );
}

export default App;
