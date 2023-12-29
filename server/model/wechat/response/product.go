package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/wechat"
	"time"
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
	//HotProductList interface{} `json:"hotProductList"`
}

type HomeRecommendResponse struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Pic      string  `json:"pic"`
	SubTitle string  `json:"subTitle"`
	Price    float32 `json:"price"`
}

type HomeFlashResponse struct {
	StartTime     time.Time `json:"startTime"`
	EndTime       time.Time `json:"endTime"`
	NextStartTime time.Time `json:"nextStartTime"`
	NextEndTime   time.Time `json:"nextEndTime"`
	//HotProductList interface{} `json:"hotProductList"`
	ProductList []wechat.Product `json:"productList"`
}
