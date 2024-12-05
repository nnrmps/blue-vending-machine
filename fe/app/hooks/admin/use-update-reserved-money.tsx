import { AdminService } from '@/service/admin-service.apis';
import { useMutation } from '@tanstack/react-query';

export const useUpdateReservedMoney = () => {
  const { isError, isPending, mutateAsync } = useMutation({
    mutationFn: AdminService.updateReservedMoney,
  });

  return {
    mutateAsync,
    isError,
    isLoading: isPending,
  };
};
