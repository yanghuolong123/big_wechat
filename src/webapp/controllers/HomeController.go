package controllers

import (
	"fmt"
	"webapp/models"
	"yhl/help"
)

type HomeController struct {
	help.BaseController
}

func (this *HomeController) Get() {
	//u, _ := models.GetById(1)
	u := this.GetSession("user")
	if u == nil {
		//u = models.User{}
		u = new(models.User)
	}

	help.Log.Info("===============================")
	//help.Log.Info(u.Username)
	fmt.Println(u)

	this.Data["welcome"] = "hello, welcome to bigwechart! thank you!"
	this.Data["user"] = u
	this.TplName = "home/index.tpl"
}

func (this *HomeController) Login() {
	email := this.GetString("email")
	passwd := this.GetString("password")
	u, err := models.Login(email, passwd)
	if err != nil {
		this.SendRes(-1, err.Error(), nil)
	}
	this.SetSession("user", u)

	this.SendRes(0, "success", u)
}

func (this *HomeController) Logout() {
	this.DelSession("user")
	this.SendRes(0, "success", nil)
}

func (this *HomeController) Register() {
	this.SendRes(0, "success", nil)
}
