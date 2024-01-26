package pay

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type PayRouter struct{}

func (e *PayRouter) InitPayRouter(Router *gin.RouterGroup, RouterPub *gin.RouterGroup) {
	payPublicRouterWithoutRecord := RouterPub.Group("pay")
	payRouterWithoutRecord := Router.Group("pay")
	payApi := v1.ApiGroupApp.PayApiGroup.PayApi
	{
		payRouterWithoutRecord.POST("generateOrder", payApi.GenerateOrder)
		payRouterWithoutRecord.GET("detail", payApi.GetOrderDetail)
	}
	{
		payPublicRouterWithoutRecord.GET("notify", payApi.OrderNotify)

	}
}
