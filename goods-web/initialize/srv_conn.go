package initialize

import (
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"google.golang.org/grpc"
	"log"
	"mxshop-api/goods-web/global"
	"mxshop-api/goods-web/proto"
)

func InitSrvConn() {
	conn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s&tag=python", global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port, global.ServerConfig.GoodsSrvInfo.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		log.Fatal(err)
	}
	//defer conn.Close()
	global.GoodsSrvClient = proto.NewGoodsClient(conn)
}
