<template>
    <div style="margin-top: 50px">
        <el-form
            ref="ruleFormRef"
            :model="productInfoForm"
            status-icon
            label-width="120px"
            class="demo-ruleForm"
            :rules="rules"
            style="max-width: 860px;margin-left: 20px;"
            >
            <el-form-item label="商品分类" prop="productCategoryName">
                    <el-cascader v-model="productType" :options="productCategoryOptions" 
                        placeholder="请选择" clearable @change="handleProductTypeChange"/>
                </el-form-item>
                <el-form-item label="商品名称" prop="name">
                    <el-input v-model="productInfoForm.name" type="tel" autocomplete="off" />
                </el-form-item>
                <el-form-item label="副标题" prop="subTitle">
                    <el-input v-model="productInfoForm.subTitle" type="text" autocomplete="off" />
                </el-form-item>
                <el-form-item label="商品品牌" prop="brand">
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
                    <el-input v-model="productInfoForm.description" type="textarea" autocomplete="off" />
                </el-form-item>
                <el-form-item label="商品货号" prop="collection">
                    <el-input v-model="productInfoForm.productSN" type="text" autocomplete="off" />
                </el-form-item>
                <el-form-item label="商品售价" prop="date">
                    <el-input-number v-model="productInfoForm.price" :precision="2"></el-input-number>
                </el-form-item>
                <el-form-item label="市场价" prop="cardId">
                    <el-input-number v-model="productInfoForm.originalPrice" :precision="2"></el-input-number>
                </el-form-item>
                <el-form-item label="商品库存" prop="telephone">
                    <el-input v-model.number="productInfoForm.stock" type="number" autocomplete="off" style="width: 100px"/>
                </el-form-item>
                <el-form-item label="计量单位" prop="name">
                    <el-input v-model="productInfoForm.unit" type="text" autocomplete="off" />
                </el-form-item>
                <el-form-item label="商品重量" prop="name">
                    <el-col :span="8">
                        <el-input-number v-model="productInfoForm.weight" :precision="2"></el-input-number>
                    </el-col>
                    <a style=" margin-left: 10px;">克</a>
                </el-form-item>
                <el-form-item label="排序" prop="name">
                    <el-input v-model="productInfoForm.sort" autocomplete="off" />
                </el-form-item>
        </el-form>
    </div>
</template>

<script lang="ts" setup>
    import { createProduct, getProductAttributeList, getProductDetail } from '@/api/product'
    import { reactive, ref, onBeforeMount, watch, toRefs } from 'vue'
    import { FormInstance, FormRules, ElMessage } from 'element-plus'
    import { ProductStore } from '@/pinia/modules/product' 
    const productStore = ProductStore()
    export interface Form {
        albumPics: string,
            brandId: number,
            brandName: string,
            deleteStatus: number,
            description: string,
            detailDesc: string,
            detailHTML: string,
            detailMobileHTML: string,
            detailTitle: string,
            feightTemplateId: number,    // ??
            flashPromotionCount: number,
            flashPromotionId: number,
            flashPromotionPrice: number,
            flashPromotionSort: number,
            giftPoint: number,
            giftGrowth: number,
            keywords: string,
            lowStock: number,
            name: string,
            newStatus: number,
            note: string,
            originalPrice: number,
            pic: string,
            prefrenceAreaProductRelationList: [],
            previewStatus: number,
            price: number,
            productAttributeCategoryId: number,
            productCategoryId: number,
            productCategoryName: string,
            
            productSN: string,
            promotionEndTime: string,
            promotionPerLimit: number,
            promotionPrice: null,
            promotionStartTime: string,
            promotionType: number,
            publishStatus: number,
            recommandStatus: number,
            sale: number,
            serviceIds: string,
            skuStockList: [],
            sort: number,
            stock: number,
            subTitle: string,
            subjectProductRelationList: [],
            unit: string,
            usePointLimit: number,
            verifyStatus: number,
            weight: number,
    }
    const props = defineProps<{
        //子组件接收父组件传递过来的值
        productInfoForm: Form,
        isEdit: Boolean,
    }>()
    //使用父组件传递过来的值
    const { isEdit, productInfoForm } = toRefs(props)

    const ruleFormRef = ref<FormInstance>()
    const rules = reactive<FormRules>({
        productCategoryName: [
            { required: true, message: "请选择分类", trigger: "blur" }
        ],
        name: [
            { required: true, message: "请输入商品名称", trigger: "blur" }
        ],
        brand: [
            { required: true, message: "请选择品牌", trigger: "blur" }
        ],
    })

    onBeforeMount(async() => {
        if (isEdit.value) {
            getProductCategoryData()
            getProductBrandData()
        }
    })

    const productAttributeOption = ref<elementItem>()
    const productAttributeOptions = ref<elementItem[]>([])
    const productCategoryOptions = ref([])

    const getProductCategoryData = async() => {
        await productStore.BuildProductAttributeData()
        productCategoryOptions.value = productStore.ProductCategoryOptions
        let productAttributeCategoryList = productStore.ProductAttributeCategoryList
        
        productAttributeOptions.value = productAttributeCategoryList.map((item) => {
            return {key: item.id, value: item.name}
        })
    }
    const productType = ref()
    const handleProductTypeChange = (value) => {
        productInfoForm.value.productAttributeCategoryId = value.at(-1)
    }
    watch(() => productAttributeOption.value, () => {
        productInfoForm.value.productAttributeCategoryId = productAttributeOption.value.key
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
        productInfoForm.value.brandId = productBrandOption.value.key
        productInfoForm.value.brandName = productBrandOption.value.value
    })

</script>