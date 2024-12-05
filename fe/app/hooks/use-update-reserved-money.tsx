import { Service } from '@/service';
import { useMutation } from '@tanstack/react-query';

export const useUpdateReservedMoney = () => {
  const { isError, isPending, mutateAsync } = useMutation({
    mutationFn: Service.updateReservedMoney,
  });

  return {
    mutateAsync,
    isError,
    isLoading: isPending,
  };
};
