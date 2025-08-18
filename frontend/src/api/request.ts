import axios from 'axios'
import type { AxiosInstance, AxiosRequestConfig, AxiosResponse } from 'axios'
import { useAuthStore } from '@/store/auth'

// 创建axios实例
const request: AxiosInstance = axios.create({
  baseURL: '/api',
  timeout: 30000,
  headers: {
    'Content-Type': 'application/json'
  }
})

// 请求拦截器
request.interceptors.request.use(
  (config: AxiosRequestConfig) => {
    const authStore = useAuthStore()
    
    // 添加认证token
    if (authStore.token) {
      config.headers = {
        ...config.headers,
        Authorization: `Bearer ${authStore.token}`
      }
    }

    return config
  },
  (error) => {
    console.error('请求拦截器错误:', error)
    return Promise.reject(error)
  }
)

// 响应拦截器
request.interceptors.response.use(
  (response: AxiosResponse) => {
    return response.data
  },
  async (error) => {
    const authStore = useAuthStore()
    
    if (error.response?.status === 401) {
      // token过期，尝试刷新
      try {
        await authStore.refreshAccessToken()
        // 重新发送原请求
        const originalRequest = error.config
        originalRequest.headers.Authorization = `Bearer ${authStore.token}`
        return request(originalRequest)
      } catch (refreshError) {
        // 刷新失败，跳转到登录页
        authStore.clearAuth()
        window.location.href = '/login'
        return Promise.reject(refreshError)
      }
    }

    // 其他错误处理
    const errorMessage = error.response?.data?.message || error.message || '请求失败'
    console.error('API错误:', errorMessage)
    
    return Promise.reject({
      code: error.response?.status || 500,
      message: errorMessage,
      data: error.response?.data
    })
  }
)

export default request