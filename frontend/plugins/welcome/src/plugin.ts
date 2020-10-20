import { createPlugin } from '@backstage/core';
import User from './components/User'


export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', User);
    

  },
});
