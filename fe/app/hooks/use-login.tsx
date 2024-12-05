import { Service } from '@/service';
import { useMutation } from '@tanstack/react-query';

export const useLogin = () => {
  const { isError, isPending, mutateAsync } = useMutation({
    mutationFn: Service.login,
  });

  return {
    mutateAsync,
    isError,
    isLoading: isPending,
  };
};
