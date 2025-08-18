<template>
  <div class="customer-form">
    <div class="page-header">
      <h1>{{ isEdit ? '编辑客户' : '新建客户' }}</h1>
      <div class="header-actions">
        <d-button @click="handleCancel">取消</d-button>
        <d-button type="primary" @click="handleSave" :loading="saving">
          {{ isEdit ? '更新' : '创建' }}
        </d-button>
      </div>
    </div>

    <d-card>
      <d-form
        ref="formRef"
        :model="formData"
        :rules="rules"
        label-position="left"
      >
        <div class="form-sections">
          <!-- 基本信息 -->
          <div class="form-section">
            <h3>基本信息</h3>
            
            <div class="form-row">
              <div class="form-item">
                <d-form-item field="customer_code" label="客户编码">
                  <d-input
                    v-model="formData.customer_code"
                    placeholder="请输入客户编码"
                    :disabled="isEdit"
                  />
                </d-form-item>
              </div>
              
              <div class="form-item">
                <d-form-item field="company_name" label="公司名称">
                  <d-input
                    v-model="formData.company_name"
                    placeholder="请输入公司名称"
                  />
                </d-form-item>
              </div>
            </div>

            <div class="form-row">
              <div class="form-item">
                <d-form-item field="contact_name" label="联系人">
                  <d-input
                    v-model="formData.contact_name"
                    placeholder="请输入联系人姓名"
                  />
                </d-form-item>
              </div>
              
              <div class="form-item">
                <d-form-item field="contact_phone" label="联系电话">
                  <d-input
                    v-model="formData.contact_phone"
                    placeholder="请输入联系电话"
                  />
                </d-form-item>
              </div>
            </div>

            <d-form-item field="contact_email" label="联系邮箱">
              <d-input
                v-model="formData.contact_email"
                placeholder="请输入联系邮箱"
              />
            </d-form-item>

            <d-form-item field="address" label="公司地址">
              <d-textarea
                v-model="formData.address"
                placeholder="请输入公司详细地址"
                :rows="3"
              />
            </d-form-item>
          </div>

          <!-- 业务信息 -->
          <div class="form-section">
            <h3>业务信息</h3>
            
            <div class="form-row">
              <div class="form-item">
                <d-form-item field="level" label="客户级别">
                  <d-select
                    v-model="formData.level"
                    placeholder="选择客户级别"
                  >
                    <d-option 
                      v-for="option in levelOptions" 
                      :key="option.value"
                      :label="option.label" 
                      :value="option.value"
                    />
                  </d-select>
                </d-form-item>
              </div>
              
              <div class="form-item">
                <d-form-item field="business_type" label="业务类型">
                  <d-input
                    v-model="formData.business_type"
                    placeholder="如：互联网、制造业等"
                  />
                </d-form-item>
              </div>
            </div>

            <d-form-item field="credit_limit" label="信用额度">
              <d-input
                v-model="formData.credit_limit"
                placeholder="请输入信用额度"
                type="number"
              >
                <template #prepend>¥</template>
              </d-input>
            </d-form-item>

            <d-form-item field="parent_id" label="上级客户" v-if="formData.level > 1">
              <d-select
                v-model="formData.parent_id"
                placeholder="选择上级客户"
                filterable
              >
                <d-option 
                  v-for="option in parentOptions" 
                  :key="option.value"
                  :label="option.label" 
                  :value="option.value"
                />
              </d-select>
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

const isEdit = computed(() => route.name === 'EditCustomer')

const formData = reactive({
  customer_code: '',
  company_name: '',
  contact_name: '',
  contact_phone: '',
  contact_email: '',
  address: '',
  level: 1,
  business_type: '',
  credit_limit: '',
  parent_id: ''
})

const rules = {
  customer_code: [
    { required: true, message: '请输入客户编码' },
    { pattern: /^[A-Z0-9]{4,20}$/, message: '客户编码必须为4-20位大写字母和数字' }
  ],
  company_name: [
    { required: true, message: '请输入公司名称' },
    { minLength: 2, message: '公司名称至少2个字符' }
  ],
  contact_name: [
    { required: true, message: '请输入联系人姓名' }
  ],
  contact_phone: [
    { pattern: /^1[3-9]\d{9}$/, message: '请输入正确的手机号码' }
  ],
  contact_email: [
    { pattern: /^[^\s@]+@[^\s@]+\.[^\s@]+$/, message: '请输入正确的邮箱地址' }
  ],
  level: [
    { required: true, message: '请选择客户级别' }
  ],
  credit_limit: [
    { required: true, message: '请输入信用额度' },
    { min: 0, message: '信用额度不能为负数' }
  ]
}

const levelOptions = [
  { label: '一级客户', value: 1 },
  { label: '二级代理', value: 2 },
  { label: '三级代理', value: 3 }
]

const parentOptions = ref([
  // TODO: 从API加载上级客户选项
  { label: '示例上级客户', value: 'parent-id-1' }
])

const loadCustomer = async () => {
  if (isEdit.value) {
    const customerId = route.params.id
    
    // TODO: 从API加载客户数据
    // const response = await customerAPI.getCustomer(customerId)
    // Object.assign(formData, response.data)
    
    console.log('加载客户数据:', customerId)
    
    // 模拟数据
    Object.assign(formData, {
      customer_code: 'CUST001',
      company_name: '北京科技有限公司',
      contact_name: '张三',
      contact_phone: '13800138000',
      contact_email: 'zhangsan@example.com',
      address: '北京市朝阳区',
      level: 1,
      business_type: '互联网',
      credit_limit: '1000000',
      parent_id: ''
    })
  } else {
    // 新建客户，生成客户编码
    formData.customer_code = generateCustomerCode()
  }
}

const generateCustomerCode = () => {
  const date = new Date()
  const year = date.getFullYear()
  const month = String(date.getMonth() + 1).padStart(2, '0')
  const day = String(date.getDate()).padStart(2, '0')
  const random = Math.floor(Math.random() * 1000).toString().padStart(3, '0')
  return `CUST${year}${month}${day}${random}`
}

const handleSave = async () => {
  try {
    const isValid = await formRef.value.validate()
    if (!isValid) return

    saving.value = true

    if (isEdit.value) {
      // TODO: 调用更新API
      // await customerAPI.updateCustomer(route.params.id, formData)
      console.log('更新客户:', formData)
    } else {
      // TODO: 调用创建API
      // await customerAPI.createCustomer(formData)
      console.log('创建客户:', formData)
    }

    // 模拟延迟
    await new Promise(resolve => setTimeout(resolve, 1000))

    router.push('/customers')
  } catch (error) {
    console.error('保存失败:', error)
  } finally {
    saving.value = false
  }
}

const handleCancel = () => {
  router.push('/customers')
}

onMounted(() => {
  loadCustomer()
})
</script>

<style scoped>
.customer-form {
  max-width: 1000px;
  margin: 0 auto;
}

.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
}

.page-header h1 {
  font-size: 24px;
  font-weight: 600;
  color: #333;
  margin: 0;
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
}
</style>