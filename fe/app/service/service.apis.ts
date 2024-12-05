import { axiosInstance, ServiceAdminAPI, ServiceAPI } from '@/service';
import { LoginSchemaType } from '@/types/login.type';
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
import { adminAxiosInstance } from './admin-axios-instance';

export const Service = {
  //user
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
  checkoutProduct: async (
    payload: CheckoutPayload
  ): Promise<CheckoutResponse> => {
    const { data } = await axiosInstance.post(ServiceAPI.checkout, payload);
    return data.data;
  },
  login: async (payload: LoginSchemaType) => {
    const { data } = await axiosInstance.post(ServiceAPI.login, payload);
    return data;
  },

  //admin
  getAdminProductList: async (): Promise<GetProductListResponse> => {
    const { data } = await adminAxiosInstance.get(`${ServiceAdminAPI.product}`);
    return data.data;
  },
  getAdminProductById: async (productId: string): Promise<GetProductDetail> => {
    const { data } = await adminAxiosInstance.get(
      `${ServiceAdminAPI.product}/${productId}`
    );
    return data.data;
  },
  updateProductById: async (
    payload: UpdateProductPayload
  ): Promise<GetProductDetail> => {
    const { data } = await adminAxiosInstance.put(
      `${ServiceAdminAPI.product}/${payload.productId}`,
      payload
    );
    return data;
  },
  createProduct: async (
    payload: CreateProductPayload
  ): Promise<GetProductDetail> => {
    const { data } = await adminAxiosInstance.post(
      ServiceAdminAPI.product,
      payload
    );
    return data;
  },

  deleteProductById: async (productId: string) => {
    const { data } = await adminAxiosInstance.delete(
      `${ServiceAdminAPI.product}/${productId}`
    );
    return data.data;
  },

  getReservedMoneyList: async (): Promise<ReservedMoneyResponse> => {
    const { data } = await adminAxiosInstance.get(
      `${ServiceAdminAPI.reservedMoney}`
    );
    return data.data;
  },
  updateReservedMoney: async (payload: ReservedMoneyResponse) => {
    const { data } = await adminAxiosInstance.put(
      ServiceAdminAPI.reservedMoney,
      payload
    );
    return data;
  },
};
