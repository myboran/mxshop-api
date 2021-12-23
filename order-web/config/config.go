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

type AlipayConfig struct {
	AppID        string `json:"app_id"`
	PrivateKey   string `json:"private_key"`
	AliPublicKey string `json:"ali_public_key"`
	NotifyURL    string `json:"notify_url"`
	ReturnURL    string `json:"return_url"`
}

type ServerConfig struct {
	Name         string         `json:"name"`
	Host         string         `json:"host"`
	Port         int            `json:"port"`
	Tags         []string       `json:"tags"`
	GoodsSrvInfo GoodsSrvConfig `json:"goods_Srv"`

	OrderSrvInfo GoodsSrvConfig `json:"order_srv"`
	JWTInfo      JWTConfig      `json:"jwt"`
	ConsulInfo   ConsulConfig   `json:"consul"`
	AliPayInfo   AlipayConfig   `json:"alipay"`
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
