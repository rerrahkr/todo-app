import type { Todo } from "@/types/todo";
import Button from "@mui/material/Button";
import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
import DialogTitle from "@mui/material/DialogTitle";
import TextField from "@mui/material/TextField";
import type React from "react";
import { useState, useTransition } from "react";
import { z } from "zod";

const editTodoPayloadSchema = z.object({
	content: z.string().min(1),
});

export type EditTodoDialogProps = {
	todo: Todo;
	open: boolean;
	onClose: () => void;
};

export function EditTodoDialog({
	todo,
	open,
	onClose,
}: EditTodoDialogProps): React.JSX.Element {
	const [failedValidation, setFailedValidation] = useState<boolean>(false);

	const [isPending, startTransition] = useTransition();

	function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
		e.preventDefault();
		startTransition(async () => {
			const formData = new FormData(e.currentTarget);
			const formObject = Object.fromEntries(formData.entries());
			const parseResult = editTodoPayloadSchema.safeParse(formObject);

			if (!parseResult.success) {
				setFailedValidation(true);
				return;
			}

			await new Promise((resolve) => setTimeout(resolve, 1000));

			setFailedValidation(false);
			onClose();
		});
	}

	const hasError = !isPending && failedValidation;

	return (
		<Dialog open={open} onClose={onClose}>
			<form onSubmit={handleSubmit}>
				<DialogTitle>Edit Todo Item</DialogTitle>
				<DialogContent>
					<TextField
						id="content"
						label="Content"
						name="content"
						variant="standard"
						defaultValue={todo.content}
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
