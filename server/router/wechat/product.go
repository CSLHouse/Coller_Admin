package wechat

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WechatRouter struct{}

func (e *WechatRouter) InitWechatRouter(Router *gin.RouterGroup, RouterPub *gin.RouterGroup) {

	homePublicRouterWithoutRecord := RouterPub.Group("product")

	wechatRouter := Router.Group("product").Use(middleware.OperationRecord())
	wechatRouterWithoutRecord := Router.Group("product")
	homeApi := v1.ApiGroupApp.WechatApiGroup.HomeApi
	{
		wechatRouter.POST("advertise", homeApi.CreateHomeAdvertise)
		wechatRouter.PUT("advertise", homeApi.UpdateHomeAdvertise)
		wechatRouter.DELETE("advertise", homeApi.DeleteHomeAdvertise)
		wechatRouter.PUT("advertiseState", homeApi.UpdateHomeAdvertiseOnlineState)
		wechatRouter.POST("create", homeApi.CreateProduct)
		wechatRouter.PUT("update", homeApi.UpdateProduct)
		wechatRouter.PUT("updateKeyword", homeApi.UpdateProductForKeyword)

		wechatRouter.POST("brand", homeApi.CreateProductBrand)
		wechatRouter.PUT("brand", homeApi.UpdateProductBrand)
		wechatRouter.DELETE("brand", homeApi.DeleteHomeProductBrand)
		wechatRouter.PUT("brandState", homeApi.UpdateProductBrandOnlineStat)

		wechatRouter.POST("recommendProduct", homeApi.CreateRecommendProducts)
		wechatRouter.PUT("recommendProduct", homeApi.UpdateRecommendProducts)
		wechatRouter.DELETE("recommendProduct", homeApi.DeleteRecommendProducts)

		wechatRouter.POST("attributeCategory", homeApi.CreateProductAttributeCategory)
		wechatRouter.PUT("attributeCategory", homeApi.UpdateProductAttributeCategory)
		wechatRouter.DELETE("attributeCategory", homeApi.DeleteProductAttributeCategory)
		wechatRouter.POST("attribute", homeApi.CreateProductAttribute)
		wechatRouter.PUT("attribute", homeApi.UpdateProductAttribute)
		wechatRouter.DELETE("attribute", homeApi.DeleteProductAttribute)
		wechatRouter.POST("productCategory", homeApi.CreateProductCategory)
		wechatRouter.PUT("productCategory", homeApi.UpdateProductCategory)
		wechatRouter.DELETE("productCategory", homeApi.DeleteProductCategory)
		wechatRouter.PUT("sku", homeApi.UpdateSKUStock)
	}
	{
		wechatRouterWithoutRecord.POST("cart", homeApi.CreateProductCart)
		wechatRouterWithoutRecord.PUT("cart", homeApi.UpdateProductCartQuantity)
		wechatRouterWithoutRecord.DELETE("cart", homeApi.DeleteProductCartById)
		wechatRouterWithoutRecord.DELETE("cart/clear", homeApi.ClearProductCart)
		wechatRouterWithoutRecord.GET("cart/list", homeApi.GetProductCartList)
		wechatRouterWithoutRecord.DELETE("carts", homeApi.DeleteProductCartByIds)
		wechatRouterWithoutRecord.DELETE("deletes", homeApi.DeleteProducts)
	}
	{
		homePublicRouterWithoutRecord.GET("content", homeApi.GetAllWechatContent)
		homePublicRouterWithoutRecord.GET("advertiseList", homeApi.GetHomeAdvertiseList)
		homePublicRouterWithoutRecord.GET("recommendProductList", homeApi.GetRecommendProductList) // 获取推荐商品列表
		homePublicRouterWithoutRecord.GET("list", homeApi.GetProductList)                          // 获取商品列表
		homePublicRouterWithoutRecord.GET("brand", homeApi.GetProductBrandList)

		homePublicRouterWithoutRecord.GET("brandDetail", homeApi.GetProductBrandByID)
		homePublicRouterWithoutRecord.GET("productDetail", homeApi.GetProductByID)
		homePublicRouterWithoutRecord.GET("attributeCategory", homeApi.GetProductAttributeCategoryList)
		homePublicRouterWithoutRecord.GET("attribute", homeApi.GetProductAttributeListByCategoryId)
		homePublicRouterWithoutRecord.GET("productCategory", homeApi.GetProductCategoryList)
		homePublicRouterWithoutRecord.GET("productList", homeApi.GetProductListByOnlyIDWithSort) // 获取商品列表
		homePublicRouterWithoutRecord.GET("allCategory", homeApi.GetProductAllCategory)          // 获取商品分类列表
		homePublicRouterWithoutRecord.GET("sku", homeApi.GetSkuStockByProductID)
	}
}
