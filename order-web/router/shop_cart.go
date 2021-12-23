package router

import (
	"github.com/gin-gonic/gin"
	"mxshop-api/order-web/api/shop_cart"
)

func InitShopCartRouter(Router *gin.RouterGroup) {
	ShopCartRouter := Router.Group("shopcarts")
	{
		ShopCartRouter.GET("", shop_cart.List)
		ShopCartRouter.DELETE("/:id", shop_cart.Delete)
		ShopCartRouter.POST("", shop_cart.New)
		ShopCartRouter.PATCH("/:id", shop_cart.Update)
	}
}
