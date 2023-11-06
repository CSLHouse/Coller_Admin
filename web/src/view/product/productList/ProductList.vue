<template>
    <div>
      <div>
        <el-card class="box-card">
          <el-form :inline="true" :model="searchData" class="demo-form-inline">
            <el-form-item label="商品名称：" class="form-item">
              <el-input v-model.number="searchData.name" placeholder="商品名称" clearable />
            </el-form-item>
            <el-form-item label="商品货号：" class="form-item">
              <el-input v-model.number="searchData.productSN" placeholder="商品货号" clearable />
            </el-form-item>
            <el-form-item label="商品分类：" class="form-item">
                <el-cascader v-model="searchData.productCategoryName" :options="productTypeOptions" placeholder="请选择" clearable />
            </el-form-item>
            <el-form-item label="商品品牌：" class="form-item">
              <el-select v-model="searchData.brandName" value-key="id" class="m-2" 
                placeholder="请选择" size="large" clearable>
                <el-option
                  v-for="item in brandOptions"
                  :key="item.value"
                  :label="item.label"
                  :value="item"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="上架状态：" class="form-item">
              <el-select v-model="searchData.publishStatus" value-key="id" class="m-2" 
                placeholder="请选择" size="large" clearable @clear="onSearch">
                <el-option
                  v-for="item in publishStatusOptions"
                  :key="item.value"
                  :label="item.label"
                  :value="item"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="审核状态：" class="form-item">
              <el-select v-model="searchData.verifyStatus" value-key="id" class="m-2" 
                placeholder="请选择" size="large" clearable @clear="onSearch">
                <el-option
                  v-for="item in verifyStatusOptions"
                  :key="item.value"
                  :label="item.label"
                  :value="item"
                />
              </el-select>
            </el-form-item>

            <el-form-item class="form-item">
              <el-button type="primary" @click="onSearch">搜索结果</el-button>
            </el-form-item>
            <el-form-item class="form-item">
              <el-button  @click="onCancel">重置</el-button>
            </el-form-item>
            <el-form-item class="form-item">
              <el-button type="primary" @click="onExport">导出
                <el-icon class="el-icon--right"><Download /></el-icon>
              </el-button>
            </el-form-item>
          </el-form>
        </el-card>
      </div>
      <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="handleAddNewProduct">新增</el-button>
        </div>
        <el-table
          ref="multipleTable"
          :data="productData"
          style="width: 100%"
          tooltip-effect="dark"
          row-key="Id"
          @selection-change="handleSelectionChange"
        >
          <el-table-column type="selection" width="40" />
          <el-table-column align="left" label="编号" prop="ID" width="60"></el-table-column>
          <el-table-column align="left" label="商品图片" prop="pic" width="100" >
            <template #default="scope">
                <img  :src="scope.row.pic" alt="" style="width: 50px;height: 50px">
            </template>
          </el-table-column>
          <el-table-column align="left" label="商品名称" prop="name" width="180" />
          <el-table-column align="left" label="价格/货号" prop="priceSN" width="160" />
          <el-table-column align="left" label="标签" prop="status" width="140" >
            <template #default="scope">
                <el-switch
                    v-model="scope.row.publishStatus"
                    :active-value="1"
                    :inactive-value="0"
                    inactive-text="上架:">
                </el-switch>
                <br />
                <el-switch
                    v-model="scope.row.newStatus"
                    :active-value="1"
                    :inactive-value="0"
                    inactive-text="新品:">
                </el-switch>
                <br />
                <el-switch
                    v-model="scope.row.recommandStatus"
                    :active-value="1"
                    :inactive-value="0"
                    inactive-text="推荐:">
                </el-switch>
            </template>
          </el-table-column>
          <el-table-column align="left" label="排序" prop="sort" width="120" />
          <el-table-column align="left" label="SKU库存" prop="stock" width="120">
            <template #default="scope">
              <el-button type="primary" :icon="Edit" circle @click="handleEditStock(scope.row)"/>
            </template>
          </el-table-column>
          <el-table-column align="left" label="销量" prop="sale" width="160" />
          <el-table-column align="left" label="审核状态" prop="verifyStatus" width="80" />
          <el-table-column align="left" label="操作" min-width="160">
            <template #default="scope">
              <el-button type="primary" link icon="edit" @click="handleUpdateProduct(scope.row)">变更</el-button>
              <el-popover v-model="scope.row.visible" placement="top" width="160">
                <p>确定要删除吗？</p>
                <div style="text-align: right; margin-top: 8px;">
                  <el-button type="primary" link @click="scope.row.visible = false">取消</el-button>
                  <el-button type="primary" @click="handleDeleteProduct(scope.row)">确定</el-button>
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
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[5, 5, 5, 5]"
            :total="total"
            layout="total, sizes, prev, pager, next, jumper"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
          />
        </div>
      </div>
      <el-dialog v-model="isShowStockEditDialog" title="编辑货品信息">
        <el-row>
          <el-text>商品货号：{{dialogProductSN}}</el-text>
          <el-input
            v-model="selectedSkuCode"
            placeholder="按sku编号搜索"
            style="width: 220px;margin-left:100px;"
            size="small"
            clearable
            @clear="handleClearSkuCode"
          >
            <template #append>
              <el-button :icon="Search" @click="handleSearchSkuByCode"/>
            </template>
          </el-input>
        </el-row>
        <el-table :data="skuStockTableData"   border 
            :cell-style="{ textAlign: 'center' }" :header-cell-style="{ 'text-align': 'center' }" >
            <el-table-column v-for="(item, index) in skuColumsData" :key="index" :prop="item.prop" :label="item.label" :width="item.width" >
                <template #default="scope">
                    <div v-if="item.canEdit">
                        <el-input  v-model="scope.row[item.prop]"  autocomplete="off" @change="handleChangeInput(scope.row)"/>
                    </div>
                    <el-text v-else>{{scope.row[item.prop]}}</el-text>
                </template>
            </el-table-column>
        </el-table>

        <template #footer>
            <div class="dialog-footer">
            <el-button @click="handleCloseDialog">取 消</el-button>
            <el-button type="primary" @click="handleComfirmDialog">确 定</el-button>
            </div>
        </template>
      </el-dialog>
    </div>
  </template>
  
  <script setup lang="ts">
  import {
    getProductList,
    updateProducts,
    getProductSKUStockByProductID,
    getProductAttributeList,
    updateProductSKUStock,
  } from '@/api/product'
  import { ref, reactive, onBeforeMount, watch } from 'vue'
  import { ElMessage, ElTable } from 'element-plus'
  import { Edit, Search } from '@element-plus/icons-vue'
  import { ProductStore } from '@/pinia/modules/product'
  import { useRouter } from "vue-router";
  const router = useRouter()

  const productStore = ProductStore()
  const searchData = reactive({
    name: null,
    productSN: null,
    productCategoryName: null,
    brandName: null,
    publishStatus: null,
    verifyStatus: null,
  })
  const memberData = ref([])

  onBeforeMount(() => {
    getProductAttributeData()
    getProductBrandData()
  })
  const getProductAttributeData = async() => {
    await productStore.BuildProductAttributeData()
    productTypeOptions.value = productStore.ProductCategoryOptions
  }

  const getProductBrandData = async() => {
    await productStore.BuildBrandData()
    let productBrandList = productStore.RandData['list']
    
    brandOptions.value = productBrandList.map((item) => {
      return {value: item.id, label: item.name}
    })
  }

  const productTypeOptions = ref()
  const brandOptions = ref()
  const publishStatusOptions = [
    {
        value: '上架',
        label: '上架'
    },
    {
        value: '下架',
        label: '下架'
    },
  ]
  const verifyStatusOptions = [
    {
        value: '审核通过',
        label: '审核通过'
    },
    {
        value: '未审核',
        label: '未审核'
    },
  ]
  const onSearch = async() => {
    getTableData()
  }
  
  const onCancel = () => {
    getTableData()
  }
  
  const onExport = () => {
  
  }
  
  const page = ref(1)
  const total = ref(0)
  const pageSize = ref(5)
  const productData = ref([])
  
  // 分页
  const handleSizeChange = (val) => {
    pageSize.value = val
    getTableData()
  }
  
  const handleCurrentChange = (val) => {
    page.value = val
    getTableData()
  }
  
  // 查询
  const getTableData = async() => {
    const res = await getProductList({ name: searchData.name, productSN:searchData.productSN,
        productCategoryName:searchData.productCategoryName, brandName:searchData.brandName,
        publishStatus: searchData.publishStatus, verifyStatus:searchData.verifyStatus,
        page: page.value, pageSize: pageSize.value })
    if ('code' in res && res.code === 0) {
        productData.value = res.data.list
        productData.value.forEach(element => {
            element.priceSN = "价格：￥" + element.price + "\n 货号：" + element.productSN
        });
        console.log("----productData----", productData)
        total.value = res.data.total
        page.value = res.data.page
        pageSize.value = res.data.pageSize
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
    { id: 0, label: '商品上架', key: "publishstatus", dbKey: "publish_status", value: 1 },
    { id: 1, label: '商品下架', key: "publishStatus", dbKey: "publish_status",  value: 0 },
    { id: 2, label: '设为推荐', key: "recommandStatus", dbKey: "recommand_status",  value: 1 },
    { id: 3, label: '取消推荐', key: "recommandStatus", dbKey: "recommand_status", value: 0 },
    { id: 4, label: '设为新品', key: "newStatus", dbKey: "new_status", value: 1 },
    { id: 5, label: '取消新品', key: "newStatus", dbKey: "new_status", value: 0 },
    { id: 6, label: '移入回收站', key: "", dbKey: "", value: 0 },
  ]

  const multipleSelection = ref()
  const handleSelectionChange = (val) => {
    multipleSelection.value = val
  }
  var updateList:number[] = new Array() 
  const toggleSelection = async() => {
    if (!multipleSelection.value || multipleSelection.value.length < 1) return
    multipleSelection.value.forEach((item) => {
        item[stateOption.value.key] = stateOption.value.value
        updateList.push(item.id)
    })
    const res = await updateProducts({ids: updateList, key: stateOption.value.dbKey, value: stateOption.value.value })
    if ('code' in res && res.code !== 0) {
        productData.value.forEach(element => {
            element[stateOption.value.key] = stateOption.value.value
        });
    }
    updateList = []
  }
  
  const handleAddNewProduct = () => {
    const router = useRouter()
    router.push('NewProduct')
  }

  const dialogProductSN = ref("")
  const isShowStockEditDialog = ref(false)
  const skuColumsData = ref([])
  const skuStockTableData = ref([])
  const selectedProductId = ref(0)
  const selectedSkuCode = ref("")

  const handleEditStock = async(item) => {
    isShowStockEditDialog.value = true
    console.log("------item----", item)
    dialogProductSN.value = item.productSN
    selectedProductId.value = item.id
    skuColumsData.value = []
    skuStockTableData.value = []
    selectedSkuCode.value = ""
    const attributeRes = await getProductAttributeList({tag: item.productAttributeCategoryId, state: 0})
    if ('code' in attributeRes && attributeRes.code === 0) {
      console.log("---skuStockTableData-", attributeRes.data)
      skuColumsData.value = [{label: 'SKU编号', prop: 'skuCode', width: "120px", canEdit: true},]
      attributeRes.data.list.forEach(element => {
        let attributeitem = {label: element.name, prop: element.name, width: "120px"}
        skuColumsData.value.push(attributeitem)
      });
      skuColumsData.value.push(
        {label: '销售价格', prop: 'price', width: "100px", canEdit: true},
        {label: '促销价格', prop: 'promotionPrice', width: "100px", canEdit: true},
        {label: '商品库存', prop: 'stock', width: "100px", canEdit: true},
        {label: '库存预警值', prop: 'lowStock', width: "120px", canEdit: true},
      )
    }
    const res = await getProductSKUStockByProductID({id: item.id})
    if ('code' in res && res.code === 0) {
      console.log("---skuStockTableData-", res.data)
      parseSkuStockTable(res.data)
    }
  }
  const parseSkuStockTable = (data) => {
    data.forEach(element => {
        let attributeJson = JSON.parse(element.spData)
        let skuStock = {}
        skuStock["skuCode"] = element.skuCode
        attributeJson.forEach(data => {
          skuStock[data.key] = data.value
        });
        skuStock["price"] = element.price
        skuStock["promotionPrice"] = element.promotionPrice
        skuStock["stock"] = element.stock
        skuStock["lowStock"] = element.lowStock
        skuStock["spData"] = element.spData
        skuStockTableData.value.push(skuStock)
      });
  }
  const handleCloseDialog = () => {
    isShowStockEditDialog.value = false
  }
  const hadChange = ref(false)
  const handleChangeInput = (item) => {
    hadChange.value = true
  }

  const handleComfirmDialog = async() => {
    isShowStockEditDialog.value = false
    if (hadChange) {
      const res = await updateProductSKUStock(skuStockTableData.value)
      if ('code' in res && res.code !== 0) {
        ElMessage({
          type: 'error',
          message: '更新成功'
        })
      }
    }
  }
  const handleSearchSkuByCode = async (params) => {
    skuStockTableData.value = []
    const res = await getProductSKUStockByProductID({id: selectedProductId.value, keyword: selectedSkuCode.value})
    if ('code' in res && res.code === 0) {
      console.log("---skuStockTableData-", res.data)
      parseSkuStockTable(res.data)
    }
  }
  const handleClearSkuCode = () => {
    selectedSkuCode.value = ""
  }
  
  const handleUpdateProduct = (row) => {
    router.push({ path: 'updateProduct', query: {id: row.id}})
  }
  const handleDeleteProduct = (row) => {

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
  