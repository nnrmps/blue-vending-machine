import { Card } from 'antd';
import { ProductDetailType } from '@/types/use-get-product-list.type';
import { ShoppingOutlined, StopOutlined } from '@ant-design/icons';
import { GradientButton } from '../button/gradient-button';
import { useCheckoutProduct } from '@/hooks/use-checkout-product';
import { totalPriceType } from '@/types/use-checkout-product.type';
import { AxiosError } from 'axios';
import { useToast } from '../snackbar/use-toast';
type ProductDetailProps = {
  data: ProductDetailType;
  isLoading: boolean;
  totalPrice: totalPriceType;
  sumPrice: number;
  handleClearData: () => void;
};

export const ProductDetail = ({
  data,
  isLoading,
  totalPrice,
  sumPrice,
  handleClearData,
}: ProductDetailProps) => {
  const { Meta } = Card;
  const { mutateAsync: mutateCheckout, isLoading: isLoadingCheckout } =
    useCheckoutProduct();

  const { toast } = useToast();

  const outOfStock = data?.stock === 0;

  const handleCheckout = () => {
    mutateCheckout({
      productId: data?.productId,
      total: totalPrice,
    })
      .then((res) => {
        toast.success({
          content: 'Total Change: ' + res?.data?.totalChange,
        });
      })
      .catch((error) => {
        if (error instanceof AxiosError) {
          toast.error({
            content: error?.response?.data,
          });
        }
      })
      .finally(() => {
        handleClearData();
      });
  };

  return (
    <>
      <Card
        loading={isLoading}
        cover={<img alt='product-image' src={data?.imageUrl} />}
        actions={[
          <GradientButton
            icon={outOfStock ? <StopOutlined /> : <ShoppingOutlined />}
            title={outOfStock ? 'Out of Stock!' : 'Buy Now!'}
            variant='filled'
            type='primary'
            disabled={sumPrice < data?.price || outOfStock}
            onClick={handleCheckout}
          />,
        ]}
      >
        <Meta title={data?.name} description={data?.price} />
      </Card>
    </>
  );
};
