import { useCreateProduct } from '@/hooks/use-create-product';
import { useGetProductById } from '@/hooks/use-get-product-by-id';
import { useGetProductList } from '@/hooks/use-get-product-list';
import { useUpdateProductById } from '@/hooks/use-update-product-by-id';
import { GetProductDetail } from '@/types/use-get-product-list.type';
import { Button, Form, Input, InputNumber, Modal, Space } from 'antd';
import { useEffect } from 'react';

type ModalEditProductProps = {
  productId: string;
  isOpen: boolean;
  onClose: () => void;
};

export const ModalEditProduct = ({
  productId,
  isOpen,
  onClose,
}: ModalEditProductProps) => {
  const { data: productDetail, isLoading } = useGetProductById(productId);
  const { refetch: refetchProductList } = useGetProductList();
  const {
    mutateAsync: mutateUpdateProduct,
    isLoading: isLoadingUpdateProduct,
  } = useUpdateProductById();
  const {
    mutateAsync: mutateCreateProduct,
    isLoading: isLoadingCreateProduct,
  } = useCreateProduct();
  const [form] = Form.useForm();
  const handleEditProduct = (value: GetProductDetail) => {
    if (productId) {
      mutateUpdateProduct({
        productId: productId,
        name: value.name,
        imageUrl: value.imageUrl,
        price: value.price,
        stock: value.stock,
      }).then(() => {
        refetchProductList();
        handleClose();
      });
    } else {
      mutateCreateProduct({
        name: value.name,
        imageUrl: value.imageUrl,
        price: value.price,
        stock: value.stock,
      }).then(() => {
        refetchProductList();
        handleClose();
      });
    }
  };

  const handleClose = () => {
    onClose();
    form.setFieldsValue({
      name: '',
      imageUrl: '',
      price: '',
      stock: '',
    });
  };

  useEffect(() => {
    if (!productDetail) return;
    form.setFieldsValue({
      name: productDetail?.name,
      imageUrl: productDetail?.imageUrl,
      price: productDetail?.price,
      stock: productDetail?.stock,
    });
  }, [productDetail, form]);

  return (
    <div>
      <Modal
        title={productId ? 'Edit Product' : 'Add Product'}
        open={isOpen && !isLoading}
        closable={false}
        footer={null}
      >
        <Form
          name='product-detail'
          form={form}
          labelCol={{ span: 8 }}
          wrapperCol={{ span: 12 }}
          className='my-[24px]'
          initialValues={{
            name: productDetail?.name,
            imageUrl: productDetail?.imageUrl,
            price: productDetail?.price,
            stock: productDetail?.stock,
          }}
          onFinish={handleEditProduct}
          preserve={false}
          autoComplete='off'
        >
          <Form.Item
            label='Product Name'
            name='name'
            rules={[{ required: true, message: 'Please input your username!' }]}
          >
            <Input />
          </Form.Item>

          <Form.Item
            label='Image Url'
            name='imageUrl'
            rules={[
              { required: true, message: 'Please input your image url!' },
              { type: 'url', message: 'Incorrect format!' },
            ]}
          >
            <Input />
          </Form.Item>

          <Form.Item
            label='Price'
            name='price'
            rules={[{ required: true, message: 'Please input your price!' }]}
          >
            <InputNumber />
          </Form.Item>
          <Form.Item
            label='Stock'
            name='stock'
            rules={[{ required: true, message: 'Please input your stock!' }]}
          >
            <InputNumber />
          </Form.Item>

          <Form.Item wrapperCol={{ offset: 8, span: 16 }}>
            <Space direction='horizontal'>
              <Button
                disabled={isLoadingUpdateProduct || isLoadingCreateProduct}
                type='default'
                onClick={handleClose}
              >
                cancel
              </Button>
              <Button
                disabled={isLoadingUpdateProduct || isLoadingCreateProduct}
                type='primary'
                htmlType='submit'
              >
                Submit
              </Button>
            </Space>
          </Form.Item>
        </Form>
      </Modal>
    </div>
  );
};
