package request

type OrderCreateRequest struct {
	AppId                  string `json:"appId" gorm:"null;default null"`
	OpenId                 string `json:"openId" gorm:"not null;" binding:"required"`
	IP                     string `json:"ip" gorm:"not null;" `
	Ids                    []int  `json:"ids" gorm:"not null;"`
	CouponId               int    `json:"couponId" gorm:"null;default null"`
	MemberReceiveAddressId int    `json:"memberReceiveAddressId" gorm:"not null;"`
	PayType                int    `json:"payType" gorm:"null;default null;comment:支付方式"`
	UseIntegration         int    `json:"useIntegration" gorm:"null;default null"`
	Note                   string `json:"note" gorm:"null;default null"`
	BuyType                int    `json:"buyType" gorm:"null;default null;comment:购买类型，0直接购买，1购物车购买"`
}

type PaySuccessRequest struct {
	OrderId int `json:"orderId" gorm:"not null;" binding:"required"`
	PayType int `json:"payType" gorm:"not null;"`
}
