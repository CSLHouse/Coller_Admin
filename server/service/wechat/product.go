package wechat

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wechat"
	wechatRequest "github.com/flipped-aurora/gin-vue-admin/server/model/wechat/request"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strings"
)

type HomeService struct{}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateHomeAdvertise
//@description: 创建套餐
//@param: e model.HomeAdvertise
//@return: err error

func (exa *HomeService) CreateHomeAdvertise(e wechat.HomeAdvertise) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteHomeAdvertise
//@description: 删除套餐
//@param: e model.HomeAdvertise
//@return: err error

func (exa *HomeService) DeleteHomeAdvertise(e wechat.HomeAdvertise) (err error) {
	err = global.GVA_DB.Delete(&e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateHomeAdvertise
//@description: 更新套餐
//@param: e *model.HomeAdvertise
//@return: err error

func (exa *HomeService) UpdateHomeAdvertise(e *wechat.HomeAdvertise) (err error) {
	err = global.GVA_DB.Save(e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetHomeAdvertise
//@description: 获取套餐信息
//@param: id uint
//@return: customer model.HomeAdvertise, err error

func (exa *HomeService) GetHomeAdvertise(id int) (customer wechat.HomeAdvertise, err error) {
	err = global.GVA_DB.Where("combo_id = ?", id).First(&customer).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetHomeAdvertiseInfoList
//@description: 获取套餐列表
//@param: sysUserAuthorityID string, info request.PageInfo
//@return: list interface{}, total int64, err error

func (exa *HomeService) GetHomeAdvertiseInfoList(info request.PageInfo) (list interface{}, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&wechat.HomeAdvertise{})
	if err != nil {
		return
	}

	var ComboList []wechat.HomeAdvertise
	err = db.Limit(limit).Offset(offset).Find(&ComboList).Error
	return ComboList, err
}

func (exa *HomeService) GetAllHomeAdvertiseInfoList() (list interface{}, err error) {
	db := global.GVA_DB.Model(&wechat.HomeAdvertise{})
	//var a system.SysAuthority
	//a.AuthorityId = sysUserAuthorityID
	//auth, err := systemService.AuthorityServiceApp.GetAuthorityInfo(a)
	//if err != nil {
	//	return
	//}
	//var dataId []uint
	//for _, v := range auth.DataAuthorityId {
	//	dataId = append(dataId, v.AuthorityId)
	//}
	var advertiseList []wechat.HomeAdvertise
	err = db.Find(&advertiseList).Error
	return advertiseList, err
}

func (exa *HomeService) GetAllHomeBrandInfoList() (list interface{}, err error) {
	db := global.GVA_DB.Model(&wechat.HomeBrand{})
	//var a system.SysAuthority
	//a.AuthorityId = sysUserAuthorityID
	//auth, err := systemService.AuthorityServiceApp.GetAuthorityInfo(a)
	//if err != nil {
	//	return
	//}
	//var dataId []uint
	//for _, v := range auth.DataAuthorityId {
	//	dataId = append(dataId, v.AuthorityId)
	//}
	var brandList []wechat.HomeBrand
	err = db.Find(&brandList).Error
	return brandList, err
}

func (exa *HomeService) GetAllHomeFlashPromotionInfoList() (list interface{}, err error) {
	db := global.GVA_DB.Model(&wechat.HomeFlashPromotion{})
	//var a system.SysAuthority
	//a.AuthorityId = sysUserAuthorityID
	//auth, err := systemService.AuthorityServiceApp.GetAuthorityInfo(a)
	//if err != nil {
	//	return
	//}
	//var dataId []uint
	//for _, v := range auth.DataAuthorityId {
	//	dataId = append(dataId, v.AuthorityId)
	//}
	var flashPromotionList []wechat.HomeFlashPromotion
	err = db.Find(&flashPromotionList).Error
	return flashPromotionList, err
}

func (exa *HomeService) CreateNewProduct(e *wechat.HomeNewProduct) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

func (exa *HomeService) GetAllNewProductInfoList() (list interface{}, err error) {
	db := global.GVA_DB.Model(&wechat.HomeNewProduct{})
	//var a system.SysAuthority
	//a.AuthorityId = sysUserAuthorityID
	//auth, err := systemService.AuthorityServiceApp.GetAuthorityInfo(a)
	//if err != nil {
	//	return
	//}
	//var dataId []uint
	//for _, v := range auth.DataAuthorityId {
	//	dataId = append(dataId, v.AuthorityId)
	//}
	var newProduct []wechat.HomeNewProduct
	err = db.Find(&newProduct).Error
	return newProduct, err
}

func (exa *HomeService) GetAllHomeHotProductListInfoList() (list interface{}, err error) {
	db := global.GVA_DB.Model(&wechat.HomeHotProduct{})
	//var a system.SysAuthority
	//a.AuthorityId = sysUserAuthorityID
	//auth, err := systemService.AuthorityServiceApp.GetAuthorityInfo(a)
	//if err != nil {
	//	return
	//}
	//var dataId []uint
	//for _, v := range auth.DataAuthorityId {
	//	dataId = append(dataId, v.AuthorityId)
	//}
	var hotProductList []wechat.HomeHotProduct
	err = db.Find(&hotProductList).Error
	return hotProductList, err
}

func (exa *HomeService) CreateRecommendProduct(e *wechat.HomeRecommendProduct) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

func (exa *HomeService) GetRecommendProductList(pageInfo request.PageInfo) (recommendProductList []wechat.HomeRecommendProduct, err error) {
	limit := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)

	db := global.GVA_DB.Model(&wechat.HomeRecommendProduct{})
	//var a system.SysAuthority
	//a.AuthorityId = sysUserAuthorityID
	//auth, err := systemService.AuthorityServiceApp.GetAuthorityInfo(a)
	//if err != nil {
	//	return
	//}
	//var dataId []uint
	//for _, v := range auth.DataAuthorityId {
	//	dataId = append(dataId, v.AuthorityId)
	//}
	err = db.Limit(limit).Offset(offset).Preload("Product").Find(&recommendProductList).Error
	return recommendProductList, err
}

func (exa *HomeService) CreateHomeProduct(e *wechat.HomeProduct) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

func (exa *HomeService) UpdateHomeProducts(e *wechatRequest.ProductUpdateRequest) (err error) {
	db := global.GVA_DB.Model(&wechat.HomeProduct{})
	db.Where("id IN ?", e.Products).Updates(map[string]interface{}{e.Key: e.Value})
	return err
}

// 获取查询数据库的命令
func (exa *HomeService) getProductSearchCmd(searchInfo request.ProductSearchInfo) string {
	var cmdName string
	var cmdSN string
	var cmdCategoryName string
	var cmdBrand string
	var cmdPublishStatus string
	var cmdVerifyStatus string
	if len(searchInfo.Name) > 0 {
		cmdName += fmt.Sprintf("name like '%%%s%%'", searchInfo.Name)
	}
	if searchInfo.BrandId > 0 {
		cmdSN += fmt.Sprintf("brand_id = %d", searchInfo.BrandId)
	}
	if len(searchInfo.ProductSN) > 0 {
		cmdSN += fmt.Sprintf("product_sn like '%%%s%%'", strings.TrimSpace(searchInfo.ProductSN))
	}
	if len(searchInfo.ProductCategoryName) > 0 {
		cmdCategoryName += fmt.Sprintf("product_category_name like '%%%s%%'", strings.TrimSpace(searchInfo.ProductCategoryName))
	}
	if len(searchInfo.BrandName) > 0 {
		cmdBrand += fmt.Sprintf("brand_name like '%%%s%%'", strings.TrimSpace(searchInfo.BrandName))
	}
	if searchInfo.PublishStatus > 0 {
		cmdPublishStatus += fmt.Sprintf("publish_status = %d", searchInfo.PublishStatus)
	}
	if searchInfo.VerifyStatus > 0 {
		cmdVerifyStatus += fmt.Sprintf("verify_status = %d", searchInfo.VerifyStatus)
	}

	cmdSearch := ""
	cmds := [6]string{cmdName, cmdSN, cmdCategoryName, cmdBrand, cmdPublishStatus, cmdVerifyStatus}
	isFirst := true
	for _, cmd := range cmds {
		if len(cmd) > 0 {
			if isFirst {
				cmdSearch += cmd
				isFirst = false
			} else {
				cmdSearch += " and " + cmd
			}
		}
	}
	return cmdSearch
}

func (exa *HomeService) GetProductByID(id int) (product wechat.HomeProduct, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&product).Error
	return product, err
}

func (exa *HomeService) GetProductList(searchInfo request.ProductSearchInfo) (list interface{}, total int64, err error) {
	limit := searchInfo.PageSize
	offset := searchInfo.PageSize * (searchInfo.Page - 1)

	var productList []wechat.HomeProduct
	homeServiceApp := new(HomeService)
	cmd := homeServiceApp.getProductSearchCmd(searchInfo)

	db := global.GVA_DB.Model(&wechat.HomeProduct{})
	err = db.Where(cmd).Count(&total).Error
	if err != nil {
		return productList, total, err
	} else {
		err = db.Debug().Limit(limit).Offset(offset).Debug().Where(cmd).Find(&productList).Error
	}

	return productList, total, err
}

func (exa *HomeService) GetProductListByOnlyID(searchInfo *request.KeySearchInfo) (productList []wechat.HomeProduct, total int64, err error) {
	limit := searchInfo.PageSize
	offset := searchInfo.PageSize * (searchInfo.Page - 1)

	db := global.GVA_DB.Model(&wechat.HomeProduct{})
	cmd := fmt.Sprintf("%s = %d", searchInfo.Key, searchInfo.ID)
	err = db.Where(cmd).Count(&total).Error
	if err != nil {
		return productList, total, err
	} else {
		err = db.Debug().Limit(limit).Offset(offset).Debug().Where(cmd).Find(&productList).Error
	}
	return productList, total, err
}

func (exa *HomeService) GetProductListByOnlyIDWithSort(searchInfo *request.SortSearchInfo) (productList []wechat.HomeProduct, total int64, err error) {
	limit := searchInfo.PageSize
	offset := searchInfo.PageSize * (searchInfo.Page - 1)

	db := global.GVA_DB.Model(&wechat.HomeProduct{})

	cmd := fmt.Sprintf("%s = %d", searchInfo.Key, searchInfo.ID)
	var orderCmd string
	switch searchInfo.Sort {
	case 0:
		orderCmd = fmt.Sprintf("`sale` desc, `price` asc")
	case 2:
		orderCmd = fmt.Sprintf("`sale` desc")
	case 3:
		orderCmd = fmt.Sprintf("`price` asc")
	case 4:
		orderCmd = fmt.Sprintf("`price` desc")
	default:
		global.GVA_LOG.Error(fmt.Sprintf("sort参数错误,sort:%d", searchInfo.Sort), zap.Error(err))
		return productList, total, err
	}
	err = db.Where(cmd).Count(&total).Error
	if err != nil {
		return productList, total, err
	} else {
		err = db.Debug().Limit(limit).Offset(offset).Where(cmd).Order(orderCmd).Find(&productList).Error
	}
	return productList, total, err
}

func (exa *HomeService) GetProductBrand(id int) (brand wechat.HomeBrand, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&brand).Error
	return
}

func (exa *HomeService) GetProductBrandList(searchInfo request.PageInfo) (list interface{}, total int64, err error) {
	limit := searchInfo.PageSize
	offset := searchInfo.PageSize * (searchInfo.Page - 1)

	var productList []wechat.HomeBrand

	db := global.GVA_DB.Model(&wechat.HomeBrand{})
	err = db.Count(&total).Error
	if err != nil {
		return productList, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Find(&productList).Error
	}

	return productList, total, err
}

func (exa *HomeService) CreateHomeProductBrand(e *wechat.HomeBrand) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

func (exa *HomeService) UpdateHomeBrand(e *wechat.HomeBrand) (err error) {
	err = global.GVA_DB.Save(e).Error
	return err
}

func (exa *HomeService) DeleteHomeProductBrand(e wechat.HomeBrand) (err error) {
	err = global.GVA_DB.Delete(&e).Error
	return err
}

func (exa *HomeService) CreateHotProductBrand(e *wechat.HomeHotProduct) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

func (exa *HomeService) UpdateHotProducts(e *wechatRequest.ProductUpdateRequest) (err error) {
	db := global.GVA_DB.Model(&wechat.HomeHotProduct{})
	db.Debug().Where("id IN ?", e.Products).Updates(map[string]interface{}{e.Key: e.Value})
	return err
}

func (exa *HomeService) DeleteHotProducts(e *wechatRequest.ProductDeleteRequest) (err error) {
	db := global.GVA_DB.Model(&wechat.HomeHotProduct{})
	db.Debug().Where("id IN ?", e.Products).Delete(&wechat.HomeHotProduct{})
	return err
}

func (exa *HomeService) GetHotProductList(searchInfo request.PageInfo) (list interface{}, total int64, err error) {
	limit := searchInfo.PageSize
	offset := searchInfo.PageSize * (searchInfo.Page - 1)

	var productList []wechat.HomeHotProduct

	db := global.GVA_DB.Model(&wechat.HomeHotProduct{})
	err = db.Count(&total).Error
	if err != nil {
		return productList, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Find(&productList).Error
	}

	return productList, total, err
}

func (exa *HomeService) CreateProductAttributeCategory(e *wechat.HomeProductAttributeCategory) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

func (exa *HomeService) UpdateProductAttributeCategory(e *wechat.HomeProductAttributeCategory) (err error) {
	err = global.GVA_DB.Save(e).Error
	return err
}

func (exa *HomeService) DeleteProductAttributeCategory(id int) (err error) {
	var attribute wechat.HomeProductAttributeCategory
	err = global.GVA_DB.Where("id = ?", id).Delete(&attribute).Error
	return err
}

func (exa *HomeService) GetProductAttributeCategoryList(searchInfo request.PageInfo) (list interface{}, total int64, err error) {
	limit := searchInfo.PageSize
	offset := searchInfo.PageSize * (searchInfo.Page - 1)

	var productList []wechat.HomeProductAttributeCategory

	db := global.GVA_DB.Model(&wechat.HomeProductAttributeCategory{})
	err = db.Count(&total).Error
	if err != nil {
		return productList, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Find(&productList).Error
	}

	return productList, total, err
}

func (exa *HomeService) GetProductAttributeListById(id int) (list []wechat.HomeProductAttribute, err error) {
	var attribute []wechat.HomeProductAttribute
	db := global.GVA_DB.Model(&wechat.HomeProductAttribute{})
	err = db.Where("product_attribute_category_id = ?", id).Find(&attribute).Error
	return attribute, err
}

func (exa *HomeService) CreateProductAttribute(e *wechat.HomeProductAttribute) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

func (exa *HomeService) UpdateProductAttribute(e *wechat.HomeProductAttribute) (err error) {
	err = global.GVA_DB.Save(e).Error
	return err
}

func (exa *HomeService) DeleteProductAttribute(id int) (err error) {
	var attribute wechat.HomeProductAttribute
	err = global.GVA_DB.Where("id = ?", id).Delete(&attribute).Error
	return err
}

func (exa *HomeService) GetProductAttributeList(searchInfo request.TagSearchInfo) (list interface{}, total int64, err error) {
	limit := searchInfo.PageSize
	offset := searchInfo.PageSize * (searchInfo.Page - 1)

	var productList []wechat.HomeProductAttribute

	db := global.GVA_DB.Model(&wechat.HomeProductAttribute{})
	err = db.Where("product_attribute_category_id=? and type=?", searchInfo.Tag, searchInfo.State).Count(&total).Error
	if err != nil {
		return productList, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Debug().Where("product_attribute_category_id=? and type=?", searchInfo.Tag, searchInfo.State).Find(&productList).Error
	}

	return productList, total, err
}

func (exa *HomeService) CreateProductCategory(e *wechat.HomeProductCategory) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

func (exa *HomeService) UpdateProductCategory(e *wechat.HomeProductCategory) (err error) {
	err = global.GVA_DB.Save(e).Error
	return err
}

func (exa *HomeService) DeleteProductCategory(id int) (err error) {
	var category wechat.HomeProductCategory
	err = global.GVA_DB.Where("id = ?", id).Delete(&category).Error
	return err
}

func (exa *HomeService) GetProductCategoryList(searchInfo request.TagSearchInfo) (productList []wechat.HomeProductCategory, total int64, err error) {
	limit := searchInfo.PageSize
	offset := searchInfo.PageSize * (searchInfo.Page - 1)

	db := global.GVA_DB.Model(&wechat.HomeProductCategory{})
	err = db.Where("parent_id = ?", searchInfo.Tag).Count(&total).Error
	if err != nil {
		return productList, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Debug().Where("parent_id = ?", searchInfo.Tag).Find(&productList).Error
	}

	return productList, total, err
}

func (exa *HomeService) GetAllProductCategoryList() (productList []*wechat.HomeProductCategory, err error) {
	db := global.GVA_DB.Model(&wechat.HomeProductCategory{})
	err = db.Find(&productList).Error
	return productList, err
}

func (exa *HomeService) GetProductAttributeValueByProductId(id int) (list []wechat.HomeProductAttributeValue, err error) {
	db := global.GVA_DB.Model(&wechat.HomeProductAttributeValue{})
	err = db.Where("product_id = ?", id).Find(&list).Error
	return list, err
}

func (exa *HomeService) CreateProductMemberPrice(e *wechat.MemberPrice) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

func (exa *HomeService) CreateProductAttributeValue(e *wechat.HomeProductAttributeValue) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

func (exa *HomeService) GetProductFullReductionByProductId(id int) (list []wechat.ProductFullReduction, err error) {
	db := global.GVA_DB.Model(&wechat.ProductFullReduction{})
	err = db.Where("product_id = ?", id).Find(&list).Error
	return list, err
}

func (exa *HomeService) CreateProductFullReduction(e *wechat.ProductFullReduction) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

func (exa *HomeService) CreateProductLadder(e *wechat.ProductLadder) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

func (exa *HomeService) CreateProductSKUStock(e *wechat.SKUStock) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

func (exa *HomeService) GetProductSKUStockByProductId(id int, keyword string) (list []wechat.SKUStock, err error) {
	db := global.GVA_DB.Model(&wechat.SKUStock{})
	cmd := fmt.Sprintf("product_id = %d", id)
	if len(keyword) > 0 {
		cmd += fmt.Sprintf("and sku_code = %s", keyword)
	}
	err = db.Where(cmd).Find(&list).Error
	return list, err
}

func (exa *HomeService) GetProductSKUStockById(id int) (stock wechat.SKUStock, err error) {
	db := global.GVA_DB.Model(&wechat.SKUStock{})
	err = db.Where("id = ?", id).First(&stock).Error
	return stock, err
}

func (exa *HomeService) UpdateSKUStock(e *wechat.SKUStock) (err error) {
	err = global.GVA_DB.Save(e).Error
	return err
}

func (exa *HomeService) GetProductCartList() (list []wechat.CartItem, err error) {
	db := global.GVA_DB.Model(&wechat.CartItem{})
	err = db.Find(&list).Error
	return list, err
}

func (exa *HomeService) CreateProductCart(e *wechat.CartItem) (err error) {
	db := global.GVA_DB.Model(&wechat.CartItem{})
	var cart wechat.CartItem
	result := db.Where("user_id = ? and product_id = ? and product_sku_id = ?", e.UserId, e.ProductId, e.ProductSkuId).First(&cart)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			err = global.GVA_DB.Create(&e).Error
			return err
		}
		err = result.Error
	} else {
		err = db.Debug().Where("id = ? and user_id = ?", cart.ID, e.UserId).Update("quantity", cart.Quantity+1).Error
		return err
	}
	return err
}

func (exa *HomeService) UpdateProductCartQuantity(userId int, id int, quantity int) (err error) {
	db := global.GVA_DB.Model(&wechat.CartItem{})
	db.Debug().Where("user_id = ? and id = ?", userId, id).UpdateColumn("quantity", quantity)
	return err
}

func (exa *HomeService) DeleteProductCartById(userId int, id int) (err error) {
	var cart wechat.CartItem
	err = global.GVA_DB.Where("user_id = ? and id = ?", userId, id).Delete(&cart).Error
	return err
}

func (exa *HomeService) ClearProductCartUserId(id int) (err error) {
	var cart wechat.CartItem
	err = global.GVA_DB.Where("user_id = ?", id).Delete(&cart).Error
	return err
}
