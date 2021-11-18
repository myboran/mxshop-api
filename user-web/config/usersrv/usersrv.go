package usersrv

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"mxshop-api/user-web/global"
)

func GetUserSrvPort() *grpc.ClientConn {
	consulinfo := global.ServerConfig.ConsulInfo
	usersrvinfo := global.ServerConfig.UserSrvInfo
	// 从注册中心获取到用户服务的信息
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", consulinfo.Host, consulinfo.Port)
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	data, err := client.Agent().ServicesWithFilter(fmt.Sprintf("Service== \"%s\"", usersrvinfo.Name))
	if err != nil {
		panic(err)
	}
	userSrvHost := ""
	userSrvPort := 0
	for _, value := range data {
		println(value)
		userSrvHost = value.Address
		userSrvPort = value.Port
	}
	if userSrvHost == "" {
		zap.S().Errorw("[GetUserList] 连接 【用户服务失败】", "msg", err.Error())
	}
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", userSrvHost, userSrvPort), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[GetUserList]链接[用户服务失败]", "msg", err.Error())
	}
	return userConn
}
