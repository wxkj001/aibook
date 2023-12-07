import { createRouter, createWebHistory,createWebHashHistory } from 'vue-router'
import HomeView from '../components/LayoutView.vue'
import { useUserStore } from '../stores/user'

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
      beforeEnter: (to, from) => {
        const user = useUserStore()
        if (user.$state.user?.token==undefined){
          return {name:'login'}
        }
        return true
      },
      children:[
        {
          path: '/home/my',
          name: 'homeMy',
          component: () => import('../views/my/MyView.vue')
        },
        {
          path: '/home/setting',
          name: 'homeSetting',
          component: () => import('../views/setting/SettingView.vue')
        },
        {
          path: '/home/chat',
          name: 'homeChat',
          component: () => import('../views/chat/ChatView.vue')
        },
        {
          path: '/home/color',
          name: 'homeColor',
          component: () => import('../views/color/ColorView.vue')
        }
      ]
    },
    {
      path: '/login',
      name: 'login',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/LoginView.vue')
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/AboutView.vue')
    }
  ]
})

export default router
