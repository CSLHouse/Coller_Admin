<template>
  <div>
    <warning-bar
      title="支持模糊查找，左键选中自动填充！"
    />
  </div>
  <el-menu class="el-menu-demo" mode="horizontal">
    <el-menu-item index="1">消费登记</el-menu-item>
  </el-menu>
  <div class="h-6" />
  <el-form
    ref="ruleFormRef"
    :model="memberForm"
    status-icon
    label-width="120px"
    class="demo-ruleForm"
  >
    <el-form-item label="手机/会员卡" prop="cardId" >
      <el-col :span="8">
        <el-input v-model.number="memberForm.onlyId" type="number" autocomplete="off"
        clearable @input="onTriggerSearch" @clear="onSearch"/>
      </el-col>
    </el-form-item>
    <el-form-item label="" prop="cardId">
      <div v-if="tableVisible">
        <el-table :data="tableData" style="width: 100%"
        highlight-current-row
        @row-click="OnConfirm">
          <el-table-column prop="cardId" label="会员卡号" width="120" height="40" />
          <el-table-column prop="telephone" label="手机号" width="120" height="40" />
          <el-table-column prop="memberName" label="姓名" width="80" />
          <el-table-column prop="comboType" label="套餐类型" width="80"/>
          <el-table-column prop="remainTimes" label="当前剩余" width="80"/>
          <el-table-column prop="deadline" label="截止日期" width="100"/>
          <el-table-column prop="stateStr" label="状态" width="100"/>
        </el-table>
      </div>
      
    </el-form-item>
    <el-form-item label="会员姓名" prop="memberName">
      <el-col :span="8">
        <el-input v-model="memberForm.memberName" type="text" autocomplete="off" disabled/>
      </el-col>
    </el-form-item>
    <el-form-item label="VIP卡" prop="comboType">
      <el-col :span="8">
        <el-input v-model="memberForm.comboType" type="text" autocomplete="off" disabled/>
      </el-col>
    </el-form-item>
    <el-form-item label="剩余" prop="remainTimes">
      <el-col :span="8">
        <el-input v-model.number="memberForm.remainTimes" type="number" autocomplete="off" disabled/>
      </el-col>
    </el-form-item>
    <el-form-item label="到期日期" prop="deadline">
      <el-col :span="8">
        <el-input v-model="memberForm.deadline" type="text" autocomplete="off" disabled/>
      </el-col>
    </el-form-item>
    <el-form-item label="划卡" prop="deadline">
      <el-col :span="8">
        <el-input v-model.number="memberForm.number" type="number" autocomplete="off" />
      </el-col>
      <a style="color: red; margin-left: 10px;">默认划卡一次，只有次卡可多划</a>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="submitForm(ruleFormRef)"
        >确认</el-button
      >
      <el-button @click="resetForm(ruleFormRef)">清空</el-button>
    </el-form-item>
  </el-form>
</template>
    
<script lang="ts" setup>
  import WarningBar from '@/components/warningBar/warningBar.vue'
  import { searchVIPCards, consumeVIPCard } from '@/api/member'
  import { reactive, ref, onBeforeMount, watch } from 'vue'
  import { FormInstance, ElMessage } from 'element-plus'
  import { comboStore } from '@/pinia/modules/combo'
  import { trim } from '@/utils/stringFun'
  import config from '@/core/config'

  const ruleFormRef = ref<FormInstance>()
  const tableVisible = ref(false)
  const tableData = ref([])
  // 查询会员卡类型
  const vipComboStore = comboStore()

  const memberForm = reactive({
    onlyId: null,
    Id: 0,
    cardId: null,
    telephone: null,
    memberName: '',
    comboType: '',
    remainTimes: null,
    startDate: '',
    deadline: '',
    state: null,
    comboId: 0,
    stateStr: '',
    number: null,
  })

  const onTriggerSearch = () => {
    if (memberForm.onlyId >= 1000 ) {
      onSearch()
    }
  }

  const onSearch = async() => {
    const res = await searchVIPCards({ onlyId: memberForm.onlyId})
    if ('code' in res && res.code === 0) {
      if (res.data.list) {
        tableVisible.value = true
        let comboList = vipComboStore.comboList
        if (comboList['length'] < 1) {
          await vipComboStore.GetAllVIPCombos()
          comboList = vipComboStore.comboList
        }

        tableData.value = res.data.list
        tableData.value.forEach(memberElement => {
          memberElement.comboType = memberElement.combo.comboName
          // comboList.forEach(comboElement => {
          //   if (memberElement.comboId == comboElement.Id) {
          //     memberElement.comboType = comboElement.comboName
          //   }
          // });
          config.memberStateOptions.forEach(stateElement => {
            if (memberElement.state == stateElement.id) {
              memberElement.stateStr = stateElement.label
            }
          })
        });
        console.log("----tableData----", tableData)
      }
    }
  }
  
  const OnConfirm = async(row) => {
    memberForm.Id = row.Id
    memberForm.cardId = row.cardId
    memberForm.telephone = row.telephone
    memberForm.memberName = row.memberName
    memberForm.comboId = row.comboId
    memberForm.comboType = row.comboType
    memberForm.remainTimes = row.remainTimes
    memberForm.startDate = row.startDate
    memberForm.deadline = row.deadline
    memberForm.state = row.state
    memberForm.number = 1
    console.log("---memberForm:", memberForm)
    // let comboList = vipComboStore.comboList
    // if (comboList['length'] < 1) {
    //   await vipComboStore.GetAllVIPCombos()
    //   comboList = vipComboStore.comboList
    // }
    // comboList.forEach(element => {
    //   if (element.Id == memberForm.comboId) {
    //     memberForm.comboType = element.comboName
    //   }
    // });
  }

  // const getComboData = async() => {
  //   let comboList = vipComboStore.comboList
  //   if (comboList['length'] < 1) {
  //     await vipComboStore.GetAllVIPCombos()
  //     comboList = vipComboStore.comboList
  //   }

  //   comboList.forEach(element => {
  //     comboOptions.value.push({key: element.Id, value: element.comboName, price: element.comboPrice})
  //   });
  // }

  // onBeforeMount(() => {
  //   getComboData()
  // })

  const submitForm = async(formEl: FormInstance | undefined) => {
    if (memberForm.memberName.length < 1) {
      ElMessage({
        type: 'error',
        message: '信息填写错误'
      })
      return
    }
    const requestData = {
      Id: memberForm.Id,
      cardId: memberForm.cardId,
      number: memberForm.number,
    }

    let res = await consumeVIPCard(requestData)
    if ('code' in res && res.code === 0) {
      ElMessage({
        type: 'success',
        message: '恭喜，登记成功'
      })
      tableVisible.value = false
      memberForm.onlyId = null
      memberForm.memberName = ''
      memberForm.comboId = 0
      memberForm.remainTimes = null
      memberForm.deadline = ''
      memberForm.comboType = ''
      memberForm.number = null
      tableData.value = []
      return
    }

  }
  
  const resetForm = (formEl: FormInstance | undefined) => {
    if (!formEl) return
    formEl.resetFields()
    tableVisible.value = false
    memberForm.onlyId = null
    memberForm.memberName = ''
    memberForm.comboId = 0
    memberForm.remainTimes = null
    memberForm.deadline = ''
    memberForm.comboType = ''
    memberForm.number = null
    tableData.value = []
  }
</script>
  