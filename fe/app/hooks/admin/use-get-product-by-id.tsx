import { AdminService } from '@/service/admin-service.apis';
import { useQuery } from '@tanstack/react-query';

export const useGetProductById = (productId: string) => {
  const { data, isLoading, refetch } = useQuery({
    queryKey: ['productById', productId],
    queryFn: () => AdminService.getAdminProductById(productId),
    enabled: !!productId,
  });

  return {
    data,
    isLoading,
    refetch,
  };
};
