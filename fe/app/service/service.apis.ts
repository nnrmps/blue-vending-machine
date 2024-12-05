import { axiosInstance, ServiceAPI } from '@/service';
import {
  CheckoutPayload,
  CheckoutResponse,
  ReservedMoneyResponse,
} from '@/types/use-checkout-product.type';
import {
  GetProductDetail,
  GetProductListResponse,
} from '@/types/use-get-product-list.type';
import {
  CreateProductPayload,
  UpdateProductPayload,
} from '@/types/use-update-product-by-id.type';

export const Service = {
  getProductList: async (): Promise<GetProductListResponse> => {
    const { data } = await axiosInstance.get(`${ServiceAPI.product}`);
    return data.data;
  },
  getProductById: async (productId: string): Promise<GetProductDetail> => {
    const { data } = await axiosInstance.get(
      `${ServiceAPI.product}/${productId}`
    );
    return data.data;
  },
  updateProductById: async (
    payload: UpdateProductPayload
  ): Promise<GetProductDetail> => {
    const { data } = await axiosInstance.put(
      `${ServiceAPI.product}/${payload.productId}`,
      payload
    );
    return data;
  },
  createProduct: async (
    payload: CreateProductPayload
  ): Promise<GetProductDetail> => {
    const { data } = await axiosInstance.post(ServiceAPI.product, payload);
    return data;
  },

  deleteProductById: async (productId: string) => {
    const { data } = await axiosInstance.delete(
      `${ServiceAPI.product}/${productId}`
    );
    return data.data;
  },
  checkoutProduct: async (
    payload: CheckoutPayload
  ): Promise<CheckoutResponse> => {
    const { data } = await axiosInstance.post(ServiceAPI.checkout, payload);
    return data;
  },
  getReservedMoneyList: async (): Promise<ReservedMoneyResponse> => {
    const { data } = await axiosInstance.get(`${ServiceAPI.reservedMoney}`);
    return data.data;
  },
  updateReservedMoney: async (payload: ReservedMoneyResponse) => {
    const { data } = await axiosInstance.put(ServiceAPI.reservedMoney, payload);
    return data;
  },
};
