import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    name: 'index',
    component: () => import('../views/Index.vue'),
    redirect: '/login',
    children: [
      { path: '/login', component: () => import('../views/Login.vue') },
    ],
  },
  {
    path: '/home',
    name: 'home',
    component: () => import('../views/Home.vue'),
    redirect: '/dashboard',
    children: [
      {
        path: '/dashboard',
        name: 'dashboard',
        component: () => import('../views/Dashboard.vue'),
      },
      {
        path: '/customer',
        name: 'customer',
        component: () => import('../views/Customer.vue'),
      },
      {
        path: '/contract',
        name: 'contract',
        component: () => import('../views/Contract.vue'),
      },
      {
        path: '/product',
        name: 'product',
        component: () => import('../views/Product.vue'),
      },
      {
        path: '/config',
        name: 'config',
        component: () => import('../views/Config.vue'),
      },
      {
        path: '/subscribe',
        name: 'subscribe',
        component: () => import('../views/Subscribe.vue'),
      },
      {
        path: '/result',
        name: 'result',
        component: () => import('../views/Result.vue'),
      },
    ],
  },
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: routes,
})

export default router
