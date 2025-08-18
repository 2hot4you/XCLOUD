<template>
  <div class="login-container">
    <div class="login-box">
      <div class="login-header">
        <h1>XCloud多云对账平台</h1>
        <p>请输入您的登录凭据</p>
      </div>
      
      <d-form
        ref="loginFormRef"
        :data="loginForm"
        :rules="rules"
        label-size="md"
        label-align="start"
        @submit="handleLogin"
      >
        <d-form-item field="username" label="用户名">
          <d-input
            v-model="loginForm.username"
            placeholder="请输入用户名"
            :disabled="loading"
            size="md"
          />
        </d-form-item>

        <d-form-item field="password" label="密码">
          <d-input
            v-model="loginForm.password"
            type="password"
            placeholder="请输入密码"
            :disabled="loading"
            size="md"
            @keyup.enter="handleLogin"
          />
        </d-form-item>

        <d-form-item>
          <d-button
            type="primary"
            :loading="loading"
            block
            size="md"
            @click="handleLogin"
          >
            登录
          </d-button>
        </d-form-item>
      </d-form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/store/auth'

const router = useRouter()
const authStore = useAuthStore()

const loginFormRef = ref()
const loading = ref(false)

const loginForm = reactive({
  username: '',
  password: ''
})

const rules = {
  username: [
    { required: true, message: '请输入用户名' }
  ],
  password: [
    { required: true, message: '请输入密码' },
    { minLength: 6, message: '密码长度至少6位' }
  ]
}

const handleLogin = async () => {
  try {
    const isValid = await loginFormRef.value.validate()
    if (!isValid) return

    loading.value = true
    
    const success = await authStore.login(loginForm)
    if (success) {
      router.push('/dashboard')
    }
  } catch (error: any) {
    console.error('登录失败:', error)
    // TODO: 显示错误消息
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  padding: 20px;
}

.login-box {
  width: 100%;
  max-width: 400px;
  padding: 40px;
  background: white;
  border-radius: 8px;
  box-shadow: 0 10px 25px rgba(0, 0, 0, 0.1);
}

.login-header {
  text-align: center;
  margin-bottom: 30px;
}

.login-header h1 {
  font-size: 24px;
  font-weight: 600;
  color: #333;
  margin-bottom: 8px;
}

.login-header p {
  color: #666;
  font-size: 14px;
}

:deep(.devui-form-item) {
  margin-bottom: 20px;
}

:deep(.devui-button) {
  height: 44px;
  font-size: 16px;
}
</style>