
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="网点ID:" prop="networkId">
    <el-input v-model.number="formData.networkId" :clearable="true" placeholder="请输入网点ID" />
</el-form-item>
        <el-form-item label="规则名称:" prop="ruleName">
    <el-input v-model="formData.ruleName" :clearable="true" placeholder="请输入规则名称" />
</el-form-item>
        <el-form-item label="1-计时收费 2-按日收费:" prop="feeType">
    <el-select v-model="formData.feeType" placeholder="请选择1-计时收费 2-按日收费" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in fee_typeOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="1-小柜子 2-大柜子:" prop="lockerType">
    <el-select v-model="formData.lockerType" placeholder="请选择1-小柜子 2-大柜子" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in locker_typeOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="免费时长(小时):" prop="freeDuration">
    <el-input-number v-model="formData.freeDuration" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
        <el-form-item label="是否启用押金:" prop="isDepositEnabled">
    <el-select v-model="formData.isDepositEnabled" placeholder="请选择是否启用押金" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in is_deposit_enabledOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="是否启用预付:" prop="isAdvancePay">
    <el-select v-model="formData.isAdvancePay" placeholder="请选择是否启用预付" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in is_advance_payOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="每小时费用:" prop="hourlyRate">
    <el-input-number v-model="formData.hourlyRate" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
        <el-form-item label="24小时封顶价:" prop="dailyCap">
    <el-input-number v-model="formData.dailyCap" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
        <el-form-item label="每日费用:" prop="dailyRate">
    <el-input-number v-model="formData.dailyRate" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
        <el-form-item label="预付金额:" prop="advanceAmount">
    <el-input-number v-model="formData.advanceAmount" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
        <el-form-item label="押金金额:" prop="depositAmount">
    <el-input-number v-model="formData.depositAmount" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
        <el-form-item label="1-生效 0-失效:" prop="status">
    <el-select v-model="formData.status" placeholder="请选择1-生效 0-失效" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in rule_statusOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createLockerPricingRules,
  updateLockerPricingRules,
  findLockerPricingRules
} from '@/api/ito_user/lockerPricingRules'

defineOptions({
    name: 'LockerPricingRulesForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const is_deposit_enabledOptions = ref([])
const is_advance_payOptions = ref([])
const fee_typeOptions = ref([])
const locker_typeOptions = ref([])
const rule_statusOptions = ref([])
const formData = ref({
            networkId: undefined,
            ruleName: '',
            feeType: '',
            lockerType: '',
            freeDuration: 0,
            isDepositEnabled: '',
            isAdvancePay: '',
            hourlyRate: 0,
            dailyCap: 0,
            dailyRate: 0,
            advanceAmount: 0,
            depositAmount: 0,
            status: '',
        })
// 验证规则
const rule = reactive({
               networkId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               feeType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               lockerType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               freeDuration : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findLockerPricingRules({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    is_deposit_enabledOptions.value = await getDictFunc('is_deposit_enabled')
    is_advance_payOptions.value = await getDictFunc('is_advance_pay')
    fee_typeOptions.value = await getDictFunc('fee_type')
    locker_typeOptions.value = await getDictFunc('locker_type')
    rule_statusOptions.value = await getDictFunc('rule_status')
}

init()
// 保存按钮
const save = async() => {
      btnLoading.value = true
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return btnLoading.value = false
            let res
           switch (type.value) {
             case 'create':
               res = await createLockerPricingRules(formData.value)
               break
             case 'update':
               res = await updateLockerPricingRules(formData.value)
               break
             default:
               res = await createLockerPricingRules(formData.value)
               break
           }
           btnLoading.value = false
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
