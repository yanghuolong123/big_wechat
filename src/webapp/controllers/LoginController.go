package controllers

import (
	"fmt"
	"strconv"
	"time"
	//	"webapp/models"
	"yhl/help"
	"yhl/wechat"
)

type LoginController struct {
	help.BaseController
}

func (this *LoginController) LoginGet() {
	sceneId := "login_" + time.Now().Format(help.DatetimeNumFormat) + strconv.Itoa(help.RandNum(10000, 99999))
	qrImgUrl := wechat.GetTmpStrQrImg(sceneId)

	this.Data["qrImgUrl"] = qrImgUrl
	this.Data["sceneId"] = sceneId
	this.TplName = "login/index.tpl"
	s, _ := this.RenderString()

	this.SendRes(0, "success", s)
}

func (this *LoginController) LoginPost() {
	sceneId := this.GetString("sceneId")
	cache := help.Cache
	c := cache.Get(sceneId)
	fmt.Println(c)
	if c != nil {
		u := c.([]uint8)
		this.SendRes(0, "success", u)
	}

	this.SendRes(0, "failed", nil)
}

func (this *LoginController) LogOut() {
	this.SendRes(0, "success", nil)
}
