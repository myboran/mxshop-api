package initialize

import (
	"encoding/json"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"mxshop-api/user-web/global"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig() {
	debug := GetEnvInfo("MXSHOP_DEBUG")
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("user-web/%s-debug.yaml", configFilePrefix)
	if !debug {
		configFileName = fmt.Sprintf("user-web/%s-production.yaml", configFilePrefix)
	}

	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := v.Unmarshal(global.NacosConfig); err != nil {
		panic(err)
	}

	zap.S().Infof("配置信息： %v", global.NacosConfig)

	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		zap.S().Infof("config file changed: %v", e.Name)
		_ = v.ReadInConfig()
		_ = v.Unmarshal(global.NacosConfig)
		zap.S().Infof("配置信息： %v", global.NacosConfig)
	})

	sc := []constant.ServerConfig{
		{
			IpAddr: global.NacosConfig.Host,
			Port:   uint64(global.NacosConfig.Port),
		},
	}

	cc := constant.ClientConfig{
		NamespaceId:         global.NacosConfig.NameSpace, // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "debug",
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		panic(err)
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: global.NacosConfig.DataId,
		Group:  global.NacosConfig.Group})

	if err != nil {
		panic(err)
	}

	json.Unmarshal([]byte(content), &global.ServerConfig)

	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: global.NacosConfig.DataId,
		Group:  global.NacosConfig.Group,
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("配置文件变化")
			fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
			json.Unmarshal([]byte(content), &global.ServerConfig)
		},
	})
}
