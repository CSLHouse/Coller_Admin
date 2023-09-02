<template>
    <el-dialog v-model="dialogVisible" :before-close="closeDialog" title="Vip套餐">
      <el-form  :model="comboForm" label-width="140px" >
        <el-form-item label="会员套餐名称">
          <el-input v-model="comboForm.comboName" autocomplete="off" />
        </el-form-item>
        <el-form-item label="会员套餐类型">
          <el-select v-model="comboTypeValue" class="m-2" placeholder="请选择会员卡" size="large">
              <el-option
              v-for="item in config.comboTypeOptions"
              :key="item.id"
              :label="item.label"
              :value="item"
              />
          </el-select>
        </el-form-item>
        <el-form-item label="价格">
          <el-input v-model.number="comboForm.comboPrice" autocomplete="off" />
        </el-form-item>
        <el-form-item label="天数/次数/金额">
          <el-input v-model.number="comboForm.amount" autocomplete="off" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
</template>
<script setup lang="ts">
import {
  createExaVIPCombo,
  updateExaVIPCombo,
  deleteExaVIPCombo,
  getExaVIPComboList
} from '@/api/combo'
import { ElMessage } from 'element-plus'
import { ref, watch, toRefs } from 'vue'
import config from '@/core/config'

const props = defineProps({
  //子组件接收父组件传递过来的值
  dialogVisible: Boolean,
  operateType: String
})
//使用父组件传递过来的值
const { dialogVisible, operateType } = toRefs(props)
const emit = defineEmits(['cancel'])

const comboForm = ref({
  ID: 0,
  storeName: '',
  comboName: '',
  comboType: 1,
  comboPrice: 0,
  amount: 0,
  state: 0,
})

type Option = {
  id: number
  label: string
}
const comboTypeValue = ref<Option>()

watch(() => comboTypeValue.value, () => {
    comboForm.value.comboType = comboTypeValue.value.id
})

const closeDialog = () => {
    dialogVisible.value = false
    emit('cancel', true)
}

const enterDialog = async() => {
  let res
  switch (operateType.value) {
    case 'create':
      res = await createExaVIPCombo(comboForm.value)
      break
    case 'update':
      res = await updateExaVIPCombo(comboForm.value)
      break
    default:
      res = await createExaVIPCombo(comboForm.value)
      break
  }

  if (res.code === 0) {
    dialogVisible.value = false
  }
}
</script>
<style></style>
