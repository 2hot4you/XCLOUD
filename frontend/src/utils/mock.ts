// Mock数据和API模拟
export const mockAPI = {
  // 模拟登录API
  login: async (data: { username: string; password: string }) => {
    await new Promise(resolve => setTimeout(resolve, 1000)) // 模拟网络延迟
    
    if (data.username === 'admin' && data.password === 'admin123') {
      return {
        code: 200,
        message: '登录成功',
        data: {
          access_token: 'mock_access_token_' + Date.now(),
          refresh_token: 'mock_refresh_token_' + Date.now(),
          expires_in: 86400,
          token_type: 'Bearer'
        }
      }
    } else {
      throw new Error('用户名或密码错误')
    }
  },

  // 模拟用户信息API
  getUserInfo: async () => {
    await new Promise(resolve => setTimeout(resolve, 500))
    
    return {
      code: 200,
      message: '获取成功',
      data: {
        id: 'user-1',
        username: 'admin',
        email: 'admin@xcloud.com',
        role: 'admin',
        is_active: true,
        created_at: '2024-01-01T00:00:00Z'
      }
    }
  },

  // 模拟客户列表API
  getCustomers: async (params: any) => {
    await new Promise(resolve => setTimeout(resolve, 800))
    
    return {
      code: 200,
      message: '获取成功',
      data: {
        customers: [
          {
            id: 'customer-1',
            customer_code: 'CUST001',
            company_name: '北京科技有限公司',
            contact_name: '张三',
            contact_phone: '13800138000',
            contact_email: 'zhangsan@example.com',
            status: 'active',
            level: 1,
            business_type: '互联网',
            credit_limit: 1000000,
            created_at: '2024-01-15T10:30:00Z'
          },
          {
            id: 'customer-2',
            customer_code: 'CUST002',
            company_name: '上海创新公司',
            contact_name: '李四',
            contact_phone: '13900139000',
            contact_email: 'lisi@example.com',
            status: 'active',
            level: 2,
            business_type: '制造业',
            credit_limit: 500000,
            created_at: '2024-01-16T14:20:00Z'
          }
        ],
        total: 2,
        page: 1,
        page_size: 20
      }
    }
  }
}

// 模拟网络请求拦截器
export const enableMockAPI = () => {
  // 可以在这里拦截axios请求并返回mock数据
  console.log('Mock API已启用')
}