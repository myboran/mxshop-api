package router

import (
	"github.com/gin-gonic/gin"

	"mxshop-api/goods-web/api/goods"
)

func InitGoodsRouter(Router *gin.RouterGroup) {
	GoodsRouter := Router.Group("goods")
	{
		GoodsRouter.GET("", goods.List)
	}
}
