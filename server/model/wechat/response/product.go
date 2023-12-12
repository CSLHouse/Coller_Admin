package response

type FlashPromotionResModel struct {
	NextStartTime int         `json:"nextStartTime"`
	EndTime       int         `json:"endTime"`
	ProductList   interface{} `json:"productList"`
}

type HomeContentResponse struct {
	AdvertiseList  interface{} `json:"advertiseList"`
	BrandList      interface{} `json:"brandList"`
	FlashPromotion interface{} `json:"homeFlashPromotion"`
	NewProductList interface{} `json:"newProductList"`
	//HotProductList interface{} `json:"hotProductList"`
}

type HomeRecommendResponse struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Pic      string  `json:"pic"`
	SubTitle string  `json:"subTitle"`
	Price    float32 `json:"price"`
}
