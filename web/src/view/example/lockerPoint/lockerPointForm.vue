
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="寄存点ID:" prop="id">
    <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入寄存点ID" />
</el-form-item>
        <el-form-item label="所属地点ID:" prop="locationId">
    <el-input v-model.number="formData.locationId" :clearable="true" placeholder="请输入所属地点ID" />
</el-form-item>
        <el-form-item label="寄存点名称:" prop="name">
    <el-input v-model="formData.name" :clearable="true" placeholder="请输入寄存点名称" />
</el-form-item>
        <el-form-item label="详细地址:" prop="address">
    <el-input v-model="formData.address" :clearable="true" placeholder="请输入详细地址" />
</el-form-item>
        <el-form-item label="纬度:" prop="latitude">
    <el-input-number v-model="formData.latitude" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
        <el-form-item label="经度:" prop="longitude">
    <el-input-number v-model="formData.longitude" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
        <el-form-item label="可用大柜数量:" prop="availableLarge">
    <el-input v-model.number="formData.availableLarge" :clearable="true" placeholder="请输入可用大柜数量" />
</el-form-item>
        <el-form-item label="可用中柜数量:" prop="availableMedium">
    <el-input v-model.number="formData.availableMedium" :clearable="true" placeholder="请输入可用中柜数量" />
</el-form-item>
        <el-form-item label="可用小柜数量:" prop="availableSmall">
    <el-input v-model.number="formData.availableSmall" :clearable="true" placeholder="请输入可用小柜数量" />
</el-form-item>
        <el-form-item label="营业时间:" prop="openTime">
    <el-input v-model="formData.openTime" :clearable="true" placeholder="请输入营业时间" />
</el-form-item>
        <el-form-item label="联系电话:" prop="mobile">
    <el-input v-model="formData.mobile" :clearable="true" placeholder="请输入联系电话" />
</el-form-item>
        <el-form-item label="管理员id:" prop="adminId">
    <el-input v-model.number="formData.adminId" :clearable="true" placeholder="请输入管理员id" />
</el-form-item>
        <el-form-item label="网点图片:" prop="pointImage">
    <SelectImage
     v-model="formData.pointImage"
     file-type="image"
    />
</el-form-item>
        <el-form-item label="状态:" prop="status">
    <el-input v-model="formData.status" :clearable="true" placeholder="请输入状态" />
</el-form-item>
        <el-form-item label="网点类型:" prop="pointType">
    <el-input v-model="formData.pointType" :clearable="true" placeholder="请输入网点类型" />
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
  createLockerPoint,
  updateLockerPoint,
  findLockerPoint
} from '@/api/example/lockerPoint'

defineOptions({
    name: 'LockerPointForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const formData = ref({
            id: undefined,
            locationId: undefined,
            name: '',
            address: '',
            latitude: 0,
            longitude: 0,
            availableLarge: undefined,
            availableMedium: undefined,
            availableSmall: undefined,
            openTime: '',
            mobile: '',
            adminId: undefined,
            pointImage: "",
            status: '',
            pointType: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findLockerPoint({ ID: route.query.id })
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
               res = await createLockerPoint(formData.value)
               break
             case 'update':
               res = await updateLockerPoint(formData.value)
               break
             default:
               res = await createLockerPoint(formData.value)
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
