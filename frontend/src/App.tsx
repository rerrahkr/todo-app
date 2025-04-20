import Masonry from "@mui/lab/Masonry";
import Button from "@mui/material/Button";
import Card from "@mui/material/Card";
import CardActionArea from "@mui/material/CardActionArea";
import CardActions from "@mui/material/CardActions";
import CardContent from "@mui/material/CardContent";
import Checkbox from "@mui/material/Checkbox";
import Container from "@mui/material/Container";
import CssBaseline from "@mui/material/CssBaseline";
import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
import DialogTitle from "@mui/material/DialogTitle";
import TextField from "@mui/material/TextField";
import Typography from "@mui/material/Typography";
import type { SxProps, Theme } from "@mui/material/styles";
import type React from "react";
import { useState, useTransition } from "react";
import { z } from "zod";

type Todo = {
	id: number;
	content: string;
	createdAt: Date;
	updatedAt: Date;
};

type CardItemProps = {
	cardId: number;
	content: string;
	onCheck: (id: number) => void;
	onClick: (id: number) => void;
	sx?: SxProps<Theme>;
};

function CardItem({
	cardId,
	content,
	onCheck,
	onClick,
	sx,
}: CardItemProps): React.JSX.Element {
	return (
		<Card
			sx={{
				...sx,
				display: "flex",
			}}
		>
			<CardActions
				sx={{
					display: "flex",
					flexDirection: "column",
					justifyContent: "center",
				}}
			>
				<Checkbox onChange={() => onCheck(cardId)} />
			</CardActions>
			<CardActionArea onClick={() => onClick(cardId)}>
				<CardContent>
					<Typography variant="body1">{content}</Typography>
				</CardContent>
			</CardActionArea>
		</Card>
	);
}

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
	}),
);

type EditTodoDialogProps = {
	todo: Todo;
	open: boolean;
	onClose: () => void;
};

const editTodoPayloadSchema = z.object({
	content: z.string().min(1),
});

function EditTodoDialog({
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
				<Masonry columns={4} spacing={2}>
					{TODOS.map((todo) => (
						<CardItem
							key={todo.id}
							cardId={todo.id}
							content={todo.content}
							onCheck={handleChecked}
							onClick={handleClicked}
						/>
					))}
				</Masonry>
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
