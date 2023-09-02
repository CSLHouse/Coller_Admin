package business

type VIPStatement struct {
	Date               string `json:"date" form:"date" gorm:"primary_key;comment:日期;"`                          // 分店名称
	Recharge           uint   `json:"recharge" form:"recharge" gorm:"comment:会员卡流水"`                            // 套餐名称
	CardNumber         uint   `json:"cardNumber" form:"cardNumber" gorm:"comment:会员卡单量"`                        // 套餐类型
	NewMember          uint   `json:"newMember" form:"newMember" gorm:"comment:新增会员"`                           // 套餐价格
	ConsumeNumber      uint   `json:"consumeNumber" form:"consumeNumber" gorm:"comment:入店统计"`                   // 天数/次数/金额
	SysUserAuthorityID uint   `json:"sys_user_authority_id" form:"sys_user_authority_id" gorm:"comment:管理角色ID"` // 管理角色ID
}

func (VIPStatement) TableName() string {
	return "bus_statement"
}

type VIPStatistics struct {
	TotalStream        float64 `json:"totalStream" form:"totalStream" gorm:"comment:总流水;"`
	TotalOrder         uint    `json:"totalOrder" form:"totalOrder" gorm:"comment:会员卡总单量"`
	TotalMember        uint    `json:"totalMember" form:"totalMember" gorm:"comment:总会员"`
	TotalConsumer      uint    `json:"totalConsumer" form:"totalConsumer" gorm:"comment:入店总人数"`
	SysUserAuthorityID uint    `json:"sys_user_authority_id" form:"sys_user_authority_id" gorm:"primary_key;comment:管理角色ID"`
}

func (VIPStatistics) TableName() string {
	return "bus_statistics"
}
