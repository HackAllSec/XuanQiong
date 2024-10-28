// router/index.ts 文件

import { createRouter, createWebHashHistory, RouterOptions, Router, RouteRecordRaw } from 'vue-router'
import { jwtDecode } from 'jwt-decode'
//由于router的API默认使用了类型进行初始化，内部包含类型定义，所以本文内部代码中的所有数据类型是可以省略的
//RouterRecordRaw是路由组件对象
const routes: RouteRecordRaw[] = [
 { path: '/', name: 'Home', component: () => import('../views/Index.vue') },
 { path: '/login', name: 'Login', component: () => import('../views/Login.vue') },
 { path: '/submit', name: 'Submit', component: () => import('../views/Submit.vue') },
 { path: '/ranklist', name: 'Ranklist', component: () => import('../views/Ranklist.vue') },
 { path: '/search', name: 'Search', component: () => import('../views/Search.vue') },
 { path: '/register', name: 'Register', component: () => import('../views/Register.vue') },
 { path: '/forgotpwd', name: 'Forgotpwd', component: () => import('../views/Forgotpwd.vue') },
 { path: '/vulnlist', name: 'Vulnlist', component: () => import('../views/Vulnlist.vue') },
 { path: '/profile', name: 'Profile', component: () => import('../views/Profile.vue') },
 { path: '/modifypwd', name: 'Modifypwd', component: () => import('../views/Modifypasswd.vue') },
 { path: '/myvulns', name: 'Myvulns', component: () => import('../views/Myvulns.vue') },
]

// RouterOptions是路由选项类型
const options: RouterOptions = {
 history: createWebHashHistory(),
 routes,
}

// Router是路由对象类型
const router: Router = createRouter(options)
router.beforeEach((to, from, next) => {
    // 每次路由变化时调用
    to.redirectedFrom = from
    next();
  });
export default router