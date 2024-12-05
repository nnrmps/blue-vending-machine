import { AdminService } from '@/service/admin-service.apis';
import { useQuery } from '@tanstack/react-query';

export const useGetProductList = () => {
  const { data, isLoading, refetch } = useQuery({
    queryKey: ['productList'],
    queryFn: () => AdminService.getAdminProductList(),
  });

  return {
    data,
    isLoading,
    refetch,
  };
};
