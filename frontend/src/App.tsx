import Masonry from "@mui/lab/Masonry";
import Card from "@mui/material/Card";
import CardActionArea from "@mui/material/CardActionArea";
import CardActions from "@mui/material/CardActions";
import CardContent from "@mui/material/CardContent";
import Checkbox from "@mui/material/Checkbox";
import Container from "@mui/material/Container";
import CssBaseline from "@mui/material/CssBaseline";
import Typography from "@mui/material/Typography";
import type { SxProps, Theme } from "@mui/material/styles";
import type React from "react";

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

function handleChecked(cardId: number) {}

function handleClicked(cardId: number) {}

function App() {
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
		</>
	);
}

export default App;
