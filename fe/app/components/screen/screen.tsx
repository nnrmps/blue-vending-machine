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
    <Carousel arrows infinite={false}>
      {data?.map((item: ProductDetailType) => (
        <div
          key={item.productId}
          className='flex bg-[#B3CEE5] p-[24px] md:p-[48px]'
        >
          <ProductDetail
            data={item}
            isLoading={isLoading}
            totalPrice={totalPrice}
            sumPrice={sumPrice}
            handleClearData={handleClearData}
          />
        </div>
      ))}
    </Carousel>
  );
};
