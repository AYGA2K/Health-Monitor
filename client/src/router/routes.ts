import { RouteRecordRaw } from 'vue-router';

const routes: RouteRecordRaw[] = [
  {
    path: '/sign-in',
    component: () => import('pages/client/SignIn.vue')
  },
  {
    path: '/sign-up',
    component: () => import('pages/client/SignUp.vue')
  },
  {
    path: '/',
    meta: {
      needsAuth: true
    },
    component: () => import('layouts/HomeLayout.vue'),
    children: [
      {
        path: '',
        component: () => import('pages/client/HomePage.vue')
      }
    ]
  },
  {
    path: '/dashboard',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      {
        path: '',
        component: () => import('pages/IndexPage.vue')
      },
      {
        path: '/add-client',
        component: () => import('pages/AddClient.vue')

      },
      {
        path: '/edit-client/:id',
        component: () => import('pages/EditClient.vue')
      },
      {
        path: '/add-client',
        component: () => import('pages/AddClient.vue')
      },
      {
        path: '/edit-client/:id',
        component: () => import('pages/EditClient.vue')
      },
      {
        path: '/view-client/:id',
        component: () => import('pages/ViewClient.vue')

      },
      {
        path: '/socket',
        component: () => import('pages/SocketClient.vue')
      },


      {
        path: '/view-client/:id',
        component: () => import('layouts/MainLayout.vue'),
        children: [{ path: '', component: () => import('pages/ViewClient.vue') }],

      },

    ],
  },
  // Always leave this as last one,
  // but you can also remove it
  {
    path: '/:catchAll(.*)*',
    component: () => import('pages/ErrorNotFound.vue'),
  },
];

export default routes;
