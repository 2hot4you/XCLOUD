<template>
  <div class="commission-view">
    <div class="page-header">
      <h1>返佣管理</h1>
      <div class="header-actions">
        <d-button type="primary">
          <template #icon>
            <span>🧮</span>
          </template>
          计算返佣
        </d-button>
        <d-button>
          <template #icon>
            <span>📥</span>
          </template>
          导出记录
        </d-button>
      </div>
    </div>

    <!-- 统计概览 -->
    <div class="stats-grid">
      <d-card class="stat-card">
        <div class="stat-content">
          <div class="stat-info">
            <h3>¥156,789</h3>
            <p>待处理返佣</p>
          </div>
          <div class="stat-icon">
            <span style="font-size: 20px;">⏱️</span>
          </div>
        </div>
      </d-card>
      <d-card class="stat-card">
        <div class="stat-content">
          <div class="stat-info">
            <h3>¥1,234,567</h3>
            <p>本月已发放</p>
          </div>
          <div class="stat-icon">
            <span style="font-size: 20px;">✅</span>
          </div>
        </div>
      </d-card>
      <d-card class="stat-card">
        <div class="stat-content">
          <div class="stat-info">
            <h3>89</h3>
            <p>返佣记录</p>
          </div>
          <div class="stat-icon">
            <span style="font-size: 20px;">📋</span>
          </div>
        </div>
      </d-card>
    </div>

    <!-- 功能区域 -->
    <d-card class="feature-card">
      <template #header>
        <h3>返佣管理功能</h3>
      </template>
      <div class="feature-grid">
        <div class="feature-item">
          <div class="feature-icon">
            <span style="font-size: 24px;">🧮</span>
          </div>
          <h4>自动计算</h4>
          <p>基于合同规则的自动返佣计算</p>
        </div>
        <div class="feature-item">
          <div class="feature-icon">
            <span style="font-size: 24px;">🔍</span>
          </div>
          <h4>记录查询</h4>
          <p>返佣记录查询和审核</p>
        </div>
        <div class="feature-item">
          <div class="feature-icon">
            <span style="font-size: 24px;">💰</span>
          </div>
          <h4>支付管理</h4>
          <p>返佣支付状态管理</p>
        </div>
        <div class="feature-item">
          <div class="feature-icon">
            <span style="font-size: 24px;">📊</span>
          </div>
          <h4>统计分析</h4>
          <p>返佣报表和统计分析</p>
        </div>
      </div>
      <div class="feature-note">
        <p><small>需要实现返佣计算引擎后才能显示实际功能</small></p>
      </div>
    </d-card>
  </div>
</template>

<script setup lang="ts">
// 返佣管理页面逻辑
</script>

<style scoped>
.commission-view {
  padding: 32px;
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
}

.page-header h1 {
  font-size: 32px;
  font-weight: 700;
  color: #2d3748;
  margin: 0;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.header-actions {
  display: flex;
  gap: 12px;
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 24px;
  margin-bottom: 24px;
}

.stat-card {
  border: none;
  border-radius: 16px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.08);
  transition: all 0.3s ease;
}

.stat-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 8px 40px rgba(0, 0, 0, 0.12);
}

.stat-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 24px;
}

.stat-info h3 {
  font-size: 24px;
  font-weight: 700;
  color: #2d3748;
  margin: 0 0 6px 0;
}

.stat-info p {
  color: #718096;
  margin: 0;
  font-size: 14px;
}

.stat-icon {
  width: 48px;
  height: 48px;
  border-radius: 12px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-size: 20px;
}

.feature-card {
  border: none;
  border-radius: 16px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.08);
}

.feature-card h3 {
  font-size: 18px;
  font-weight: 600;
  color: #2d3748;
  margin: 0;
}

.feature-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
  gap: 24px;
  padding: 24px;
}

.feature-item {
  text-align: center;
  padding: 24px;
  border-radius: 12px;
  background: #f7fafc;
  transition: all 0.3s ease;
}

.feature-item:hover {
  background: #edf2f7;
  transform: translateY(-2px);
}

.feature-icon {
  width: 64px;
  height: 64px;
  margin: 0 auto 16px;
  border-radius: 16px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  font-size: 24px;
}

.feature-item h4 {
  font-size: 16px;
  font-weight: 600;
  color: #2d3748;
  margin: 0 0 8px 0;
}

.feature-item p {
  color: #718096;
  margin: 0;
  font-size: 14px;
}

.feature-note {
  text-align: center;
  padding: 16px 24px;
  border-top: 1px solid #e2e8f0;
  background: #f7fafc;
}

.feature-note small {
  color: #a0aec0;
  font-size: 12px;
}

@media (max-width: 768px) {
  .commission-view {
    padding: 16px;
  }
  
  .page-header {
    flex-direction: column;
    gap: 16px;
    align-items: stretch;
  }
  
  .stats-grid {
    grid-template-columns: 1fr;
  }
  
  .feature-grid {
    grid-template-columns: 1fr;
  }
}
</style>