import { z } from 'zod';

const TotalSchema = z.object({
  coins1: z.number(),
  coins5: z.number(),
  coins10: z.number(),
  bank20: z.number(),
  bank50: z.number(),
  bank100: z.number(),
  bank500: z.number(),
  bank1000: z.number(),
});

const CheckoutPayloadSchema = z.object({
  productId: z.string(),
  total: TotalSchema,
});

const CheckoutResponseSchema = z.object({
  totalChange: z.number(),
  coins1: z.number(),
  coins5: z.number(),
  coins10: z.number(),
  bank20: z.number(),
  bank50: z.number(),
  bank100: z.number(),
  bank500: z.number(),
  bank1000: z.number(),
});

export type totalPriceType = z.infer<typeof TotalSchema>;
export type CheckoutPayload = z.infer<typeof CheckoutPayloadSchema>;
export type CheckoutResponse = z.infer<typeof CheckoutResponseSchema>;
