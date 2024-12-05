import { AdminService } from '@/service/admin-service.apis';
import { useMutation } from '@tanstack/react-query';

export const useUpdateProductById = () => {
  const { isError, isPending, mutateAsync } = useMutation({
    mutationFn: AdminService.updateProductById,
  });

  return {
    mutateAsync,
    isError,
    isLoading: isPending,
  };
};
