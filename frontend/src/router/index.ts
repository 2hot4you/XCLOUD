import { createRouter, createWebHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/auth/LoginView.vue'),
    meta: {
      title: '登录',
      requiresAuth: false
    }
  },
  {
    path: '/',
    name: 'Layout',
    component: () => import('@/views/layout/MainLayout.vue'),
    redirect: '/dashboard',
    meta: {
      requiresAuth: true
    },
    children: [
      {
        path: '/dashboard',
        name: 'Dashboard',
        component: () => import('@/views/dashboard/DashboardView.vue'),
        meta: {
          title: '控制台',
          icon: 'icon-home'
        }
      },
      {
        path: '/customers',
        name: 'Customers',
        component: () => import('@/views/customer/CustomerList.vue'),
        meta: {
          title: '客户管理',
          icon: 'icon-user'
        }
      },
      {
        path: '/customers/create',
        name: 'CreateCustomer',
        component: () => import('@/views/customer/CustomerForm.vue'),
        meta: {
          title: '新建客户',
          parentName: 'Customers'
        }
      },
      {
        path: '/customers/:id/edit',
        name: 'EditCustomer',
        component: () => import('@/views/customer/CustomerForm.vue'),
        meta: {
          title: '编辑客户',
          parentName: 'Customers'
        }
      },
      {
        path: '/contracts',
        name: 'Contracts',
        component: () => import('@/views/contract/ContractList.vue'),
        meta: {
          title: '合同管理',
          icon: 'icon-file'
        }
      },
      {
        path: '/contracts/create',
        name: 'CreateContract',
        component: () => import('@/views/contract/ContractForm.vue'),
        meta: {
          title: '新建合同',
          parentName: 'Contracts'
        }
      },
      {
        path: '/contracts/:id',
        name: 'ContractDetail',
        component: () => import('@/views/contract/ContractForm.vue'),
        meta: {
          title: '合同详情',
          parentName: 'Contracts'
        }
      },
      {
        path: '/contracts/:id/edit',
        name: 'EditContract',
        component: () => import('@/views/contract/ContractForm.vue'),
        meta: {
          title: '编辑合同',
          parentName: 'Contracts'
        }
      },
      {
        path: '/billing',
        name: 'Billing',
        component: () => import('@/views/billing/BillingView.vue'),
        meta: {
          title: '账单数据',
          icon: 'icon-list'
        }
      },
      {
        path: '/commission',
        name: 'Commission',
        component: () => import('@/views/commission/CommissionView.vue'),
        meta: {
          title: '返佣管理',
          icon: 'icon-money'
        }
      },
      {
        path: '/reports',
        name: 'Reports',
        component: () => import('@/views/reports/ReportsView.vue'),
        meta: {
          title: '数据报表',
          icon: 'icon-chart'
        }
      },
      {
        path: '/settings',
        name: 'Settings',
        component: () => import('@/views/settings/SettingsView.vue'),
        meta: {
          title: '系统设置',
          icon: 'icon-setting'
        }
      }
    ]
  },
  {
    path: '/:pathMatch(.*)*',
    name: 'NotFound',
    component: () => import('@/views/error/NotFound.vue'),
    meta: {
      title: '页面不存在'
    }
  }
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes
})

// 路由守卫
router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('access_token')
  
  // 设置页面标题
  if (to.meta.title) {
    document.title = `${to.meta.title} - XCloud多云对账平台`
  }
  
  // 检查是否需要登录
  if (to.meta.requiresAuth !== false && !token) {
    next('/login')
  } else if (to.path === '/login' && token) {
    next('/dashboard')
  } else {
    next()
  }
})

export default router