package business

import (
	"gorm.io/gorm"
	"time"
)

type VIPCombo struct {
	CreatedAt   time.Time      // 创建时间
	UpdatedAt   time.Time      // 更新时间
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`                                                  // 删除时间
	ComboId     int            `json:"comboId" gorm:"not null;unique;primary_key;comment:套餐ID;size:90"` // 套餐ID
	ComboName   string         `json:"comboName" form:"comboName" gorm:"comment:套餐名称"`                  // 套餐名称
	ComboType   int            `json:"comboType" form:"comboType" gorm:"comment:套餐类型"`                  // 套餐类型
	ComboPrice  int            `json:"comboPrice" form:"comboPrice" gorm:"comment:套餐价格"`                // 套餐价格
	Times       int            `json:"times" form:"times" gorm:"comment:天数/次数/金额"`                      // 天数/次数/金额
	State       int            `json:"state" form:"state" gorm:"comment:状态"`                            // 状态
	AuthorityId int            `json:"-" gorm:"comment:角色ID;column:authority_id"`
}

func (VIPCombo) TableName() string {
	return "bus_combo"
}
