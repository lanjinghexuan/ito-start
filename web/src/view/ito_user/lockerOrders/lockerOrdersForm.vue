
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="业务订单号（唯一标识）:" prop="orderNumber">
    <el-input v-model="formData.orderNumber" :clearable="true" placeholder="请输入业务订单号（唯一标识）" />
</el-form-item>
        <el-form-item label="用户ID（关联用户表）:" prop="userId">
    <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入用户ID（关联用户表）" />
</el-form-item>
        <el-form-item label="寄存开始时间:" prop="startTime">
    <el-date-picker v-model="formData.startTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="计划寄存时长（小时）:" prop="scheduledDuration">
    <el-input v-model.number="formData.scheduledDuration" :clearable="true" placeholder="请输入计划寄存时长（小时）" />
</el-form-item>
        <el-form-item label="实际寄存时长（小时）:" prop="actualDuration">
    <el-input v-model.number="formData.actualDuration" :clearable="true" placeholder="请输入实际寄存时长（小时）" />
</el-form-item>
        <el-form-item label="基础费用:" prop="price">
    <el-input-number v-model="formData.price" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
        <el-form-item label="优惠金额:" prop="discount">
    <el-input-number v-model="formData.discount" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
        <el-form-item label="实付金额:" prop="amountPaid">
    <el-input-number v-model="formData.amountPaid" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
        <el-form-item label="寄存网点名称:" prop="storageLocationName">
    <el-input v-model="formData.storageLocationName" :clearable="true" placeholder="请输入寄存网点名称" />
</el-form-item>
        <el-form-item label="柜子ID:" prop="cabinetId">
    <el-input v-model.number="formData.cabinetId" :clearable="true" placeholder="请输入柜子ID" />
</el-form-item>
        <el-form-item label="订单状态：1-待支付、2-寄存中、3-已完成、4-已取消、5-超时、6-异常:" prop="status">
    <el-select v-model="formData.status" placeholder="请选择订单状态：1-待支付、2-寄存中、3-已完成、4-已取消、5-超时、6-异常" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="创建时间:" prop="createTime">
    <el-date-picker v-model="formData.createTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="更新时间:" prop="updateTime">
    <el-date-picker v-model="formData.updateTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="押金状态：1-已支付、2-已退还、3-已扣除:" prop="depositStatus">
    <el-select v-model="formData.depositStatus" placeholder="请选择押金状态：1-已支付、2-已退还、3-已扣除" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in depositStatusOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
        <el-form-item label="寄存点id:" prop="lockerPointId">
    <el-input v-model.number="formData.lockerPointId" :clearable="true" placeholder="请输入寄存点id" />
</el-form-item>
        <el-form-item label="title字段:" prop="title">
    <el-input v-model="formData.title" :clearable="true" placeholder="请输入title字段" />
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
  createLockerOrders,
  updateLockerOrders,
  findLockerOrders
} from '@/api/ito_user/lockerOrders'

defineOptions({
    name: 'LockerOrdersForm'
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
const statusOptions = ref([])
const depositStatusOptions = ref([])
const formData = ref({
            orderNumber: '',
            userId: undefined,
            startTime: new Date(),
            scheduledDuration: undefined,
            actualDuration: undefined,
            price: 0,
            discount: 0,
            amountPaid: 0,
            storageLocationName: '',
            cabinetId: undefined,
            status: '',
            createTime: new Date(),
            updateTime: new Date(),
            depositStatus: '',
            lockerPointId: undefined,
            title: '',
        })
// 验证规则
const rule = reactive({
               orderNumber : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               userId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               scheduledDuration : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               amountPaid : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               storageLocationName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               cabinetId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               status : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               createTime : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               updateTime : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               depositStatus : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               title : [{
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
      const res = await findLockerOrders({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    statusOptions.value = await getDictFunc('status')
    depositStatusOptions.value = await getDictFunc('depositStatus')
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
               res = await createLockerOrders(formData.value)
               break
             case 'update':
               res = await updateLockerOrders(formData.value)
               break
             default:
               res = await createLockerOrders(formData.value)
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
