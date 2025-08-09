
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAtRange">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>

      <el-date-picker
            v-model="searchInfo.createdAtRange"
            class="w-[380px]"
            type="datetimerange"
            range-separator="至"
            start-placeholder="开始时间"
            end-placeholder="结束时间"
          />
       </el-form-item>
      
            <el-form-item label="网点ID" prop="networkId">
  <el-input v-model.number="searchInfo.networkId" placeholder="搜索条件" />
</el-form-item>

            <el-form-item label="状态" prop="status">
  <el-select v-model="searchInfo.status" clearable filterable placeholder="请选择" @clear="()=>{searchInfo.status=undefined}">
    <el-option v-for="(item,key) in rule_statusOptions" :key="key" :label="item.label" :value="item.value" />
  </el-select>
</el-form-item>
            

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button v-auth="btnAuth.add" type="primary" icon="plus" @click="openDialog()">新增</el-button>
            <el-button v-auth="btnAuth.batchDelete" icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
            <ExportTemplate v-auth="btnAuth.exportTemplate" template-id="ito_user_LockerPricingRules" />
            <ExportExcel v-auth="btnAuth.exportExcel" template-id="ito_user_LockerPricingRules" filterDeleted/>
            <ImportExcel v-auth="btnAuth.importExcel" template-id="ito_user_LockerPricingRules" @on-success="getTableData" />
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column sortable align="left" label="日期" prop="CreatedAt"width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
            <el-table-column align="left" label="网点ID" prop="networkId" width="120" />

            <el-table-column align="left" label="规则名称" prop="ruleName" width="120" />

            <el-table-column align="left" label="收费类型" prop="feeType" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.feeType,fee_typeOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="柜子类型" prop="lockerType" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.lockerType,locker_typeOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="免费时长" prop="freeDuration" width="120" />

            <el-table-column align="left" label="是否启用押金" prop="isDepositEnabled" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.isDepositEnabled,is_deposit_enabledOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="是否启用预付" prop="isAdvancePay" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.isAdvancePay,is_advance_payOptions) }}
    </template>
</el-table-column>
            <el-table-column align="left" label="每小时费用" prop="hourlyRate" width="120" />

            <el-table-column align="left" label="24小时封顶价" prop="dailyCap" width="120" />

            <el-table-column align="left" label="每日费用" prop="dailyRate" width="120" />

            <el-table-column align="left" label="预付金额" prop="advanceAmount" width="120" />

            <el-table-column align="left" label="押金金额" prop="depositAmount" width="120" />

            <el-table-column align="left" label="状态" prop="status" width="120">
    <template #default="scope">
    {{ filterDict(scope.row.status,rule_statusOptions) }}
    </template>
</el-table-column>
        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
            <template #default="scope">
            <el-button v-auth="btnAuth.info" type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button v-auth="btnAuth.edit" type="primary" link icon="edit" class="table-button" @click="updateLockerPricingRulesFunc(scope.row)">编辑</el-button>
            <el-button  v-auth="btnAuth.delete" type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'新增':'编辑'}}</span>
                <div>
                  <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="网点ID:" prop="networkId">
    <el-input v-model.number="formData.networkId" :clearable="true" placeholder="请输入网点ID" />
</el-form-item>
            <el-form-item label="规则名称:" prop="ruleName">
    <el-input v-model="formData.ruleName" :clearable="true" placeholder="请输入规则名称" />
</el-form-item>
            <el-form-item label="收费类型:" prop="feeType">
    <el-select v-model="formData.feeType" placeholder="请选择收费类型" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in fee_typeOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
            <el-form-item label="柜子类型:" prop="lockerType">
    <el-select v-model="formData.lockerType" placeholder="请选择柜子类型" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in locker_typeOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
            <el-form-item label="免费时长:" prop="freeDuration">
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
            <el-form-item label="状态:" prop="status">
    <el-select v-model="formData.status" placeholder="请选择状态" style="width:100%" filterable :clearable="true">
        <el-option v-for="(item,key) in rule_statusOptions" :key="key" :label="item.label" :value="item.value" />
    </el-select>
</el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="网点ID">
    {{ detailForm.networkId }}
</el-descriptions-item>
                    <el-descriptions-item label="规则名称">
    {{ detailForm.ruleName }}
</el-descriptions-item>
                    <el-descriptions-item label="收费类型">
    {{ detailForm.feeType }}
</el-descriptions-item>
                    <el-descriptions-item label="柜子类型">
    {{ detailForm.lockerType }}
</el-descriptions-item>
                    <el-descriptions-item label="免费时长">
    {{ detailForm.freeDuration }}
</el-descriptions-item>
                    <el-descriptions-item label="是否启用押金">
    {{ detailForm.isDepositEnabled }}
</el-descriptions-item>
                    <el-descriptions-item label="是否启用预付">
    {{ detailForm.isAdvancePay }}
</el-descriptions-item>
                    <el-descriptions-item label="每小时费用">
    {{ detailForm.hourlyRate }}
</el-descriptions-item>
                    <el-descriptions-item label="24小时封顶价">
    {{ detailForm.dailyCap }}
</el-descriptions-item>
                    <el-descriptions-item label="每日费用">
    {{ detailForm.dailyRate }}
</el-descriptions-item>
                    <el-descriptions-item label="预付金额">
    {{ detailForm.advanceAmount }}
</el-descriptions-item>
                    <el-descriptions-item label="押金金额">
    {{ detailForm.depositAmount }}
</el-descriptions-item>
                    <el-descriptions-item label="状态">
    {{ detailForm.status }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createLockerPricingRules,
  deleteLockerPricingRules,
  deleteLockerPricingRulesByIds,
  updateLockerPricingRules,
  findLockerPricingRules,
  getLockerPricingRulesList
} from '@/api/ito_user/lockerPricingRules'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
// 引入按钮权限标识
import { useBtnAuth } from '@/utils/btnAuth'
import { useAppStore } from "@/pinia"

// 导出组件
import ExportExcel from '@/components/exportExcel/exportExcel.vue'
// 导入组件
import ImportExcel from '@/components/exportExcel/importExcel.vue'
// 导出模板组件
import ExportTemplate from '@/components/exportExcel/exportTemplate.vue'


defineOptions({
    name: 'LockerPricingRules'
})
// 按钮权限实例化
    const btnAuth = useBtnAuth()

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const fee_typeOptions = ref([])
const is_deposit_enabledOptions = ref([])
const rule_statusOptions = ref([])
const locker_typeOptions = ref([])
const is_advance_payOptions = ref([])
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
               },
              ],
               feeType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               lockerType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               freeDuration : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getLockerPricingRulesList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    fee_typeOptions.value = await getDictFunc('fee_type')
    is_deposit_enabledOptions.value = await getDictFunc('is_deposit_enabled')
    rule_statusOptions.value = await getDictFunc('rule_status')
    locker_typeOptions.value = await getDictFunc('locker_type')
    is_advance_payOptions.value = await getDictFunc('is_advance_pay')
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteLockerPricingRulesFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
      const res = await deleteLockerPricingRulesByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateLockerPricingRulesFunc = async(row) => {
    const res = await findLockerPricingRules({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteLockerPricingRulesFunc = async (row) => {
    const res = await deleteLockerPricingRules({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
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
        }
}
// 弹窗确定
const enterDialog = async () => {
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
                closeDialog()
                getTableData()
              }
      })
}

const detailForm = ref({})

// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findLockerPricingRules({ ID: row.ID })
  if (res.code === 0) {
    detailForm.value = res.data
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailForm.value = {}
}


</script>

<style>

</style>
