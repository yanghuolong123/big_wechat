package controllers

import (
	"fmt"
	"webapp/models"
	"yhl/help"
)

type PrivateGroupController struct {
	help.BaseController
}

func (this *PrivateGroupController) Get() {
	pgroups := models.GetPrivateGroupByLimit(16)
	fmt.Println(pgroups)

	this.Data["pgroups"] = pgroups
	this.Data["welcome"] = "welcome to add group, accessToken:"
	this.TplName = "privateGroup/index.tpl"
}

func (this *PrivateGroupController) Post() {

	this.SendRes(0, "success", nil)
}
