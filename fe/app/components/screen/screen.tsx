import { Carousel } from 'antd';
import { useGetProductList } from '@/hooks/use-get-product-list';
import { ProductDetailType } from '@/types/use-get-product-list.type';
import { ProductDetail } from '../product-detail/product-detail';
import { totalPriceType } from '@/types/use-checkout-product.type';

type ScreenProps = {
  totalPrice: totalPriceType;
  sumPrice: number;
  handleClearData: () => void;
};

export const Screen = ({
  totalPrice,
  sumPrice,
  handleClearData,
}: ScreenProps) => {
  const { data, isLoading } = useGetProductList();

  if (!data || !data?.length) {
    return null;
  }

  return (
    <div className='flex snap-x snap-mandatory overflow-y-auto bg-[#B3CEE5]'>
      {data?.map((item: ProductDetailType) => (
        <div key={item.productId} className='snap-always snap-center p-[24px]'>
          <ProductDetail
            data={item}
            isLoading={isLoading}
            totalPrice={totalPrice}
            sumPrice={sumPrice}
            handleClearData={handleClearData}
          />
        </div>
      ))}
    </div>
  );
};
