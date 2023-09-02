package response

type VipOrderResModel struct {
	ID         uint   `gorm:"primarykey"`
	OrderID    int64  `json:"orderId" form:"orderId" gorm:"comment:订单编号"`       // 客户名
	Telephone  uint   `json:"telephone" form:"telephone" gorm:"comment:会员手机号"`  // 客户手机号
	MemberName string `json:"memberName" form:"memberName" gorm:"comment:会员名"`  // 客户名
	ComboId    uint   `json:"comboId" form:"comboId" gorm:"comment:套餐ID"`       // 管理ID
	ComboType  string `json:"comboType" form:"comboType" gorm:"comment:套餐类型"`   // 管理ID
	ComboPrice uint   `json:"comboPrice" form:"comboPrice" gorm:"comment:套餐类型"` // 管理ID
	BuyDate    string `json:"buyDate" form:"buyDate" gorm:"comment:购买日期"`
	State      uint   `json:"state" form:"state" gorm:"comment:状态"` // 管理角色ID
	IsNew      bool   `json:"isNew" form:"isNew" gorm:"comment:新会员"`
	Type       uint   `json:"type" form:"type" gorm:"comment:订单类型"`
	Collection uint   `json:"collection" form:"collection" gorm:"comment:实付金额"`
}
type VipOrderResponse struct {
	Order VipOrderResModel `json:"order"`
}
