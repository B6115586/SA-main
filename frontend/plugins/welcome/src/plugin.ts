import { createPlugin, createRouteRef } from '@backstage/core';
import Problem from './components/Problem';

export const rootRouteRef = createRouteRef({
  path: '/problem',
  title: 'problem',
});

export const plugin = createPlugin({
  id: 'problem',
  register({ router }) {
    router.addRoute(rootRouteRef, Problem);
  },
});
