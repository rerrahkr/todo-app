import { EditTodoDialog } from "@/components/EditTodoDialog";
import { TodoList } from "@/components/TodoList";
import type { Todo } from "@/types/todo";
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

const TODOS: Todo[] = [...range(1, 20)].map(
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

  function handleItemChecked(cardId: number) {}

  function handleItemClicked(cardId: number) {
    setTodo(TODOS.find((todo) => todo.id === cardId));
    setDialogIsOpened(true);
  }

  function handleAddClicked() {}

  function handleClosedDialog() {
    setDialogIsOpened(false);
  }

  return (
    <>
      <CssBaseline />
      <Container sx={{ my: 4 }}>
        <TodoList
          todos={TODOS}
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
      {todo && (
        <EditTodoDialog
          todo={todo}
          open={dialogIsOpened}
          onClose={handleClosedDialog}
        />
      )}
    </>
  );
}

export default App;
