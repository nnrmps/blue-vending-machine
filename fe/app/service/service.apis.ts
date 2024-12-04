import { axiosInstance, ServiceAPI } from '@/service';
import {
  CheckoutPayload,
  CheckoutResponse,
} from '@/types/use-checkout-product.type';
import { GetProductListResponse } from '@/types/use-get-product-list.type';

export const Service = {
  getProductList: async (): Promise<GetProductListResponse> => {
    const { data } = await axiosInstance.get(`${ServiceAPI.getProduct}`);
    return data.data;
  },
  checkoutProduct: async (
    payload: CheckoutPayload
  ): Promise<CheckoutResponse> => {
    const { data } = await axiosInstance.post(ServiceAPI.checkout, payload);
    return data;
  },
};
