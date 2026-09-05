import { lazy } from 'react';
import { createHashRouter } from 'react-router';
import Root from '@/components/root';

const router = createHashRouter([
  {
    path: '/',
    Component: Root,
    children: [
      { index: true, Component: lazy(() => import('@/pages/home')) },
      { path: '/articles/new', Component: lazy(() => import('@/pages/article-editor')) },
      { path: '/articles/:id', Component: lazy(() => import('@/pages/article-detail')) },
      { path: '/articles/:id/edit', Component: lazy(() => import('@/pages/article-editor')) },
    ],
  },
]);

export default router;
