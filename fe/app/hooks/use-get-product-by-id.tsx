import { Service } from '@/service';
import { useQuery } from '@tanstack/react-query';

export const useGetProductById = (productId: string) => {
  const { data, isLoading, refetch } = useQuery({
    queryKey: ['productById', productId],
    queryFn: () => Service.getProductById(productId),
    enabled: !!productId,
  });

  return {
    data,
    isLoading,
    refetch,
  };
};
