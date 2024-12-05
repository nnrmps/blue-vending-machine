import React, { useState } from 'react';
import { Routes, Route, Navigate, useNavigate } from 'react-router-dom';
import { HomePage } from './pages/home-page/home-page';
import { AdminPage } from './pages/admin-page/admin-page';
import { Button, Input, Layout, Modal } from 'antd';
import { Header } from 'antd/es/layout/layout';
import { SettingOutlined } from '@ant-design/icons';
export const App: React.FC = () => {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [userName, setUserName] = useState('');
  const [password, setPassword] = useState('');
  const navigate = useNavigate();

  const handleLogin = () => {
    setIsModalOpen(false);
    setUserName('');
    setPassword('');
    navigate('/admin');
  };
  const handleCancelLogin = () => {
    setIsModalOpen(false);
    setUserName('');
    setPassword('');
  };

  return (
    <div className='relative bg-[#EBECEF] h-[100vh]'>
      <Layout>
        <Header style={{ backgroundColor: '#95b9c7', padding: 0 }}>
          <Button
            type='default'
            variant='filled'
            icon={<SettingOutlined />}
            style={{
              float: 'right',
              margin: '12px',
            }}
            onClick={() => setIsModalOpen(true)}
          >
            Maintenance Mode
          </Button>
        </Header>
        <div className='sm:p-[24px] md:p-[48px] lg:p-[128px] xl:p-[256px] '>
          <Routes>
            <Route path='/' element={<HomePage />} />
            <Route path='/admin' element={<AdminPage />} />
            <Route path='*' element={<Navigate to='/' replace />} />
          </Routes>
        </div>
      </Layout>
      <Modal
        title='Login'
        open={isModalOpen}
        onOk={handleLogin}
        onCancel={handleCancelLogin}
        okButtonProps={{ disabled: !userName || !password }}
        okText='Login'
        closable={false}
        width={250}
      >
        <div className='flex flex-col gap-[16px] my-[24px] justify-center'>
          <Input
            placeholder='Username'
            value={userName}
            onChange={(e) => setUserName(e.target.value)}
          />
          <Input.Password
            placeholder='Password'
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
        </div>
      </Modal>
    </div>
  );
};
