<template>
  <div class="dashboard">
    <h1 class="page-title">控制台</h1>
    
    <!-- 统计卡片 -->
    <div class="stats-grid">
      <d-card class="stat-card" v-for="stat in stats" :key="stat.title">
        <div class="stat-content">
          <div class="stat-info">
            <h3>{{ stat.value }}</h3>
            <p>{{ stat.title }}</p>
          </div>
          <div class="stat-icon">
            <span>{{ stat.icon }}</span>
          </div>
        </div>
        <div class="stat-trend" :class="stat.trend">
          <span>{{ stat.change }}</span>
        </div>
      </d-card>
    </div>

    <!-- 图表区域 -->
    <div class="charts-section">
      <div class="chart-row">
        <d-card class="chart-card">
          <template #header>
            <h3>收入趋势</h3>
          </template>
          <v-chart 
            class="chart" 
            :option="revenueChartOption"
            :style="{ height: '300px' }"
          />
        </d-card>
        
        <d-card class="chart-card">
          <template #header>
            <h3>云平台分布</h3>
          </template>
          <v-chart 
            class="chart" 
            :option="platformChartOption"
            :style="{ height: '300px' }"
          />
        </d-card>
      </div>
    </div>

    <!-- 最近活动 -->
    <div class="recent-section">
      <d-card>
        <template #header>
          <h3>最近活动</h3>
        </template>
        
        <d-table
          :data="recentActivities"
          :columns="activityColumns"
          size="md"
        />
      </d-card>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { use } from 'echarts/core'
import { CanvasRenderer } from 'echarts/renderers'
import { LineChart, PieChart } from 'echarts/charts'
import {
  TitleComponent,
  TooltipComponent,
  LegendComponent,
  GridComponent
} from 'echarts/components'
import VChart from 'vue-echarts'

// 注册必需的组件
use([
  CanvasRenderer,
  LineChart,
  PieChart,
  TitleComponent,
  TooltipComponent,
  LegendComponent,
  GridComponent
])

interface StatCard {
  title: string
  value: string
  change: string
  trend: 'up' | 'down'
  icon: string
}

interface Activity {
  time: string
  action: string
  target: string
  status: string
}

const stats = ref<StatCard[]>([
  {
    title: '总客户数',
    value: '156',
    change: '+12%',
    trend: 'up',
    icon: '👥'
  },
  {
    title: '活跃合同',
    value: '89',
    change: '+8%',
    trend: 'up',
    icon: '📋'
  },
  {
    title: '本月收入',
    value: '¥2,345,678',
    change: '+15%',
    trend: 'up',
    icon: '💰'
  },
  {
    title: '待处理返佣',
    value: '¥156,789',
    change: '-5%',
    trend: 'down',
    icon: '⏱️'
  }
])

const recentActivities = ref<Activity[]>([
  {
    time: '2024-01-15 10:30',
    action: '新建客户',
    target: '北京科技有限公司',
    status: '已完成'
  },
  {
    time: '2024-01-15 09:15',
    action: '合同签署',
    target: '合同编号: CON202401001',
    status: '已完成'
  },
  {
    time: '2024-01-15 08:45',
    action: '返佣计算',
    target: '2024年1月账期',
    status: '进行中'
  },
  {
    time: '2024-01-14 16:20',
    action: '数据同步',
    target: '腾讯云账单数据',
    status: '已完成'
  }
])

const activityColumns = [
  {
    field: 'time',
    header: '时间',
    width: '180px'
  },
  {
    field: 'action',
    header: '操作',
    width: '120px'
  },
  {
    field: 'target',
    header: '对象',
    minWidth: '200px'
  },
  {
    field: 'status',
    header: '状态',
    width: '100px',
    cellClass: (rowData: Activity) => {
      return rowData.status === '已完成' ? 'status-success' : 
             rowData.status === '进行中' ? 'status-warning' : 'status-error'
    }
  }
]

// 收入趋势图配置
const revenueChartOption = computed(() => ({
  tooltip: {
    trigger: 'axis'
  },
  grid: {
    left: '3%',
    right: '4%',
    bottom: '3%',
    containLabel: true
  },
  xAxis: {
    type: 'category',
    data: ['1月', '2月', '3月', '4月', '5月', '6月', '7月', '8月', '9月', '10月', '11月', '12月'],
    axisLine: {
      lineStyle: {
        color: '#E5E7EB'
      }
    }
  },
  yAxis: {
    type: 'value',
    axisLine: {
      lineStyle: {
        color: '#E5E7EB'
      }
    }
  },
  series: [
    {
      name: '收入',
      type: 'line',
      smooth: true,
      data: [1200, 1320, 1010, 1340, 1890, 2300, 2100, 2400, 2180, 2650, 2890, 3100],
      lineStyle: {
        color: '#1890ff',
        width: 3
      },
      itemStyle: {
        color: '#1890ff'
      },
      areaStyle: {
        color: {
          type: 'linear',
          x: 0,
          y: 0,
          x2: 0,
          y2: 1,
          colorStops: [
            { offset: 0, color: 'rgba(24, 144, 255, 0.3)' },
            { offset: 1, color: 'rgba(24, 144, 255, 0.05)' }
          ]
        }
      }
    }
  ]
}))

// 云平台分布图配置
const platformChartOption = computed(() => ({
  tooltip: {
    trigger: 'item',
    formatter: '{a} <br/>{b}: {c} ({d}%)'
  },
  legend: {
    bottom: '10',
    left: 'center'
  },
  series: [
    {
      name: '云平台使用量',
      type: 'pie',
      radius: ['30%', '70%'],
      center: ['50%', '45%'],
      avoidLabelOverlap: false,
      label: {
        show: false,
        position: 'center'
      },
      emphasis: {
        label: {
          show: true,
          fontSize: '16',
          fontWeight: 'bold'
        }
      },
      labelLine: {
        show: false
      },
      data: [
        { value: 1048, name: '腾讯云', itemStyle: { color: '#1890ff' } },
        { value: 735, name: '阿里云', itemStyle: { color: '#52c41a' } },
        { value: 580, name: '华为云', itemStyle: { color: '#faad14' } },
        { value: 484, name: 'AWS', itemStyle: { color: '#f5222d' } }
      ]
    }
  ]
}))

onMounted(() => {
  // TODO: 从API获取实际数据
  console.log('Dashboard mounted')
})
</script>

<style scoped>
.dashboard {
  padding: 32px;
}

.page-title {
  font-size: 32px;
  font-weight: 700;
  color: #2d3748;
  margin-bottom: 8px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 24px;
  margin-bottom: 32px;
}

.stat-card {
  border: none;
  border-radius: 16px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.08);
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  overflow: hidden;
  background: white;
}

.stat-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 8px 40px rgba(0, 0, 0, 0.12);
}

.stat-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
}

.stat-info h3 {
  font-size: 32px;
  font-weight: 700;
  color: #2d3748;
  margin: 0 0 6px 0;
  letter-spacing: -0.5px;
}

.stat-info p {
  color: #718096;
  margin: 0;
  font-size: 14px;
  font-weight: 500;
}

.stat-icon {
  width: 48px;
  height: 48px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 20px;
  color: white;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.3);
}

.stat-trend {
  display: flex;
  align-items: center;
  gap: 4px;
  font-size: 14px;
  font-weight: 600;
  padding: 4px 12px;
  border-radius: 20px;
}

.stat-trend.up {
  color: #38a169;
  background: rgba(56, 161, 105, 0.1);
}

.stat-trend.down {
  color: #e53e3e;
  background: rgba(229, 62, 62, 0.1);
}

.stat-trend::before {
  content: '';
  width: 0;
  height: 0;
  border-left: 4px solid transparent;
  border-right: 4px solid transparent;
}

.stat-trend.up::before {
  border-bottom: 6px solid #38a169;
}

.stat-trend.down::before {
  border-top: 6px solid #e53e3e;
}

.charts-section {
  margin-bottom: 32px;
}

.chart-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 24px;
}

.chart-card {
  border: none;
  border-radius: 16px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.08);
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  background: white;
}

.chart-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 30px rgba(0, 0, 0, 0.1);
}

.chart-card :deep(.devui-card-header) {
  padding: 24px 24px 0;
  border-bottom: none;
}

.chart-card :deep(.devui-card-body) {
  padding: 16px 24px 24px;
}

.chart-card h3 {
  font-size: 18px;
  font-weight: 600;
  color: #2d3748;
  margin: 0;
}

.recent-section {
  margin-bottom: 32px;
}

.recent-section .devui-card {
  border: none;
  border-radius: 16px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.08);
  background: white;
}

.recent-section :deep(.devui-card-header) {
  padding: 24px 24px 0;
  border-bottom: none;
}

.recent-section :deep(.devui-card-body) {
  padding: 16px 24px 24px;
}

.recent-section h3 {
  font-size: 18px;
  font-weight: 600;
  color: #2d3748;
  margin: 0;
}

/* 表格样式增强 */
:deep(.devui-table) {
  border-radius: 8px;
  overflow: hidden;
}

:deep(.devui-table-header) {
  background: #f7fafc;
}

:deep(.devui-table-body tr:hover) {
  background: rgba(102, 126, 234, 0.05);
}

:deep(.status-success) {
  color: #38a169;
  background: rgba(56, 161, 105, 0.1);
  padding: 4px 8px;
  border-radius: 12px;
  font-size: 12px;
  font-weight: 600;
}

:deep(.status-warning) {
  color: #d69e2e;
  background: rgba(214, 158, 46, 0.1);
  padding: 4px 8px;
  border-radius: 12px;
  font-size: 12px;
  font-weight: 600;
}

:deep(.status-error) {
  color: #e53e3e;
  background: rgba(229, 62, 62, 0.1);
  padding: 4px 8px;
  border-radius: 12px;
  font-size: 12px;
  font-weight: 600;
}

/* 图表容器 */
.chart {
  border-radius: 8px;
}

/* 动画效果 */
@keyframes slideInUp {
  from {
    opacity: 0;
    transform: translateY(30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.stat-card,
.chart-card,
.recent-section .devui-card {
  animation: slideInUp 0.6s ease-out;
}

.stat-card:nth-child(2) {
  animation-delay: 0.1s;
}

.stat-card:nth-child(3) {
  animation-delay: 0.2s;
}

.stat-card:nth-child(4) {
  animation-delay: 0.3s;
}

.chart-card:first-child {
  animation-delay: 0.4s;
}

.chart-card:last-child {
  animation-delay: 0.5s;
}

.recent-section .devui-card {
  animation-delay: 0.6s;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .dashboard {
    padding: 16px;
  }
  
  .page-title {
    font-size: 24px;
  }
  
  .stats-grid {
    grid-template-columns: 1fr;
    gap: 16px;
    margin-bottom: 24px;
  }
  
  .chart-row {
    grid-template-columns: 1fr;
    gap: 16px;
  }
  
  .charts-section,
  .recent-section {
    margin-bottom: 24px;
  }
}
</style>