<template>
    <div>
      <div>
        <el-card class="box-card" title="筛选搜索">
          <el-form :inline="true" :model="searchData" class="demo-form-inline">
            <el-form-item label="输入搜索：" class="form-item">
              <el-input v-model.number="searchData.name" placeholder="品牌名称/关键字" clearable />
            </el-form-item>
          </el-form>
          <el-button type="primary" @click="onSearch">搜索结果</el-button>
        </el-card>
      </div>
      <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
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
          <el-table-column align="left" label="商品名称" prop="name" width="120" />
          <el-table-column align="left" label="品牌首字母" prop="firstLetter" width="100" />
          <el-table-column align="left" label="排序" prop="sort" width="60" />
          <el-table-column align="left" label="品牌制造商" prop="status" width="100" >
            <template #default="scope">
                <el-switch
                    v-model="scope.row.factoryStatus"
                    :active-value="1"
                    :inactive-value="0">
                </el-switch>
            </template>
          </el-table-column>
          <el-table-column align="left" label="是否显示" prop="status" width="100" >
            <template #default="scope">
                <el-switch
                    v-model="scope.row.showStatus"
                    :active-value="1"
                    :inactive-value="0">
                </el-switch>
            </template>
          </el-table-column>
          
          <el-table-column align="left" label="相关" prop="content" width="180" />
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
      <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="编辑品牌">
        <el-form  :model="productForm" label-width="140px" >
            <el-form-item label="品牌名称">
                <el-input v-model="productForm.name" autocomplete="off" />
            </el-form-item>
            <el-form-item label="品牌首字母">
                <el-input v-model.number="productForm.firstLetter" autocomplete="off" />
            </el-form-item>
            <el-form-item label="品牌LOGO">
                <el-upload
                    list-type="picture"
                    :limit=1
                    :action="`${path}/fileUploadAndDownload/upload`"
                    :headers="{ 'x-token': userStore.token }"
                    :on-preview="handlePreview"
                    :on-remove="handleRemove"
                    :on-error="uploadError"
                    :on-success="uploadLogoSuccess"
                    :before-upload="beforeAvatarUpload">
                    <el-button type="primary">点击上传</el-button>
                    <template #tip>
                        <div class="el-upload__tip">
                            只能上传jpg/png文件，且不超过10MB
                        </div>
                    </template>
                </el-upload>
            </el-form-item>
            <el-form-item label="品牌专区大图">
                <el-upload
                    list-type="picture"
                    :limit=1
                    :action="`${path}/fileUploadAndDownload/upload`"
                    :headers="{ 'x-token': userStore.token }"
                    :on-preview="handlePreview"
                    :on-remove="handleRemove"
                    :on-error="uploadError"
                    :on-success="uploadBigPicSuccess"
                    :before-upload="beforeAvatarUpload">
                    <el-button type="primary">点击上传</el-button>
                    <template #tip>
                        <div class="el-upload__tip">
                            只能上传jpg/png文件，且不超过10MB
                        </div>
                    </template>
                </el-upload>
            </el-form-item>
            <el-form-item label="品牌故事">
                <el-input v-model="productForm.brandStory" autocomplete="off" />
            </el-form-item>
            <el-form-item label="排序">
                <el-input v-model.number="productForm.sort" autocomplete="off" />
            </el-form-item>
            <el-form-item label="是否显示">
                <el-radio-group v-model="showState" class="ml-4" @change="HandleShowRadioChanged">
                    <el-radio label="1" size="large">是</el-radio>
                    <el-radio label="0" size="large">否</el-radio>
                </el-radio-group>
            </el-form-item>
            <el-form-item label="品牌制造商">
                <el-radio-group v-model="brandState" class="ml-4" @change="HandleBrandRadioChanged">
                    <el-radio label="1" size="large">是</el-radio>
                    <el-radio label="0" size="large">否</el-radio>
                </el-radio-group>
            </el-form-item>
        </el-form>
        <template #footer>
            <div class="dialog-footer">
            <el-button @click="closeDialog">取 消</el-button>
            <el-button type="primary" @click="enterDialog">确 定</el-button>
            </div>
        </template>
        </el-dialog>
    </div>
  </template>
  
  <script setup lang="ts">
  import {
    createProductBrand,
    updateProductBrand,
  } from '@/api/product'
  import { ref, reactive, onBeforeMount, watch } from 'vue'
  import { ElMessage } from 'element-plus'
  import type { UploadProps, UploadUserFile, UploadInstance  } from 'element-plus'
  import config from '@/core/config'
  import { trim } from '@/utils/stringFun'
  import { number } from 'echarts'
  import { useRouter } from "vue-router";
  import { ProductStore } from '@/pinia/modules/product'  
  import {getAliOSSCreds} from '@/api/system'
  import OSS from 'ali-oss'
  import { useUserStore } from '@/pinia/modules/user'
  import { createID } from '@/utils/uuid'

  const userStore = useUserStore()

  const searchData = reactive({
    name: null,
    
  })
  
  const onSearch = async() => {
    // TODO: 根据品牌名或关键词搜索
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
  
  const productStore = ProductStore()
  // 查询
  const getTableData = async() => {
    productData.value = productStore.RandData['list']
    if (!productData.value) {
      await productStore.BuildBrandData()
      productData.value = productStore.RandData['list']
      productData.value.forEach(element => {
        element.content = "商品：" + element.productCount + "  评价：" + element.productCommentCount
      });
      total.value = productStore.RandData['total']
      page.value = productStore.RandData['page']
      pageSize.value = productStore.RandData['pageSize']
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
    { id: 0, label: '显示品牌', key: "showStatus", dbKey: "show_status", value: 1 },
    { id: 1, label: '隐藏品牌', key: "showStatus", dbKey: "show_status",  value: 0 },
  ]

  const multipleSelection = ref()
  const handleSelectionChange = (val) => {
    multipleSelection.value = val
  }
  var updateList:number[] = new Array() 
  const toggleSelection = async() => {
    // if (!multipleSelection.value || multipleSelection.value.length < 1) return
    // multipleSelection.value.forEach((item) => {
    //     item[stateOption.value.key] = stateOption.value.value
    //     updateList.push(item.id)
    // })
    // const res = await updateProducts({ids: updateList, key: stateOption.value.dbKey, value: stateOption.value.value })
    // if ('code' in res && res.code !== 0) {
    //     productData.value.forEach(element => {
    //         element[stateOption.value.key] = stateOption.value.value
    //     });
    // }
    // updateList = []
  }
  
  const type = ref('')
  const dialogFormVisible = ref(false)
  const productForm = ref({
    name: '',
    firstLetter: '',
    sort: 0,
    factoryStatus: 0,
    showStatus: 0,
    logo: '',
    bigPic: '',
    brandStory: '',
  })
  const updateProduct = async(row) => {
    dialogFormVisible.value = true
    type.value = 'update'
    productForm.value = row
  }
  const openDialog = () => {
    dialogFormVisible.value = true
    type.value = 'create'
  }

  const enterDialog = async() => {
    let res
    switch (type.value) {
        case 'create':
            res = await createProductBrand(productForm.value)
            break
        case 'update':
            res = await updateProductBrand(productForm.value)
            break
        default:
            res = await createProductBrand(productForm.value)
            break
    }

    if (res.code === 0) {
        closeDialog()
        await productStore.BuildBrandData(true)
        getTableData()
    }
  }
  const closeDialog = () => {
    dialogFormVisible.value = false
  }

  const uploadRef = ref<UploadInstance>()
  const path = ref(import.meta.env.VITE_BASE_API)
  // 图片上传前验证
  const beforeAvatarUpload = (file) => {
    const isJPG = file.type === 'image/jpeg'
    const isPng = file.type === 'image/png'
    const isLt2M = file.size / 1024 / 1024 < 2
    if (!isJPG && !isPng) {
        ElMessage.error('上传图片只能是 jpg或png 格式!')
    }
    if (!isLt2M) {
        ElMessage({
          type: 'error',
          message: '上传头像图片大小不能超过 2MB!'
        })
    }
    return (isPng || isJPG) && isLt2M
  }

  const upload = async(item) => {
    // let res = await getAliOSSCreds()
    // let creds = res.data.data
    // TODO: stsToken上传更安全
    const client = new OSS({
        region: 'oss-cn-beijing', // 服务器集群地区
        accessKeyId: "LTAI5tJcdwf4dTPNt1e5xQke",// creds.accessKeyId, // OSS帐号
        accessKeySecret: "nHEtlOvJ1wqWMXBiZheJh083YFZKmr",//creds.accessKeySecret, // OSS 密码
        // stsToken: "acs:ram::1573486857476997:role/ramosstest",//creds.securityToken, // 签名token
        bucket: 'cooller' // 阿里云上存储的 Bucket
    })
    let imageID = createID()
    let key = 'resource/' + userStore.userInfo.nickName + 'images/' + imageID + '.jpg'  // 存储路径，并且给图片改成唯一名字
    return client.put(key, item.file) // OSS上传
  }
  const uploadError = () => {
    ElMessage({
        type: 'error',
        message: '图片上传失败'
    })
  }
  const uploadLogoSuccess = (res) => {
    const { data } = res
    if (data.file) {
        productForm.value.logo = data.file.url
        console.log("-----logo.url-----", data.file.url)
    }
  }
  
  const uploadBigPicSuccess = (res) => {
    const { data } = res
    if (data.file) {
        productForm.value.bigPic = data.file.url
        console.log("-----bigPic.url-----", data.file.url)
    }
  }
  const handleRemove: UploadProps['onRemove'] = (uploadFile, uploadFiles) => {
    console.log(uploadFile, uploadFiles)
  }
  const handlePreview: UploadProps['onPreview'] = (file) => {
    console.log(file)
  }

  const showState = ref("1")
  const brandState = ref("1")
  const HandleShowRadioChanged = () => {
    productForm.value.showStatus = Number(showState.value)
  }
  const HandleBrandRadioChanged = () => {
    productForm.value.factoryStatus = Number(brandState.value)
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
  