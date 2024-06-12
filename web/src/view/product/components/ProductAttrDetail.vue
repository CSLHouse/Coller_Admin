<template>
  <div style="margin-top: 50px">
    <el-form :model="modelValue" ref="productAttrForm" label-width="120px" class="form-inner-container" size="small">
      <el-form-item label="属性类型：">
        <el-select v-model="productAttributeCategoryIdSelected"
                   placeholder="请选择属性类型"
                   @change="handleProductAttrChange">
          <el-option
            v-for="item in productAttributeCategoryOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value">
          </el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="商品规格：">
        <el-card shadow="never" class="cardBg">
          <div >
            <!-- {{item.name}}： -->
            <el-form :model="item" v-for="(item, index) in standardTable" :key="index" label-position="top" style="display: flex; flex-wrap: wrap;">
                <el-form-item :label="item.name + '：'" style="width: 600px; margin-bottom:10px">
                  <!-- <div v-if="item.handAddStatus === 0">
                    <el-checkbox-group v-if="item.handAddStatus === 0" v-model="item.selectList" >
                      <el-checkbox  style="margin-right: 10px;" 
                        v-for="data in item.inputArray" :label="data" :key="data"/>
                    </el-checkbox-group>
                  </div>
                  <div v-else>

                  </div> -->
                    <el-checkbox-group v-model="item.selectList" v-for="(data, index) in item.inputArray" :key="index" @change="handleSelectProductStandard()">
                        <el-checkbox :label="data" style="margin-right: 10px;"/>
                    </el-checkbox-group>
                    <div v-if="item.handAddStatus == 1" style="width: 400px; margin-bottom:10px">
                        <el-input v-model="item.newSelect"  style="width: 160px;" type="text" autocomplete="off" />
                        <el-button type="primary" style="margin-left: 12px" @click="handleAddProductAttrValue(index)">增加</el-button>
                    </div>
                </el-form-item>
            </el-form>

            <!-- <el-checkbox-group v-if="productAttr.handAddStatus===0" v-model="item.values">
              <el-checkbox v-for="item in getInputListArr(productAttr.inputList)" :label="item" :key="item"
                           class="littleMarginLeft"></el-checkbox>
            </el-checkbox-group>
            <div v-else>
              <el-checkbox-group v-model="productAttr.values" @change="handleSelectProductStandard()">
                <div v-for="(item,index) in productAttr.options" style="display: inline-block"
                     class="littleMarginLeft">
                  <el-checkbox :label="item" :key="item"></el-checkbox>
                  <el-button type="primary" link class="littleMarginLeft" @click="handleRemoveProductAttrValue(idx,index)">删除
                  </el-button>
                </div>
              </el-checkbox-group>
              <el-input v-model="addProductAttrValue" style="width: 160px;margin-left: 10px" clearable></el-input>
              <el-button class="littleMarginLeft" @click="handleAddProductAttrValue(idx)">增加</el-button>
            </div> -->
          </div>
        </el-card>
        <el-table style="width: 100%;margin-top: 20px"
                  :data="skuTableData"
                  border>
          <el-table-column
            v-for="(item,index) in skuColumsData"
            :label="item.label"
            :key="item.id"
            :prop="item.prop"
            :width="item.width"
            align="center">
            <template #default="scope">
              <el-text>{{scope.row[item.prop]}}</el-text>
              <!-- {{getProductSkuSp(scope.row,index)}} -->
            </template>
          </el-table-column>
          <el-table-column
            label="销售价格"
            width="140"
            align="center">
            <template #default="scope">
              <el-input-number size="small" v-model="scope.row.price" :precision="2"></el-input-number>
            </template>
          </el-table-column>
          <el-table-column
            label="促销价格"
            width="140"
            align="center">
            <template #default="scope">
              <el-input-number size="small" v-model="scope.row.promotionPrice" :precision="2"></el-input-number>
            </template>
          </el-table-column>
          <el-table-column label="属性图片" width="120" align="center">
              <template #default="scope">
                  <cooller-single-upload v-model="scope.row.pic" ></cooller-single-upload>
              </template>
          </el-table-column>
          <el-table-column
            label="商品库存"
            width="140"
            align="center">
            <template #default="scope">
              <el-input-number size="small" v-model="scope.row.stock" :precision="2"></el-input-number>
            </template>
          </el-table-column>
          <el-table-column
            label="库存预警值"
            width="140"
            align="center">
            <template #default="scope">
              <el-input-number size="small" v-model="scope.row.lowStock" :precision="2"></el-input-number>
            </template>
          </el-table-column>
          <el-table-column
            label="SKU编号"
            width="160"
            align="center">
            <template #default="scope">
              <el-input v-model="scope.row.skuCode"></el-input>
            </template>
          </el-table-column>
          <el-table-column
            label="操作"
            width="80"
            align="center">
            <template #default="scope">
              <el-button
              type="danger" link
                @click="handleRemoveProductSku(scope.$index, scope.row)">删除
              </el-button>
            </template>
          </el-table-column>
        </el-table>
        <el-button
          type="primary"
          style="margin-top: 20px"
          @click="handleRefreshProductSkuList">刷新列表
        </el-button>
        <el-button
          type="primary"
          style="margin-top: 20px"
          @click="handleSyncProductSkuPrice">同步价格
        </el-button>
        <el-button
          type="primary"
          style="margin-top: 20px"
          @click="handleSyncProductSkuStock">同步库存
        </el-button>
      </el-form-item>
      <el-form-item label="商品参数：">
        <el-card shadow="never" class="cardBg">
          <div v-for="(item,index) in selectProductParam" :class="{littleMarginTop:index!==0}">
            <div class="paramInputLabel">{{item.name}}:</div>
            <el-select v-if="item.inputType===1" class="paramInput" v-model="selectProductParam[index].value">
              <el-option
                v-for="item1 in getParamInputList(item.inputList)"
                :key="item1"
                :label="item1"
                :value="item1">
              </el-option>
            </el-select>
            <el-input v-else class="paramInput" v-model="selectProductParam[index].value"></el-input>
          </div>
        </el-card>
      </el-form-item>
      <el-form-item label="商品相册：">
        <cooller-image v-model="selectProductPics"></cooller-image>
      </el-form-item>
      <el-form-item label="商品详情：">
        <el-tabs v-model="activeHtmlName" type="card">
          <el-tab-pane label="电脑端详情" name="pc">
            <!-- <tinymce :width="595" :height="300" v-model="value.detailHtml"></tinymce> -->
          </el-tab-pane>
          <el-tab-pane label="移动端详情" name="mobile">
            <!-- <tinymce :width="595" :height="300" v-model="value.detailMobileHtml"></tinymce> -->
          </el-tab-pane>
        </el-tabs>
      </el-form-item>
      <el-form-item style="text-align: center">
        <el-button size="small" @click="handlePrev">上一步，填写商品促销</el-button>
        <el-button type="primary" size="small" @click="handleNext">下一步，选择商品关联</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script setup lang="ts">
import { getProductAttributeList } from '@/api/product'
// import SingleUpload from '@/components/upload/singleUpload.vue'
// import MultiUpload from '@/components/upload/multiUpload.vue'
import coollerSingleUpload from '@/components/upload/coollerSingleUpload.vue'
import CoollerImage from '@/components/upload/coollerImage.vue'
// import Tinymce from '@/components/Tinymce'
import { ref, computed, watch, toRefs, onBeforeMount, reactive } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ProductStore } from '@/pinia/modules/product'
const productStore = ProductStore()

interface ProductAttributeValue {
  productAttributeId: number, value: string
}

interface SkuStockValue {
  spData: string, pic: string, stock: number, lowStock: number, price: number,promotionPrice: number
}

const props = defineProps({
  modelValue: Object as () => ({
    id: number,
    productAttributeCategoryId: number,
    productAttributeValueList: ProductAttributeValue[],
    skuStockList: SkuStockValue[],
    pic: string,
    albumPics: string,
    detailHtml: string,
    detailMobileHtml: string,
  }),
  isEdit: Boolean
})

const { modelValue, isEdit } = toRefs(props);

const pickerOptions1 = ref({
  disabledDate(time) {
    return time.getTime() < Date.now();
  }
})

//编辑模式时是否初始化成功
const hasEditCreated = ref(false)
const productAttributeCategoryOptions = ref([])
//选中的商品属性
const selectProductAttr = ref([])
//选中的商品参数
const selectProductParam = ref([])
//选中的商品属性图片
const selectProductAttrPics = ref([])
//可手动添加的商品属性
const addProductAttrValue = ref('')
//商品富文本详情激活类型
const activeHtmlName = ref('pc')

//是否有商品属性图片
const hasAttrPic = computed(() => {
  if (selectProductAttrPics.value.length < 1) {
    return false;
  }
  return true;
})

//商品的编号
const productId = computed(() => modelValue.value.id)
let pics = ref([])
//商品的主图和画册图片
const selectProductPics = computed({
  get() {
    // let pics = []
    pics.value = []
    if(modelValue.value.pic===undefined||modelValue.value.pic==null||modelValue.value.pic===''){
      return pics;
    }
    pics.value.push(modelValue.value.pic);
    if(modelValue.value.albumPics===undefined||modelValue.value.albumPics==null||modelValue.value.albumPics===''){
      return pics;
    }
    let albumPics = modelValue.value.albumPics.split(',');
    for(let i=0;i<albumPics.length;i++){
      pics.value.push(albumPics[i]);
    }
    return pics;
  },
  set(newValue) {
    if (newValue == null || newValue.length === 0) {
      modelValue.value.pic = null;
      modelValue.value.albumPics = null;
    } else {
      modelValue.value.pic = newValue[0].url;
      modelValue.value.albumPics = '';
      if (newValue.length > 1) {
        for (let i = 1; i < newValue.length; i++) {
          modelValue.value.albumPics += newValue[i].url;
          if (i !== newValue.length - 1) {
            modelValue.value.albumPics += ',';
          }
        }
      }
    }
  }
})

const skuTableData = ref([])
const standardTable = ref([])
const parameterTable = ref([])
const skuColumsData = ref([])

onBeforeMount(() => {
  getProductAttrCateList()
  if (isEdit.value) {
    handleEditCreated()
  }
})


const productAttributeCategoryIdSelected = ref(0)
watch(() => productId.value, (newValue) => {
  
  if(!isEdit.value)return;
  if(hasEditCreated.value)return;
  if(newValue===undefined||newValue==null||newValue===0)return;
  handleEditCreated();
})

const handleEditCreated = async() => {
  //根据商品属性分类id获取属性和参数
  if(modelValue.value.productAttributeCategoryId != null){
    productAttributeCategoryIdSelected.value = modelValue.value.productAttributeCategoryId
    handleProductAttrChange(modelValue.value.productAttributeCategoryId);
    hasEditCreated.value = true;
  }
}
const getProductAttrCateList = async() => {
  await productStore.BuildProductAttributeData()
  let productAttributeCategoryList = productStore.ProductAttributeCategoryList
  for (let i = 0; i < productAttributeCategoryList.length; i++) {
    productAttributeCategoryOptions.value.push({label: productAttributeCategoryList[i].name, value: productAttributeCategoryList[i].id});
  }
}
const getProductAttrList = async(type, cid) => {
  getProductAttributeList({ tag: cid, page: 1, pageSize: 100, state: type }).then(response => {
    let list = response.data.list;
    console.log("---------getProductAttributeList------list---", list)
    if (type === 0) {
      for (let i = 0; i < list.length; i++) {
        let element = list[i]
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
        if (isEdit.value) {
          //编辑状态下获取选中属性
          if (modelValue.value.productAttributeCategoryId == cid) {
            element.selectList = getEditAttrValues(i);
          }
        }
        skuColumsData.value.push({label: element.name, prop: element.name, width: "80px"})
      }
      standardTable.value = list
      console.log("----standardTable----", standardTable.value)
      if (isEdit.value) {
        refreshSkuColumsData()
      }
      
      // selectProductAttr.value = [];
      // for (let i = 0; i < list.length; i++) {
      //   let options = [];
      //   let values = [];
      //   if (isEdit.value) {
      //     //编辑状态下获取手动添加编辑属性
      //     options = getEditAttrOptions(list[i].id);

      //     // if (list[i].handAddStatus === 1) {
      //     //   //编辑状态下获取手动添加编辑属性
      //     //   options = getEditAttrOptions(list[i].id);
      //     // }
      //     //编辑状态下获取选中属性
      //     if (modelValue.value.productAttributeCategoryId == cid) {
      //       values = getEditAttrValues(i);
      //     }
      //   }
      //   selectProductAttr.value.push({
      //     id: list[i].id,
      //     name: list[i].name,
      //     handAddStatus: list[i].handAddStatus,
      //     inputList: list[i].inputList,
      //     values: values,
      //     options: options,
      //     selectType: list[i].selectType
      //   });
      // }
      // if(isEdit.value){
      //   //编辑模式下刷新商品属性图片
      //   refreshProductAttrPics();
      // }
      // console.log("------selectProductAttr-------", selectProductAttr.value)
    } else {
      for (let i = 0; i < list.length; i++) {
        let element = list[i]
        let input_arrary = []
          element.inputList.split(",").forEach((item)=>{
              if (item !== "") {
                  input_arrary.push(item)
              }
          })
          if (input_arrary.length > 0) {
              element.inputOption = ""
              element.inputOptions = ref([])
              let input_options = []
              for (let index = 0; index < input_arrary.length; index++) {
                  let option = {}
                  option['key'] = index
                  option['value'] = input_arrary[index]
                  input_options.push(option)
              }
              element.inputOptions.value = input_options
              if(isEdit.value){
                element.inputOption = getEditParamValue(list[i].id);
              }
          }
          element.inputArray = input_arrary
      }
      parameterTable.value = list

      // selectProductParam.value = [];
      // for (let i = 0; i < list.length; i++) {
      //   let value=null;
      //   if(isEdit.value){
      //     //编辑模式下获取参数属性
      //     value= getEditParamValue(list[i].id);
      //   }
      //   selectProductParam.value.push({
      //     id: list[i].id,
      //     name: list[i].name,
      //     value: value,
      //     inputType: list[i].inputType,
      //     inputList: list[i].inputList
      //   });
      // }
    }
  });
}
//获取设置的可手动添加属性值
// const getEditAttrOptions = (id) => {
//   let options = [];
//   console.log("---------productAttributeValueList:", modelValue.value.productAttributeValueList)
//   for (let i = 0; i < modelValue.value.productAttributeValueList.length; i++) {
//     let attrValue = modelValue.value.productAttributeValueList[i];
//     if (attrValue.productAttributeId === id) {
//       let strArr = attrValue.value.split(',');
//       for (let j = 0; j < strArr.length; j++) {
//         options.push(strArr[j]);
//       }
//       break;
//     }
//   }
//   console.log("---------productAttributeValueList--options:", options)
//   return options;
// }
//获取选中的属性值
const getEditAttrValues = (index) => {
  console.log("-------modelValue.value.skuStockList-----", modelValue.value.skuStockList)
  let values = new Set();
  if (index === 0) {
    for (let i = 0; i < modelValue.value.skuStockList.length; i++) {
      let sku = modelValue.value.skuStockList[i];
      let spData = JSON.parse(sku.spData);
      if (spData!= null && spData.length>=1) {
        values.add(spData[0].value);
      }
    }
  } else if (index === 1) {
    for (let i = 0; i < modelValue.value.skuStockList.length; i++) {
      let sku = modelValue.value.skuStockList[i];
      let spData = JSON.parse(sku.spData);
      if (spData!= null && spData.length>=2) {
        values.add(spData[1].value);
      }
    }
  } else {
    for (let i = 0; i < modelValue.value.skuStockList.length; i++) {
      let sku = modelValue.value.skuStockList[i];
      let spData = JSON.parse(sku.spData);
      if (spData!= null && spData.length>=3) {
        values.add(spData[2].value);
      }
    }
  }
  return Array.from(values);
}
//获取属性的值
const getEditParamValue = (id) => {
  for(let i=0; i<modelValue.value.productAttributeValueList.length;i++){
    if(id === modelValue.value.productAttributeValueList[i].productAttributeId){
      return modelValue.value.productAttributeValueList[i].value;
    }
  }
}
const handleProductAttrChange = async(value) => {
  skuTableData.value = []
  standardTable.value = []
  parameterTable.value = []
  skuColumsData.value = []
  getProductAttrList(0, value);
  getProductAttrList(1, value);
}
const getInputListArr = (inputList) => {
  return inputList.split(',');
}
const handleAddProductAttrValue = (idx) => {
  let content = standardTable[idx].newSelect
  if ( content == null ||  content == '') {
      ElMessage({
          type: 'error',
          message: '内容不可为空'
      })
      return
  }

  if (content.indexOf(standardTable[idx].inputArray)) {
    ElMessage({
      message: '属性值不能重复',
      type: 'warning',
      duration: 1000
    })
    return;
  }

  if (standardTable[idx].selectType == 0 && standardTable[idx].inputArray.length >= 1) {
    ElMessage({
      message: '此规格只能单选，请重新添加',
      type: 'warning',
      duration: 1000
    })
    return
  }

  standardTable[idx].inputArray.push(standardTable[idx].newSelect)
  standardTable[idx].newSelect = ''


  // let options = selectProductAttr.value[idx].options;
  // if (addProductAttrValue.value == null || addProductAttrValue.value == '') {
  //   ElMessage({
  //     message: '属性值不能为空',
  //     type: 'warning',
  //     duration: 1000
  //   })
  //   return
  // }
  
  // if (options.indexOf(addProductAttrValue.value) !== -1) {
  //   ElMessage({
  //     message: '属性值不能重复',
  //     type: 'warning',
  //     duration: 1000
  //   })
  //   return;
  // }
  
  // if (selectProductAttr.value[idx].selectType == 0 && selectProductAttr.value[idx].options.length >= 1) {
  //   ElMessage({
  //     message: '此规格只能单选，请重新添加',
  //     type: 'warning',
  //     duration: 1000
  //   })
  //   return
  // }
  // selectProductAttr.value[idx].options.push(addProductAttrValue.value);
  // addProductAttrValue.value = null;
}
const handleRemoveProductAttrValue = (idx, index) => {
  selectProductAttr.value[idx].options.splice(index, 1);
}
// const getProductSkuSp = (row, index) => {
//   let spData = JSON.parse(row.spData);
//   if(spData!=null&&index<spData.length){
//     return spData[index].value;
//   }else{
//     return null;
//   }
// }

const refreshSkuColumsData = () => {
  skuTableData.value = []
  let skuTableDataTmp
  for (let index = 0; index < standardTable.value.length; index++) {
    const element = standardTable.value[index];
    console.log("---------element-----", element)
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
    console.log("--skuTableData-", skuTableData.value)
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
}
// TODO: 选择规格后刷新
const handleSelectProductStandard = async() => {
  console.log("--------standardTable.value------", standardTable.value)
  refreshSkuColumsData()
  
  skuTableData.value.forEach(element => {
      element["price"] = 0
      element["promotionPrice"] = 0
      element["stock"] = 0
      element["lowStock"] = 0
      element["pic"] = ""
      element["operation"] = true
  });
}

const handleRefreshProductSkuList = () => {
  ElMessageBox.confirm('刷新列表将导致sku信息重新生成，是否要刷新', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    skuTableData.value.forEach(element => {
        element["price"] = 0
        element["promotionPrice"] = 0
        element["stock"] = 0
        element["lowStock"] = 0
        element["pic"] = "pic"
        element["operation"] = true
    });

    // refreshProductAttrPics();
    // refreshProductSkuList();
  });
}
const handleSyncProductSkuPrice = () => {
  ElMessageBox.confirm('将同步第一个sku的价格到所有sku,是否继续', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
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

    // if(modelValue.value.skuStockList!==null&&modelValue.value.skuStockList.length>0){
    //   let tempSkuList = [];
    //   tempSkuList = tempSkuList.concat(tempSkuList,modelValue.value.skuStockList);
    //   let price=modelValue.value.skuStockList[0].price;
    //   for(let i=0;i<tempSkuList.length;i++){
    //     tempSkuList[i].price=price;
    //   }
    //   modelValue.value.skuStockList=[];
    //   modelValue.value.skuStockList=modelValue.value.skuStockList.concat(modelValue.value.skuStockList,tempSkuList);
    // }
  });
}
const handleSyncProductSkuStock = () => {
  ElMessageBox.confirm('将同步第一个sku的库存到所有sku,是否继续', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
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

    // if(modelValue.value.skuStockList!==null&&modelValue.value.skuStockList.length>0){
    //   let tempSkuList = [];
    //   tempSkuList = tempSkuList.concat(tempSkuList,modelValue.value.skuStockList);
    //   let stock=modelValue.value.skuStockList[0].stock;
    //   let lowStock=modelValue.value.skuStockList[0].lowStock;
    //   for(let i=0;i<tempSkuList.length;i++){
    //     tempSkuList[i].stock=stock;
    //     tempSkuList[i].lowStock=lowStock;
    //   }
    //   modelValue.value.skuStockList=[];
    //   modelValue.value.skuStockList=modelValue.value.skuStockList.concat(modelValue.value.skuStockList,tempSkuList);
    // }
  });
}

const confirmData = () => {
  // 选择的商品规格
  standardTable.value.forEach((item) => {
        if (item.selectList.length > 0) {
            
            let attriValue = ""
            let count = 0
            item.selectList.forEach((attribute) => {
                count += 1
                attriValue +=  attribute
                if (count < item.selectList.length) {
                    attriValue += ','
                }
            })
            let attributeMap:ProductAttributeValue = {
                "productAttributeId": item.id,
                "value": attriValue
            }
            modelValue.value.productAttributeValueList.push(attributeMap)
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
        let parameterValue:ProductAttributeValue = {"productAttributeId": item.id, "value": inputValue}
        modelValue.value.productAttributeValueList.push(parameterValue)
    })
    // 库存
    skuTableData.value.forEach((item) => {
        let skuStock:SkuStockValue = {
          "spData": item.spData,
          "price": item.price,
          "promotionPrice": item.promotionPrice,
          "stock": item.stock,
          "lowStock": item.lowStock,
          "pic": item.pic
        }
        modelValue.value.skuStockList.push(skuStock)
    })
}

// const refreshProductSkuList = () => {
//   modelValue.value.skuStockList = [];
//   let skuList = modelValue.value.skuStockList;
//   //只有一个属性时
//   if (selectProductAttr.value.length === 1) {
//     let attr = selectProductAttr.value[0];
//     for (let i = 0; i < attr.values.length; i++) {
//       skuList.push({
//         spData: JSON.stringify([{key:attr.name,value:attr.values[i]}]), pic: "", stock: 0, lowStock: 0, price: 0
//       });
//     }
//   } else if (selectProductAttr.value.length === 2) {
//     let attr0 = selectProductAttr.value[0];
//     let attr1 = selectProductAttr.value[1];
//     for (let i = 0; i < attr0.values.length; i++) {
//       if (attr1.values.length === 0) {
//         skuList.push({
//           spData: JSON.stringify([{key:attr0.name,value:attr0.values[i]}]), pic: "", stock: 0, lowStock: 0, price: 0
//         });
//         continue;
//       }
//       for (let j = 0; j < attr1.values.length; j++) {
//         let spData = [];
//         spData.push({key:attr0.name,value:attr0.values[i]});
//         spData.push({key:attr1.name,value:attr1.values[j]});
//         skuList.push({
//           spData: JSON.stringify(spData), pic: "", stock: 0, lowStock: 0, price: 0
//         });
//       }
//     }
//   } else {
//     let attr0 = selectProductAttr.value[0];
//     let attr1 = selectProductAttr.value[1];
//     let attr2 = selectProductAttr.value[2];
//     for (let i = 0; i < attr0.values.length; i++) {
//       if (attr1.values.length === 0) {
//         skuList.push({
//           spData: JSON.stringify([{key:attr0.name,value:attr0.values[i]}]), pic: "", stock: 0, lowStock: 0, price: 0
//         });
//         continue;
//       }
//       for (let j = 0; j < attr1.values.length; j++) {
//         if (attr2.values.length === 0) {
//           let spData = [];
//           spData.push({key:attr0.name,value:attr0.values[i]});
//           spData.push({key:attr1.name,value:attr1.values[j]});
//           skuList.push({
//             spData: JSON.stringify(spData), pic: "", stock: 0, lowStock: 0, price: 0
//           });
//           continue;
//         }
//         for (let k = 0; k < attr2.values.length; k++) {
//           let spData = [];
//           spData.push({key:attr0.name,value:attr0.values[i]});
//           spData.push({key:attr1.name,value:attr1.values[j]});
//           spData.push({key:attr2.name,value:attr2.values[k]});
//           skuList.push({
//             spData: JSON.stringify(spData), pic: "", stock: 0, lowStock: 0, price: 0
//           });
//         }
//       }
//     }
//   }
// }
// const refreshProductAttrPics = () => {
//   selectProductAttrPics.value = [];
//   if (selectProductAttr.value.length >= 1) {
//     let values = selectProductAttr.value[0].values;
//     for (let i = 0; i < values.length; i++) {
//       let pic=null;
//       if(isEdit.value){
//         //编辑状态下获取图片
//         pic= getProductSkuPic(values[i]);
//       }
//       selectProductAttrPics.value.push({name: values[i], pic: pic})
//     }
//   }
// }
// //获取商品相关属性的图片
// const getProductSkuPic = (name) => {
//   for(let i=0;i<modelValue.value.skuStockList.length;i++){
//     let spData = JSON.parse(modelValue.value.skuStockList[i].spData);
//     if(name===spData[0].value){
//       return modelValue.value.skuStockList[i].pic;
//     }
//   }
//   return null;
// }
// //合并商品属性
// const mergeProductAttrValue = () => {
//   modelValue.value.productAttributeValueList = [];
//   console.log("------selectProductAttr.value-----", selectProductAttr.value)
//   for (let i = 0; i < selectProductAttr.value.length; i++) {
//     let attr = selectProductAttr.value[i];
//     if (attr.handAddStatus === 1 && attr.options != null && attr.options.length > 0) {
//       modelValue.value.productAttributeValueList.push({
//         productAttributeId: attr.id,
//         value: getOptionStr(attr.options)
//       });
//     }
//   }
//   for (let i = 0; i < selectProductParam.value.length; i++) {
//     let param = selectProductParam.value[i];
//     modelValue.value.productAttributeValueList.push({
//       productAttributeId: param.id,
//       value: param.value
//     });
//   }
// }
// //合并商品属性图片
// const mergeProductAttrPics = () => {
//   for (let i = 0; i < selectProductAttrPics.value.length; i++) {
//     for (let j = 0; j < modelValue.value.skuStockList.length; j++) {
//       let spData = JSON.parse(modelValue.value.skuStockList[j].spData);
//       if (spData[0].value === selectProductAttrPics.value[i].name) {
//         modelValue.value.skuStockList[j].pic = selectProductAttrPics.value[i].pic;
//       }
//     }
//   }
// }
// const getOptionStr = (arr) => {
//   let str = '';
//   for (let i = 0; i < arr.length; i++) {
//     str += arr[i];
//     if (i != arr.length - 1) {
//       str += ',';
//     }
//   }
//   return str;
// }
const handleRemoveProductSku = (index, row) => {
  let list = modelValue.value.skuStockList;
  if (list.length === 1) {
    list.pop();
  } else {
    list.splice(index, 1);
  }
}
const getParamInputList = (inputList) => {
  return inputList.split(',');
}

const emits = defineEmits(["prevStep", "nextStep"]);
const handlePrev = () => {
  emits('prevStep')
}
const handleNext = () => {
  // mergeProductAttrValue();
  // mergeProductAttrPics();
  confirmData()
  emits('nextStep')
}

</script>


<style scoped>
  .littleMarginLeft {
    margin-left: 10px;
  }

  .littleMarginTop {
    margin-top: 10px;
  }

  .paramInput {
    width: 250px;
  }

  .paramInputLabel {
    display: inline-block;
    width: 100px;
    text-align: right;
    padding-right: 10px
  }

  .cardBg {
    background: #F8F9FC;
  }
</style>
