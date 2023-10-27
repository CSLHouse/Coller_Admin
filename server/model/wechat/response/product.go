package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/wechat"
)

type FlashPromotionResModel struct {
	NextStartTime int         `json:"nextStartTime"`
	EndTime       int         `json:"endTime"`
	ProductList   interface{} `json:"productList"`
}

type HomeContentResponse struct {
	AdvertiseList      interface{} `json:"advertiseList"`
	BrandList          interface{} `json:"brandList"`
	HomeFlashPromotion interface{} `json:"homeFlashPromotion"`
	NewProductList     interface{} `json:"newProductList"`
	HotProductList     interface{} `json:"hotProductList"`
}

type HomeRecommendResponse struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Pic      string  `json:"pic"`
	SubTitle string  `json:"subTitle"`
	Price    float32 `json:"price"`
}

type ProductDetailResponse struct {
	Product                   wechat.HomeProduct                 `json:"product"`
	Brand                     wechat.HomeBrand                   `json:"brand"`
	ProductAttributeList      []wechat.HomeProductAttribute      `json:"productAttributeList"`
	ProductAttributeValueList []wechat.HomeProductAttributeValue `json:"productAttributeValueList"`
	SkuStockList              []wechat.SKUStock                  `json:"skuStockList"`
	ProductFullReductionList  []wechat.ProductFullReduction      `json:"productFullReductionList"`
	//Brand                     wechat.HomeBrand              `json:"couponList"`
}

type CategoryResponse struct {
	Category wechat.HomeProductCategory   `json:"category" `
	Children []wechat.HomeProductCategory `json:"children"`
}
