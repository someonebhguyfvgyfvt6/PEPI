import { createRouter, createWebHistory } from 'vue-router'

export default createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/', name: 'inicio', component: () => import('@/views/HomeView.vue') },
    { path: '/entrar', name: 'entrar', component: () => import('@/views/LoginView.vue') },
    { path: '/registro', name: 'registro', component: () => import('@/views/RegisterView.vue') },
    { path: '/explorar', name: 'explorar', component: () => import('@/views/ExploreView.vue') },
    { path: '/perfil', name: 'perfil', component: () => import('@/views/ProfileView.vue') },
    { path: '/mensajes', name: 'mensajes', component: () => import('@/views/MessagesView.vue') },
    { path: '/panel', name: 'panel', component: () => import('@/views/DashboardView.vue') },
  ],
  scrollBehavior: () => ({ top: 0 }),
})
