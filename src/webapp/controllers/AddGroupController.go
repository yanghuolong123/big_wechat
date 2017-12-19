package controllers

import (
	//	"time"
	"yhl/help"
	"yhl/wechat"
)

type AddGroupController struct {
	help.BaseController
}

func (this *AddGroupController) Get() {
	/*cache := help.Cache
	token := cache.Get("access_token_")

	if token == nil {
		token = "ddsdsdsdssddsds"
		cache.Put("access_token_", token, 100*time.Second)
	} else {
		token = string(token.([]uint8))
	}

	accessToken := token.(string) //wechat.GetAccessToken()
	*/
	accessToken := wechat.GetAccessToken()
	m := map[string]interface{}{}
	m["touser"] = "oTbmFxG5r1WRrHdb32O5y2aSAIkc"
	m["msgtype"] = "text"
	m["text"] = map[string]string{"content": "Hello World!"}
	wechat.SendMsg(m)
	wechat.SendTextMsg("oTbmFxG5r1WRrHdb32O5y2aSAIkc", "I love you")
	this.Data["welcome"] = "welcome to add group, accessToken:" + accessToken
	this.TplName = "addGroup/index.tpl"
}

func (this *AddGroupController) Post() {

	this.SendRes(0, "success", nil)
}
