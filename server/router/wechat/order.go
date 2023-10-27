package wechat

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type OrderRouter struct{}

func (s *OrderRouter) InitOrderRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	wxOrderRouter := Router.Group("order").Use(middleware.OperationRecord())
	wxOrderApi := v1.ApiGroupApp.WechatApiGroup.OrderApi
	{
		wxOrderRouter.POST("generateConfirmOrder", wxOrderApi.GenerateConfirmOrder)
		wxOrderRouter.POST("generateOrder", wxOrderApi.GenerateOrder)
		wxOrderRouter.GET("detail", wxOrderApi.GetOrderDetail)
		wxOrderRouter.GET("list", wxOrderApi.GetOrderList)
		wxOrderRouter.POST("paySuccess", wxOrderApi.PaySuccess)
	}
	return wxOrderRouter
}
