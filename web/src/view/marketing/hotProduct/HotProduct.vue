<template>
    <div>
      <div>
        <el-card class="box-card" title="筛选搜索">
          <el-form :inline="true" :model="searchData" class="demo-form-inline">
            <el-form-item label="商品名称：" class="form-item">
              <el-input v-model="searchData.productName" placeholder="商品名称" clearable />
            </el-form-item>
            <el-form-item label="推荐状态：" class="form-item">
              <el-input v-model="searchData.recommendStatus" placeholder="广告位置" clearable />
            </el-form-item>
          </el-form>
          <el-button type="primary" @click="onSearch">搜索结果</el-button>
        </el-card>
      </div>
      <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openDialog">选择商品</el-button>
        </div>
        <el-table
          ref="multipleTable"
          :data="hotProductData"
          style="width: 100%"
          tooltip-effect="dark"
          row-key="Id"
          @selection-change="handleSelectionChange"
        >
          <el-table-column type="selection" width="40" />
          <el-table-column align="left" label="编号" prop="ID" width="60"></el-table-column>
          <el-table-column align="left" label="商品名称" prop="productName" width="460" />
          <el-table-column align="left" label="是否推荐" prop="state" width="120" >
            <template #default="scope">
                <el-switch
                    v-model="scope.row.recommendStatus"
                    :active-value="1"
                    :inactive-value="0">
                </el-switch>
            </template>
          </el-table-column>
          <el-table-column align="left" label="排序" prop="sort" width="120" />
          <el-table-column align="left" label="操作" min-width="160">
            <template #default="scope">
              <el-button type="primary" link icon="edit" @click="updateProduct(scope.row)">变更</el-button>
              <el-popover v-model="scope.row.visible" placement="top" width="160">
                <p>确定要删除吗？</p>
                <div style="text-align: right; margin-top: 8px;">
                  <el-button type="primary" link @click="scope.row.visible = false">取消</el-button>
                  <el-button type="primary" @click="deleteProduct(scope.row)">确定</el-button>
                </div>
                <template #reference>
                  <el-button type="danger" link icon="delete" @click="scope.row.visible = true">删除</el-button>
                </template>
              </el-popover>
            </template>
          </el-table-column>
        </el-table>
        <div class="gva-pagination">
          <div style="margin-top: 15px;float: left;">
            <el-select v-model="stateOption" value-key="id" class="m-2" placeholder="批量操作" size="large">
                <el-option
                    v-for="item in stateOptions"
                    :key="item.id"
                    :label="item.label"
                    :value="item"
                />
            </el-select>
            <el-button type="primary" @click="toggleSelection()">确定</el-button>
          </div>
          <el-pagination
            :current-page="pageHot"
            :page-size="pageSizeHot"
            :page-sizes="[5, 5, 5, 5]"
            :total="totalHot"
            layout="total, sizes, prev, pager, next, jumper"
            @current-change="handleCurrentChangeHot"
            @size-change="handleSizeChangeHot"
          />
        </div>
      </div>
      <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="选择商品">
        <el-input
            v-model="productSearchData.name"
            placeholder="商品名称搜索"
            class="input-with-select"
            >
            <template #append>
                <el-button :icon="Search" @click="onSearchProduct" />
            </template>
        </el-input>
        <el-table
          ref="multipleTable"
          :data="productData"
          style="width: 100%"
          tooltip-effect="dark"
          row-key="Id"
          @selection-change="handleSelectionProductChange"
        >
          <el-table-column type="selection" width="40" />
          <el-table-column align="left" label="商品名称" prop="name" width="200" />
          <el-table-column align="left" label="货号" prop="productSN" width="200" />
          <el-table-column align="left" label="价格" prop="price" width="80" />
        </el-table>
        <el-pagination
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[5, 5, 5, 5]"
            :total="total"
            layout=" prev, pager, next"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
          />
        <el-button  @click="closeDialog()">取消</el-button>
        <el-button type="primary" @click="toggleSelectionProduct()">确定</el-button>
      </el-dialog>
    </div>
  </template>
  
  <script setup lang="ts">
  import {
    getProductList,
    getHotProductList,
    addHotProductList,
    updateHotProduct,
    deleteHotProduct,
  } from '@/api/product'
  import { ref, reactive, onBeforeMount, watch } from 'vue'
  import { ElMessage } from 'element-plus'
  import config from '@/core/config'
  import { trim } from '@/utils/stringFun'
  import { number } from 'echarts'
  import { useRouter } from "vue-router";
  import { ProductStore } from '@/pinia/modules/product'  
  import { Search } from '@element-plus/icons-vue'

  const searchData = reactive({
    productName: null,
    recommendStatus: null,
  })
  
  const onSearch = async() => {
    // TODO: 根据广告名称或到期时间搜索
    getTableData()
  }
  
  const onCancel = () => {
    getTableData()
  }
  
  const onExport = () => {
  
  }
  
  const pageHot = ref(1)
  const totalHot = ref(0)
  const pageSizeHot = ref(5)
  const hotProductData = ref([])
  
  // 分页
  const handleSizeChangeHot = (val) => {
    pageSizeHot.value = val
    getTableData()
  }
  
  const handleCurrentChangeHot = (val) => {
    pageHot.value = val
    getTableData()
  }
  
  const productStore = ProductStore()
  // 查询
  const getTableData = async() => {
    const res = await getHotProductList({ productName: searchData.productName,
         recommendStatus:searchData.recommendStatus,
        page: pageHot.value, pageSize: pageSizeHot.value })
    if ('code' in res && res.code === 0) {
        hotProductData.value = res.data.list
        totalHot.value = res.data.total
        pageHot.value = res.data.page
        pageSizeHot.value = res.data.pageSize
    }
  }

  getTableData()
  interface stateItem {
    id: number,
    label: string,
    key: string,
    dbKey: string,
    value: number,
  }
  const stateOption = ref<stateItem>()
  const stateOptions = ref<stateItem[]>([])
  stateOptions.value = [
    { id: 0, label: '设为推荐', key: "recommendStatus", dbKey: "recommend_status", value: 1 },
    { id: 1, label: '取消推荐', key: "recommendStatus", dbKey: "recommend_status",  value: 0 },
    { id: 2, label: '删除', key: "delete", dbKey: "delete",  value: 0 },
  ]

  const multipleSelection = ref()
  const handleSelectionChange = (val) => {
    multipleSelection.value = val
  }
  var updateList:number[] = new Array() 
  const toggleSelection = async() => {
    if (!stateOption.value || !multipleSelection.value || multipleSelection.value.length < 1) return
    multipleSelection.value.forEach((item) => {
        item[stateOption.value.key] = stateOption.value.value
        updateList.push(item.id)
    })
    if (stateOption.value.id < 2) {
        const res = await updateHotProduct({products: updateList, key: stateOption.value.dbKey, value: stateOption.value.value })
        if ('code' in res && res.code == 0) {
            hotProductData.value.forEach(element => {
                element[stateOption.value.key] = stateOption.value.value
            });
        }
    } else if (stateOption.value.id == 2) {
        const res = await deleteHotProduct({products: updateList })
        if ('code' in res && res.code == 0) {
            getTableData()
        }
    } else {
        ElMessage.error('没有此操作')
    }
    
    updateList = []
  }
  
  const type = ref('')
  const dialogFormVisible = ref(false)
  const productForm = ref({
    productId: 0,
    productName: '',
    recommendStatus: 0,
    sort: 0,
  })
  const updateProduct = async(row) => {
    dialogFormVisible.value = true
    type.value = 'update'
    productForm.value = row
  }

  const page = ref(1)
  const total = ref(0)
  const pageSize = ref(5)
  const productData = ref([])
  const getProductListTableData = async() => {
    const res = await getProductList({ name: productSearchData.name, page: page.value, pageSize: pageSize.value })
    if ('code' in res && res.code === 0) {
        productData.value = res.data.list
        productData.value.forEach(element => {
            element.productSN = "NO." + element.productSN
            element.price = "￥" + element.price
        });
        total.value = res.data.total
        page.value = res.data.page
        pageSize.value = res.data.pageSize
    }
  }
   // 分页
   const handleSizeChange = (val) => {
    pageSize.value = val
    getProductListTableData()
  }
  
  const handleCurrentChange = (val) => {
    page.value = val
    getProductListTableData()
  }

  const openDialog = () => {
    dialogFormVisible.value = true
    type.value = 'create'
    getProductListTableData()
  }

  const closeDialog = () => {
    dialogFormVisible.value = false
  }

  const productSearchData = reactive({
    name: null,
  })

  const onSearchProduct = async() => {
    getProductListTableData()
  }

  const multipleProductSelection = ref()
  const handleSelectionProductChange = (val) => {
    multipleProductSelection.value = val
  }
  var updateProductList:object[] = new Array() 
  const toggleSelectionProduct = async() => {
    if (!multipleProductSelection.value || multipleProductSelection.value.length < 1) return
    multipleProductSelection.value.forEach((item) => {
        let product = {productId: item.id, productName: item.name}
        updateProductList.push(product)
    })
    
    const res = await addHotProductList({products: updateProductList })
    if ('code' in res && res.code == 0) {
        getTableData()
        closeDialog()
    }
    updateProductList = []
    
  }
  </script>
  
  <style scoped>
  .form-item {
    margin: 2px 5px 0 2px;
  }
  .box-card {
    margin: 10px 0 10px 0;
    border: 1px 0 1px 0;
  }

  :deep(.gva-table-box) .el-table .cell {
    white-space: pre-line !important;
  }
  </style>
  