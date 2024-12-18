import { TableProduct } from '@/components/table-product/table-product';
import { TableReserveMoney } from '@/components/table-reserved-money/table-reserved-money';
import { CreditCardOutlined, ShoppingOutlined } from '@ant-design/icons';
import { Tabs } from 'antd';
import { useEffect } from 'react';
import { useNavigate } from 'react-router-dom';

export const AdminPage = () => {
  const navigate = useNavigate();
  const token = localStorage.getItem('token');

  useEffect(() => {
    if (!token) {
      navigate('/');
    }
  }, [token]);

  return (
    <Tabs>
      <Tabs.TabPane tab='Product' key='1' icon={<ShoppingOutlined />}>
        <TableProduct />
      </Tabs.TabPane>
      <Tabs.TabPane tab='Reserved Money' key='2' icon={<CreditCardOutlined />}>
        <TableReserveMoney />
      </Tabs.TabPane>
    </Tabs>
  );
};
