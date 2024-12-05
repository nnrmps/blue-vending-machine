import { useGetReservedMoneyList } from '@/hooks/admin/use-get-reserved-money-list';
import { useUpdateReservedMoney } from '@/hooks/admin/use-update-reserved-money';
import { ReservedMoneyResponse } from '@/types/use-checkout-product.type';
import { SaveOutlined } from '@ant-design/icons';
import { Button, Form, InputNumber, Spin } from 'antd';
import Title from 'antd/es/typography/Title';
import { useEffect } from 'react';

export const TableReserveMoney = () => {
  const {
    data,
    isLoading,
    refetch: refetchReservedMoneyList,
  } = useGetReservedMoneyList();
  const {
    mutateAsync: mutateUpdateReserveMoney,
    isLoading: isLoadingUpdateReservedMoney,
  } = useUpdateReservedMoney();
  const [form] = Form.useForm();

  const handleUpdateReservedMoney = (value: ReservedMoneyResponse) => {
    mutateUpdateReserveMoney(value).then(() => {
      refetchReservedMoneyList();
    });
  };

  useEffect(() => {
    if (!data) return;
    form.setFieldsValue({
      coins1: data?.coins1,
      coins5: data?.coins5,
      coins10: data?.coins10,
      bank20: data?.bank20,
      bank50: data?.bank50,
      bank100: data?.bank100,
      bank500: data?.bank500,
      bank1000: data?.bank1000,
    });
  }, [data, form]);

  return (
    <Spin spinning={isLoading || isLoadingUpdateReservedMoney}>
      <Title level={5}>Reserved Money</Title>
      <Form
        name='reserved-money-list'
        form={form}
        labelCol={{ span: 10 }}
        wrapperCol={{ span: 24 }}
        className='mt-[32px]'
        initialValues={{
          coins1: data?.coins1,
          coins5: data?.coins5,
          coins10: data?.coins10,
          bank20: data?.bank20,
          bank50: data?.bank50,
          bank100: data?.bank100,
          bank500: data?.bank500,
          bank1000: data?.bank1000,
        }}
        onFinish={handleUpdateReservedMoney}
        preserve={false}
        autoComplete='off'
      >
        <Form.Item
          label='1 THB'
          name='coins1'
          rules={[
            {
              required: true,
              message: 'Please input your coins!',
            },
          ]}
        >
          <InputNumber min={0} />
        </Form.Item>
        <Form.Item
          label='5 THB'
          name='coins5'
          rules={[{ required: true, message: 'Please input your coins!' }]}
        >
          <InputNumber min={0} />
        </Form.Item>
        <Form.Item
          label='10 THB'
          name='coins10'
          rules={[{ required: true, message: 'Please input your coins!' }]}
        >
          <InputNumber min={0} />
        </Form.Item>
        <Form.Item
          label='20 THB'
          name='bank20'
          rules={[
            {
              required: true,
              message: 'Please input your banknotes!',
            },
          ]}
        >
          <InputNumber min={0} />
        </Form.Item>
        <Form.Item
          label='50 THB'
          name='bank50'
          rules={[
            {
              required: true,
              message: 'Please input your banknotes!',
            },
          ]}
        >
          <InputNumber min={0} />
        </Form.Item>
        <Form.Item
          label='100 THB'
          name='bank100'
          rules={[
            {
              required: true,
              message: 'Please input your banknotes!',
            },
          ]}
        >
          <InputNumber min={0} />
        </Form.Item>
        <Form.Item
          label='500 THB'
          name='bank500'
          rules={[
            {
              required: true,
              message: 'Please input your banknotes!',
            },
          ]}
        >
          <InputNumber min={0} />
        </Form.Item>
        <Form.Item
          label='1,000 THB'
          name='bank1000'
          rules={[
            {
              required: true,
              message: 'Please input your banknotes!',
            },
          ]}
        >
          <InputNumber min={0} />
        </Form.Item>
        <Form.Item wrapperCol={{ offset: 8, span: 16 }}>
          <Button
            style={{ float: 'right' }}
            icon={<SaveOutlined />}
            htmlType='submit'
            variant='solid'
            type='primary'
          >
            Update Reserved Money
          </Button>
        </Form.Item>
      </Form>
    </Spin>
  );
};
