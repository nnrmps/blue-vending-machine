import { GetProductDetail } from '@/types/use-get-product-list.type';
import {
  DeleteOutlined,
  EditOutlined,
  PlusCircleOutlined,
} from '@ant-design/icons';
import { Button, Spin, Table } from 'antd';
import { useState } from 'react';
import { ModalEditProduct } from '../modal-edit-product/modal-edit-product';
import { useDeleteProductById } from '@/hooks/admin/use-delete-product-by-id';
import { useGetProductList } from '@/hooks/admin/use-get-product-list';

export const TableProduct = () => {
  const {
    data: productList,
    isLoading,
    refetch: refetchProductList,
  } = useGetProductList();
  const { mutateAsync, isLoading: isLoadingDeleteProduct } =
    useDeleteProductById();
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [productId, setProductId] = useState('');

  const columns = [
    {
      title: 'Product Name',
      dataIndex: 'name',
      key: 'name',
    },
    {
      title: 'Price',
      dataIndex: 'price',
      key: 'price',
    },
    {
      title: 'Stock',
      dataIndex: 'stock',
      key: 'stock',
    },
    {
      title: '',
      dataIndex: 'action',
      key: 'action',
      render: (_index: number, record: GetProductDetail) => (
        <div className='flex gap-[12px]'>
          <Button
            size='small'
            onClick={() => handleEditProduct(record.productId)}
            disabled={!record.productId}
            icon={<EditOutlined />}
          />
          <Button
            size='small'
            onClick={() => {
              mutateAsync(record.productId).then(() => {
                refetchProductList();
              });
            }}
            disabled={!record.productId}
            danger
            type='default'
            icon={<DeleteOutlined />}
          />
        </div>
      ),
    },
  ];

  const handleEditProduct = (productId: string) => {
    setIsModalOpen(true);
    setProductId(productId);
  };

  const handleCloseModal = () => {
    setIsModalOpen(false);
    setProductId('');
  };

  return (
    <Spin spinning={isLoading || isLoadingDeleteProduct}>
      <div className='flex flex-col gap-[16px]'>
        <div>
          <Button
            size='small'
            style={{ float: 'right' }}
            type='primary'
            icon={<PlusCircleOutlined />}
            onClick={() => handleEditProduct('')}
          >
            Add Product
          </Button>
        </div>
        <Table columns={columns} dataSource={productList} pagination={false} />
        <ModalEditProduct
          productId={productId}
          isOpen={isModalOpen}
          onClose={handleCloseModal}
        />
      </div>
    </Spin>
  );
};
