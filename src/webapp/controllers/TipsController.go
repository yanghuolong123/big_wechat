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
	this.TplName = "tips/tips.tpl"
}

func (this *TipsController) PgIndex() {
	this.Layout = "layout/addwechat.tpl"
	this.TplName = "tips/pgindex.tpl"
}

func (this *TipsController) PgList() {
	this.Layout = "layout/addwechat.tpl"
	this.TplName = "tips/pglist.tpl"
}
