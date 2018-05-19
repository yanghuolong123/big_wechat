package controllers

import (
	"miaopost/backend/models"
	//"yhl/help"
)

type SiteController struct {
	BaseController
}

func (this *SiteController) Get() {

	this.Layout = "layout/main.tpl"
	this.TplName = "site/index.tpl"
}

func (this *SiteController) LoginGet() {
	user := this.GetSession("admin")
	if user != nil {
		this.Redirect("/", 302)
	}

	this.TplName = "site/login.tpl"
}

func (this *SiteController) LoginPost() {
	username := this.GetString("username")
	password := this.GetString("password")
	admin, err := models.Login(username, password)
	if err != nil {
		this.SendRes(-1, err.Error(), nil)
	}

	this.SetSession("admin", admin)
	this.SendRes(0, "success", nil)

}

func (this *SiteController) Logout() {
	this.DelSession("admin")

	this.Redirect("/", 302)
}
