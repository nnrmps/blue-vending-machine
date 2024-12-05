import { AdminService } from '@/service/admin-service.apis';
import { useMutation } from '@tanstack/react-query';

export const useCreateProduct = () => {
  const { isError, isPending, mutateAsync } = useMutation({
    mutationFn: AdminService.createProduct,
  });

  return {
    mutateAsync,
    isError,
    isLoading: isPending,
  };
};
