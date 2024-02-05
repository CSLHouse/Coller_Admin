<template>
  <div>
    <div>
      <el-card class="box-card" shadow="never">
        <div>
          <!-- <i class="el-icon-search"></i> -->
          <span>筛选搜索</span>
          <el-button
            style="float: right"
            @click="handleSearchList()"
            type="primary"
            size="small">
            查询结果
          </el-button>
          <el-button
            style="float: right;margin-right: 15px"
            @click="handleResetSearch()"
            size="small">
            重置
          </el-button>
        </div>
        <div style="margin-top: 15px">

        </div>

        <el-form :inline="true" :model="searchData" class="demo-form-inline" size="small" label-width="140px">
          <el-form-item label="商品名称：" class="form-item">
            <el-input v-model.number="searchData.keyword" placeholder="商品名称" clearable />
          </el-form-item>
          <el-form-item label="商品货号：" class="form-item">
            <el-input v-model.number="searchData.productSN" placeholder="商品货号" clearable />
          </el-form-item>
          <el-form-item label="商品分类：" class="form-item">
              <el-cascader v-model="searchData.productCategoryId" :options="productTypeOptions" 
              placeholder="请选择" clearable @change="handleCategoryIdChange"/>
          </el-form-item>
          <el-form-item label="商品品牌：" class="form-item">
            <el-select v-model="searchData.brandId" value-key="id"
              placeholder="请选择" size="default" clearable>
              <el-option
                v-for="item in brandOptions"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="上架状态：" class="form-item">
            <el-select v-model="searchData.publishStatus" value-key="id"
              placeholder="全部" size="default" clearable>
              <el-option
                v-for="item in publishStatusOptions"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="审核状态：" class="form-item">
            <el-select v-model="searchData.verifyStatus" value-key="id"
              placeholder="全部" size="default" clearable>
              <el-option
                v-for="item in verifyStatusOptions"
                :key="item.value"
                :label="item.label"
                :value="item"
              />
            </el-select>
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
        border
      >
        <el-table-column type="selection" width="40" />
        <el-table-column align="left" label="编号" prop="id" width="60"></el-table-column>
        <el-table-column align="left" label="商品图片" prop="pic" width="100" >
          <template #default="scope">
              <img  :src="scope.row.pic" alt="" style="width: 50px;height: 50px">
          </template>
        </el-table-column>
        <el-table-column align="left" label="商品名称" prop="name" width="180">
          <template #default="scope">
            <p>{{scope.row.name}}</p>
            <p>品牌：{{scope.row.brandName}}</p>
          </template>
        </el-table-column>
        <el-table-column align="left" label="价格/货号" prop="priceSN" width="140">
          <template #default="scope">
            <p>价格：￥{{scope.row.price}}</p>
            <p>货号：{{scope.row.productSn}}</p>
          </template>
        </el-table-column>
        <el-table-column align="left" label="标签" prop="status" width="120" >
          <template #default="scope">
              <el-switch
                  @change="handlePublishStatusChange(scope.$index, scope.row)"
                  v-model="scope.row.publishStatus"
                  :active-value="1"
                  :inactive-value="0"
                  inactive-text="上架:">
              </el-switch>
              <br />
              <el-switch
                  @change="handleNewStatusChange(scope.$index, scope.row)"
                  v-model="scope.row.newStatus"
                  :active-value="1"
                  :inactive-value="0"
                  inactive-text="新品:">
              </el-switch>
              <br />
              <el-switch
                  @change="handleRecommendStatusChange(scope.$index, scope.row)"
                  v-model="scope.row.recommandStatus"
                  :active-value="1"
                  :inactive-value="0"
                  inactive-text="推荐:">
              </el-switch>
          </template>
        </el-table-column>
        <el-table-column align="left" label="排序" prop="sort" width="60" />
        <el-table-column align="left" label="SKU库存" prop="stock" width="120">
          <template #default="scope">
            <el-button type="primary" :icon="Edit" circle @click="handleEditStock(scope.row)"/>
          </template>
        </el-table-column>
        <el-table-column align="left" label="销量" prop="sale" width="60" />
        <el-table-column align="left" label="审核状态" prop="verifyStatus" width="100" />
        <el-table-column align="left" label="操作" min-width="120">
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
      <div>
        <div style="margin-top: 15px;float: left;">
          <el-select v-model="stateOption" value-key="id" class="m-2" placeholder="批量操作" size="large">
              <el-option
                  v-for="item in operates"
                  :key="item.key"
                  :label="item.label"
                  :value="item.key"
              />
          </el-select>
          <el-button style="margin-left: 20px" type="primary" @click="toggleSelection()">确定</el-button>
        </div>
        <div style="float: right;">
          <el-pagination
            background
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[5, 5, 5, 5]"
            :total.number="+total"
            layout="total, sizes, prev, pager, next, jumper"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
          />
        </div>
      </div>
    </div>
    <el-dialog v-model="editSkuInfo.dialogVisible" title="编辑货品信息" width="40%">
      <el-row>
        <el-text>商品货号：{{editSkuInfo.productSn}}</el-text>
        <el-input
          v-model="editSkuInfo.keyword"
          placeholder="按sku编号搜索"
          style="width: 50%;margin-left: 20px"
          size="small"
          clearable
          @clear="handleClearSkuCode"
        >
          <template #append>
            <el-button :icon="Search" @click="handleSearchSkuByCode"/>
          </template>
        </el-input>
      </el-row>
      <el-table :data="editSkuInfo.stockList"  border  style="width: 100%;margin-top: 20px"
          :cell-style="{ textAlign: 'center' }" :header-cell-style="{ 'text-align': 'center' }" >
          
          <el-table-column
            label="SKU编号"
            align="center">
            <template #default="scope">
              <el-input v-model="scope.row.skuCode"></el-input>
            </template>
          </el-table-column>

          <el-table-column v-for="(item, index) in editSkuInfo.productAttr"
            :label="item.name" :key="item.id" align="center">
              <template #default="scope">
                  {{getProductSkuSp(scope.row, index)}}
              </template>
          </el-table-column>
          <el-table-column
            label="销售价格"
            width="80"
            align="center">
            <template #default="scope">
              <el-input v-model="scope.row.price"></el-input>
            </template>
          </el-table-column>
          <el-table-column
            label="商品库存"
            width="80"
            align="center">
            <template #default="scope">
              <el-input v-model="scope.row.stock"></el-input>
            </template>
          </el-table-column>
          <el-table-column
            label="库存预警值"
            width="100"
            align="center">
            <template #default="scope">
              <el-input v-model="scope.row.lowStock"></el-input>
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
  updateProductKeyword,
  getProductSKUStockByProductID,
  getProductAttributeList,
  updateProductSKUStock,
  deleteProducts,
} from '../../../api/product'
import { ref, reactive, onBeforeMount, watch } from 'vue'
import { ElMessage, ElMessageBox, ElTable } from 'element-plus'
import { Edit, Search } from '@element-plus/icons-vue'
import { ProductStore } from '../../../pinia/modules/product'
import { useRouter } from "vue-router";
import { el } from 'element-plus/es/locale'
const router = useRouter()

const productStore = ProductStore()
const searchData = reactive({
  keyword: null,
  productSN: null,
  productCategoryId: null,
  brandId: null,
  publishStatus: null,
  verifyStatus: null,
})

onBeforeMount(() => {
  getProductCategoryData()
  getProductBrandData()
})
const getProductCategoryData = async() => {
  await productStore.BuildProductCategoryData()
  productTypeOptions.value = productStore.ProductCategoryOptions
}

const handleCategoryIdChange = () => {
  searchData.productCategoryId = searchData.productCategoryId.at(-1)
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
      value: 101,
      label: '上架'
  },
  {
      value: 100,
      label: '下架'
  },
]
const verifyStatusOptions = [
  {
      value: 101,
      label: '审核通过'
  },
  {
      value: 100,
      label: '未审核'
  },
]
const handleSearchList = async() => {
  getTableData()
}

const handleResetSearch = () => {
  searchData.keyword = null
  searchData.productSN = null
  searchData.productCategoryId = null
  searchData.brandId = null
  searchData.publishStatus = null
  searchData.verifyStatus = null
  getTableData()
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
  const res = await getProductList({ keyword: searchData.keyword, productSN:searchData.productSN,
      productCategoryId:searchData.productCategoryId, brandId:searchData.brandId,
      publishStatus: searchData.publishStatus, verifyStatus:searchData.verifyStatus,
      page: page.value, pageSize: pageSize.value })
  if ('code' in res && res.code === 0) {
      productData.value = res.data.list
      console.log("----productData----", productData)
      total.value = res.data.total
      page.value = res.data.page
      pageSize.value = res.data.pageSize
  }
}

getTableData()

const stateOption = ref()
const operates = [
  {
    label: "商品上架",
    key: "publishOn",
    dbKey: "publish_status", 
    value: 1
  },
  {
    label: "商品下架",
    key: "publishOff",
    dbKey: "publish_status",
    value: 0
  },
  {
    label: "设为推荐",
    key: "recommendOn",
    dbKey: "recommand_status",
    value: 1
  },
  {
    label: "取消推荐",
    key: "recommendOff",
    dbKey: "recommand_status",
    value: 0
  },
  {
    label: "设为新品",
    key: "newOn",
    dbKey: "new_status",
    value: 1
  },
  {
    label: "取消新品",
    key: "newOff",
    dbKey:"new_status",
    value: 0
  },
  {
    label: "移入回收站",
    key: "recycle",
    dbKey: "",
    value: 0
  }
]

const multipleSelection = ref()
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}
var updateList:number[] = new Array() 
const toggleSelection = async() => {
  if (!multipleSelection.value || multipleSelection.value.length < 1){
    ElMessage({
      message: '请选择要操作的商品',
      type: 'warning',
      duration: 1000
    });
    return
  }
  multipleSelection.value.forEach((item) => {
      updateList.push(item.id)
  })
  if (stateOption.value == 'recycle') {
    const res = await deleteProducts({ids: updateList})
    if ('code' in res && res.code === 0) {
      ElMessage({
        message: '删除成功',
        type: 'success',
        duration: 1000
      });
      getTableData()
    }
  } else {
    let dbkey
    let value
    for (let index = 0; index < operates.length; index++) {
      const element = operates[index];
      if (element.key == stateOption.value) {
        dbkey = element.dbKey
        value = element.value
        break
      }
    }
    const res = await updateProductKeyword({ids: updateList, key: dbkey, value: value })
    if ('code' in res && res.code === 0) {
      ElMessage({
        message: '修改成功',
        type: 'success',
        duration: 1000
      });
      getTableData()
    }
  }
  updateList = []
}
const handlePublishStatusChange = async(_, row) => {
  let ids:number[] = new Array();
  ids.push(row.id);
  const res = updateProductKeyword({ids: ids, key: 'publish_status', value: row.publishStatus })
  if ('code' in res && res.code === 0) {
    ElMessage({
      message: '修改成功',
      type: 'success',
      duration: 1000
    });
  }
}

const handleNewStatusChange = async(_, row) => {
  let ids:number[] = new Array();
  ids.push(row.id);
  const res = await updateProductKeyword({ids: ids, key: 'new_status', value: row.newStatus })
  if ('code' in res && res.code === 0) {
    ElMessage({
      message: '修改成功',
      type: 'success',
      duration: 1000
    });
  }
}

const handleRecommendStatusChange = async(_, row) => {
  let ids:number[] = new Array();
  ids.push(row.id);
  const res = await updateProductKeyword({ids: ids, key: 'recommand_status', value: row.recommandStatus })
  if ('code' in res && res.code === 0) {
    ElMessage({
      message: '修改成功',
      type: 'success',
      duration: 1000
    });
  }
}

const handleAddNewProduct = () => {
  router.push({path: '/layout/product/newProduct'});
}

const editSkuInfo = ref({
  dialogVisible:false,
  productId:null,
  productSn:'',
  productAttributeCategoryId:null,
  stockList:[],
  productAttr:[],
  keyword:null
})
const getProductSkuSp = (row, index) => {
  let spData = JSON.parse(row.spData);
  if(spData !=null && index < spData.length){
    return spData[index].value;
  }else{
    return null;
  }
}

const handleEditStock = async(row) => {
  console.log("--", row)
  editSkuInfo.value.dialogVisible = true;
  editSkuInfo.value.productId = row.id;
  editSkuInfo.value.productSn = row.productSn;
  editSkuInfo.value.productAttributeCategoryId = row.productAttributeCategoryId;
  editSkuInfo.value.keyword = null;
  const res = await getProductSKUStockByProductID({id: row.id})
  if ('code' in res && res.code === 0) {
    console.log("---skuStockTableData-", res.data)
    editSkuInfo.value.stockList = res.data
  } 
  
  if(row.productAttributeCategoryId != null){
    const attributeRes = await getProductAttributeList({tag: row.productAttributeCategoryId, state: 0})
    if ('code' in attributeRes && attributeRes.code === 0) {
      console.log("---skuStockTableData-", attributeRes.data)
      editSkuInfo.value.productAttr = attributeRes.data.list
    }
  }
}

const handleCloseDialog = () => {
  editSkuInfo.value.dialogVisible = false;
}
const hadChange = ref(false)
const handleChangeInput = (item) => {
  hadChange.value = true
}

const handleComfirmDialog = async() => {
  if(editSkuInfo.value.stockList == null||editSkuInfo.value.stockList.length <= 0){
    ElMessage({
      message: '暂无sku信息',
      type: 'warning',
      duration: 1000
    });
    return
  }
  ElMessageBox.confirm('是否要进行修改', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    const res = updateProductSKUStock(editSkuInfo.value.stockList)
    if ('code' in res && res.code !== 0) {
      ElMessage({
        type: 'success',
        message: '更新成功'
      })
      editSkuInfo.value.dialogVisible = false;
    }
  })
}
const handleSearchSkuByCode = async (params) => {
  editSkuInfo.value.stockList = []
  const res = await getProductSKUStockByProductID({id: editSkuInfo.value.productId, keyword: editSkuInfo.value.keyword})
  if ('code' in res && res.code === 0) {
    console.log("---skuStockTableData-", res.data)
    editSkuInfo.value.stockList = res.data
  }
}
const handleClearSkuCode = () => {
  editSkuInfo.value.keyword = null
}

const handleUpdateProduct = (row) => {
  router.push({path: '/layout/product/updateProduct', query:{id: row.id}});
}
const handleDeleteProduct = async(row) => {
  const res = await deleteProducts({ids: [row.id]})
    if ('code' in res && res.code === 0) {
      ElMessage({
        message: '删除成功',
        type: 'success',
        duration: 1000
      });
    }
    getTableData()
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
