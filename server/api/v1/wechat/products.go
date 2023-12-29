package wechat

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wechat"
	wechatRequest "github.com/flipped-aurora/gin-vue-admin/server/model/wechat/request"
	wechatRes "github.com/flipped-aurora/gin-vue-admin/server/model/wechat/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

type HomeApi struct{}

func (e *HomeApi) CreateHomeAdvertise(c *gin.Context) {
	var home wechat.Advertise
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
	var reqId request.GetById
	err := c.ShouldBindQuery(&reqId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	advertise, err := wechatService.GetHomeAdvertiseById(reqId.ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	var file example.ExaFileUploadAndDownload
	file.Url = advertise.Pic
	if err := fileUploadAndDownloadService.DeleteFile(file); err != nil {
		global.GVA_LOG.Error(fmt.Sprintf("阿里云图片删除失败:%s", file.Url), zap.Error(err))
	}

	err = wechatService.DeleteHomeAdvertise(reqId.ID)
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
	var home wechat.Advertise
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

func (e *HomeApi) UpdateHomeAdvertiseOnlineState(c *gin.Context) {
	var advertise wechatRequest.UpdateIdsKeywordRequest
	err := c.ShouldBindJSON(&advertise)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = wechatService.UpdateHomeAdvertiseByIdForKeyword(&advertise)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
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
	homeList, total, err := wechatService.GetHomeAdvertiseInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     homeList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

func (e *HomeApi) FindValidFlashPromotion(promotionList []wechat.FlashPromotion) interface{} {
	for _, promotion := range promotionList {
		if promotion.Status > 0 {
			t := time.Now().Unix()
			startDate := promotion.StartDate.Unix()
			endDate := promotion.EndDate.Unix()
			if t > startDate && t < endDate {
				return promotion
			}
		}
	}
	return nil
}

func (e *HomeApi) FindCurrentFlashPromotionSession(sessionList []wechat.FlashPromotionSession) (sessionRes wechat.FlashPromotionSession, err error) {
	t := time.Now().Unix()
	for _, session := range sessionList {
		if session.Status > 0 {
			startDate := session.StartTime.Unix()
			endDate := session.EndTime.Unix()
			if t > startDate && t < endDate {
				return session, nil
			}
		}
	}
	return sessionRes, fmt.Errorf("not found!")
}

func (e *HomeApi) FindNextFlashPromotionSession(sessionList []wechat.FlashPromotionSession) (sessionRes wechat.FlashPromotionSession, err error) {
	t := time.Now().Unix()
	mini := 0
	miniTime := int64(0)
	for i, session := range sessionList {
		if session.Status > 0 {
			startDate := session.StartTime.Unix()
			if t < startDate && miniTime < startDate {
				mini = i
				miniTime = startDate
			}
		}
	}
	if mini != 0 {
		return sessionList[mini], nil
	}
	return sessionRes, fmt.Errorf("not found!")
}

func (e *HomeApi) GetAllWechatContent(c *gin.Context) {
	homeList, err := wechatService.GetOnlineHomeAdvertiseInfoList()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	brandList, err := wechatService.GetOnlineHomeBrandInfoList()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}

	newProductList, err := wechatService.GetOnlineNewProductInfoList()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}

	flashPromotionList, err := wechatService.GetOnlineHomeFlashPromotionInfoList()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}

	flashPromotion := e.FindValidFlashPromotion(flashPromotionList)
	var homeFlashPromotion wechatRes.HomeFlashResponse
	if flashPromotion == nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
	} else {
		flashSessionList, err := wechatService.GetFlashSessionList()
		if err != nil {
			global.GVA_LOG.Error("GetFlashSessionList获取失败!", zap.Error(err))
		} else {
			currentFlash, err := e.FindCurrentFlashPromotionSession(flashSessionList)
			nextFlash, err1 := e.FindNextFlashPromotionSession(flashSessionList)
			if err == nil {
				homeFlashPromotion.StartTime = currentFlash.StartTime
				homeFlashPromotion.EndTime = currentFlash.EndTime
				sessionProductList, err := wechatService.GetFlashPromotionProductRelationListById(currentFlash.ID, flashPromotion.(wechat.FlashPromotion).ID)
				if err != nil {
					global.GVA_LOG.Error("GetFlashPromotionProductRelationListById获取失败!", zap.Error(err))
				} else {
					productList := make([]wechat.Product, 0)
					for _, sessionProduct := range sessionProductList {
						productList = append(productList, sessionProduct.Product)
					}
					homeFlashPromotion.ProductList = productList
				}
			}
			if err1 == nil {
				homeFlashPromotion.NextStartTime = nextFlash.StartTime
				homeFlashPromotion.NextEndTime = nextFlash.EndTime
			}
		}

	}
	//hotProductList, err := wechatService.GetOnlineRecommendProductListInfoList()
	//if err != nil {
	//	global.GVA_LOG.Error("获取失败!", zap.Error(err))
	//	response.FailWithMessage("获取失败"+err.Error(), c)
	//	return
	//}
	response.OkWithDetailed(wechatRes.HomeContentResponse{
		AdvertiseList:  homeList,
		BrandList:      brandList,
		NewProductList: newProductList,
		//HotProductList: hotProductList,
		HomeFlashPromotion: homeFlashPromotion,
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
	recommendProductList, err := wechatService.GetOnlineRecommendProductListInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	var list []wechat.Product
	for _, recommend := range recommendProductList {
		if recommend.Product.ID != 0 {
			list = append(list, recommend.Product)
		}
	}
	response.OkWithDetailed(response.AllResult{
		List: list,
	}, "获取成功", c)
}

func (e *HomeApi) CreateProduct(c *gin.Context) {
	var product wechat.Product
	err := c.ShouldBindJSON(&product)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = wechatService.CreateHomeProduct(&product)
	if err != nil {
		global.GVA_LOG.Error("创建产品失败!", zap.Error(err))
		response.FailWithMessage("创建产品失败", c)
		return
	}
	// 新品
	if product.NewStatus == 1 {
		var newProduct wechat.NewProduct
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
		var ewcommendProduct wechat.RecommendProduct
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

	response.OkWithData(product, c)
}

func (e *HomeApi) GetProductList(c *gin.Context) {
	var pageInfo wechatRequest.ProductSearchInfo
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
func (e *HomeApi) DeleteProducts(c *gin.Context) {
	var idsInfo request.IdsReq
	err := c.ShouldBindJSON(&idsInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = wechatService.DeleteProducts(idsInfo.Ids)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}

	response.OkWithMessage("删除成功", c)
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
func (e *HomeApi) UpdateProduct(c *gin.Context) {
	var product wechat.Product
	var err = c.ShouldBindJSON(&product)
	if err != nil {
		response.FailWithMessage("Json Parse Error!!", c)
		return
	}

	err3 := wechatService.UpdateHomeProduct(&product)
	if err3 != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err3))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// UpdateProductForKeyword 更新商品
func (e *HomeApi) UpdateProductForKeyword(c *gin.Context) {
	var product wechatRequest.UpdateIdsKeywordRequest
	var err = c.ShouldBindJSON(&product)
	if err != nil {
		response.FailWithMessage("Json Parse Error!!", c)
		return
	}

	err3 := wechatService.UpdateProductForKeyword(&product)
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
	var pageInfo wechatRequest.BrandSearchInfo
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
	brandList, total, err := wechatService.GetProductBrandList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:     brandList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

func (e *HomeApi) CreateProductBrand(c *gin.Context) {
	var brand wechat.Brand

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
	var brand wechat.Brand
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

func (e *HomeApi) UpdateProductBrandOnlineStat(c *gin.Context) {
	var brand wechatRequest.UpdateIdsKeywordRequest
	err := c.ShouldBindJSON(&brand)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = wechatService.UpdateHomeBrandByIdForKeyword(&brand)
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

// CreateRecommendProduct 创建推荐商品
func (e *HomeApi) CreateRecommendProducts(c *gin.Context) {
	var recommendProducts wechatRequest.AddRecommendProductRequest
	err := c.ShouldBindJSON(&recommendProducts)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	hots := recommendProducts.Products
	for _, hotProduct := range hots {
		product := wechat.RecommendProduct{}
		product.ProductId = hotProduct.ProductId
		product.ProductName = hotProduct.ProductName
		product.RecommendStatus = 0
		product.Sort = 0
		err = wechatService.CreateRecommendProduct(&product)
		if err != nil {
			global.GVA_LOG.Error("创建失败!", zap.Error(err))
			response.FailWithMessage("创建失败", c)
			return
		}
	}
	response.OkWithMessage("创建成功", c)
}

// UpdateRecommendProducts 更新人气推荐商品
func (e *HomeApi) UpdateRecommendProducts(c *gin.Context) {
	var product wechatRequest.UpdateIdsKeywordRequest
	err := c.ShouldBindJSON(&product)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err3 := wechatService.UpdateRecommendProducts(&product)
	if err3 != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err3))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// DeleteHotProducts 删除猜你喜欢商品
func (e *HomeApi) DeleteRecommendProducts(c *gin.Context) {
	var idsReq request.IdsReq
	err := c.ShouldBindJSON(&idsReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err3 := wechatService.DeleteRecommendProducts(idsReq.Ids)
	if err3 != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err3))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
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
	var attribute wechat.ProductAttributeCategory

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
	var attribute wechat.ProductAttributeCategory
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
	var attribute wechat.ProductAttribute

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
	var attribute wechat.ProductAttribute
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
	var category wechat.ProductCategory
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
	var category wechat.ProductCategory
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
func (e *HomeApi) GetProductAllCategory(c *gin.Context) {
	categoryList, err := wechatService.GetAllProductCategoryList()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}

	response.OkWithData(categoryList, c)
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
	var stockList []wechat.SkuStock
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

func (e *HomeApi) DeleteProductCartByIds(c *gin.Context) {
	var reqIds request.IdsReq
	err := c.ShouldBindQuery(&reqIds)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	for _, id := range reqIds.Ids {
		err = wechatService.DeleteProductCartById(userId, id)
		if err != nil {
			global.GVA_LOG.Error("删除失败!", zap.Error(err))
			response.FailWithMessage("删除失败", c)
			return
		}
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
