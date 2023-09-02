package business

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type VIPMember struct {
	global.GVA_MODEL
	CardID     uint   `json:"cardId" form:"cardId" gorm:"unique;comment:会员卡号"`        // 客户名
	Telephone  uint   `json:"telephone" form:"telephone" gorm:"unique;comment:会员手机号"` // 客户手机号
	MemberName string `json:"memberName" form:"memberName" gorm:"comment:会员名"`        // 客户名
	ComboId    uint   `json:"comboId" form:"comboId" gorm:"comment:套餐ID"`             // 管理ID
	//ComboType          uint     `json:"comboType" form:"comboType" gorm:"comment:套餐类型"`         // 管理ID
	RemainTimes        uint     `json:"remainTimes" form:"remainTimes" gorm:"comment:剩余次数/金额"` // 管理ID
	StartDate          string   `json:"startDate" form:"startDate" gorm:"comment:开始日期"`
	Deadline           string   `json:"deadline" form:"deadline" gorm:"comment:截止日期"` // 管理ID
	State              uint     `json:"state" form:"state" gorm:"comment:状态"`         // 管理角色ID
	IsNew              bool     `json:"isNew" form:"isNew" gorm:"comment:新会员"`
	Collection         uint     `json:"collection" form:"collection" gorm:"comment:实付金额"`                         // 管理角色ID
	SysUserAuthorityID uint     `json:"sys_user_authority_id" form:"sys_user_authority_id" gorm:"comment:管理角色ID"` // 管理角色ID
	Combo              VIPCombo `json:"combo" gorm:"foreignKey:ComboId;references:ComboId;comment:套餐"`
}

func (VIPMember) TableName() string {
	return "bus_member"
}
