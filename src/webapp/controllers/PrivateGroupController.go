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

func (this *PrivateGroupController) Post() {

	this.SendRes(0, "success", nil)
}
