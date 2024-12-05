import { Card, Modal, Space, Typography } from 'antd';
import { ShoppingOutlined, StopOutlined } from '@ant-design/icons';
import { GradientButton } from '../button/gradient-button';
import { useCheckoutProduct } from '@/hooks/use-checkout-product';
import { totalPriceType } from '@/types/use-checkout-product.type';
import { AxiosError } from 'axios';
import { GetProductDetail } from '@/types/use-get-product-list.type';
type ProductDetailProps = {
  data: GetProductDetail;
  isLoading: boolean;
  totalPrice: totalPriceType;
  sumPrice: number;
  handleClearData: () => void;
  onCheckout: (val: boolean) => void;
};

export const ProductDetail = ({
  data,
  isLoading,
  totalPrice,
  sumPrice,
  handleClearData,
  onCheckout,
}: ProductDetailProps) => {
  const { Meta } = Card;
  const { mutateAsync: mutateCheckout, isLoading: isLoadingCheckout } =
    useCheckoutProduct();
  const [modal, contextHolder] = Modal.useModal();

  const { Text } = Typography;

  const outOfStock = data?.stock === 0;

  const handleConfirmCheckout = () => {
    Modal.confirm({
      title: `Do you want to checkout "${data?.name}" ?`,
      content: (
        <Space direction='horizontal'>
          <Text style={{ fontSize: '14px' }} strong>
            Total Price:
          </Text>
          <Text style={{ fontSize: '14px' }}>
            <Space direction='horizontal'>
              <span>{data?.price}</span>
              <span>THB</span>
            </Space>
          </Text>
        </Space>
      ),
      okText: 'Confirm',
      okType: 'primary',
      cancelText: 'Cancel',
      onOk: () => {
        onCheckout(true);
        handleCheckout();
      },
    });
  };

  const handleCheckout = () => {
    mutateCheckout({
      productId: data?.productId,
      total: totalPrice,
    })
      .then((res) => {
        modal.success({
          title: (
            <div style={{ fontSize: '20px', marginTop: '-4px' }}>
              Checkout Success
            </div>
          ),
          content: (
            <div className='flex flex-col gap-[12px]'>
              <div>
                <Space direction='horizontal'>
                  <Text style={{ fontSize: '16px' }} strong>
                    Total Change:
                  </Text>
                  <Text style={{ fontSize: '16px' }}>
                    <Space direction='horizontal'>
                      <span>{res?.totalChange}</span>
                      <span>THB</span>
                    </Space>
                  </Text>
                </Space>
              </div>
              <div className='flex flex-col gap-[4px] ml-[2px]'>
                <Space direction='horizontal'>
                  <Text style={{ fontSize: '12px' }} strong>
                    1 THB:
                  </Text>
                  <Text style={{ fontSize: '12px' }}>
                    <Space direction='horizontal'>
                      <span>{res?.coins1}</span>
                      <span>coins</span>
                    </Space>
                  </Text>
                </Space>
                <Space direction='horizontal'>
                  <Text style={{ fontSize: '12px' }} strong>
                    5 THB:
                  </Text>
                  <Text style={{ fontSize: '12px' }}>
                    <Space direction='horizontal'>
                      <span>{res?.coins5}</span>
                      <span>coins</span>
                    </Space>
                  </Text>
                </Space>
                <Space direction='horizontal'>
                  <Text style={{ fontSize: '12px' }} strong>
                    10 THB:
                  </Text>
                  <Text style={{ fontSize: '12px' }}>
                    <Space direction='horizontal'>
                      <span>{res?.coins10}</span>
                      <span>coins</span>
                    </Space>
                  </Text>
                </Space>
                <Space direction='horizontal'>
                  <Text style={{ fontSize: '12px' }} strong>
                    20 THB:
                  </Text>
                  <Text style={{ fontSize: '12px' }}>
                    <Space direction='horizontal'>
                      <span>{res?.bank20}</span>
                      <span>banknotes</span>
                    </Space>
                  </Text>
                </Space>
                <Space direction='horizontal'>
                  <Text style={{ fontSize: '12px' }} strong>
                    50 THB:
                  </Text>
                  <Text style={{ fontSize: '12px' }}>
                    <Space direction='horizontal'>
                      <span>{res?.bank50}</span>
                      <span>banknotes</span>
                    </Space>
                  </Text>
                </Space>
                <Space direction='horizontal'>
                  <Text style={{ fontSize: '12px' }} strong>
                    100 THB:
                  </Text>
                  <Text style={{ fontSize: '12px' }}>
                    <Space direction='horizontal'>
                      <span>{res?.bank100}</span>
                      <span>banknotes</span>
                    </Space>
                  </Text>
                </Space>
                <Space direction='horizontal'>
                  <Text style={{ fontSize: '12px' }} strong>
                    500 THB:
                  </Text>
                  <Text style={{ fontSize: '12px' }}>
                    <Space direction='horizontal'>
                      <span>{res?.bank500}</span>
                      <span>banknotes</span>
                    </Space>
                  </Text>
                </Space>
                <Space direction='horizontal'>
                  <Text style={{ fontSize: '12px' }} strong>
                    1000 THB:
                  </Text>
                  <Text style={{ fontSize: '12px' }}>
                    <Space direction='horizontal'>
                      <span>{res?.bank1000}</span>
                      <span>banknotes</span>
                    </Space>
                  </Text>
                </Space>
              </div>
            </div>
          ),
          onOk() {
            handleClearData();
          },
          okText: 'Close',
          okButtonProps: { className: 'bg-[#ff4d4f]' },
        });
      })
      .catch((error) => {
        if (error instanceof AxiosError) {
          modal.error({
            title: (
              <div style={{ fontSize: '20px', marginTop: '-4px' }}>
                Checkout Failed!
              </div>
            ),
            content: (
              <Text style={{ fontSize: '14px' }}>{error?.response?.data}</Text>
            ),
            onOk() {
              handleClearData();
            },
            okText: 'Close',
            okButtonProps: { className: 'bg-[#ff4d4f]' },
          });
        }
      })
      .finally(() => {
        onCheckout(false);
      });
  };

  return (
    <>
      <Card
        loading={isLoading}
        className='w-[250px]'
        cover={<img alt='product-image' src={data?.imageUrl} />}
        actions={[
          <GradientButton
            icon={outOfStock ? <StopOutlined /> : <ShoppingOutlined />}
            title={outOfStock ? 'Out of Stock!' : 'Buy Now!'}
            variant='filled'
            type='primary'
            disabled={sumPrice < data?.price || outOfStock || isLoadingCheckout}
            onClick={handleConfirmCheckout}
          />,
        ]}
      >
        <Meta
          title={
            <Text ellipsis style={{ fontSize: '20px' }} strong>
              {data?.name}
            </Text>
          }
          description={
            <div className='flex flex-col gap-[6px]'>
              <Space direction='horizontal'>
                <Text style={{ fontSize: '16px' }} strong>
                  Price:
                </Text>
                <Text style={{ fontSize: '16px' }}>
                  <Space direction='horizontal'>
                    <span>{data?.price}</span>
                    <span>THB</span>
                  </Space>
                </Text>
              </Space>
              <Space direction='horizontal'>
                <Text style={{ fontSize: '16px' }} strong>
                  Total Stock:
                </Text>
                <Text style={{ fontSize: '16px' }}>
                  <Space direction='horizontal'>
                    <span>
                      {data?.stock > 9999
                        ? '9,999+'
                        : data?.stock
                            .toString()
                            .replace(/\B(?=(\d{3})+(?!\d))/g, ',')}
                    </span>
                    <span>Piece</span>
                  </Space>
                </Text>
              </Space>
            </div>
          }
        />
      </Card>
      {contextHolder}
    </>
  );
};
