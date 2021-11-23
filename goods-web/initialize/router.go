package initialize

import (
	"github.com/gin-gonic/gin"
	"mxshop-api/goods-web/middlewares"
	"mxshop-api/goods-web/router"
	"net/http"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	Router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})
	// 配置跨越
	Router.Use(middlewares.Cors())
	ApiGroup := Router.Group("/g/v1")
	router.InitGoodsRouter(ApiGroup)

	return Router
}
