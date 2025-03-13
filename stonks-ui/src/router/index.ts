import { createRouter, createWebHistory } from 'vue-router';
import StocksView from '@/views/StocksView.vue';
import RecommendationsView from '@/views/RecommendationsView.vue';

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      redirect: '/stocks'
    },
    {
      path: '/stocks',
      name: 'stocks',
      component: StocksView
    },
    {
      path: '/recommendations',
      name: 'recommendations',
      component: RecommendationsView
    }
  ]
});

export default router; 