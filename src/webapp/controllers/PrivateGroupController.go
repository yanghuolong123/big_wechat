package controllers

import (
	"webapp/models"
	"yhl/help"
)

type PrivateGroupController struct {
	help.BaseController
}

func (this *PrivateGroupController) Get() {
	pgroups := models.GetPrivateGroupByLimit(16)
	user := this.GetSession("user")

	this.Data["pgroups"] = pgroups
	this.Data["user"] = user
	this.TplName = "privateGroup/index.tpl"
}

func (this *PrivateGroupController) CreateGet() {
	user := this.GetSession("user")
	this.Data["user"] = user
	this.Layout = "layout/addwechat.tpl"
	this.TplName = "privateGroup/create.tpl"
}

func (this *PrivateGroupController) CreatePost() {

	this.SendRes(0, "success", nil)
}
