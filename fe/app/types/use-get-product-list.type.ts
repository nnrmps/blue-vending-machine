import { z } from 'zod';

const ProductDetail = z.object({
  productId: z.string(),
  name: z.string(),
  imageUrl: z.string().optional(),
  stock: z.number(),
  price: z.number(),
});

export type GetProductDetail = z.infer<typeof ProductDetail>;
export type GetProductListResponse = GetProductDetail[];
