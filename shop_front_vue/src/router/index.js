import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import ViewPost from '../views/ViewPost.vue'
import NotFound from '../views/errors/NotFound.vue'
import Collect from '../views/Collect.vue'
import CollectInfo from '../views/CollectInfo.vue'
import Seckill from '../views/Seckill.vue'
import Payment from '../views/Payment.vue'
Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path:'/product/:id',
    name:'ProductInfo',
    component:ViewPost
  },
  {
    path: '*',
    name: 'NotFound',
    component: NotFound
  },
  {
    path:'/collect',
    name:'collect',
    component:Collect
  },
  {
    path:'/collect/:id',
    name:'collectInfo',
    component:CollectInfo
  },
  {
    path:'/seckill',
    name:'seckill',
    component:Seckill
  },
  {
    path:'/payment',
    name:'payment',
    component:Payment
  },
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
