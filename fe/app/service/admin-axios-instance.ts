import axios from 'axios';

const API_ENDPOINT = '/admin-api';
export const adminAxiosInstance = axios.create({
  baseURL: API_ENDPOINT,
  withCredentials: true,
});

adminAxiosInstance.interceptors.request.use((config) => {
  config.headers['Authorization'] = `Bearer ${localStorage.getItem('token')}`;
  return config;
});

adminAxiosInstance.interceptors.response.use(
  (response) => {
    return response;
  },
  (error) => {
    if (error.status === 401) {
      window.location.replace('/');
      localStorage.removeItem('token');
    }
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
