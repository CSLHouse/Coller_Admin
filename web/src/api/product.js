import service from '@/utils/request'

export const getProductList = (params) => {
    return service({
        url: '/product/list',
        method: 'get',
        params
    })
}

export const getProductDetail = (params) => {
    return service({
        url: '/product/productDetail',
        method: 'get',
        params
    })
}

export const updateProducts = (data) => {
    return service({
        url: '/product/update',
        method: 'put',
        data
    })
}

export const getBrandList = (params) => {
    return service({
        url: '/product/brand',
        method: 'get',
        params
    })
}

export const createProduct = (data) => {
    return service({
        url: '/product/create',
        method: 'post',
        data
    })
}

export const createProductBrand = (data) => {
    return service({
        url: '/product/brand',
        method: 'post',
        data
    })
}

export const updateProductBrand = (data) => {
    return service({
        url: '/product/brand',
        method: 'put',
        data
    })
}

export const deleteProductBrand = (params) => {
    return service({
        url: '/product/brand',
        method: 'delete',
        params: params,
    })
}

// 获取首页轮播广告表
export const getAdvertiseList = (params) => {
    return service({
        url: '/product/advertiseList',
        method: 'get',
        params
    })
}

// 创建首页轮播广告
export const createAdvertise = (data) => {
    return service({
      url: '/product/content',
      method: 'post',
      data
    })
}

// 更新首页轮播广告
export const updateAdvertise = (data) => {
    return service({
        url: '/product/content',
        method: 'put',
        data
    })
}

// 获取首页推荐专题表 猜你喜欢
export const getHotProductList = (params) => {
    return service({
        url: '/product/hotProduct',
        method: 'get',
        params
    })
}

// 获取首页推荐专题表 猜你喜欢
export const addHotProductList = (data) => {
    return service({
        url: '/product/hotProduct',
        method: 'post',
        data
    })
}
// 更新首页推荐专题表
export const updateHotProduct = (data) => {
    return service({
        url: '/product/hotProduct',
        method: 'put',
        data
    })
}

// 删除首页推荐专题表
export const deleteHotProduct = (data) => {
    return service({
        url: '/product/hotProduct',
        method: 'delete',
        data
    })
}

// 获取商品属性分类列表
export const getProductAttributeCategoryList = (params) => {
    return service({
        url: '/product/attributeCategory',
        method: 'get',
        params
    })
}

// 创建商品属性分类列表
export const createProductAttributeCategory = (data) => {
    return service({
      url: '/product/attributeCategory',
      method: 'post',
      data
    })
}

// 更新商品属性分类列表
export const updateProductAttributeCategory = (data) => {
    return service({
        url: '/product/attributeCategory',
        method: 'put',
        data
    })
}

// 删除商品属性分类列表
export const deleteProductAttributeCategory = (data) => {
    return service({
      url: '/product/attributeCategory',
      method: 'delete',
      data
    })
}

// 获取商品属性参数列表
export const getProductAttributeList = (params) => {
    return service({
        url: '/product/attribute',
        method: 'get',
        params
    })
}

// 创建商品属性参数列表
export const createProductAttribute = (data) => {
    return service({
      url: '/product/attribute',
      method: 'post',
      data
    })
}

// 更新商品属性参数列表
export const updateProductAttribute = (data) => {
    return service({
        url: '/product/attribute',
        method: 'put',
        data
    })
}

// 删除商品属性参数列表
export const deleteProductAttribute = (data) => {
    return service({
      url: '/product/attribute',
      method: 'delete',
      data
    })
}

// 获取商品分类列表
export const getProductCategoryList = (params) => {
    return service({
        url: '/product/productCategory',
        method: 'get',
        params
    })
}

// 创建商品分类列表
export const createProductCategory= (data) => {
    return service({
      url: '/product/productCategory',
      method: 'post',
      data
    })
}

// 更新商品分类列表
export const updateProductCategory = (data) => {
    return service({
        url: '/product/productCategory',
        method: 'put',
        data
    })
}

// 删除商品分类列表
export const deleteProductCategory = (data) => {
    return service({
      url: '/product/productCategory',
      method: 'delete',
      data
    })
}

// 获取商品sku库存
export const getProductSKUStockByProductID = (params) => {
    return service({
        url: '/product/sku',
        method: 'get',
        params
    })
}

// 更新商品sku库存
export const updateProductSKUStock = (data) => {
    return service({
        url: '/product/sku',
        method: 'put',
        data
    })
}