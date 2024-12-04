import axios from 'axios';

const API_ENDPOINT = '/api';

export const axiosInstance = axios.create({
  baseURL: API_ENDPOINT,
  withCredentials: true,
});

axiosInstance.interceptors.request.use((config) => {
  return config;
});

axiosInstance.interceptors.response.use(
  (response) => {
    return response;
  },
  (error) => {
    return Promise.reject(error);
  }
);

export type ErrorResponse<T = any> = {
  id: string;
  code: string;
  title: string;
  meta?: T;
};

export type BaseResponse<T, V = any, U = any> = {
  data?: T;
  errors?: ErrorResponse<U>[];
  meta?: V;
};

export type BaseGetResponse<T> = {
  data: T;
};
