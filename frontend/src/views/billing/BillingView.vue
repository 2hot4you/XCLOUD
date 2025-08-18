<template>
  <div class="billing-view">
    <div class="page-header">
      <h1>账单数据</h1>
      <div class="header-actions">
        <d-button type="primary">
          <template #icon>
            <span>📥</span>
          </template>
          导出数据
        </d-button>
        <d-button>
          <template #icon>
            <span>🔄</span>
          </template>
          同步数据
        </d-button>
      </div>
    </div>

    <!-- 统计卡片 -->
    <div class="stats-grid">
      <d-card class="stat-card">
        <div class="stat-content">
          <div class="stat-info">
            <h3>¥2,345,678</h3>
            <p>本月账单总额</p>
          </div>
          <div class="stat-icon">
            <span style="font-size: 20px;">💰</span>
          </div>
        </div>
      </d-card>
      <d-card class="stat-card">
        <div class="stat-content">
          <div class="stat-info">
            <h3>1,234</h3>
            <p>账单记录数</p>
          </div>
          <div class="stat-icon">
            <span style="font-size: 20px;">📋</span>
          </div>
        </div>
      </d-card>
      <d-card class="stat-card">
        <div class="stat-content">
          <div class="stat-info">
            <h3>4</h3>
            <p>云平台数量</p>
          </div>
          <div class="stat-icon">
            <span style="font-size: 20px;">☁️</span>
          </div>
        </div>
      </d-card>
    </div>

    <!-- 筛选区域 -->
    <d-card class="filter-card">
      <div class="filter-form">
        <d-date-picker
          v-model="dateRange"
          type="daterange"
          placeholder="选择日期范围"
          style="width: 240px;"
        />
        <d-select
          v-model="selectedProvider"
          placeholder="选择云平台"
          style="width: 120px;"
        >
          <d-option label="全部" value=""></d-option>
          <d-option label="腾讯云" value="tencent"></d-option>
          <d-option label="阿里云" value="alibaba"></d-option>
          <d-option label="华为云" value="huawei"></d-option>
          <d-option label="AWS" value="aws"></d-option>
        </d-select>
        <d-button type="primary">查询</d-button>
      </div>
    </d-card>

    <!-- 图表区域 -->
    <d-card class="chart-card">
      <template #header>
        <h3>账单趋势</h3>
      </template>
      <div class="chart-placeholder">
        <span style="font-size: 48px;" class="chart-icon">📊</span>
        <p>账单数据图表</p>
        <small>集成后端数据后显示真实图表</small>
      </div>
    </d-card>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const dateRange = ref([])
const selectedProvider = ref('')
</script>

<style scoped>
.billing-view {
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

.filter-card {
  margin-bottom: 24px;
  border: none;
  border-radius: 16px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.08);
}

.filter-form {
  display: flex;
  gap: 16px;
  align-items: center;
  padding: 24px;
}

.chart-card {
  border: none;
  border-radius: 16px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.08);
}

.chart-card h3 {
  font-size: 18px;
  font-weight: 600;
  color: #2d3748;
  margin: 0;
}

.chart-placeholder {
  height: 300px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  color: #a0aec0;
}

.chart-icon {
  font-size: 48px;
  margin-bottom: 16px;
}

.chart-placeholder p {
  margin: 0 0 8px 0;
  font-size: 16px;
}

.chart-placeholder small {
  font-size: 12px;
}

@media (max-width: 768px) {
  .billing-view {
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
  
  .filter-form {
    flex-direction: column;
    align-items: stretch;
  }
}
</style>