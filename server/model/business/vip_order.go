package business

type VIPOrder struct {
	ID        uint  `gorm:"primarykey"`
	OrderID   int64 `json:"orderId" form:"orderId" gorm:"unique;comment:订单编号"` // 客户名
	Telephone uint  `json:"telephone" form:"telephone" gorm:"comment:会员手机号"`   // 客户手机号
	//MemberName string `json:"memberName" form:"memberName" gorm:"comment:会员名"`   // 客户名
	//ComboId uint `json:"comboId" form:"comboId" gorm:"comment:套餐ID"` // 管理ID
	//ComboType          uint   `json:"comboType" form:"comboType" gorm:"comment:套餐类型"`    // 管理ID
	BuyDate string `json:"buyDate" form:"buyDate" gorm:"comment:购买日期"`
	State   uint   `json:"state" form:"state" gorm:"comment:状态"` // 管理角色ID
	IsNew   bool   `json:"isNew" form:"isNew" gorm:"comment:新会员"`
	Type    uint   `json:"type" form:"type" gorm:"comment:订单类型"`
	//Collection         uint      `json:"collection" form:"collection" gorm:"comment:实付金额"`                         // 管理角色ID
	SysUserAuthorityID uint      `json:"sys_user_authority_id" form:"sys_user_authority_id" gorm:"comment:管理角色ID"` // 管理角色ID
	Member             VIPMember `json:"member" gorm:"foreignKey:Telephone;references:Telephone;comment:会员信息"`
}

func (VIPOrder) TableName() string {
	return "bus_order"
}
