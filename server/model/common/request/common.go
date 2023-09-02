package request

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page     int    `json:"page" form:"page"`         // 页码
	PageSize int    `json:"pageSize" form:"pageSize"` // 每页大小
	Keyword  string `json:"keyword" form:"keyword"`   //关键字
}

// GetById Find by id structure
type GetById struct {
	ID int `json:"id" form:"id"` // 主键ID
}

func (r *GetById) Uint() uint {
	return uint(r.ID)
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}

// GetAuthorityId Get role by id structure
type GetAuthorityId struct {
	AuthorityId uint `json:"authorityId" form:"authorityId"` // 角色ID
}

type Empty struct{}

type MemberSearchInfo struct {
	Telephone  uint   `json:"telephone" form:"telephone"`   // 联系电话
	MemberName string `json:"memberName" form:"memberName"` // 姓名
	Deadline   string `json:"deadline" form:"deadline"`     // 截止时间
	State      uint   `json:"state" form:"state"`           // 状态
	Page       int    `json:"page" form:"page"`             // 页码
	PageSize   int    `json:"pageSize" form:"pageSize"`     // 每页大小
}

type CardInfo struct {
	OnlyId uint `json:"onlyId" form:"onlyId"` // 唯一标识
}

type ConsumeSearchInfo struct {
	Telephone uint `json:"telephone" form:"telephone"` // 联系电话
	State     uint `json:"state" form:"state"`         // 状态
	Page      int  `json:"page" form:"page"`           // 页码
	PageSize  int  `json:"pageSize" form:"pageSize"`   // 每页大小
}

type OrderSearchInfo struct {
	OrderId   string `json:"orderId" form:"orderId" gorm:"comment:订单编号"`
	Telephone uint   `json:"telephone" form:"telephone"` // 联系电话
	Type      uint   `json:"type" form:"type" gorm:"comment:订单类型"`
	BuyDate   string `json:"buyDate" form:"buyDate" gorm:"comment:购买日期"`
	Page      int    `json:"page" form:"page"`         // 页码
	PageSize  int    `json:"pageSize" form:"pageSize"` // 每页大小
}

type StatisticsSearchInfo struct {
	StartDate string `json:"startDate" form:"startDate" gorm:"comment:开始日期"`
	EndDate   string `json:"endDate" form:"endDate" gorm:"comment:结束日期"`
}
