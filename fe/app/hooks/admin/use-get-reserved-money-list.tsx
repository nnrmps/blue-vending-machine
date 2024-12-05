import { AdminService } from '@/service/admin-service.apis';
import { useQuery } from '@tanstack/react-query';

export const useGetReservedMoneyList = () => {
  const { data, isLoading, refetch } = useQuery({
    queryKey: ['reserveMoneyList'],
    queryFn: () => AdminService.getReservedMoneyList(),
  });

  return {
    data,
    isLoading,
    refetch,
  };
};
