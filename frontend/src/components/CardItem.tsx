import Card from "@mui/material/Card";
import CardActionArea from "@mui/material/CardActionArea";
import CardActions from "@mui/material/CardActions";
import CardContent from "@mui/material/CardContent";
import Checkbox from "@mui/material/Checkbox";
import Typography from "@mui/material/Typography";
import type { SxProps, Theme } from "@mui/material/styles";

export type CardItemProps = {
  cardId: number;
  content: string;
  onCheck: (id: number) => void;
  onClick: (id: number) => void;
  sx?: SxProps<Theme>;
};

export function CardItem({
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
