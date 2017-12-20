package controllers

import (
	//	"time"
	//	"fmt"
	"yhl/help"
	"yhl/wechat"
)

type TestController struct {
	help.BaseController
}

func (this *TestController) Get() {
	/*cache := help.Cache
	token := cache.Get("access_token_")

	if token == nil {
		token = "ddsdsdsdssddsds"
		cache.Put("access_token_", token, 100*time.Second)
	} else {
		token = string(token.([]uint8))
	}

	accessToken := token.(string) //wechat.GetAccessToken()
	accessToken := wechat.GetAccessToken()
	m := map[string]interface{}{}
	m["touser"] = "oTbmFxG5r1WRrHdb32O5y2aSAIkc"
	m["msgtype"] = "text"
	m["text"] = map[string]string{"content": "Hello World!"}
	wechat.SendMsg(m)
	wechat.SendTextMsg("oTbmFxG5r1WRrHdb32O5y2aSAIkc", "I love you")
	userinfo := wechat.GetWxUserinfo("oTbmFxG5r1WRrHdb32O5y2aSAIkc", "")
	fmt.Println(userinfo)
	*/
	this.Data["welcome"] = "welcome to add group, accessToken:"
	qrImgUrl := wechat.GetTmpStrQrImg("login_123")
	this.Data["qrImgUrl"] = qrImgUrl
	this.TplName = "test/index.tpl"
}

func (this *TestController) Post() {

	this.SendRes(0, "success", nil)
}
