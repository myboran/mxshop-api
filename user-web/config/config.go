package config

type UserSrvConfig struct {
	Name string `json:"name"`
	Msg  string `json:"msg"`
}

type JWTConfig struct {
	SigningKey string `json:"key"`
}

type RedisConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type ConsulConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type ServerConfig struct {
	Name        string        `json:"name"`
	Host        string        `json:"host"`
	Port        int           `json:"port"`
	Tags        []string      `json:"tags"`
	UserSrvInfo UserSrvConfig `json:"userSrv"`
	JWTInfo     JWTConfig     `json:"jwt"`
	RedisInfo   RedisConfig   `json:"redis"`
	ConsulInfo  ConsulConfig  `json:"consul"`
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
