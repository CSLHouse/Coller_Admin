package pay

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type PayRouter struct{}

func (e *PayRouter) InitPayRouter(Router *gin.RouterGroup) {
	payRouterWithoutRecord := Router.Group("pay")
	payApi := v1.ApiGroupApp.PayApiGroup.PayApi
	{
		payRouterWithoutRecord.POST("generateOrder", payApi.GenerateOrder)

	}
}
