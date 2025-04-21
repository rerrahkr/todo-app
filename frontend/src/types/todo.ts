import { z } from "zod";

export const todoEditableFieldsSchema = z.object({
  content: z.string().min(1),
});

export type TodoEditableFields = z.infer<typeof todoEditableFieldsSchema>;

export type Todo = {
  id: number;
  createdAt: Date;
  updatedAt: Date;
} & TodoEditableFields;
