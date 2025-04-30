import { z } from "zod";

export const todoEditableFieldsSchema = z.object({
  content: z.string().min(1),
});

export type TodoEditableFields = z.infer<typeof todoEditableFieldsSchema>;

export const todoSchema = todoEditableFieldsSchema.extend({
  id: z.number().min(1),
  createdAt: z.coerce.date(),
  updatedAt: z.coerce.date(),
});

export type Todo = z.infer<typeof todoSchema>;

export const getTodosResponseSchema = z.object({
  todos: z.array(todoSchema),
});

export type GetTodosResponse = z.infer<typeof getTodosResponseSchema>;
