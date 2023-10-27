<template>
    <el-card>
        <el-steps :active="active" finish-status="success">
            <el-step :title="item" v-for="item in titleList" :key="item"/>
        </el-steps>
        <div class="h-6" />
        <el-form
            ref="ruleFormRef"
            :model="productForm"
            status-icon
            label-width="120px"
            class="demo-ruleForm"
            style="max-width: 860px;margin-left: 20px;"
            >
            <div v-if="active==0">
                <el-form-item label="商品分类" prop="productCategoryName">
                    <el-cascader v-model="productType" :options="productCategoryOptions" 
                        placeholder="请选择" clearable @change="handleProductTypeChange"/>
                </el-form-item>
                <el-form-item label="商品名称" prop="name">
                    <el-input v-model="productForm.name" type="tel" autocomplete="off" />
                </el-form-item>
                <el-form-item label="副标题" prop="subTitle">
                    <el-input v-model="productForm.subTitle" type="text" autocomplete="off" />
                </el-form-item>
                <el-form-item label="商品品牌" prop="comboId">
                    <el-select v-model="productBrandOption" class="m-2" placeholder="请选择品牌" size="large">
                        <el-option
                            v-for="item in productBrandOptions"
                            :key="item.key"
                            :label="item.value"
                            :value="item"
                        />
                    </el-select>
                </el-form-item>
                <el-form-item label="商品介绍" prop="gift">
                    <el-input v-model="productForm.description" type="textarea" autocomplete="off" />
                </el-form-item>
                <el-form-item label="商品货号" prop="collection">
                    <el-input v-model="productForm.productSN" type="text" autocomplete="off" />
                </el-form-item>
                <el-form-item label="商品售价" prop="date">
                    <el-input-number v-model="productForm.price" :precision="2"></el-input-number>
                </el-form-item>
                <el-form-item label="市场价" prop="cardId">
                    <el-input-number v-model="productForm.originalPrice" :precision="2"></el-input-number>
                </el-form-item>
                <el-form-item label="商品库存" prop="telephone">
                    <el-input v-model.number="productForm.stock" type="number" autocomplete="off" style="width: 100px"/>
                </el-form-item>
                <el-form-item label="计量单位" prop="name">
                    <el-input v-model="productForm.unit" type="text" autocomplete="off" />
                </el-form-item>
                <el-form-item label="商品重量" prop="name">
                    <el-col :span="8">
                        <el-input-number v-model="productForm.weight" :precision="2"></el-input-number>
                    </el-col>
                    <a style=" margin-left: 10px;">克</a>
                </el-form-item>
                <el-form-item label="排序" prop="name">
                    <el-input v-model="productForm.sort" autocomplete="off" />
                </el-form-item>
            </div>
            <div v-if="active==1">
                <el-form-item label="赠送积分" prop="date" >
                    <el-input v-model.number="productForm.giftPoint" type="number" autocomplete="off" style="width: 100px"/>
                </el-form-item>
                <el-form-item label="赠送成长值" prop="cardId">
                    <el-input v-model.number="productForm.giftGrowth" type="number" autocomplete="off" style="width: 100px"/>
                </el-form-item>
                <el-form-item label="积分购买限制" prop="telephone">
                    <el-input v-model.number="productForm.usePointLimit" type="number" autocomplete="off" style="width: 100px"/>
                </el-form-item>
                <el-form-item label="预告商品" prop="telephone">
                    <el-switch
                        v-model="productForm.previewStatus"
                        :active-value="1"
                        :inactive-value="0">
                    </el-switch>
                </el-form-item>
                <el-form-item label="商品上架" prop="telephone">
                    <el-switch
                        v-model="productForm.publishStatus"
                        :active-value="1"
                        :inactive-value="0">
                    </el-switch>
                </el-form-item>
                <el-form-item label="商品推荐" prop="telephone">
                    <el-switch
                        v-model="productForm.newStatus"
                        :active-value="1"
                        :inactive-value="0"
                        inactive-text="新品:">
                    </el-switch>
                    <el-switch
                        v-model="productForm.recommandStatus"
                        :active-value="1"
                        :inactive-value="0"
                        inactive-text="推荐:">
                    </el-switch>
                </el-form-item>
                <el-form-item label="服务保证" prop="telephone">
                    <el-checkbox-group v-model="checkList" @change="HandleServiceIdsRadioChanged">
                        <el-checkbox label="无忧退货" />
                        <el-checkbox label="快速退款" />
                        <el-checkbox label="免费包邮" />
                    </el-checkbox-group>
                </el-form-item>
                <el-form-item label="详细页标题" prop="name">
                    <el-input v-model="productForm.detailTitle" type="text" autocomplete="off" />
                </el-form-item>
                <el-form-item label="详细页描述" prop="name">
                    <el-input v-model="productForm.detailDesc" type="text" autocomplete="off" />
                </el-form-item>
                <el-form-item label="商品关键字" prop="name">
                    <el-input v-model="productForm.keywords" type="text" autocomplete="off" />
                </el-form-item>
                <el-form-item label="商品备注" prop="gift">
                    <el-input v-model="productForm.note" type="textarea" autocomplete="off" />
                </el-form-item>
                <el-form-item label="选择优惠方式" prop="name">
                    <el-radio-group v-model="promotionTypeState" @change="HandlePromotionTypeRadioChanged">
                        <el-radio-button label="0" size="large">无优惠</el-radio-button>
                        <el-radio-button label="1" size="large">特惠促销</el-radio-button>
                        <el-radio-button label="2" size="large">会员价格</el-radio-button>
                        <el-radio-button label="3" size="large">阶梯价格</el-radio-button>
                        <el-radio-button label="4" size="large">满减价格</el-radio-button>
                        <el-radio-button label="5" size="large">限时购</el-radio-button>
                    </el-radio-group>
                    <el-row>
                        
                    </el-row>
                    <div v-if="promotionTypeState == 1">
                        <el-form
                            ref="ruleFormRef"
                            :model="productForm"
                            status-icon
                            label-width="120px"
                            class="demo-ruleForm"
                            style="max-width: 760px;"
                            >
                            <el-form-item label="开始时间" prop="gift" style="margin-top:10px;">
                                <el-date-picker
                                    v-model="productForm.promotionStartTime"
                                    type="datetime"
                                    placeholder="选择开始时间"
                                    value-format="YYYY-MM-DD h:m:s"
                                />
                            </el-form-item>
                            <el-form-item label="结束时间" prop="gift" style="margin-top:10px;">
                                <el-date-picker
                                    v-model="productForm.promotionEndTime"
                                    type="datetime"
                                    placeholder="选择结束时间"
                                    value-format="YYYY-MM-DD h:m:s"
                                />
                            </el-form-item>
                            <el-form-item label="促销价格" prop="gift" style="margin-top:10px;">
                                <el-input-number v-model="productForm.promotionPrice" :precision="2" placeholder="输入促销价格"></el-input-number>
                            </el-form-item>
                        </el-form>
                    </div>
                    <div v-if="promotionTypeState == 2">
                        <el-form
                            ref="ruleFormRef"
                            :model="item"
                            v-for="(item, index) in memberPriceList"
                            :key="index"
                            status-icon
                            label-width="120px"
                            class="demo-ruleForm"
                            style="max-width: 760px;"
                            >
                            <el-form-item :label="item.memberLevelName" prop="gift" style="margin-top:10px;">
                                <el-input-number v-model="item.memberPrice" :precision="2"></el-input-number>
                            </el-form-item>

                            <!-- <el-form-item label="黄金会员" prop="gift" style="margin-top:10px;">
                                <el-input-number v-model="item.promotionPrice" :precision="2"></el-input-number>
                            </el-form-item>
                            <el-form-item label="白金会员" prop="gift" style="margin-top:10px;">
                                <el-input-number v-model="item.promotionPrice" :precision="2"></el-input-number>
                            </el-form-item>
                            <el-form-item label="钻石会员" prop="gift" style="margin-top:10px;">
                                <el-input-number v-model="item.promotionPrice" :precision="2"></el-input-number>
                            </el-form-item> -->
                        </el-form>
                    </div>
                    <div v-if="promotionTypeState == 3">
                        <el-table :data="productLadderList" style="width: 600px;margin-top: 10px;">
                            <el-table-column label="数量" >
                                <template #default="scope">
                                    <div style="display: flex; align-items: center">
                                        <el-input-number v-model="scope.row.count" :precision="2"></el-input-number>
                                    </div>
                                </template>
                            </el-table-column>
                            <el-table-column label="折扣" >
                                <template #default="scope">
                                    <div style="display: flex; align-items: center">
                                        <el-input-number v-model="scope.row.discount" :precision="2"></el-input-number>
                                    </div>
                                </template>
                            </el-table-column>
                            <el-table-column label="操作">
                                <template #default="scope">
                                    <el-button size="small" @click="handleLadderDelete(scope.$index, scope.row)" >删除</el-button >
                                    <el-button size="small" @click="handleLadderAdd(scope.$index, scope.row)" >添加</el-button >
                                </template>
                            </el-table-column>
                        </el-table>
                    </div>
                    <div v-if="promotionTypeState == 4">
                        <el-table :data="productFullReductionList" style="width: 600px;margin-top: 10px;">
                            <el-table-column label="满" >
                                <template #default="scope">
                                    <div style="display: flex; align-items: center">
                                        <el-input-number v-model="scope.row.fullPrice" :precision="2"></el-input-number>
                                    </div>
                                </template>
                            </el-table-column>
                            <el-table-column label="立减" >
                                <template #default="scope">
                                    <div style="display: flex; align-items: center">
                                        <el-input-number v-model="scope.row.reducePrice" :precision="2"></el-input-number>
                                    </div>
                                </template>
                            </el-table-column>
                            <el-table-column label="操作">
                                <template #default="scope">
                                    <el-button
                                        size="small"
                                        @click="handleFullDiscountDelete(scope.$index, scope.row)" >删除</el-button >
                                    <el-button size="small" @click="handleFullDiscountAdd(scope.$index, scope.row)" >添加</el-button >
                                </template>
                            </el-table-column>
                        </el-table>
                    </div>
                </el-form-item>
            </div>
            <div v-if="active==2">
                <el-form-item label="属性类型" prop="comboId">
                    <el-select v-model="productAttributeOption" class="m-2" placeholder="请选择品牌" size="large" @change="handleSelectProductAttribute">
                        <el-option
                        v-for="item in productAttributeOptions"
                        :key="item.key"
                        :label="item.value"
                        :value="item"
                        />
                    </el-select>
                </el-form-item>
                <el-form-item label="商品规格" >
                    <el-card style="width: 740px;background-color: rgb(246, 248, 250);min-width: 100px;" shadow="never">
                        <el-form :model="item" v-for="(item, index) in standardTable" :key="index" label-position="top" style="display: flex; flex-wrap: wrap;">
                            <el-form-item :label="item.name + '：'" style="width: 600px; margin-bottom:10px">
                                <el-checkbox-group v-model="item.selectList" v-for="(data, index) in item.inputArray" :key="index" @change="handleSelectProductStandard(item)">
                                    <el-checkbox :label="data" style="margin-right: 10px;"/>
                                </el-checkbox-group>
                                <div v-if="item.inputType == 0" style="width: 400px; margin-bottom:10px">
                                    <el-input v-model="item.newSelect"  style="width: 160px;" type="text" autocomplete="off" />
                                    <el-button type="primary" style="margin-left: 12px" @click="addStandard(item)">增加</el-button>
                                </div>
                            </el-form-item>
                        </el-form>
                    </el-card>
                    <div >
                        <el-table :data="skuTableData" style="width: 840px;margin:10px 0 10px 0;"  border 
                            :cell-style="{ textAlign: 'center' }" :header-cell-style="{ 'text-align': 'center' }" >
                            <el-table-column v-for="(item, index) in skuColumsData" :key="index" :prop="item.prop" :label="item.label" :width="item.width" >
                                <template #default="scope">
                                    <div v-if="item.operation">
                                        <el-button link type="primary" size="small" @click.prevent="handleDeleteRow(scop.$index)">删除</el-button>
                                    </div>
                                    <div v-if="item.canEdit">
                                        <el-input-number v-model="scope.row[item.prop]" :precision="2" size="small"></el-input-number>
                                    </div>
                                    <el-text v-else>{{scope.row[item.prop]}}</el-text>
                                </template>
                            </el-table-column>
                        </el-table>
                    </div>
                    <div>
                        <el-button type="primary" style="margin-left: 12px" @click="handleRefresh">刷新列表</el-button>
                        <el-button type="primary" style="margin-left: 12px" @click="handleSynchronizePrice">同步价格</el-button>
                        <el-button type="primary" style="margin-left: 12px" @click="handleSynchronizeStock">同步库存</el-button>
                    </div>
                </el-form-item>
                <el-form-item label="商品参数">
                    <el-card style="width: 740px;height: 200px;" shadow="never">
                        <el-form :model="item" label-width="100px" v-for="(item, index) in parameterTable" :key="index" >
                            <el-form-item :label="item.name" >
                                <view v-if="item.inputArray.length > 0">
                                    <el-select style="width: 240px" v-model="item.inputOption"  placeholder="请选择" size="large">
                                        <el-option
                                            v-for="data in item.inputOptions"
                                            :key="data.key"
                                            :label="data.value"
                                            :value="data.value"
                                        />
                                    </el-select>
                                </view>
                                <view v-else>
                                    <el-input v-model="item.inputList" type="text" autocomplete="off" style="width: 240px;"/>
                                </view>
                            </el-form-item>
                        </el-form>
                    </el-card>
                </el-form-item>
                <el-form-item label="商品相册" prop="name">
                    <el-input v-model="productForm.unit" type="text" autocomplete="off" />
                </el-form-item>
                <el-form-item label="商品详情" prop="name">
                    <el-input v-model="productForm.detailHTML" type="textarea" autocomplete="off" />
                </el-form-item>
            </div>
            <div v-if="active==3">
                <el-form-item label="关联专题" prop="name">
                    <el-input v-model="productForm.unit" type="text" autocomplete="off" />
                </el-form-item>
                <el-form-item label="关联优选" prop="name">
                    <el-input v-model="productForm.unit" type="text" autocomplete="off" />
                </el-form-item>
            </div>
            
        </el-form>
        <el-row>
            <el-col v-if="active>0" :span="8">
                <el-button style="margin-top: 12px" @click="back">上一步，{{titleList[active-1]}}</el-button>
            </el-col>
            <el-col v-if="active<3" :span="8">
                <el-button type="primary" style="margin-top: 12px" @click="next">下一步，{{titleList[active+1]}}</el-button>
            </el-col>
            <el-col v-if="active==3" :span="8">
                <el-button type="primary" style="margin-top: 12px" @click="submitForm(ruleFormRef)">完成，提交商品</el-button>
            </el-col>
        </el-row>
    </el-card>
    
        
  </template>
  
  <script lang="ts" setup>
  import { createProduct, getProductAttributeList } from '@/api/product'
  import { reactive, ref, onBeforeMount, watch } from 'vue'
  import { FormInstance, FormRules, ElMessage } from 'element-plus'
  import { Plus } from '@element-plus/icons-vue'
  import { ProductStore } from '@/pinia/modules/product'  
import { number } from 'echarts'
  const titleList = ["填写商品信息", "填写商品促销", "填写商品属性", "选择商品关联"]
  const active = ref(0)
  const back = () => {
    if (active.value-- < 0) active.value = 0
  }
  const next = () => {
    if (active.value++ > 3) active.value = 0
  }

  const ruleFormRef = ref<FormInstance>()
  
  const productStore = ProductStore()
  
  const productAttributeOption = ref<elementItem>()
  const productAttributeOptions = ref<elementItem[]>([])
  const productCategoryOptions = ref([])
//   const parseProductAttributeCategory = (productAttributeCategoryList) => {
//     let productAttributeMap = {}
//     productAttributeCategoryList.forEach((item) => {
//         let splitted = item.name.split("-")
//         if ( !productAttributeMap[splitted[0]]) {
//             if (splitted.length > 1) {
//                 productAttributeMap[splitted[0]] = [{"id": item.id, "data": splitted[1]}]
//             } else {
//                 productAttributeMap[splitted[0]] = {"id": item.id, "data": splitted[0]}
//             }
//         } else {
//             if (splitted.length > 1) {
//                 productAttributeMap[splitted[0]].push({"id": item.id, "data": splitted[1]})
//             }
//         }
//     })
//     for (let key in productAttributeMap) {
//         let productAttribute = {}
//         productAttribute["label"] = key
//         productAttribute["value"] = productAttributeMap[key].id
//         if (Array.isArray(productAttributeMap[key])) {
//             productAttribute["children"] = []
//             productAttributeMap[key].forEach((item) => {
//                 let productAttributeItem = {}
//                 productAttributeItem["label"] = item.data
//                 productAttributeItem["value"] = item.id
//                 productAttribute["children"].push(productAttributeItem)
//             })
//         }
//         productCategoryOptions.value.push(productAttribute)
//     }
//     console.log("----productCategoryOptions---", productCategoryOptions.value)
//   }
  const getProductAttributeData = async() => {
    await productStore.BuildProductAttributeData()
    productCategoryOptions.value = productStore.ProductCategoryOptions
    console.log("--productCategoryOptions-", productCategoryOptions.value )

    let productAttributeCategoryList = productStore.ProductAttributeCategoryList
    
    productAttributeOptions.value = productAttributeCategoryList.map((item) => {
        return {key: item.id, value: item.name}
    })
    // parseProductAttributeCategory(productAttributeCategoryList)
  }

  const productType = ref()
  const handleProductTypeChange = (value) => {
    console.log("-productType--", productType.value)
    productForm.value.productAttributeCategoryId = productType.value.at(-1)
  }
  watch(() => productAttributeOption.value, () => {
    console.log("---productAttributeOption--", productAttributeOption.value)
    productForm.value.productAttributeCategoryId = productAttributeOption.value.key
  })
  
  interface elementItem {
    key: number,
    value: string,
  }
  const productBrandOption = ref<elementItem>()
  const productBrandOptions = ref<elementItem[]>([])
  const getProductBrandData = async() => {
    await productStore.BuildBrandData()
    let productBrandList = productStore.RandData['list']
    
    productBrandOptions.value = productBrandList.map((item) => {
      return {key: item.id, value: item.name}
    })
  }
  
  watch(() => productBrandOption.value, () => {
    productForm.value.brandId = productBrandOption.value.key
    productForm.value.brandName = productBrandOption.value.value
  })
  

  const productForm = ref({
    albumPics: "",
    brandId: 0,
    brandName: 'brandName',
    deleteStatus: 0,
    description: '这里是商品介绍',
    detailDesc: '这里是详细页描述',
    detailHTML: '',
    detailMobileHTML: '',
    detailTitle: '这里是详细页标题',
    feightTemplateId: 0,    // ??
    flashPromotionCount: 0,
    flashPromotionId: 0,
    flashPromotionPrice: 0,
    flashPromotionSort: 0,
    giftPoint: 0,
    giftGrowth: 0,
    keywords: '关键字',
    lowStock: 0,
    name: '小米6',
    newStatus: 0,
    note: '商品备注',
    originalPrice: 0,
    pic: "",
    prefrenceAreaProductRelationList: [],
    previewStatus: 0,
    price: 0,
    productAttributeCategoryId: 0,
    productCategoryId: 0,
    productCategoryName: 'productCategoryName',
    
    productSN: 'mi2023',
    promotionEndTime: "",
    promotionPerLimit: 0,
    promotionPrice: null,
    promotionStartTime: "",
    promotionType: 0,
    publishStatus: 0,
    recommandStatus: 0,
    sale: 0,
    serviceIds: '1,2',
    skuStockList: [],
    sort: 0,
    stock: 0,
    subTitle: '这里是副标题',
    subjectProductRelationList: [],
    unit: '件',
    usePointLimit: 0,
    verifyStatus: 0,
    weight: 0,
  })
  const memberPriceList =  ref([
        {memberLevelId: 1, memberLevelName: "黄金会员", memberPrice: null},
        {memberLevelId: 2, memberLevelName: "白金会员", memberPrice: null},
        {memberLevelId: 3, memberLevelName: "钻石会员", memberPrice: null}])
  const productAttributeValueList = ref([])
  // 满减
  const productFullReductionList  = ref([
    {
        fullPrice: 0,
        reducePrice: 0,
    }
  ])
  const handleFullDiscountDelete = (index: number, row) =>{
    productFullReductionList.value.splice(index, 1)
  }
  const handleFullDiscountAdd = (index: number, row) =>{
    productFullReductionList.value.push({
        fullPrice: 0,
        reducePrice: 0,
    })
  }

  interface LadderPrice {
    count: 0,
    price: 0,
  }
  // 阶梯价格
  const productLadderList  = ref([
    {
        count: 0,
        discount: 0,
        price: 0,
    }
  ])
  const handleLadderDelete = (index: number, row: LadderPrice) =>{
    productLadderList.value.splice(index, 1)
  }
  const handleLadderAdd = (index: number, row: LadderPrice) =>{
    productLadderList.value.push({
        count: 0,
        discount: 0,
        price: 0,
    })
  }

  onBeforeMount(() => {
    getProductAttributeData()
    getProductBrandData()
  })

  watch(() => productForm.value.productCategoryName, () => {
    if (Array.isArray(productForm.value.productCategoryName)) {
        let categoryNameLength = productForm.value.productCategoryName.length
        productForm.value.productCategoryName = productForm.value.productCategoryName[categoryNameLength-1]
    }
  })

  const skuStockValueList = ref([])
  const submitForm = async(formEl: FormInstance | undefined) => {
    if (!formEl) return false
    // 选择的商品规格
    standardTable.value.forEach((item) => {
        if (item.selectList.length > 0) {
            let attributeMap = {}
            attributeMap['productAttributeId'] = item.id
            let attriValue = ""
            let count = 0
            item.selectList.forEach((attribute) => {
                count += 1
                attriValue +=  attribute
                if (count < item.selectList.length) {
                    attriValue += ','
                }
            })
            attributeMap['value'] = attriValue
            productAttributeValueList.value.push(attributeMap)
        }
    })
    // 商品参数
    parameterTable.value.forEach((item) => {
        let inputValue
        if (item.inputArray.length > 0) {
            if (item.inputOption != "") {
                inputValue = item.inputOption
            }
        } else {
            inputValue = item.inputList
        }
        let parameterValue = {}
        parameterValue['productAttributeId'] = item.id
        parameterValue['value'] = inputValue
        productAttributeValueList.value.push(parameterValue)
    })
    // 库存
    skuTableData.value.forEach((item) => {
        let skuStock = {}
        skuStock["spData"] = item.spData
        skuStock["price"] = item.price
        skuStock["promotionPrice"] = item.promotionPrice
        skuStock["stock"] = item.stock
        skuStock["lowStock"] = item.lowStock
        skuStock["pic"] = item.pic
        skuStockValueList.value.push(skuStock)
    })
    let res
    formEl.validate(async(valid) => {
      if (valid) {
        res = await createProduct({
            "product": productForm.value,
            "memberPriceList": memberPriceList.value,
            "productAttributeValueList": productAttributeValueList.value,
            "productFullReductionList": productFullReductionList.value,
            "productLadderList": productLadderList.value,
            "skuStockList": skuTableData.value
        })
        if (res.code === 0) {
          ElMessage({
            type: 'success',
            message: '添加成功'
          })
          formEl.resetFields()
          resetData()
        }
      } else {
        return false
      } 
    })
  }
  
  const resetData = () =>{
    productForm.value.productCategoryName = null
    productForm.value.subTitle = null
    productForm.value.brandName = ''
    productForm.value.description = null
    productForm.value.productSN = null
    productForm.value.price = 0
    productForm.value.originalPrice = 0
    productForm.value.stock = 0
    productForm.value.unit = null
    productForm.value.sort = null

    productBrandOption.value = {key: 0, value: ""}
  }
  const resetForm = (formEl: FormInstance | undefined) => {
    if (!formEl) return
    formEl.resetFields()
    resetData()
  }
  const checkList = ref([])
  const HandleServiceIdsRadioChanged = () => {
    let serviceIdsStr = ""
    let count = 0
    checkList.value.forEach(element => {
        count += 1
        serviceIdsStr += element 
        if (count < checkList.value.length) {
            serviceIdsStr += ","
        }
    });
    productForm.value.serviceIds = serviceIdsStr
  }
  const promotionTypeState = ref("0")
  const HandlePromotionTypeRadioChanged = () => {
    productForm.value.promotionType = Number(promotionTypeState.value)
  }

  const standardTable = ref([])
  const parameterTable = ref([])
  const handleSelectProductAttribute = async() => {
    skuTableData.value = []
    standardTable.value = []
    parameterTable.value = []
    skuColumsData.value = []
    const attributeId = productAttributeOption.value
    const standardRes = await getProductAttributeList({ tag: attributeId.key, page: 1, pageSize: 100, state: 0 })
    if ("code" in standardRes && standardRes.code === 0) {
        standardRes.data.list.forEach((element)=>{
            let input_arrary = []
            element.inputList.split(",").forEach((item)=>{
                if (item !== "") {
                    input_arrary.push(item)
                }
            })
            element.inputArray = input_arrary
            if (!element.selectList) {

                element.selectList = []
                element.newSelect = ""
            }
            skuColumsData.value.push({label: element.name, prop: element.name, width: "80px"})
        })
        standardTable.value = standardRes.data.list
    }
    const parameterRes = await getProductAttributeList({ tag: attributeId.key, page: 1, pageSize: 100, state: 1 })
    if ("code" in parameterRes && parameterRes.code === 0) {
        parameterRes.data.list.forEach((element)=>{
            let input_arrary = []
            element.inputList.split(",").forEach((item)=>{
                if (item !== "") {
                    input_arrary.push(item)
                }
                
            })
            if (input_arrary.length > 0) {
                element.inputOption = ""
                element.inputOptions = ref<elementItem[]>([])
                let input_options = []
                for (let index = 0; index < input_arrary.length; index++) {
                    let option = {}
                    option['key'] = index
                    option['value'] = input_arrary[index]
                    input_options.push(option)
                }
                element.inputOptions.value = input_options
            }
            element.inputArray = input_arrary
        })
        parameterTable.value = parameterRes.data.list
        skuColumsData.value.push(
            {label: '销售价格', prop: 'price', width: "160px", canEdit: true},
            {label: '促销价格', prop: 'promotionPrice', width: "160px", canEdit: true},
            {label: '商品库存', prop: 'stock', width: "160px", canEdit: true},
            {label: '库存预警值', prop: 'lowStock', width: "160px", canEdit: true},
            {label: '属性图片', prop: 'pic', width: "120px"},
            {label: '操作', operation: true, width: "80px"},
        )
    }
  }
  const attributeValue = {}

  const addStandard = async(item) => {
    if (item.newSelect == '') {
        ElMessage({
            type: 'error',
            message: '内容不可为空'
        })
        return
    } else {
        item.inputArray.push(item.newSelect)
        item.newSelect = ''
    }
  }

  const skuColumsData = ref([])
  const skuTableData = ref([])
  const checkSelectAllStandard = () => {
    for (let index = 0; index < standardTable.value.length; index++) {
        const element = standardTable.value[index];
        if (element.selectList.length < 1) {
            return false
        }
    }
    return true
  }

  const handleSelectProductStandard = async (params) => {
    let allSelect = checkSelectAllStandard()
    if (allSelect) {
        skuTableData.value = []
        let skuTableDataTmp
        for (let index = 0; index < standardTable.value.length; index++) {
            const element = standardTable.value[index];
            skuTableDataTmp = JSON.parse(JSON.stringify(skuTableData.value))
            skuTableData.value = []
            element.selectList.forEach(item => {
                if (skuTableDataTmp.length < 1) {
                    let standard = {}
                    standard[element.name] = item
                    skuTableData.value.push(standard)
                } else {
                    for (let i = 0; i < skuTableDataTmp.length; i++) {
                        const element1 = JSON.parse(JSON.stringify(skuTableDataTmp[i]))
                        element1[element.name] = item

                        // 记录选择的规格（转json）
                        let standardJson = []
                        for (const key in element1) {
                            let standardItem = {}
                            standardItem["key"] = key
                            standardItem["value"] = element1[key]
                            standardJson.push(standardItem)
                        }
                        element1["spData"] = JSON.stringify(standardJson)
                        skuTableData.value.push(element1)
                    }
                }
            });
        }
        // 记录选择的规格（转json）
        skuTableData.value.forEach(element => {
            let standardJson = []
            standardTable.value.forEach(item => {
                let standardItem = {}
                standardItem["key"] = item.name
                standardItem["value"] = element[item.name]
                standardJson.push(standardItem)
            })
            element['spData'] = JSON.stringify(standardJson)
        })
        
        skuTableData.value.forEach(element => {
            element["price"] = 0
            element["promotionPrice"] = 0
            element["stock"] = 0
            element["lowStock"] = 0
            element["pic"] = "pic"
            element["operation"] = true
        });
        console.log("--skuTableData-", skuTableData.value)
    }
  }

  const handleRefresh = async (item) => {
    skuTableData.value.forEach(element => {
        element["price"] = 0
        element["promotionPrice"] = 0
        element["stock"] = 0
        element["lowStock"] = 0
        element["pic"] = "pic"
        element["operation"] = true
    });
  }

  const handleSynchronizePrice = async (item) => {
    let price = 0
    skuTableData.value.forEach(element => {
        if (element["price"] > 0) {
            price = element["price"]
            return
        }
    });
    skuTableData.value.forEach(element => {
        element["price"] = price
    });
  }

  const handleSynchronizeStock = async (item) => {
    let stock = 0
    let lowStock = 0
    skuTableData.value.forEach(element => {
        if (element["stock"] > 0) {
            stock = element["stock"]
        }
        if (element["lowStock"] > 0) {
            lowStock = element["lowStock"]
        }
        if (lowStock > 0 && stock > 0) {
            return
        }
    });
    skuTableData.value.forEach(element => {
        element["stock"] = stock
        element["lowStock"] = lowStock
    });
  }
  const handleDeleteRow = async (index) => {
    skuTableData.value.slice(index, 1)
  }
  </script>
