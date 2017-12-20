package controllers

import (
	"yhl/help"
)

type AddGroupController struct {
	help.BaseController
}

func (this *AddGroupController) Get() {
	this.Data["welcome"] = "welcome to add group, accessToken:"
	this.TplName = "addGroup/index.tpl"
}

func (this *AddGroupController) Post() {

	this.SendRes(0, "success", nil)
}
