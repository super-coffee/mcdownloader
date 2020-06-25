import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    redirect: { name: 'Download' }
  },
  {
    path: '/download',
    name: 'Download',
    component: () => import('../views/Download.vue')
  },
  {
    path: '/chooseMirror',
    name: 'ChooseMirror',
    component: () => import('../views/ChooseMirror.vue')
  }
]

const router = new VueRouter({
  routes
})

export default router
