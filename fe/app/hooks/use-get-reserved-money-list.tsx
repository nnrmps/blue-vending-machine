import { Service } from '@/service';
import { useQuery } from '@tanstack/react-query';

export const useGetReservedMoneyList = () => {
  const { data, isLoading, refetch } = useQuery({
    queryKey: ['reserveMoneyList'],
    queryFn: () => Service.getReservedMoneyList(),
  });

  return {
    data,
    isLoading,
    refetch,
  };
};
