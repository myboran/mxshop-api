package global

import (
	"mxshop-api/goods-web/config"
	"mxshop-api/goods-web/proto"
)

var (
	ServerConfig *config.ServerConfig = &config.ServerConfig{}

	NacosConfig *config.NacosConfig = &config.NacosConfig{}

	GoodsSrvClient proto.GoodsClient
)
