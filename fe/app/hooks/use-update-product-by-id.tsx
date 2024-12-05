import { Service } from '@/service';
import { useMutation } from '@tanstack/react-query';

export const useUpdateProductById = () => {
  const { isError, isPending, mutateAsync } = useMutation({
    mutationFn: Service.updateProductById,
  });

  return {
    mutateAsync,
    isError,
    isLoading: isPending,
  };
};
