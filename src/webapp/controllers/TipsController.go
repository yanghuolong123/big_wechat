package controllers

import (
	"yhl/help"
)

type TipsController struct {
	help.BaseController
}

func (this *TipsController) Tips() {
	user := this.GetSession("user")
	this.Data["user"] = user

	msg := this.GetString("msg")

	this.Data["msg"] = msg
	this.TplName = "tips/tips.tpl"
}

func (this *TipsController) PgIndex() {
	user := this.GetSession("user")
	this.Data["user"] = user

	this.Layout = "layout/addwechat.tpl"
	this.TplName = "tips/pgindex.tpl"
}

func (this *TipsController) PgList() {
	user := this.GetSession("user")
	this.Data["user"] = user

	this.Layout = "layout/addwechat.tpl"
	this.TplName = "tips/pglist.tpl"
}
