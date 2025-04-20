import type { Todo } from "@/types/todo";
import Masonry from "@mui/lab/Masonry";
import type React from "react";
import { CardItem, type CardItemProps } from "./CardItem";

type TodoListProps = {
  todos: Todo[];
  onCheckItem: CardItemProps["onCheck"];
  onClickItem: CardItemProps["onClick"];
  columns?: number;
  spacing?: number;
};

export function TodoList({
  todos,
  onCheckItem,
  onClickItem,
  columns = 4,
  spacing = 2,
}: TodoListProps): React.JSX.Element {
  return (
    <Masonry columns={columns} spacing={spacing}>
      {todos.map((todo) => (
        <CardItem
          key={todo.id}
          cardId={todo.id}
          content={todo.content}
          onCheck={onCheckItem}
          onClick={onClickItem}
        />
      ))}
    </Masonry>
  );
}
