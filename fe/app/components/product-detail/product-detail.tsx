import { Card, Modal } from 'antd';
import { ProductDetailType } from '@/types/use-get-product-list.type';
import { ShoppingOutlined, StopOutlined } from '@ant-design/icons';
import { GradientButton } from '../button/gradient-button';
import { useCheckoutProduct } from '@/hooks/use-checkout-product';
import { totalPriceType } from '@/types/use-checkout-product.type';
import { AxiosError } from 'axios';
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
  const { mutateAsync: mutateCheckout } = useCheckoutProduct();
  const [modal, contextHolder] = Modal.useModal();

  const outOfStock = data?.stock === 0;

  const handleCheckout = () => {
    mutateCheckout({
      productId: data?.productId,
      total: totalPrice,
    })
      .then((res) => {
        modal.success({
          title: 'Checkout Success',
          content: (
            <>
              <div>Total Change: {res?.data?.totalChange}</div>
              <div>1 THB: {res?.data?.coins1} coins</div>
              <div>5 THB: {res?.data?.coins5} coins</div>
              <div>10 THB: {res?.data?.coins10} coins</div>
              <div>20 THB: {res?.data?.bank20} coins</div>
              <div>50 THB: {res?.data?.bank50} coins</div>
              <div>100 THB: {res?.data?.bank100} coins</div>
              <div>500 THB: {res?.data?.bank500} coins</div>
              <div>1000 THB: {res?.data?.bank1000} coins</div>
            </>
          ),
          onOk() {
            handleClearData();
          },
        });
      })
      .catch((error) => {
        if (error instanceof AxiosError) {
          modal.error({
            title: 'Checkout Filed!',
            content: (
              <>
                <div>{error?.response?.data}</div>
              </>
            ),
            onOk() {
              handleClearData();
            },
          });
        }
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
      {contextHolder}
    </>
  );
};
