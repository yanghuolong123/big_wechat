package controllers

import (
	"webapp/models"
	"yhl/help"
)

type UserController struct {
	help.BaseController
}

func (this *UserController) Index() {
	user := this.GetSession("user")
	if user == nil {
		this.Redirect("/", 302)
	}
	u := user.(models.User)

	pgs := models.GetPrivateGroupByUid(u.Id)
	ugs := models.GetUnlockGroupByUid(u.Id)

	this.Data["user"], _ = models.GetUserById(u.Id)
	this.Data["pgs"] = pgs
	this.Data["ugs"] = ugs

	this.Layout = "layout/addwechat.tpl"
	this.TplName = "user/index.tpl"
}

func (this *UserController) EditGet() {
	user := this.GetSession("user")
	if user == nil {
		this.Redirect("/", 302)
	}
	u := user.(models.User)

	this.Data["user"], _ = models.GetUserById(u.Id)

	this.Layout = "layout/addwechat.tpl"
	this.TplName = "user/edit.tpl"
}

func (this *UserController) EditPost() {
	user := this.GetSession("user")
	if user == nil {
		this.Redirect("/", 302)
	}

	nickname := this.GetString("nickname")

	u, _ := models.GetUserById(user.(models.User).Id)
	u.Nickname = nickname

	err := models.UpdateUser(u)
	if err != nil {
		this.Redirect("/tips?msg="+err.Error(), 302)
	}

	this.SetSession("user", *u)

	this.Redirect("/user", 302)

}
