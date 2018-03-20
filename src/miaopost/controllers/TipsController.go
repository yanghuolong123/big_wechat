package controllers

import (
	"yhl/help"
)

type TipsController struct {
	help.BaseController
}

func (this *TipsController) Tips() {
	msg := this.GetString("msg")

	this.Data["msg"] = msg
	this.Layout = "layout/main.tpl"
	this.TplName = "tips/tips.tpl"
}
