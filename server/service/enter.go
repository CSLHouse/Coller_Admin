package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/business"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/wechat"
)

type ServiceGroup struct {
	SystemServiceGroup   system.ServiceGroup
	ExampleServiceGroup  example.ServiceGroup
	BusinessServiceGroup business.ServiceGroup
	WechatServiceGroup   wechat.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
