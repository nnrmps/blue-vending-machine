import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
import path from 'path';

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    host: '0.0.0.0',
    port: 3000,
    proxy: {
      '/api': 'http://api:8080', // Proxy API requests to backend
      '/admin-api': 'http://api:8080',
    },
  },
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './app'),
    },
  },
});
