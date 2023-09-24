import { z } from "zod";

export const todoSchema = z
  .object({
    id: z.string().default(crypto.randomUUID()),
    description: z.string().min(5),
    completed: z.boolean().default(false)
  })

export type TodoSchema = z.infer<typeof todoSchema>;
