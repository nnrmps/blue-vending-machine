import { Service } from '@/service';
import { useMutation } from '@tanstack/react-query';

export const useCheckoutProduct = () => {
  const { isError, isPending, mutateAsync } = useMutation({
    mutationFn: Service.checkoutProduct,
  });

  return {
    mutateAsync,
    isError,
    isLoading: isPending,
  };
};
