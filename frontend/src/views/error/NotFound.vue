<template>
  <div class="not-found">
    <div class="error-content">
      <div class="error-icon">
        <span style="font-size: 80px;">📋</span>
      </div>
      <h1>404</h1>
      <h2>页面未找到</h2>
      <p>抱歉，您访问的页面不存在或已被移动。</p>
      <div class="actions">
        <d-button @click="handleGoBack">返回上页</d-button>
        <d-button type="primary" @click="handleGoHome">回到首页</d-button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'

const router = useRouter()

const handleGoBack = () => {
  router.go(-1)
}

const handleGoHome = () => {
  router.push('/dashboard')
}
</script>

<style scoped>
.not-found {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background: #f5f5f5;
}

.error-content {
  text-align: center;
  padding: 40px;
}

.error-icon {
  font-size: 80px;
  color: #ddd;
  margin-bottom: 20px;
}

.error-content h1 {
  font-size: 72px;
  font-weight: 600;
  color: #333;
  margin: 0 0 16px 0;
}

.error-content h2 {
  font-size: 24px;
  font-weight: 500;
  color: #666;
  margin: 0 0 16px 0;
}

.error-content p {
  font-size: 16px;
  color: #999;
  margin: 0 0 32px 0;
}

.actions {
  display: flex;
  justify-content: center;
  gap: 16px;
}

@media (max-width: 768px) {
  .error-content h1 {
    font-size: 48px;
  }
  
  .error-content h2 {
    font-size: 20px;
  }
  
  .actions {
    flex-direction: column;
    align-items: center;
  }
  
  .actions .devui-button {
    min-width: 120px;
  }
}
</style>