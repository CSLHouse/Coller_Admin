package request

type VIPMemberRequest struct {
	CardID     uint `json:"cardId" form:"cardId" gorm:"unique;comment:会员卡号"`   // 客户名
	ComboId    uint `json:"comboId" form:"comboId" gorm:"comment:套餐ID;size:9"` // 套餐ID
	Times      uint `json:"times" form:"times" gorm:"comment:赠送次数/金额"`         // 管理ID
	Collection uint `json:"collection" form:"collection" gorm:"comment:实付金额"`  // 管理角色ID
}
