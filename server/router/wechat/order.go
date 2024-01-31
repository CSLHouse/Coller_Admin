package wechat

import (
	v1 "cooller/server/api/v1"
	"cooller/server/middleware"
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
		wxOrderRouter.POST("cancelOrder", wxOrderApi.CancelOrders)
		wxOrderRouter.GET("setting", wxOrderApi.GetOrderSetting)
		wxOrderRouter.POST("settingUpdate", wxOrderApi.UpdateOrderSetting)
		wxOrderRouter.POST("closeOrders", wxOrderApi.CloseOrders)
		wxOrderRouter.DELETE("delete", wxOrderApi.DeleteOrders)
		wxOrderRouter.POST("update/receiverInfo", wxOrderApi.UpdateOrderReceiverInfo)
		wxOrderRouter.POST("update/moneyInfo", wxOrderApi.UpdateOrderMoneyInfo)
		wxOrderRouter.POST("update/note", wxOrderApi.UpdateOrderNote)
		wxOrderRouter.POST("update/complete", wxOrderApi.UpdateOrderCompletedStatus)
	}
	return wxOrderRouter
}
