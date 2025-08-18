// 认证相关类型定义

// 登录请求
export interface LoginRequest {
  username: string
  password: string
}

// 刷新令牌请求
export interface RefreshRequest {
  refresh_token: string
}

// 令牌数据
export interface TokenData {
  access_token: string
  refresh_token: string
  expires_in: number
  token_type: string
}

// 用户信息
export interface UserInfo {
  id: string
  username: string
  email: string
  role: 'admin' | 'operator' | 'viewer'
  is_active: boolean
  last_login_at?: string
  created_at: string
}

// 基础响应
export interface BaseResponse {
  code: number
  message: string
}

// 登录响应
export interface LoginResponse extends BaseResponse {
  data: TokenData
}

// 用户信息响应
export interface UserInfoResponse extends BaseResponse {
  data: UserInfo
}