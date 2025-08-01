
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="id字段:" prop="id">
    <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入id字段" />
</el-form-item>
        <el-form-item label="网点ID:" prop="networkId">
    <el-input v-model.number="formData.networkId" :clearable="true" placeholder="请输入网点ID" />
</el-form-item>
        <el-form-item label="规则名称:" prop="ruleName">
    <el-input v-model="formData.ruleName" :clearable="true" placeholder="请输入规则名称" />
</el-form-item>
        <el-form-item label="1-计时收费 2-按日收费:" prop="feeType">
    <el-switch v-model="formData.feeType" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="1-小柜子 2-大柜子:" prop="lockerType">
    <el-switch v-model="formData.lockerType" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="免费时长(小时):" prop="freeDuration">
    <el-input-number v-model="formData.freeDuration" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
        <el-form-item label="是否启用押金:" prop="isDepositEnabled">
    <el-switch v-model="formData.isDepositEnabled" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="是否启用预付:" prop="isAdvancePay">
    <el-switch v-model="formData.isAdvancePay" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
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
    <el-switch v-model="formData.status" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="createdAt字段:" prop="createdAt">
    <el-date-picker v-model="formData.createdAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="updatedAt字段:" prop="updatedAt">
    <el-date-picker v-model="formData.updatedAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
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
} from '@/api/example/lockerPricingRules'

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
const formData = ref({
            id: undefined,
            networkId: undefined,
            ruleName: '',
            feeType: false,
            lockerType: false,
            freeDuration: 0,
            isDepositEnabled: false,
            isAdvancePay: false,
            hourlyRate: 0,
            dailyCap: 0,
            dailyRate: 0,
            advanceAmount: 0,
            depositAmount: 0,
            status: false,
            createdAt: new Date(),
            updatedAt: new Date(),
        })
// 验证规则
const rule = reactive({
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
