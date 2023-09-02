package business

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ConsumeRouter struct{}

func (e *ConsumeRouter) InitConsumeRouter(Router *gin.RouterGroup) {
	businessRouter := Router.Group("business").Use(middleware.OperationRecord())
	businessRouterWithoutRecord := Router.Group("business")
	vipConsumeApi := v1.ApiGroupApp.BusinessApiGroup.ConsumeApi
	{
		businessRouter.POST("consume", vipConsumeApi.ConsumeVIPCard) // VIP客户消费
	}
	{
		businessRouterWithoutRecord.GET("consumeList", vipConsumeApi.GetConsumeList) // 获取客户列表
	}
}
