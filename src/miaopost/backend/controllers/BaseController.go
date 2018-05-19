package controllers

import (
	"yhl/help"
)

type BaseController struct {
	help.BaseController
}

func (this *BaseController) Prepare() {
	this.BaseController.Prepare()

	user := this.GetSession("admin")
	this.Data["user"] = user
}
