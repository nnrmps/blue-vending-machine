import React from 'react';
import { Home } from './components/home/home';
import { Toaster } from 'react-hot-toast';

export const App: React.FC = () => {
  return (
    <div className='sm:p-[24px] md:p-[48px] lg:p-[128px] xl:p-[256px] bg-[#EBECEF] h-[100vh]'>
      <Toaster />
      <Home />
    </div>
  );
};
