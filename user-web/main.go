package main

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"mxshop-api/user-web/global"
	"mxshop-api/user-web/utils/addr"

	"mxshop-api/user-web/initialize"
	validators "mxshop-api/user-web/validator"
)

func main() {
	// 1. 初始化 logger
	initialize.InitLogger()

	initialize.InitConfig()
	initialize.InitSrvConn()
	// 2. 初始化routers
	Router := initialize.Routers()

	// 注册验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("mobile", validators.ValidateMobile)
	}
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
