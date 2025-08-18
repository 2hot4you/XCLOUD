<template>
  <div class="main-layout">
    <!-- 顶部导航栏 -->
    <header class="top-navbar">
      <div class="navbar-container">
        <!-- Logo区域 -->
        <div class="navbar-brand" @click="goHome">
          <div class="logo-icon">
            <span>🌍</span>
          </div>
          <h1 class="brand-title">XCloud</h1>
          <span class="brand-subtitle">多云对账平台</span>
        </div>

        <!-- 导航菜单 -->
        <nav class="navbar-nav">
          <div 
            v-for="item in menuItems" 
            :key="item.name"
            class="nav-item"
            :class="{ active: activeMenu === item.name }"
            @click="handleMenuSelect(item)"
          >
            <div class="nav-link">
              <i :class="item.icon" class="nav-icon"></i>
              <span class="nav-text">{{ item.title }}</span>
            </div>
            <div class="nav-indicator"></div>
          </div>
        </nav>

        <!-- 右侧工具栏 -->
        <div class="navbar-actions">
          <!-- 通知 -->
          <div class="action-item" @click="showNotifications">
            <span>🔔</span>
            <span class="notification-badge">3</span>
          </div>
          
          <!-- 用户菜单 -->
          <d-dropdown placement="bottom-end">
            <div class="user-profile">
              <div class="user-avatar">
                <span>👤</span>
              </div>
              <div class="user-info">
                <span class="user-name">{{ userInfo?.username || '管理员' }}</span>
                <span class="user-role">系统管理员</span>
              </div>
              <span>🔽</span>
            </div>
            
            <template #menu>
              <d-dropdown-menu>
                <d-dropdown-item @click="viewProfile">
                  <span>👤</span>
                  个人资料
                </d-dropdown-item>
                <d-dropdown-item @click="viewSettings">
                  <span>⚙️</span>
                  系统设置
                </d-dropdown-item>
                <d-dropdown-item divider></d-dropdown-item>
                <d-dropdown-item @click="handleLogout">
                  <span>🚪</span>
                  退出登录
                </d-dropdown-item>
              </d-dropdown-menu>
            </template>
          </d-dropdown>
        </div>
      </div>
    </header>

    <!-- 页面内容 -->
    <main class="page-content">
      <div class="content-wrapper">
        <!-- 路由内容 -->
        <div class="page-container">
          <Transition name="page" mode="out-in">
            <router-view />
          </Transition>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, watch } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useAuthStore } from '@/store/auth'

const router = useRouter()
const route = useRoute()
const authStore = useAuthStore()

const activeMenu = ref('')

const userInfo = computed(() => authStore.userInfo)

const menuItems = [
  { name: 'Dashboard', title: '控制台', icon: 'icon-home', path: '/dashboard' },
  { name: 'Customers', title: '客户管理', icon: 'icon-users', path: '/customers' },
  { name: 'Contracts', title: '合同管理', icon: 'icon-file-text', path: '/contracts' },
  { name: 'Billing', title: '账单数据', icon: 'icon-list', path: '/billing' },
  { name: 'Commission', title: '返佣管理', icon: 'icon-credit-card', path: '/commission' },
  { name: 'Reports', title: '数据报表', icon: 'icon-bar-chart', path: '/reports' },
  { name: 'Settings', title: '系统设置', icon: 'icon-settings', path: '/settings' }
]


const goHome = () => {
  router.push('/dashboard')
}

const handleMenuSelect = (menuItem: any) => {
  if (menuItem?.path) {
    router.push(menuItem.path)
  }
}

const showNotifications = () => {
  console.log('显示通知')
  // TODO: 实现通知功能
}

const viewProfile = () => {
  console.log('查看个人资料')
  // TODO: 实现个人资料功能
}

const viewSettings = () => {
  router.push('/settings')
}

const handleLogout = async () => {
  await authStore.logout()
  router.push('/login')
}

// 监听路由变化，更新活跃菜单
watch(
  () => route.path,
  (newPath) => {
    const currentMenu = menuItems.find(item => 
      newPath.startsWith(item.path)
    )
    if (currentMenu) {
      activeMenu.value = currentMenu.name
    }
  },
  { immediate: true }
)

onMounted(() => {
  // 初始化认证状态
  authStore.initFromStorage()
})
</script>

<style scoped>
/* 主布局 */
.main-layout {
  min-height: 100vh;
  background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
}

/* 顶部导航栏 */
.top-navbar {
  position: sticky;
  top: 0;
  z-index: 1000;
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(20px);
  border-bottom: 1px solid rgba(0, 0, 0, 0.08);
  box-shadow: 0 2px 20px rgba(0, 0, 0, 0.1);
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.navbar-container {
  max-width: 1400px;
  margin: 0 auto;
  display: flex;
  align-items: center;
  padding: 0 24px;
  height: 72px;
}

/* Logo品牌区域 */
.navbar-brand {
  display: flex;
  align-items: center;
  gap: 12px;
  cursor: pointer;
  margin-right: 48px;
  transition: transform 0.2s ease;
}

.navbar-brand:hover {
  transform: translateY(-1px);
}

.logo-icon {
  width: 40px;
  height: 40px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-size: 18px;
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.4);
}

.brand-title {
  font-size: 24px;
  font-weight: 700;
  margin: 0;
  color: #2d3748;
  letter-spacing: -0.5px;
}

.brand-subtitle {
  font-size: 12px;
  color: #718096;
  font-weight: 500;
  padding: 2px 8px;
  background: rgba(113, 128, 150, 0.1);
  border-radius: 12px;
}

/* 导航菜单 */
.navbar-nav {
  display: flex;
  align-items: center;
  gap: 8px;
  flex: 1;
}

.nav-item {
  position: relative;
  cursor: pointer;
  border-radius: 12px;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  overflow: hidden;
}

.nav-item:hover {
  background: rgba(102, 126, 234, 0.08);
  transform: translateY(-1px);
}

.nav-item.active {
  background: rgba(102, 126, 234, 0.15);
}

.nav-item.active .nav-indicator {
  transform: scaleX(1);
}

.nav-link {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 12px 16px;
  color: #4a5568;
  font-weight: 500;
  font-size: 14px;
  transition: color 0.2s ease;
}

.nav-item:hover .nav-link,
.nav-item.active .nav-link {
  color: #667eea;
}

.nav-icon {
  font-size: 16px;
  transition: transform 0.2s ease;
}

.nav-item:hover .nav-icon {
  transform: scale(1.1);
}

.nav-indicator {
  position: absolute;
  bottom: 0;
  left: 50%;
  transform: translateX(-50%) scaleX(0);
  width: 24px;
  height: 3px;
  background: linear-gradient(90deg, #667eea, #764ba2);
  border-radius: 2px 2px 0 0;
  transition: transform 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

/* 右侧操作区 */
.navbar-actions {
  display: flex;
  align-items: center;
  gap: 16px;
}

.action-item {
  position: relative;
  width: 40px;
  height: 40px;
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  color: #718096;
  transition: all 0.2s ease;
  background: rgba(113, 128, 150, 0.1);
}

.action-item:hover {
  background: rgba(102, 126, 234, 0.15);
  color: #667eea;
  transform: translateY(-1px);
}

.notification-badge {
  position: absolute;
  top: -2px;
  right: -2px;
  width: 18px;
  height: 18px;
  background: #f56565;
  color: white;
  border-radius: 50%;
  font-size: 10px;
  font-weight: 600;
  display: flex;
  align-items: center;
  justify-content: center;
  animation: pulse 2s infinite;
}

@keyframes pulse {
  0% { box-shadow: 0 0 0 0 rgba(245, 101, 101, 0.7); }
  70% { box-shadow: 0 0 0 6px rgba(245, 101, 101, 0); }
  100% { box-shadow: 0 0 0 0 rgba(245, 101, 101, 0); }
}

/* 用户资料 */
.user-profile {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 8px 16px;
  border-radius: 12px;
  cursor: pointer;
  transition: all 0.2s ease;
  background: rgba(113, 128, 150, 0.05);
  border: 1px solid rgba(113, 128, 150, 0.1);
}

.user-profile:hover {
  background: rgba(102, 126, 234, 0.1);
  border-color: rgba(102, 126, 234, 0.2);
  transform: translateY(-1px);
}

.user-avatar {
  width: 32px;
  height: 32px;
  border-radius: 8px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-size: 14px;
}

.user-info {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.user-name {
  font-size: 14px;
  font-weight: 600;
  color: #2d3748;
  line-height: 1.2;
}

.user-role {
  font-size: 12px;
  color: #718096;
  line-height: 1.2;
}

/* 页面内容 */
.page-content {
  flex: 1;
  padding: 24px;
}

.content-wrapper {
  max-width: 1400px;
  margin: 0 auto;
}


.page-container {
  background: white;
  border-radius: 16px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.08);
  border: 1px solid rgba(0, 0, 0, 0.05);
  overflow: hidden;
}

/* 页面过渡动画 */
.page-enter-active,
.page-leave-active {
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.page-enter-from {
  opacity: 0;
  transform: translateX(20px);
}

.page-leave-to {
  opacity: 0;
  transform: translateX(-20px);
}

/* 响应式设计 */
@media (max-width: 1024px) {
  .navbar-container {
    padding: 0 16px;
  }
  
  .navbar-brand {
    margin-right: 24px;
  }
  
  .brand-subtitle {
    display: none;
  }
}

@media (max-width: 768px) {
  .nav-text {
    display: none;
  }
  
  .nav-item {
    padding: 8px;
  }
  
  .user-info {
    display: none;
  }
  
  .page-content {
    padding: 16px;
  }
  
  .navbar-nav {
    gap: 4px;
  }
}

/* 下拉菜单样式增强 */
:deep(.devui-dropdown-menu) {
  margin-top: 8px;
  border-radius: 12px;
  border: none;
  box-shadow: 0 10px 40px rgba(0, 0, 0, 0.15);
  overflow: hidden;
}

:deep(.devui-dropdown-item) {
  padding: 12px 16px;
  transition: all 0.2s ease;
}

:deep(.devui-dropdown-item:hover) {
  background: rgba(102, 126, 234, 0.1);
  color: #667eea;
}

:deep(.devui-dropdown-item i) {
  margin-right: 8px;
  font-size: 14px;
}
</style>