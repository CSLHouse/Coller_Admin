package pay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service"
)

type ApiGroup struct {
	PayApi
}

var (
	jspaymentService = service.ServiceGroupApp.PaymentServiceGroup.PayMentService
	accountService   = service.ServiceGroupApp.WechatServiceGroup.AccountService
	orderService     = service.ServiceGroupApp.WechatServiceGroup.OrderService
	wechatService    = service.ServiceGroupApp.WechatServiceGroup.HomeService
)
