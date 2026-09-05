import axios, { type AxiosError, type AxiosRequestConfig } from 'axios';

export const axiosInstance = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL || '/api',
  timeout: 15_000,
});

// Return the backend body directly (including pagination), not AxiosResponse.
export const customAxios = async <T>(config: AxiosRequestConfig, options?: AxiosRequestConfig): Promise<T> => {
  // mergeConfig takes url/method/data only from its second argument.
  const response = await axiosInstance<T>(axios.mergeConfig(config, { ...config, ...options }));
  return response.data;
};

export type ErrorType<T> = AxiosError<T>;
