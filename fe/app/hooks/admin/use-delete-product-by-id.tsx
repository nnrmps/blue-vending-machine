import { AdminService } from '@/service/admin-service.apis';
import { useMutation } from '@tanstack/react-query';

export const useDeleteProductById = () => {
  const { isError, isPending, mutateAsync } = useMutation({
    mutationFn: AdminService.deleteProductById,
  });

  return {
    mutateAsync,
    isError,
    isLoading: isPending,
  };
};
