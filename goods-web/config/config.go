package config

type GoodsSrvConfig struct {
	Name string `json:"name"`
	Msg  string `json:"msg"`
}

type JWTConfig struct {
	SigningKey string `json:"key"`
}

type ConsulConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type ServerConfig struct {
	Name         string         `json:"name"`
	Host         string         `json:"host"`
	Port         int            `json:"port"`
	GoodsSrvInfo GoodsSrvConfig `json:"goodsSrv"`
	JWTInfo      JWTConfig      `json:"jwt"`
	ConsulInfo   ConsulConfig   `json:"consul"`
}

type NacosConfig struct {
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	NameSpace string `mapstructure:"namespace"`
	User      string `mapstructure:"user"`
	PassWord  string `mapstructure:"password"`
	DataId    string `mapstructure:"dataid"`
	Group     string `mapstructure:"group"`
}
