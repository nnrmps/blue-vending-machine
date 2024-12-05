import { z } from 'zod';

const LoginSchema = z.object({
  username: z.string(),
  password: z.number().or(z.string()),
});

export type LoginSchemaType = z.infer<typeof LoginSchema>;
