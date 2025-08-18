<template>
  <div class="customer-list">
    <h1 class="page-title">客户管理</h1>
    
    <d-card>
      <div class="placeholder-content">
        <div class="placeholder-icon">
          <span style="font-size: 64px;">👥</span>
        </div>
        <h2>客户管理功能</h2>
        <p>此页面将提供完整的客户管理功能，包括：</p>
        <ul>
          <li>客户信息的新增、编辑、查看和删除</li>
          <li>客户层级关系管理</li>
          <li>客户状态和权限管理</li>
          <li>客户业务数据统计</li>
          <li>客户合同和账单查询</li>
        </ul>
        <p><small>需要完善客户管理模块后才能显示实际功能</small></p>
      </div>
    </d-card>
  </div>
</template>

<script setup lang="ts">
// 客户管理页面逻辑
</script>

<style scoped>
.customer-list {
  max-width: 1400px;
  margin: 0 auto;
  padding: 32px;
}

.page-title {
  font-size: 32px;
  font-weight: 700;
  color: #2d3748;
  margin: 0 0 24px 0;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.placeholder-content {
  text-align: center;
  padding: 60px 40px;
}

.placeholder-icon {
  font-size: 64px;
  color: #ddd;
  margin-bottom: 24px;
}

.placeholder-content h2 {
  font-size: 24px;
  font-weight: 600;
  color: #333;
  margin: 0 0 16px 0;
}

.placeholder-content p {
  color: #666;
  margin: 16px 0;
  font-size: 16px;
}

.placeholder-content ul {
  text-align: left;
  max-width: 400px;
  margin: 24px auto;
  color: #666;
}

.placeholder-content li {
  margin: 8px 0;
}

.placeholder-content small {
  color: #999;
  font-size: 14px;
}

@media (max-width: 768px) {
  .customer-list {
    padding: 16px;
  }
}
</style>