import { useGetProductList } from '@/hooks/use-get-product-list';
import { ProductDetail } from '../product-detail/product-detail';
import { totalPriceType } from '@/types/use-checkout-product.type';
import { GetProductDetail } from '@/types/use-get-product-list.type';

type ScreenProps = {
  totalPrice: totalPriceType;
  sumPrice: number;
  handleClearData: () => void;
  onCheckout: (val: boolean) => void;
};

export const Screen = ({
  totalPrice,
  sumPrice,
  handleClearData,
  onCheckout,
}: ScreenProps) => {
  const { data, isLoading } = useGetProductList();

  if (!data || !data?.length) {
    return null;
  }

  return (
    <div className='flex h-[400px] items-center snap-x snap-mandatory overflow-y-auto rounded-[8px] bg-[#B3CEE5]'>
      {data?.map((item: GetProductDetail) => (
        <div key={item.productId} className='snap-always snap-center p-[24px]'>
          <ProductDetail
            data={item}
            isLoading={isLoading}
            totalPrice={totalPrice}
            sumPrice={sumPrice}
            handleClearData={handleClearData}
            onCheckout={onCheckout}
          />
        </div>
      ))}
    </div>
  );
};
