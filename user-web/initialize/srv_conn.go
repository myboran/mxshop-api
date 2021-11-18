package initialize

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"mxshop-api/user-web/global"
	"mxshop-api/user-web/proto"
)

func InitSrvConn2() {
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

	// 问题
	// 1. 后续的用户服务下线、 改端口了、 改ip了、 （负载均衡来做）
	// 3. 一个连接多个groutine共用， 性能 --- 连接池
	userSrvClient := proto.NewUserClient(userConn)
	global.UserSrvClient = userSrvClient
}

func InitSrvConn() {
	conn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s&tag=python", global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port, global.ServerConfig.UserSrvInfo.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		log.Fatal(err)
	}
	//defer conn.Close()
	userSrvClient := proto.NewUserClient(conn)
	global.UserSrvClient = userSrvClient
}
