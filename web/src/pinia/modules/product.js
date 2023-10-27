import { getProductAttributeCategoryList, getBrandList } from '@/api/product'
import { defineStore } from 'pinia'
import { ref, computed, watch } from 'vue'

export const ProductStore = defineStore('product', () => {

  const ProductAttributeCategoryList = ref()
  const RandData = ref()
  const ProductCategoryOptions = []

  const setProductAttributeCategoryList = (val) => {
    ProductAttributeCategoryList.value = val
  }

  /* 获取产品属性分类*/
  const GetProductAttributeCategoryList = async() => {
    const res = await getProductAttributeCategoryList()
    if (res.code === 0) {
      setProductAttributeCategoryList(res.data.list)
    }
    return res
  }
  
  const setBrandData = (val) => {
    RandData.value = val
  }

  /* 获取产品品牌分类*/
  const BuildBrandData = async(isRefresh) => {
    if (!RandData.value || isRefresh) {
      const res = await getBrandList()
      if (res.code === 0) {
        setBrandData(res.data)
      }
    }
  }
  
  const parseProductAttributeCategory = () => {
    let productAttributeMap = {}
    ProductAttributeCategoryList.value.forEach((item) => {
        let splitted = item.name.split("-")
        if ( !productAttributeMap[splitted[0]]) {
            if (splitted.length > 1) {
                productAttributeMap[splitted[0]] = [{"id": item.id, "data": splitted[1]}]
            } else {
                productAttributeMap[splitted[0]] = {"id": item.id, "data": splitted[0]}
            }
        } else {
            if (splitted.length > 1) {
                productAttributeMap[splitted[0]].push({"id": item.id, "data": splitted[1]})
            }
        }
    })
    console.log("---productAttributeMap-", productAttributeMap)
    for (let key in productAttributeMap) {
        let productAttribute = {}
        productAttribute["label"] = key
        productAttribute["value"] = productAttributeMap[key].id
        
        if (Array.isArray(productAttributeMap[key])) {
            productAttribute["children"] = []
            productAttributeMap[key].forEach((item) => {
                let productAttributeItem = {}
                productAttributeItem["label"] = item.data
                productAttributeItem["value"] = item.id
                productAttribute["children"].push(productAttributeItem)
            })
        }
        ProductCategoryOptions.push(productAttribute)
    }
  }
  const BuildProductAttributeData = async(isRefresh) => {
    if (!ProductAttributeCategoryList.value || isRefresh) {
      await GetProductAttributeCategoryList()
    }
    parseProductAttributeCategory()
  }

  return {
    ProductAttributeCategoryList,
    GetProductAttributeCategoryList,
    RandData,
    BuildBrandData,
    ProductCategoryOptions,
    BuildProductAttributeData,
  }
})
