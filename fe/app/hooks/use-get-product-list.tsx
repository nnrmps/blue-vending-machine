import { Service } from '@/service';
import { useQuery } from '@tanstack/react-query';

export const useGetProductList = () => {
  const { data, isLoading } = useQuery({
    queryKey: ['productList'],
    queryFn: () => Service.getProductList(),
  });

  return {
    data,
    isLoading,
  };
};
