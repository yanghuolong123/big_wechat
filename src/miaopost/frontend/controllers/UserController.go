package controllers

import (
	"miaopost/frontend/models"
)

type UserController struct {
	BaseController
}

func (this *UserController) Prepare() {
	user := this.GetSession("user")
	if user == nil {
		this.Redirect("/", 302)
	}

	this.BaseController.Prepare()
}

func (this *UserController) Index() {

	this.Layout = "layout/main.tpl"
	this.TplName = "user/index.tpl"
}

func (this *UserController) EditGet() {

	this.Layout = "layout/main.tpl"
	this.TplName = "user/edit.tpl"
}

func (this *UserController) EditPost() {
	user := this.GetSession("user")
	nickname := this.GetString("nickname")

	u, _ := models.GetUserById(user.(*models.User).Id)
	u.Nickname = nickname

	err := models.UpdateUser(u)
	if err != nil {
		this.Redirect("/tips?msg="+err.Error(), 302)
	}

	this.SetSession("user", u)

	this.Redirect("/user", 302)

}
