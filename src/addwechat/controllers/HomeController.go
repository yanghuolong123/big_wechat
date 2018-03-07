package controllers

import (
	//	"errors"
	"addwechat/models"
	"fmt"
	"yhl/help"
)

type HomeController struct {
	help.BaseController
}

func (this *HomeController) Get() {
	user, _ := models.GetUserById(1)
	fmt.Println(user)
	u := this.GetSession("user")
	if u == nil {
		//u = models.User{}
		u = new(models.User)
	}
	follow := this.GetSession("follow")
	if follow == nil {
		follow = []int{}
	}
	group := this.GetSession("group")
	if group == nil {
		group = new(models.Group)
	}

	this.Data["user"] = u
	this.Data["follow"] = follow
	this.Data["group"] = group

	this.TplName = "home/index.tpl"
}
