import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../views/Login'
import Admin from '../views/Admin'
import AddActive from '../components/Active/AddActive'
import ActiveList from '../components/Active/ActiveList'
import Users from '../components/Users/Users'
import ProductList from '../components/Product/ProductList'
import AddProduct from '../components/Product/AddProduct'

Vue.use(VueRouter)

const routes = [
  {
    path:'/',
    redirect:'/login',
  },
  {
    path: '/login',
    name: 'login',
    component: Login,
    meta:{
      title:'登录'
    }
  },
  {
    path:'/admin',
    name:'admin',
    component:Admin,
    children:[
      {
        path:'/addActive',
        component:AddActive,
        meta:{
          title:'活动管理'
        }
      },
      {
        path:'/putActive/:id',
        component:AddActive,
        props:true,
        meta:{
          title:'活动管理'
        }
      },
      {
        path:'/activeList',
        component:ActiveList,
        meta:{
          title:'活动列表'
        }
      },
      {
        path:'/productList',
        component:ProductList,
        meta:{
          title:'商品列表'
        }
      },
      {
        path:'/addProduct/',
        component:AddProduct,
        props:true,
        meta:{
          title:'商品管理'
        }
      },
      {
        path:'/putProduct/:id',
        component:AddProduct,
        props:true,
        meta:{
          title:'商品管理'
        }
      },
      {
        path:'/users',
        component:Users,
        meta:{
          title:'用户列表'
        }
      },
    ]
  }
]

const router = new VueRouter({
  mode:'history',
  routes
})

const originalPush = VueRouter.prototype.push

VueRouter.prototype.push = function push (location) {

return originalPush.call(this, location).catch(err => err)

}


router.beforeEach((to, from, next) => {

  if (to.meta.title) {
    document.title = to.meta.title
  }
  next()

  const userToken = window.sessionStorage.getItem('token')
  if (to.path === '/login') return next()
  if (!userToken) {
    next('/login')
  } else {
    next()
  }
})



export default router
