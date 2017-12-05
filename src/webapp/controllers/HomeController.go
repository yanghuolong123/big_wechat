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
	user, _ := models.GetUserById(1)
	fmt.Println(user)
	u := this.GetSession("user")
	if u == nil {
		//u = models.User{}
		u = new(models.User)
	}
	gids := this.GetSession("gids")
	if gids == nil {
		gids = []int{}
	}

	help.Log.Info("===============================")
	//help.Log.Info(u.Username)
	fmt.Println(u)

	this.Data["welcome"] = "hello, welcome to bigwechart! thank you!"
	this.Data["user"] = u
	this.Data["gids"] = gids
	this.TplName = "home/index.tpl"
}

func (this *HomeController) Login() {
	email := this.GetString("email")
	passwd := this.GetString("password")
	u, err := models.Login(email, passwd)
	if err != nil {
		this.SendRes(-1, err.Error(), nil)
	}
	gids := models.GetFollowByUid(u.Id)

	this.SetSession("user", u)
	this.SetSession("gids", gids)

	m := make(map[string]interface{})
	m["user"] = u
	m["gids"] = gids

	this.SendRes(0, "success", m)
}

func (this *HomeController) Logout() {
	this.DelSession("user")
	this.DelSession("gids")

	this.SendRes(0, "success", nil)
}

func (this *HomeController) Register() {
	var user models.User
	user.Group_id = 2
	user.Username = "yhl27ml@163.com"
	user.Password = help.Md5("123456")
	user.Email = "yhl27ml@163.com"
	user.Nickname = "Jason"
	user.Mobile = "18210189803"
	user.Avatar = "/statis/upload/avatar/1.png"
	user.Level = 1
	user.Status = 1

	models.CreateUser(&user)

	this.SendRes(0, "success", nil)
}
