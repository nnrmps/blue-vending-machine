import { Service } from '@/service';
import { useMutation } from '@tanstack/react-query';

export const useCreateProduct = () => {
  const { isError, isPending, mutateAsync } = useMutation({
    mutationFn: Service.createProduct,
  });

  return {
    mutateAsync,
    isError,
    isLoading: isPending,
  };
};
