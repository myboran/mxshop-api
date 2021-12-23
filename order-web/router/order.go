package router

import (
	"github.com/gin-gonic/gin"
	"mxshop-api/order-web/api/order"
	"mxshop-api/order-web/api/pay"
)

func InitOrderRouter(Router *gin.RouterGroup) {
	OrderRouter := Router.Group("orders")
	{
		OrderRouter.GET("", order.List)
		OrderRouter.POST("", order.New)
		OrderRouter.GET("/:id", order.Detail)
	}
	PayRouter := Router.Group("pay")
	{
		PayRouter.POST("alipay/notify", pay.Notify)
	}
}
