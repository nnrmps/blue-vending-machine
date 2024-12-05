import { Service } from '@/service';
import { useMutation } from '@tanstack/react-query';

export const useDeleteProductById = () => {
  const { isError, isPending, mutateAsync } = useMutation({
    mutationFn: Service.deleteProductById,
  });

  return {
    mutateAsync,
    isError,
    isLoading: isPending,
  };
};
