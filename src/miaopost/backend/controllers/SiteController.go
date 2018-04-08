package controllers

import (
	"fmt"
	"miaopost/backend/models"
	"yhl/help"
)

type SiteController struct {
	help.BaseController
}

func (this *SiteController) Get() {
	user := this.GetSession("user")
	fmt.Println("========= user:", user)

	this.Data["user"] = user

	this.Layout = "layout/main.tpl"
	this.TplName = "site/index.tpl"
}

func (this *SiteController) LoginGet() {
	user := this.GetSession("user")
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

	this.SetSession("user", admin)
	this.SendRes(0, "success", nil)

}

func (this *SiteController) Logout() {
	this.DelSession("user")

	this.Redirect("/", 302)
}
