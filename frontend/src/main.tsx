import { createRoot } from 'react-dom/client';
import router from './routers';
import '@/assets/styles/golbal.css';
import { RouterProvider } from 'react-router/dom';
import { applyColorTheme, readStoredColorTheme } from '@/utils/theme-helper';

applyColorTheme(readStoredColorTheme());

const root = createRoot(document.querySelector('#root')!);
root.render(
  <RouterProvider router={router} />,
);
