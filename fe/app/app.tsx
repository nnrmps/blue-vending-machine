import React from 'react';
import { Home } from './components/home/home';
import { Layout } from 'antd';
import { Content } from 'antd/es/layout/layout';
import { Toaster } from 'react-hot-toast';

export const App: React.FC = () => {
  return (
    <Layout className='sm:p-[24px] lg:px-[300px] lg:py-[48px]'>
      <Toaster />
      <Content>
        <Home />
      </Content>
    </Layout>
  );
};
