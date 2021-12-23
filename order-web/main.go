package main

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	validators "mxshop-api/order-web/validator"
	"os"
	"os/signal"
	"syscall"

	"github.com/satori/go.uuid"
	"go.uber.org/zap"
	"mxshop-api/order-web/global"
	"mxshop-api/order-web/initialize"
	"mxshop-api/order-web/utils/addr"
	"mxshop-api/order-web/utils/register/consul"
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

	register_client := consul.NewRegistryClient(global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)
	serviceIdStr := fmt.Sprintf("%s", uuid.NewV4())
	err := register_client.Register(global.ServerConfig.Host, global.ServerConfig.Port, global.ServerConfig.Name, global.ServerConfig.Tags, serviceIdStr)
	if err != nil {
		zap.S().Panic("启动失败", err.Error())
	}

	zap.S().Debugf("启动服务器， 端口： %d", global.ServerConfig.Port)

	go func() {
		if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
			zap.S().Panic("启动失败", err.Error())
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	err = register_client.DeRegister(serviceIdStr)
	if err != nil {
		zap.S().Info("注销失败:", err.Error())
	} else {
		zap.S().Info("注销成功")
	}

}
