import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { LoginRequest, TokenData, UserInfo } from '@/api/types/auth'
import { authAPI } from '@/api/auth'

export const useAuthStore = defineStore('auth', () => {
  // 状态
  const token = ref<string>('')
  const refreshToken = ref<string>('')
  const userInfo = ref<UserInfo | null>(null)
  const isLoading = ref<boolean>(false)

  // 计算属性
  const isAuthenticated = computed(() => !!token.value)
  const userRole = computed(() => userInfo.value?.role || '')

  // 从localStorage恢复状态
  const initFromStorage = () => {
    const storedToken = localStorage.getItem('access_token')
    const storedRefreshToken = localStorage.getItem('refresh_token')
    const storedUserInfo = localStorage.getItem('user_info')

    if (storedToken) token.value = storedToken
    if (storedRefreshToken) refreshToken.value = storedRefreshToken
    if (storedUserInfo) {
      try {
        userInfo.value = JSON.parse(storedUserInfo)
      } catch (error) {
        console.error('解析用户信息失败:', error)
      }
    }
  }

  // 保存token到localStorage
  const saveTokens = (tokenData: TokenData) => {
    token.value = tokenData.access_token
    refreshToken.value = tokenData.refresh_token
    
    localStorage.setItem('access_token', tokenData.access_token)
    localStorage.setItem('refresh_token', tokenData.refresh_token)
  }

  // 保存用户信息
  const saveUserInfo = (info: UserInfo) => {
    userInfo.value = info
    localStorage.setItem('user_info', JSON.stringify(info))
  }

  // 清除所有认证信息
  const clearAuth = () => {
    token.value = ''
    refreshToken.value = ''
    userInfo.value = null
    
    localStorage.removeItem('access_token')
    localStorage.removeItem('refresh_token')
    localStorage.removeItem('user_info')
  }

  // 登录
  const login = async (loginData: LoginRequest) => {
    try {
      isLoading.value = true
      const response = await authAPI.login(loginData)
      
      if (response.code === 200) {
        saveTokens(response.data)
        // TODO: 获取用户信息
        // const userResponse = await authAPI.getUserInfo()
        // saveUserInfo(userResponse.data)
        return true
      }
      
      return false
    } catch (error) {
      console.error('登录失败:', error)
      throw error
    } finally {
      isLoading.value = false
    }
  }

  // 刷新token
  const refreshAccessToken = async () => {
    try {
      if (!refreshToken.value) {
        throw new Error('没有刷新令牌')
      }

      const response = await authAPI.refresh({
        refresh_token: refreshToken.value
      })

      if (response.code === 200) {
        saveTokens(response.data)
        return true
      }

      return false
    } catch (error) {
      console.error('刷新令牌失败:', error)
      clearAuth()
      throw error
    }
  }

  // 登出
  const logout = async () => {
    try {
      if (token.value) {
        await authAPI.logout()
      }
    } catch (error) {
      console.error('登出失败:', error)
    } finally {
      clearAuth()
    }
  }

  // 检查权限
  const hasPermission = (permission: string) => {
    // TODO: 实现权限检查逻辑
    return true
  }

  // 检查角色
  const hasRole = (role: string) => {
    return userRole.value === role
  }

  return {
    // 状态
    token,
    refreshToken,
    userInfo,
    isLoading,
    // 计算属性
    isAuthenticated,
    userRole,
    // 方法
    initFromStorage,
    login,
    logout,
    refreshAccessToken,
    hasPermission,
    hasRole,
    clearAuth
  }
})