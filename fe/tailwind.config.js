/** @type {import('tailwindcss').Config} */
module.exports = {
  darkMode: ['class'],
  content: ['./public/index.html', './app/**/*.{ts,tsx,css}'],
  presets: [],
  prefix: '',
  theme: {
    screens: {
      sm: '640px',
      md: '768px',
      lg: '1024px',
      xl: '1440px',
      '2xl': '1920px',
    },
    container: {
      center: 'true',
    },
    extend: {},
    plugins: [require('tailwindcss-animate')],
  },
};
