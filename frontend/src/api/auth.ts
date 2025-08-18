import request from './request'
import type { LoginRequest, RefreshRequest, LoginResponse, BaseResponse } from './types/auth'
import { mockAPI } from '@/utils/mock'

// 是否使用Mock数据（开发环境下可以启用）
const useMock = false

export const authAPI = {
  // 用户登录
  login: (data: LoginRequest): Promise<LoginResponse> => {
    if (useMock) {
      return mockAPI.login(data)
    }
    return request.post('/v1/auth/login', data)
  },

  // 刷新令牌
  refresh: (data: RefreshRequest): Promise<LoginResponse> => {
    if (useMock) {
      return mockAPI.login({ username: 'admin', password: 'admin123' })
    }
    return request.post('/v1/auth/refresh', data)
  },

  // 用户登出
  logout: (): Promise<BaseResponse> => {
    if (useMock) {
      return Promise.resolve({ code: 200, message: '登出成功' })
    }
    return request.post('/v1/auth/logout')
  }
}