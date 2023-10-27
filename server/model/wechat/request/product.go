package request

import wechatModel "github.com/flipped-aurora/gin-vue-admin/server/model/wechat"

type ProductUpdateRequest struct {
	Products []int  `json:"products" gorm:"not null;comment:物品序号"`
	Key      string `json:"key" gorm:"not null；comment:各种状态：上架、新品、推荐"`
	Value    int    `json:"value" gorm:"not null；comment:状态值:0->不是；1->是"`
}

type ProductDeleteRequest struct {
	Products []int `json:"products" gorm:"not null;comment:物品序号"`
}

type AddHotProductRequest struct {
	Products []wechatModel.HomeHotProduct `json:"products" gorm:"not null"`
}

type MemberLevel struct {
	MemberLevelId   int     `json:"memberLevelId" gorm:"not null"`
	MemberLevelName string  `json:"memberLevelName" gorm:"not null"`
	MemberPrice     float32 `json:"memberPrice" `
}

type ProductAttributeValue struct {
	ProductAttributeId int    `json:"productAttributeId" gorm:"not null"`
	Value              string `json:"value" gorm:"not null"`
}

type FullReduction struct {
	FullPrice   float32 `json:"fullPrice" gorm:"not null"`
	ReducePrice float32 `json:"reducePrice" gorm:"not null"`
}

type LadderPrice struct {
	Count    int     `json:"count"`
	Discount float32 `json:"discount"`
	Price    float32 `json:"price" `
}

type SkuStock struct {
	LowStock       int     `json:"lowStock"`
	Pic            string  `json:"pic"`
	Price          float32 `json:"price" `
	PromotionPrice float32 `json:"promotionPrice" `
	SpData         string  `json:"spData"`
	Stock          int     `json:"stock"`
}
type ProductCreateRequest struct {
	Product                   wechatModel.HomeProduct `json:"product" gorm:"not null;comment:品牌ID"`
	MemberPriceList           []MemberLevel           `json:"memberPriceList" gorm:"not null;comment:品牌ID"`
	ProductAttributeValueList []ProductAttributeValue `json:"productAttributeValueList" gorm:"not null;comment:品牌ID"`
	ProductFullReductionList  []FullReduction         `json:"productFullReductionList" gorm:"comment:满减"` // 满减
	ProductLadderList         []LadderPrice           `json:"productLadderList" gorm:"comment:阶梯价格"`      // 阶梯价格
	SkuStockList              []SkuStock              `json:"skuStockList" gorm:"comment:阶梯价格"`
}
