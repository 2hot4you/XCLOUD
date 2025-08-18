<template>
  <div class="contract-list">
    <div class="page-header">
      <h1>合同管理</h1>
      <d-button type="primary" @click="handleCreate">
        <template #icon>
          <span>➕</span>
        </template>
        新建合同
      </d-button>
    </div>

    <!-- 搜索筛选 -->
    <d-card class="search-card">
      <div class="search-form">
        <d-input
          v-model="searchForm.keyword"
          placeholder="搜索合同编号或标题"
          style="width: 240px;"
          @keyup.enter="handleSearch"
        >
          <template #prefix>
            <span>🔍</span>
          </template>
        </d-input>
        
        <d-select
          v-model="searchForm.status"
          placeholder="选择状态"
          style="width: 120px;"
        >
          <d-option label="全部" value=""></d-option>
          <d-option label="草稿" value="draft"></d-option>
          <d-option label="待审批" value="pending"></d-option>
          <d-option label="生效中" value="active"></d-option>
          <d-option label="已到期" value="expired"></d-option>
          <d-option label="已终止" value="terminated"></d-option>
        </d-select>
        
        <d-select
          v-model="searchForm.customer_id"
          placeholder="选择客户"
          style="width: 180px;"
          filterable
        >
          <d-option label="全部客户" value=""></d-option>
          <d-option label="北京科技有限公司" value="customer-1"></d-option>
        </d-select>
        
        <div class="search-buttons">
          <d-button type="primary" @click="handleSearch">
            <template #icon>
              <span>🔍</span>
            </template>
            搜索
          </d-button>
          <d-button @click="handleReset">
            <template #icon>
              <span>🔄</span>
            </template>
            重置
          </d-button>
        </div>
      </div>
    </d-card>

    <!-- 合同列表 -->
    <d-card>
      <d-table
        :data="contracts"
        :loading="loading"
        size="md"
        stripe
      >
        <d-column field="contract_no" header="合同编号" width="140px"></d-column>
        <d-column field="title" header="合同标题" min-width="200px"></d-column>
        <d-column field="customer_name" header="客户名称" min-width="180px"></d-column>
        <d-column field="status" header="状态" width="100px">
          <template #cell="{ rowData }">
            <d-tag :type="getStatusType(rowData.status)">
              {{ getStatusText(rowData.status) }}
            </d-tag>
          </template>
        </d-column>
        <d-column field="start_date" header="开始日期" width="110px"></d-column>
        <d-column field="end_date" header="结束日期" width="110px"></d-column>
        <d-column field="contract_amount" header="合同金额" width="120px">
          <template #cell="{ rowData }">
            <span class="amount">¥{{ formatAmount(rowData.contract_amount) }}</span>
          </template>
        </d-column>
        <d-column field="discount_rate" header="折扣率" width="80px">
          <template #cell="{ rowData }">
            <span>{{ (rowData.discount_rate * 100).toFixed(1) }}%</span>
          </template>
        </d-column>
        <d-column field="actions" header="操作" width="140px">
          <template #cell="{ rowData }">
            <d-button variant="text" size="sm" @click="handleView(rowData)">查看</d-button>
            <d-button variant="text" size="sm" @click="handleEdit(rowData)">编辑</d-button>
            <d-dropdown>
              <d-button variant="text" size="sm">
                更多
                <span>🔽</span>
              </d-button>
              <template #overlay>
                <d-dropdown-menu>
                  <d-dropdown-item @click="handleCopy(rowData)">复制</d-dropdown-item>
                  <d-dropdown-item @click="handleDelete(rowData)" class="danger">删除</d-dropdown-item>
                </d-dropdown-menu>
              </template>
            </d-dropdown>
          </template>
        </d-column>
      </d-table>
      
      <!-- 分页 -->
      <div class="pagination">
        <d-pagination
          v-model:current="pagination.current"
          v-model:page-size="pagination.pageSize"
          :total="pagination.total"
          :show-size-changer="true"
          :page-size-options="[10, 20, 50, 100]"
          @current-change="handlePageChange"
          @page-size-change="handlePageSizeChange"
        />
      </div>
    </d-card>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()

interface Contract {
  id: string
  contract_no: string
  customer_id: string
  customer_name: string
  title: string
  status: 'draft' | 'pending' | 'active' | 'expired' | 'terminated'
  start_date: string
  end_date: string
  contract_amount: number
  discount_rate: number
  created_at: string
}

const loading = ref(false)
const contracts = ref<Contract[]>([])

const searchForm = reactive({
  keyword: '',
  status: '',
  customer_id: ''
})

const pagination = reactive({
  current: 1,
  pageSize: 20,
  total: 0
})

const statusOptions = [
  { label: '全部', value: '' },
  { label: '草稿', value: 'draft' },
  { label: '待审批', value: 'pending' },
  { label: '生效中', value: 'active' },
  { label: '已到期', value: 'expired' },
  { label: '已终止', value: 'terminated' }
]

const customerOptions = ref([
  { label: '全部客户', value: '' },
  { label: '北京科技有限公司', value: 'customer-1' }
])

const getStatusType = (status: string) => {
  switch (status) {
    case 'active': return 'success'
    case 'draft': return 'info'
    case 'pending': return 'warning'
    case 'expired': return 'danger'
    case 'terminated': return 'danger'
    default: return 'info'
  }
}

const getStatusText = (status: string) => {
  switch (status) {
    case 'draft': return '草稿'
    case 'pending': return '待审批'
    case 'active': return '生效中'
    case 'expired': return '已到期'
    case 'terminated': return '已终止'
    default: return '未知'
  }
}

const formatAmount = (amount: number) => {
  return new Intl.NumberFormat('zh-CN', {
    minimumFractionDigits: 2,
    maximumFractionDigits: 2
  }).format(amount)
}

const loadContracts = async () => {
  try {
    loading.value = true
    
    // TODO: 调用API获取合同列表
    // const response = await contractAPI.getContracts({
    //   page: pagination.current,
    //   page_size: pagination.pageSize,
    //   ...searchForm
    // })
    
    // 模拟数据
    contracts.value = [
      {
        id: '1',
        contract_no: 'CON202401001',
        customer_id: 'customer-1',
        customer_name: '北京科技有限公司',
        title: '云服务代理合同',
        status: 'active',
        start_date: '2024-01-01',
        end_date: '2024-12-31',
        contract_amount: 1000000.00,
        discount_rate: 0.85,
        created_at: '2024-01-15 10:30:00'
      }
    ]
    pagination.total = 1
    
  } catch (error) {
    console.error('加载合同列表失败:', error)
  } finally {
    loading.value = false
  }
}

const handleSearch = () => {
  pagination.current = 1
  loadContracts()
}

const handleReset = () => {
  searchForm.keyword = ''
  searchForm.status = ''
  searchForm.customer_id = ''
  handleSearch()
}

const handleCreate = () => {
  router.push('/contracts/create')
}

const handleView = (contract: Contract) => {
  router.push(`/contracts/${contract.id}`)
}

const handleEdit = (contract: Contract) => {
  router.push(`/contracts/${contract.id}/edit`)
}

const handleCopy = (contract: Contract) => {
  console.log('复制合同:', contract)
}

const handleDelete = async (contract: Contract) => {
  console.log('删除合同:', contract)
}

const handlePageChange = (page: number) => {
  pagination.current = page
  loadContracts()
}

const handlePageSizeChange = (size: number) => {
  pagination.pageSize = size
  pagination.current = 1
  loadContracts()
}

onMounted(() => {
  loadContracts()
})
</script>

<style scoped>
.contract-list {
  max-width: 1400px;
  margin: 0 auto;
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

.search-card {
  margin-bottom: 16px;
}

.search-form {
  display: flex;
  gap: 16px;
  align-items: end;
  flex-wrap: wrap;
}

.search-form :deep(.devui-form-item) {
  margin-bottom: 0;
}

.amount {
  font-weight: 600;
  color: #1890ff;
}

.pagination {
  display: flex;
  justify-content: center;
  margin-top: 16px;
}

:deep(.danger) {
  color: #ff4d4f;
}

@media (max-width: 768px) {
  .page-header {
    flex-direction: column;
    gap: 16px;
    align-items: stretch;
  }
  
  .search-form {
    flex-direction: column;
    gap: 12px;
  }
  
  .search-form :deep(.devui-form-item) {
    width: 100%;
  }
}
</style>