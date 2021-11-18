package api

import (
	"fmt"
	"github.com/cloopen/go-sms-sdk/cloopen"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"log"
	"math/rand"
	"mxshop-api/user-web/forms"
	"mxshop-api/user-web/global"
	"net/http"
	_ "net/http"
	"strings"
	"time"
)

func GenerateSmsCode(width int) string {

	// 生成width长度的短信验证码
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

func SendSms(c *gin.Context) {
	sendsmsform := forms.SendSmsForm{}
	if err := c.ShouldBind(&sendsmsform); err != nil {
		HandleValidatorError(c, err)
		return
	}

	mobile := sendsmsform.Mobile
	smsCode := GenerateSmsCode(6)
	cfg := cloopen.DefaultConfig().
		// 开发者主账号,登陆云通讯网站后,可在控制台首页看到开发者主账号ACCOUNT SID和主账号令牌AUTH TOKEN
		WithAPIAccount("8aaf070874af41ee0175077bfec51cda").
		// 主账号令牌 TOKEN,登陆云通讯网站后,可在控制台首页看到开发者主账号ACCOUNT SID和主账号令牌AUTH TOKEN
		WithAPIToken("7d11ad5fe5de4f50a5a318e6103ef76d")
	sms := cloopen.NewJsonClient(cfg).SMS()
	// 下发包体参数
	input := &cloopen.SendRequest{
		// 应用的APPID
		AppId: "8aaf0708751c249f01752c0ae59f072a",
		// 手机号码
		To: mobile,
		// 模版ID
		TemplateId: "1",
		// 模版变量内容 非必填
		Datas: []string{smsCode, "7"},
	}
	// 下发
	println(111)
	resp, err := sms.Send(input)
	println(222)

	if err != nil {
		println(333)
		log.Printf(err.Error())
		println(444)
		return
	}
	if resp.StatusCode != "000000" {
		c.JSON(http.StatusOK, gin.H{
			"msg": "发送失败",
		})
		return
	}
	log.Printf("Response MsgId: %s \n", resp.TemplateSMS.SmsMessageSid)

	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", global.ServerConfig.RedisInfo.Host, global.ServerConfig.RedisInfo.Port),
	})
	rdb.Set(mobile, smsCode, 600*time.Second)
	c.JSON(http.StatusOK, gin.H{
		"msg": "发送成功",
	})
}
