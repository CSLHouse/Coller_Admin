package response

type VipConsumeResModel struct {
	ID           uint
	Telephone    uint   `json:"telephone" form:"telephone" gorm:"comment:会员手机号"`  // 客户手机号
	MemberName   string `json:"memberName" form:"memberName" gorm:"comment:会员名"`  // 客户名
	MemberType   string `json:"memberType" form:"memberType" gorm:"comment:会员类型"` // 管理ID
	MemberState  uint   `json:"memberState" form:"memberState" gorm:"comment:会员状态"`
	RemainTimes  uint   `json:"remainTimes" form:"remainTimes" gorm:"comment:剩余次数/金额"` // 管理ID
	ConsumeTimes uint   `json:"consumeTimes" form:"consumeTimes" gorm:"comment:消费次数/金额"`
	PunchDate    string `json:"punchDate" form:"punchDate" gorm:"comment:登记日期"`
	Deadline     string `json:"deadline" form:"deadline" gorm:"comment:截止日期"` // 管理ID
	State        uint   `json:"state" form:"state" gorm:"comment:状态"`
}
type VipConsumeResponse struct {
	Combo VipConsumeResModel `json:"combo"`
}
