package request

type CreateCardRequest struct {
	CardID      string `json:"cardId" form:"cardId" gorm:"unique;comment:会员卡号"`        // 客户名
	Telephone   string `json:"telephone" form:"telephone" gorm:"unique;comment:会员手机号"` // 客户手机号
	UserName    string `json:"userName" form:"userName" gorm:"comment:会员名"`            // 客户名
	ComboId     int    `json:"comboId" form:"comboId" gorm:"comment:套餐ID;size:9"`      // 套餐ID
	RemainTimes int    `json:"remainTimes" form:"remainTimes" gorm:"comment:剩余次数/金额"`  // 管理ID
	StartDate   string `json:"startDate" form:"startDate" gorm:"comment:开始日期"`
	Collection  int    `json:"collection" form:"collection" gorm:"comment:实付金额"`
}

type RenewCardRequest struct {
	ID         int `json:"id" form:"id" gorm:"comment:会员卡编号"`
	CardID     int `json:"cardId" form:"cardId" gorm:"unique;comment:会员卡号"`   // 客户名
	ComboId    int `json:"comboId" form:"comboId" gorm:"comment:套餐ID;size:9"` // 套餐ID
	Times      int `json:"times" form:"times" gorm:"comment:赠送次数/金额"`         // 管理ID
	Collection int `json:"collection" form:"collection" gorm:"comment:实付金额"`  // 管理角色ID
}
