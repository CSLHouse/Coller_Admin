package wechat

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wechat"
	wechatRequest "github.com/flipped-aurora/gin-vue-admin/server/model/wechat/request"
	wechatRes "github.com/flipped-aurora/gin-vue-admin/server/model/wechat/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"time"
)

type HomeApi struct{}

func (e *HomeApi) CreateHomeAdvertise(c *gin.Context) {
	var home wechat.HomeAdvertise
	err := c.ShouldBindJSON(&home)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	//home.SysUserAuthorityID = utils.GetUserAuthorityId(c)
	err = wechatService.CreateHomeAdvertise(home)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteExaCustomer
// @Tags      ExaCustomer
// @Summary   删除客户
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      example.ExaCustomer            true  "客户ID"
// @Success   200   {object}  response.Response{msg=string}  "删除客户"
// @Router    /customer/customer [delete]
func (e *HomeApi) DeleteHomeAdvertise(c *gin.Context) {
	var home wechat.HomeAdvertise
	err := c.ShouldBindJSON(&home)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = wechatService.DeleteHomeAdvertise(home)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateExaCustomer
// @Tags      ExaCustomer
// @Summary   更新客户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      example.ExaCustomer            true  "客户ID, 客户信息"
// @Success   200   {object}  response.Response{msg=string}  "更新客户信息"
// @Router    /customer/customer [put]
func (e *HomeApi) UpdateHomeAdvertise(c *gin.Context) {
	var home wechat.HomeAdvertise
	err := c.ShouldBindJSON(&home)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//home.SysUserAuthorityID = utils.GetUserAuthorityId(c)
	err = wechatService.UpdateHomeAdvertise(&home)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// GetExaCustomer
// @Tags      ExaCustomer
// @Summary   获取单一客户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     example.ExaCustomer                                                true  "客户ID"
// @Success   200   {object}  response.Response{data=exampleRes.ExaCustomerResponse,msg=string}  "获取单一客户信息,返回包括客户详情"
// @Router    /customer/customer [get]
func (e *HomeApi) GetHomeAdvertise(c *gin.Context) {
	var combo wechat.HomeAdvertise
	err := c.ShouldBindQuery(&combo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	data, err := wechatService.GetHomeAdvertise(combo.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	response.OkWithDetailed(response.AllResult{List: data}, "获取成功", c)
}

// GetHomeAdvertiseList
// @Tags      ExaCustomer
// @Summary   分页获取权限客户列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.PageInfo                                        true  "页码, 每页大小"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取权限客户列表,返回包括列表,总数,页码,每页数量"
// @Router    /customer/customerList [get]
func (e *HomeApi) GetHomeAdvertiseList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	homeList, err := wechatService.GetHomeAdvertiseInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     homeList,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

func (e *HomeApi) GetAllHomeAdvertise(c *gin.Context) {
	homeList, err := wechatService.GetAllHomeAdvertiseInfoList()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	brandList, err := wechatService.GetAllHomeBrandInfoList()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	homeFlashPromotion, err := wechatService.GetAllHomeFlashPromotionInfoList()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	var flashPromotionList wechatRes.FlashPromotionResModel
	flashPromotionList.NextStartTime = 111
	flashPromotionList.EndTime = 222
	flashPromotionList.ProductList = homeFlashPromotion

	newProductList, err := wechatService.GetAllNewProductInfoList()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	hotProductList, err := wechatService.GetAllHomeHotProductListInfoList()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(wechatRes.HomeContentResponse{
		AdvertiseList:      homeList,
		BrandList:          brandList,
		HomeFlashPromotion: homeFlashPromotion,
		NewProductList:     newProductList,
		HotProductList:     hotProductList,
	}, "获取成功", c)
}

func (e *HomeApi) GetRecommendProductList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	recommendProductList, err := wechatService.GetRecommendProductList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	recommendDataList := make([]wechatRes.HomeRecommendResponse, 0)
	//var recommendData wechatRes.HomeRecommendResponse
	for _, recommend := range recommendProductList {
		recommendData := wechatRes.HomeRecommendResponse{}
		recommendData.ID = recommend.ProductId
		recommendData.Name = recommend.ProductName
		recommendData.Pic = recommend.Product.Pic
		recommendData.SubTitle = recommend.Product.SubTitle
		recommendData.Price = recommend.Product.Price
		recommendDataList = append(recommendDataList, recommendData)
	}

	response.OkWithDetailed(response.AllResult{
		List: recommendDataList,
	}, "获取成功", c)
}

func (e *HomeApi) CreateProduct(c *gin.Context) {
	var home wechatRequest.ProductCreateRequest
	err := c.ShouldBindJSON(&home)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var product wechat.HomeProduct
	product = home.Product
	err = wechatService.CreateHomeProduct(&product)
	if err != nil {
		global.GVA_LOG.Error("创建产品失败!", zap.Error(err))
		response.FailWithMessage("创建产品失败", c)
		return
	}
	// 新品
	if product.NewStatus == 1 {
		var newProduct wechat.HomeNewProduct
		newProduct.ProductId = product.ID
		newProduct.ProductName = product.Name
		newProduct.RecommendStatus = product.PublishStatus
		err = wechatService.CreateNewProduct(&newProduct)
		if err != nil {
			global.GVA_LOG.Error("创建新品失败!", zap.Error(err))
			response.FailWithMessage("创建新品失败", c)
			return
		}
	}
	// 推荐
	if product.RecommandStatus == 1 {
		var ewcommendProduct wechat.HomeRecommendProduct
		ewcommendProduct.ProductId = product.ID
		ewcommendProduct.ProductName = product.Name
		ewcommendProduct.RecommendStatus = product.PublishStatus
		err = wechatService.CreateRecommendProduct(&ewcommendProduct)
		if err != nil {
			global.GVA_LOG.Error("创建推荐商品失败!", zap.Error(err))
			response.FailWithMessage("创建推荐商品失败", c)
			return
		}
	}
	// 会员价格
	for _, memberPrice := range home.MemberPriceList {
		if product.ID > 0 && memberPrice.MemberPrice > 0 {
			var productMemberPrice wechat.MemberPrice
			productMemberPrice.ProductId = product.ID
			productMemberPrice.MemberLevelId = memberPrice.MemberLevelId
			productMemberPrice.MemberLevelName = memberPrice.MemberLevelName
			productMemberPrice.MemberPrice = memberPrice.MemberPrice
			err = wechatService.CreateProductMemberPrice(&productMemberPrice)
			if err != nil {
				global.GVA_LOG.Error("创建会员价格数据失败!", zap.Error(err))
				response.FailWithMessage("创建会员价格数据失败", c)
				return
			}
		}
	}

	// 创建产品参数信
	for _, attribute := range home.ProductAttributeValueList {
		if product.ID > 0 && len(attribute.Value) > 0 {
			var attributeValue wechat.HomeProductAttributeValue
			attributeValue.ProductId = product.ID
			attributeValue.ProductAttributeId = attribute.ProductAttributeId
			attributeValue.Value = attribute.Value
			err = wechatService.CreateProductAttributeValue(&attributeValue)
			if err != nil {
				global.GVA_LOG.Error("创建属性数据失败!", zap.Error(err))
				response.FailWithMessage("创建属性数据失败", c)
				return
			}
		}
	}
	//产品满减
	for _, fullReduction := range home.ProductFullReductionList {
		if product.ID != 0 && fullReduction.FullPrice > 0 && fullReduction.ReducePrice > 0 {
			var productFullReduction wechat.ProductFullReduction
			productFullReduction.ProductId = product.ID
			productFullReduction.FullPrice = fullReduction.FullPrice
			productFullReduction.ReducePrice = fullReduction.ReducePrice
			err = wechatService.CreateProductFullReduction(&productFullReduction)
			if err != nil {
				global.GVA_LOG.Error("创建产品满减数据失败!", zap.Error(err))
				response.FailWithMessage("创建产品满减数据失败", c)
				return
			}
		}
	}
	// 产品阶梯价格
	for _, ladder := range home.ProductLadderList {
		if product.ID > 0 && ladder.Count > 0 && ladder.Discount > 0 && ladder.Price > 0 {
			var productLadder wechat.ProductLadder
			productLadder.ProductId = product.ID
			productLadder.Count = ladder.Count
			productLadder.Discount = ladder.Discount
			productLadder.Price = ladder.Price
			err = wechatService.CreateProductLadder(&productLadder)
			if err != nil {
				global.GVA_LOG.Error("创建产品阶梯价格数据失败!", zap.Error(err))
				response.FailWithMessage("创建产品阶梯价格数据失败", c)
				return
			}
		}
	}
	// sku的库存
	for _, stock := range home.SkuStockList {
		if product.ID > 0 && stock.Price > 0 && stock.PromotionPrice > 0 {
			var skuStock wechat.SKUStock
			skuStock.ProductId = product.ID
			skuStock.Price = stock.Price
			skuStock.Stock = product.Stock
			skuStock.LowStock = stock.LowStock
			skuStock.PromotionPrice = stock.PromotionPrice
			skuStock.SkuCode = fmt.Sprintf("%d", time.Now().UnixNano())
			skuStock.Pic = stock.Pic
			skuStock.SpData = stock.SpData
			err = wechatService.CreateProductSKUStock(&skuStock)
			if err != nil {
				global.GVA_LOG.Error("创建sku的库存数据失败!", zap.Error(err))
				response.FailWithMessage("创建sku的库存数据失败", c)
				return
			}
		}
	}
	response.OkWithMessage("创建成功", c)
}

// GetProductByID 获取商品详情
func (e *HomeApi) GetProductByID(c *gin.Context) {
	var reqId request.GetById
	err := c.ShouldBindQuery(&reqId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	product, err := wechatService.GetProductByID(reqId.ID)
	if err != nil {
		global.GVA_LOG.Error("获取商品失败!", zap.Error(err))
		response.FailWithMessage("获取商品失败", c)
		return
	}

	attributeList, err := wechatService.GetProductAttributeListById(product.ProductAttributeCategoryId)
	if err != nil {
		global.GVA_LOG.Error("获取产品属性失败!", zap.Error(err))
		response.FailWithMessage("获取产品属性失败", c)
		return
	}
	attributeValue, err := wechatService.GetProductAttributeValueByProductId(reqId.ID)
	if err != nil {
		global.GVA_LOG.Error("获取商品参数值失败!", zap.Error(err))
		response.FailWithMessage("获取商品参数值失败", c)
		return
	}
	fullReductionList, err := wechatService.GetProductFullReductionByProductId(reqId.ID)
	if err != nil {
		global.GVA_LOG.Error("获取商品满减优惠失败!", zap.Error(err))
		response.FailWithMessage("获取商品满减优惠失败", c)
		return
	}
	skuStockList, err := wechatService.GetProductSKUStockByProductId(reqId.ID, "")
	if err != nil {
		global.GVA_LOG.Error("获取商品sku库存失败!", zap.Error(err))
		response.FailWithMessage("获取sku库存失败", c)
		return
	}
	var productRes wechatRes.ProductDetailResponse
	productRes.Product = product
	productRes.ProductAttributeList = attributeList
	productRes.ProductAttributeValueList = attributeValue
	productRes.ProductFullReductionList = fullReductionList
	productRes.SkuStockList = skuStockList

	response.OkWithData(productRes, c)
}

func (e *HomeApi) GetProductList(c *gin.Context) {
	var pageInfo request.ProductSearchInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	recommendProductList, total, err := wechatService.GetProductList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:     recommendProductList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

func (e *HomeApi) GetProductListByOnlyID(c *gin.Context) {
	var pageInfo request.KeySearchInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	productList, total, err := wechatService.GetProductListByOnlyID(&pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:     productList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

func (e *HomeApi) GetProductListByOnlyIDWithSort(c *gin.Context) {
	var pageInfo request.SortSearchInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	productList, total, err := wechatService.GetProductListByOnlyIDWithSort(&pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:     productList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// UpdateProducts 更新商品
func (e *HomeApi) UpdateProducts(c *gin.Context) {
	var product wechatRequest.ProductUpdateRequest
	var ids, err = c.GetQueryArray("ids[]")
	if !err {
		response.FailWithMessage("Json Parse Error!!", c)
		return
	}
	var key, err1 = c.GetQuery("key")
	if !err1 {
		response.FailWithMessage("Json Parse Error!!", c)
		return
	}
	var value, err2 = c.GetQuery("value")
	if !err2 {
		response.FailWithMessage("Json Parse Error!!", c)
		return
	}
	var productIds []int
	for _, id := range ids {
		productId, _ := strconv.Atoi(id)
		productIds = append(productIds, productId)
	}
	product.Products = productIds
	product.Key = key
	product.Value, _ = strconv.Atoi(value)

	err3 := wechatService.UpdateHomeProducts(&product)
	if err3 != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err3))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// GetProductAttribute 获取商品分类
//func (e *HomeApi) GetProductAttributeCategoryList(c *gin.Context) {
//	var pageInfo request.PageInfo
//	err := c.ShouldBindQuery(&pageInfo)
//	if err != nil {
//		response.FailWithMessage(err.Error(), c)
//		return
//	}
//	if pageInfo.Page == 0 {
//		pageInfo.Page = 1
//	}
//	if pageInfo.PageSize == 0 {
//		pageInfo.PageSize = 100
//	}
//	categoryList, total, err := wechatService.GetProductAttributeCategoryList(pageInfo)
//	if err != nil {
//		global.GVA_LOG.Error("获取失败!", zap.Error(err))
//		response.FailWithMessage("获取失败"+err.Error(), c)
//		return
//	}
//
//	response.OkWithDetailed(response.PageResult{
//		List:     categoryList,
//		Total:    total,
//		Page:     pageInfo.Page,
//		PageSize: pageInfo.PageSize,
//	}, "获取成功", c)
//}

// GetProductBrand 获取品牌详情
func (e *HomeApi) GetProductBrandByID(c *gin.Context) {
	var reqId request.GetById
	err := c.ShouldBindQuery(&reqId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	data, err := wechatService.GetProductBrand(reqId.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithData(data, c)
}

// GetProductBrandList 获取商品品牌列表
func (e *HomeApi) GetProductBrandList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if pageInfo.Page == 0 {
		pageInfo.Page = 1
	}
	if pageInfo.PageSize == 0 {
		pageInfo.PageSize = 100
	}
	categoryList, total, err := wechatService.GetProductBrandList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:     categoryList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

func (e *HomeApi) CreateProductBrand(c *gin.Context) {
	var brand wechat.HomeBrand

	err := c.ShouldBindJSON(&brand)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = wechatService.CreateHomeProductBrand(&brand)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

func (e *HomeApi) UpdateProductBrand(c *gin.Context) {
	var brand wechat.HomeBrand
	err := c.ShouldBindJSON(&brand)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//home.SysUserAuthorityID = utils.GetUserAuthorityId(c)
	err = wechatService.UpdateHomeBrand(&brand)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func (e *HomeApi) DeleteHomeProductBrand(c *gin.Context) {
	var reqId request.GetById
	err := c.ShouldBindQuery(&reqId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(reqId, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = wechatService.DeleteHomeProductBrand(reqId.ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// CreateHotProduct 创建首页推荐专题商品
func (e *HomeApi) CreateHotProduct(c *gin.Context) {
	var hotProducts wechatRequest.AddHotProductRequest

	err := c.ShouldBindJSON(&hotProducts)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	hots := hotProducts.Products
	//var recommendData wechatRes.HomeRecommendResponse
	for _, hotProduct := range hots {
		product := wechat.HomeHotProduct{}
		product.ProductId = hotProduct.ProductId
		product.ProductName = hotProduct.ProductName
		product.RecommendStatus = 0
		product.Sort = 0
		err = wechatService.CreateHotProductBrand(&product)
		if err != nil {
			global.GVA_LOG.Error("创建失败!", zap.Error(err))
			response.FailWithMessage("创建失败", c)
			return
		}
	}
	response.OkWithMessage("创建成功", c)
}

// UpdateHotProducts 更新猜你喜欢商品
func (e *HomeApi) UpdateHotProducts(c *gin.Context) {
	var product wechatRequest.ProductUpdateRequest
	err := c.ShouldBindJSON(&product)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err3 := wechatService.UpdateHotProducts(&product)
	if err3 != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err3))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// DeleteHotProducts 删除猜你喜欢商品
func (e *HomeApi) DeleteHotProducts(c *gin.Context) {
	var product wechatRequest.ProductDeleteRequest
	err := c.ShouldBindJSON(&product)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err3 := wechatService.DeleteHotProducts(&product)
	if err3 != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err3))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// GetHotProductList 获取商品品牌列表
func (e *HomeApi) GetHotProductList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if pageInfo.Page == 0 {
		pageInfo.Page = 1
	}
	if pageInfo.PageSize == 0 {
		pageInfo.PageSize = 100
	}
	hotList, total, err := wechatService.GetHotProductList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:     hotList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetProductAttributeCategoryList 获取商品属性列表
func (e *HomeApi) GetProductAttributeCategoryList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if pageInfo.Page == 0 {
		pageInfo.Page = 1
	}
	if pageInfo.PageSize == 0 {
		pageInfo.PageSize = 100
	}
	attributeList, total, err := wechatService.GetProductAttributeCategoryList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:     attributeList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// CreateProductAttributeCategory 创建商品属性分类
func (e *HomeApi) CreateProductAttributeCategory(c *gin.Context) {
	var attribute wechat.HomeProductAttributeCategory

	err := c.ShouldBindJSON(&attribute)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = wechatService.CreateProductAttributeCategory(&attribute)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// UpdateProductAttributeCategory 更新商品属性分类
func (e *HomeApi) UpdateProductAttributeCategory(c *gin.Context) {
	var attribute wechat.HomeProductAttributeCategory
	err := c.ShouldBindJSON(&attribute)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//home.SysUserAuthorityID = utils.GetUserAuthorityId(c)
	err = wechatService.UpdateProductAttributeCategory(&attribute)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// DeleteProductAttributeCategory 删除商品属性分类
func (e *HomeApi) DeleteProductAttributeCategory(c *gin.Context) {
	var reqId request.GetById
	err := c.ShouldBindJSON(&reqId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(reqId, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = wechatService.DeleteProductAttributeCategory(reqId.ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// CreateProductAttribute 创建商品属性参数
func (e *HomeApi) CreateProductAttribute(c *gin.Context) {
	var attribute wechat.HomeProductAttribute

	err := c.ShouldBindJSON(&attribute)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = wechatService.CreateProductAttribute(&attribute)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// UpdateProductAttribute 更新商品属性参数
func (e *HomeApi) UpdateProductAttribute(c *gin.Context) {
	var attribute wechat.HomeProductAttribute
	err := c.ShouldBindJSON(&attribute)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//home.SysUserAuthorityID = utils.GetUserAuthorityId(c)
	err = wechatService.UpdateProductAttribute(&attribute)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// DeleteProductAttribute 删除商品属性参数
func (e *HomeApi) DeleteProductAttribute(c *gin.Context) {
	var reqId request.GetById
	err := c.ShouldBindJSON(&reqId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(reqId, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = wechatService.DeleteProductAttribute(reqId.ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// GetProductAttributeListByCategoryId 根据商品分类id获取商品属性参数列表
func (e *HomeApi) GetProductAttributeListByCategoryId(c *gin.Context) {
	var pageInfo request.TagSearchInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if pageInfo.Tag < 1 {
		response.FailWithMessage("商品属性分类id不可为空", c)
		return
	}
	if pageInfo.Page == 0 {
		pageInfo.Page = 1
	}
	if pageInfo.PageSize == 0 {
		pageInfo.PageSize = 100
	}
	attributeList, total, err := wechatService.GetProductAttributeList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:     attributeList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// CreateProductCategory 创建商品分类
func (e *HomeApi) CreateProductCategory(c *gin.Context) {
	var category wechat.HomeProductCategory
	err := c.ShouldBindJSON(&category)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = wechatService.CreateProductCategory(&category)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// UpdateProductCategory 更新商品分类
func (e *HomeApi) UpdateProductCategory(c *gin.Context) {
	var category wechat.HomeProductCategory
	err := c.ShouldBindJSON(&category)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//home.SysUserAuthorityID = utils.GetUserAuthorityId(c)
	err = wechatService.UpdateProductCategory(&category)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// DeleteProductCategory 删除商品分类
func (e *HomeApi) DeleteProductCategory(c *gin.Context) {
	var reqId request.GetById
	err := c.ShouldBindJSON(&reqId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(reqId, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = wechatService.DeleteProductCategory(reqId.ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// GetProductCategoryList 根据商品分类id获取商品分类
func (e *HomeApi) GetProductCategoryList(c *gin.Context) {
	var pageInfo request.TagSearchInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if pageInfo.Page == 0 {
		pageInfo.Page = 1
	}
	if pageInfo.PageSize == 0 {
		pageInfo.PageSize = 100
	}
	attributeList, total, err := wechatService.GetProductCategoryList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:     attributeList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetProductCategoryTreeList 获取商品分类树
func (e *HomeApi) GetProductCategoryTreeList(c *gin.Context) {
	categoryList, err := wechatService.GetAllProductCategoryList()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}

	var categoryTree []wechatRes.CategoryResponse
	for _, category := range categoryList {
		if category.ParentId == 0 {
			categoryTreeItem := wechatRes.CategoryResponse{}
			categoryTreeItem.Category = *category
			for _, productCategory := range categoryList {
				if category.ID == productCategory.ParentId {
					categoryTreeItem.Children = append(categoryTreeItem.Children, *productCategory)
				}
			}
			categoryTree = append(categoryTree, categoryTreeItem)
		}

	}
	response.OkWithData(categoryTree, c)
}

func (e *HomeApi) GetSkuStockByProductID(c *gin.Context) {
	var info request.KeyWordInfo
	err := c.ShouldBindQuery(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(info, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	skuStockList, err := wechatService.GetProductSKUStockByProductId(info.ID, info.Keyword)
	if err != nil {
		global.GVA_LOG.Error("获取商品sku库存失败!", zap.Error(err))
		response.FailWithMessage("获取sku库存失败", c)
		return
	}
	response.OkWithData(skuStockList, c)
}

func (e *HomeApi) UpdateSKUStock(c *gin.Context) {
	var stockList []wechat.SKUStock
	err := c.ShouldBindJSON(&stockList)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//home.SysUserAuthorityID = utils.GetUserAuthorityId(c)
	for _, stock := range stockList {
		err = wechatService.UpdateSKUStock(&stock)
		if err != nil {
			global.GVA_LOG.Error("更新失败!", zap.Error(err))
			response.FailWithMessage("更新失败", c)
			return
		}
	}
	response.OkWithMessage("更新成功", c)
}

func (e *HomeApi) GetProductCartList(c *gin.Context) {
	cartList, err := wechatService.GetProductCartList()
	if err != nil {
		global.GVA_LOG.Error("获取商品sku库存失败!", zap.Error(err))
		response.FailWithMessage("获取sku库存失败", c)
		return
	}
	response.OkWithData(cartList, c)
}

// CreateProductCart 创建商品购物车
func (e *HomeApi) CreateProductCart(c *gin.Context) {
	var cart wechat.CartItem
	err := c.ShouldBindJSON(&cart)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	cart.UserId = utils.GetUserID(c)
	cart.AuthorityId = utils.GetUserAuthorityId(c)
	println("--UserId-", cart.UserId)
	err = wechatService.CreateProductCart(&cart)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// UpdateProductCartQuantity 更新商品购物车数量
func (e *HomeApi) UpdateProductCartQuantity(c *gin.Context) {
	var quantityInfo request.QuantityInfo
	err := c.ShouldBindQuery(&quantityInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	err = wechatService.UpdateProductCartQuantity(userId, quantityInfo.ID, quantityInfo.Quantity)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// DeleteProductCartById 删除商品购物车
func (e *HomeApi) DeleteProductCartById(c *gin.Context) {
	var reqId request.QuantityInfo
	err := c.ShouldBindQuery(&reqId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	err = wechatService.DeleteProductCartById(userId, reqId.ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

// ClearProductCart 清空商品购物车
func (e *HomeApi) ClearProductCart(c *gin.Context) {
	userId := utils.GetUserID(c)
	println("--UserId-", userId)
	err := wechatService.ClearProductCartUserId(userId)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}

	response.OkWithMessage("删除成功", c)
}
