package controllers

import (
	"yhl/help"
)

type LoginController struct {
	help.BaseController
}

func (this *LoginController) LoginGet() {
	this.Data["msg"] = "login ...."

	this.TplName = "login/index"
}

func (this *LoginController) LoginPost() {
	this.SendRes(0, "success", nil)
}

func (this *LoginController) LogOut() {
	this.SendRes(0, "success", nil)
}
