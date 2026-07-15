// router/index.ts 文件

import { createRouter, createWebHashHistory, RouterOptions, Router, RouteRecordRaw } from 'vue-router'
import { jwtDecode } from 'jwt-decode'
import { canAccessAdminPanel, clearAuthSession } from '../auth'
//由于router的API默认使用了类型进行初始化，内部包含类型定义，所以本文内部代码中的所有数据类型是可以省略的
//RouterRecordRaw是路由组件对象
const routes: RouteRecordRaw[] = [
 { path: '/', name: 'Home', component: () => import('../views/Index.vue') },
 { path: '/login', name: 'Login', component: () => import('../pages/Login.vue') },
 { path: '/forgotpwd', name: 'Forgotpwd', component: () => import('../pages/Forgotpwd.vue') },
]

// RouterOptions是路由选项类型
const options: RouterOptions = {
 history: createWebHashHistory(),
 routes,
}

// Router是路由对象类型
const router: Router = createRouter(options)
router.beforeEach((to, from, next) => {
    // 每次路由变化时调用 performAction
    to.redirectedFrom = from
    checkToken();
    if (!sessionStorage.getItem('token')) {
      sessionStorage.removeItem('force_password_change')
    }
    const forcePasswordChange = sessionStorage.getItem('force_password_change') === '1'
    if (forcePasswordChange && to.path !== '/' && to.path !== '/login' && to.path !== '/forgotpwd') {
      next('/')
      return
    }
    if (to.path !== '/login' && to.path !== '/forgotpwd' && !canAccessAdminPanel()) {
      clearAuthSession()
      next('/login')
      return
    }
    next();
  });
  
  function checkToken() {
    const token = sessionStorage.getItem('token')
    if (token) {
        try {
            const decodedToken = jwtDecode(token)
            let currentTime = Math.floor(Date.now() / 1000)
            if (currentTime > decodedToken.exp) {
                clearAuthSession()
                location.reload();
                return;
            }
        } catch (error) {
            clearAuthSession()
            location.reload();
            return;
        }
    }
  }
export default router
