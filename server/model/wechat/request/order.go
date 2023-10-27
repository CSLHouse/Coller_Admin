package request

type OrderCreateRequest struct {
	CartIds                []int  `json:"cartIds" gorm:"not null;"`
	CouponId               int    `json:"couponId" gorm:"null;default null"`
	MemberReceiveAddressId int    `json:"memberReceiveAddressId" gorm:"not null;"`
	PayType                int    `json:"payType" gorm:"null;default null"`
	UseIntegration         int    `json:"useIntegration" gorm:"null;default null"`
	Note                   string `json:"note" gorm:"null;default null"`
}

type PaySuccessRequest struct {
	OrderId int `json:"orderId" gorm:"not null;" binding:"required"`
	PayType int `json:"payType" gorm:"not null;"`
}
