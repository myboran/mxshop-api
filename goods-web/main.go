package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/satori/go.uuid"
	"go.uber.org/zap"
	"mxshop-api/goods-web/global"
	"mxshop-api/goods-web/initialize"
	"mxshop-api/goods-web/utils/addr"
	"mxshop-api/goods-web/utils/register/consul"
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
		zap.S().Panic("注销失败:", err.Error())
	} else {
		zap.S().Panic("注销成功")
	}

}
