package main

import (
	"fmt"
	"go.uber.org/zap"
	"mxshop-api/goods-web/global"
	"mxshop-api/goods-web/initialize"
	"mxshop-api/goods-web/utils/addr"
)

func main() {
	// 1. 初始化 logger
	initialize.InitLogger()

	initialize.InitConfig()
	initialize.InitSrvConn()
	// 2. 初始化routers
	Router := initialize.Routers()

	debug := initialize.GetEnvInfo("MXSHOP_DEBUG")
	if !debug {
		port, err := addr.GetFreePort()
		if err == nil {
			global.ServerConfig.Port = port

		}
	}

	zap.S().Debugf("启动服务器， 端口： %d", global.ServerConfig.Port)

	if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
		zap.S().Panic("启动失败", err.Error())
	}
}
