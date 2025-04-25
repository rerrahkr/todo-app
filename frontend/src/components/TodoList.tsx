import type { Todo } from "@/types/todo";
import Grow from "@mui/material/Grow";
import Masonry, { type MasonryProps } from "@mui/lab/Masonry";
import { TransitionGroup } from "react-transition-group";
import type React from "react";
import { CardItem, type CardItemProps } from "./CardItem";

type TodoListProps = {
  todos: Todo[];
  onCheckItem: CardItemProps["onCheck"];
  onClickItem: CardItemProps["onClick"];
  columns?: MasonryProps["columns"];
  spacing?: MasonryProps["spacing"];
};

export function TodoList({
  todos,
  onCheckItem,
  onClickItem,
  columns = 4,
  spacing = 2,
}: TodoListProps): React.JSX.Element {
  return (
    <TransitionGroup component={Masonry} columns={columns} spacing={spacing}>
      {todos.map((todo) => (
        <Grow key={todo.id}>
          <div>
            <CardItem
              cardId={todo.id}
              content={todo.content}
              onCheck={onCheckItem}
              onClick={onClickItem}
            />
          </div>
        </Grow>
      ))}
    </TransitionGroup>
  );
}
