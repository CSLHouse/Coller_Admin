package business

import (
	v1 "cooller/server/api/v1"
	"github.com/gin-gonic/gin"
)

type OrderRouter struct{}

func (e *OrderRouter) InitOrderRouter(Router *gin.RouterGroup) {
	businessRouterWithoutRecord := Router.Group("business")
	vipOrderApi := v1.ApiGroupApp.BusinessApiGroup.OrderApi

	{
		businessRouterWithoutRecord.GET("orderList", vipOrderApi.GetVIPOrderList)       // 获取客户列表
		businessRouterWithoutRecord.GET("statement", vipOrderApi.GetVIPStatementList)   // 获取流水
		businessRouterWithoutRecord.GET("statistics", vipOrderApi.GetVIPStatisticsList) // 获取统计
	}
}
