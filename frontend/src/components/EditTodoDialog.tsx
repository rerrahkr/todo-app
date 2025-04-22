import {
  type TodoEditableFields,
  todoEditableFieldsSchema,
} from "@/types/todo";
import Button from "@mui/material/Button";
import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
import DialogTitle from "@mui/material/DialogTitle";
import TextField from "@mui/material/TextField";
import type React from "react";
import { useState, useTransition } from "react";

export type EditTodoDialogProps = {
  open: boolean;
  onSubmit: (fields: TodoEditableFields) => Promise<void>;
  onClose: () => void;
  title: string;
  defaults?: TodoEditableFields;
};

export function EditTodoDialog({
  open,
  onSubmit,
  onClose,
  title,
  defaults = { content: "" },
}: EditTodoDialogProps): React.JSX.Element {
  const [failedValidation, setFailedValidation] = useState<boolean>(false);

  const [isPending, startTransition] = useTransition();

  function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault();
    startTransition(async () => {
      const formData = new FormData(e.currentTarget);
      const formObject = Object.fromEntries(formData.entries());
      const parseResult = todoEditableFieldsSchema.safeParse(formObject);

      if (!parseResult.success) {
        setFailedValidation(true);
        return;
      }
      setFailedValidation(false);

      await onSubmit(parseResult.data);

      onClose();
    });
  }

  const hasError = !isPending && failedValidation;

  return (
    <Dialog open={open} onClose={onClose}>
      <form onSubmit={handleSubmit}>
        <DialogTitle>{title}</DialogTitle>
        <DialogContent>
          <TextField
            id="content"
            label="Content"
            name="content"
            variant="standard"
            multiline
            defaultValue={defaults.content}
            error={hasError}
            helperText={hasError && "Invalid content."}
            disabled={isPending}
          />
        </DialogContent>
        <DialogActions>
          <Button onClick={onClose}>Cancel</Button>
          <Button type="submit" disabled={isPending}>
            {isPending ? "Saving..." : "Save"}
          </Button>
        </DialogActions>
      </form>
    </Dialog>
  );
}
