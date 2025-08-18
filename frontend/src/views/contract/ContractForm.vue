<template>
  <div class="contract-form">
    <div class="page-header">
      <h1>{{ isEdit ? '编辑合同' : isView ? '合同详情' : '新建合同' }}</h1>
      <div class="header-actions">
        <d-button @click="handleCancel">{{ isView ? '返回' : '取消' }}</d-button>
        <d-button @click="handleSaveDraft" v-if="!isEdit && !isView">保存草稿</d-button>
        <d-button v-if="!isView" type="primary" @click="handleSave" :loading="saving">
          {{ isEdit ? '更新' : '提交审批' }}
        </d-button>
        <d-button v-if="isView" type="primary" @click="handleEdit">
          编辑
        </d-button>
      </div>
    </div>

    <d-card>
      <d-form
        ref="formRef"
        :model="formData"
        :rules="rules"
        label-position="left"
        :disabled="isView"
      >
        <div class="form-sections">
          <!-- 基本信息 -->
          <div class="form-section">
            <h3>基本信息</h3>
            
            <div class="form-row">
              <div class="form-item">
                <d-form-item field="contract_no" label="合同编号">
                  <d-input
                    v-model="formData.contract_no"
                    placeholder="请输入合同编号"
                    :disabled="isEdit || isView"
                  />
                </d-form-item>
              </div>
              
              <div class="form-item">
                <d-form-item field="customer_id" label="客户">
                  <d-select
                    v-model="formData.customer_id"
                    placeholder="选择客户"
                    filterable
                    :disabled="isEdit"
                  >
                    <d-option 
                      v-for="customer in customerOptions" 
                      :key="customer.value"
                      :label="customer.label" 
                      :value="customer.value"
                    />
                  </d-select>
                </d-form-item>
              </div>
            </div>

            <d-form-item field="title" label="合同标题">
              <d-input
                v-model="formData.title"
                placeholder="请输入合同标题"
              />
            </d-form-item>

            <div class="form-row">
              <div class="form-item">
                <d-form-item field="start_date" label="开始日期">
                  <d-date-picker
                    v-model="formData.start_date"
                    placeholder="选择开始日期"
                    style="width: 100%"
                  />
                </d-form-item>
              </div>
              
              <div class="form-item">
                <d-form-item field="end_date" label="结束日期">
                  <d-date-picker
                    v-model="formData.end_date"
                    placeholder="选择结束日期"
                    style="width: 100%"
                  />
                </d-form-item>
              </div>
            </div>
          </div>

          <!-- 商务条款 -->
          <div class="form-section">
            <h3>商务条款</h3>
            
            <div class="form-row three-cols">
              <div class="form-item">
                <d-form-item field="contract_amount" label="合同金额">
                  <d-input
                    v-model="formData.contract_amount"
                    placeholder="请输入合同金额"
                    type="number"
                  >
                    <template #prepend>¥</template>
                  </d-input>
                </d-form-item>
              </div>
              
              <div class="form-item">
                <d-form-item field="discount_rate" label="折扣率">
                  <d-input
                    v-model="formData.discount_rate"
                    placeholder="请输入折扣率"
                    type="number"
                  >
                    <template #append>%</template>
                  </d-input>
                </d-form-item>
              </div>
              
              <div class="form-item">
                <d-form-item field="settlement_cycle" label="结算周期">
                  <d-select
                    v-model="formData.settlement_cycle"
                    placeholder="选择结算周期"
                  >
                    <d-option 
                      v-for="option in settlementOptions" 
                      :key="option.value"
                      :label="option.label" 
                      :value="option.value"
                    />
                  </d-select>
                </d-form-item>
              </div>
            </div>

            <d-form-item field="payment_terms" label="付款条款">
              <d-textarea
                v-model="formData.payment_terms"
                placeholder="请输入付款条款说明"
                :rows="3"
              />
            </d-form-item>
          </div>

          <!-- 备注信息 -->
          <div class="form-section">
            <h3>备注信息</h3>
            <d-form-item field="notes" label="备注">
              <d-textarea
                v-model="formData.notes"
                placeholder="请输入备注信息"
                :rows="4"
              />
            </d-form-item>
          </div>
        </div>
      </d-form>
    </d-card>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'

const router = useRouter()
const route = useRoute()

const formRef = ref()
const saving = ref(false)

const isEdit = computed(() => route.name === 'EditContract')
const isView = computed(() => route.name === 'ContractDetail')
const isCreate = computed(() => route.name === 'CreateContract')

const formData = reactive({
  contract_no: '',
  customer_id: '',
  title: '',
  start_date: null,
  end_date: null,
  contract_amount: '',
  discount_rate: '',
  settlement_cycle: 1,
  payment_terms: '',
  notes: ''
})

const rules = {
  contract_no: [
    { required: true, message: '请输入合同编号' }
  ],
  customer_id: [
    { required: true, message: '请选择客户' }
  ],
  title: [
    { required: true, message: '请输入合同标题' }
  ],
  start_date: [
    { required: true, message: '请选择开始日期' }
  ],
  end_date: [
    { required: true, message: '请选择结束日期' }
  ],
  contract_amount: [
    { required: true, message: '请输入合同金额' }
  ],
  discount_rate: [
    { required: true, message: '请输入折扣率' },
    { min: 0, message: '折扣率不能小于0' },
    { max: 1, message: '折扣率不能大于1' }
  ]
}

const customerOptions = ref([
  { label: '北京科技有限公司', value: 'customer-1' },
  { label: '上海创新公司', value: 'customer-2' }
])

const settlementOptions = [
  { label: '月结', value: 1 },
  { label: '季结', value: 3 },
  { label: '半年结', value: 6 },
  { label: '年结', value: 12 }
]

const providerOptions = [
  { label: '腾讯云', value: 'tencent' },
  { label: '阿里云', value: 'alibaba' },
  { label: '华为云', value: 'huawei' },
  { label: 'AWS', value: 'aws' }
]

const loadContract = async () => {
  if (isEdit.value || isView.value) {
    const contractId = route.params.id
    console.log('加载合同数据:', contractId)
    // TODO: 从API加载合同数据
    
    // 模拟数据
    Object.assign(formData, {
      contract_no: 'CON202401001',
      title: '云服务代理合同',
      customer_id: 'customer-1',
      start_date: '2024-01-01',
      end_date: '2024-12-31',
      contract_amount: '1000000.00',
      discount_rate: '0.85',
      settlement_cycle: 1,
      payment_terms: '月结30天',
      notes: '重要客户合同'
    })
  } else {
    // 新建合同，生成合同编号
    formData.contract_no = generateContractNo()
  }
}

const generateContractNo = () => {
  const date = new Date()
  const year = date.getFullYear()
  const month = String(date.getMonth() + 1).padStart(2, '0')
  const day = String(date.getDate()).padStart(2, '0')
  const random = Math.floor(Math.random() * 1000).toString().padStart(3, '0')
  return `CON${year}${month}${day}${random}`
}

const handleSaveDraft = async () => {
  try {
    saving.value = true
    console.log('保存草稿:', formData)
    // TODO: 调用API保存草稿
    router.push('/contracts')
  } catch (error) {
    console.error('保存草稿失败:', error)
  } finally {
    saving.value = false
  }
}

const handleSave = async () => {
  try {
    const isValid = await formRef.value.validate()
    if (!isValid) return

    saving.value = true

    const contractData = { ...formData }

    if (isEdit.value) {
      console.log('更新合同:', contractData)
      // TODO: 调用更新API
    } else {
      console.log('创建合同:', contractData)
      // TODO: 调用创建API
    }

    // 模拟延迟
    await new Promise(resolve => setTimeout(resolve, 1000))

    router.push('/contracts')
  } catch (error) {
    console.error('保存失败:', error)
  } finally {
    saving.value = false
  }
}

const handleCancel = () => {
  router.push('/contracts')
}

const handleEdit = () => {
  router.push(`/contracts/${route.params.id}/edit`)
}

onMounted(() => {
  loadContract()
})
</script>

<style scoped>
.contract-form {
  max-width: 1000px;
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

.header-actions {
  display: flex;
  gap: 12px;
}

.form-sections {
  max-width: 800px;
}

.form-section {
  margin-bottom: 32px;
}

.form-section h3 {
  font-size: 16px;
  font-weight: 600;
  color: #333;
  margin: 0 0 16px 0;
  padding-bottom: 8px;
  border-bottom: 1px solid #f0f0f0;
}

.form-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 24px;
  margin-bottom: 16px;
}

.form-row.three-cols {
  grid-template-columns: 1fr 1fr 1fr;
}

.form-item {
  margin-bottom: 0;
}

@media (max-width: 768px) {
  .page-header {
    flex-direction: column;
    gap: 16px;
    align-items: stretch;
  }
  
  .form-sections {
    max-width: none;
  }
  
  .form-row {
    grid-template-columns: 1fr;
    gap: 16px;
  }
  
  .form-row.three-cols {
    grid-template-columns: 1fr;
  }
}
</style>