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
	follow := this.GetSession("follow")
	if follow == nil {
		follow = []int{}
	}
	group := this.GetSession("group")
	if group == nil {
		group = new(models.Group)
	}

	this.Data["welcome"] = "hello, welcome to bigwechart! thank you!"
	this.Data["user"] = u
	this.Data["follow"] = follow
	this.Data["group"] = group

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
	group := models.GetGroupById(u.Gid)

	this.SetSession("user", u)
	this.SetSession("follow", gids)
	this.SetSession("group", group)

	m := make(map[string]interface{})
	m["user"] = u
	m["follow"] = gids
	m["group"] = group

	this.SendRes(0, "success", m)
}

func (this *HomeController) Logout() {
	this.DelSession("user")
	this.DelSession("follow")
	this.DelSession("group")

	this.SendRes(0, "success", nil)
}

func (this *HomeController) Register() {
	var user models.User
	user.Gid = 2
	user.Username = "yhl27ml@126.com"
	user.Password = help.Md5("123456")
	user.Email = "yhl27ml@126.com"
	user.Nickname = "Jason"
	user.Mobile = "18210189803"
	user.Avatar = "/statis/upload/avatar/1.png"
	user.Level = 1
	user.Status = 1

	uid := models.CreateUser(&user)
	models.CreateFollow(uid, user.Gid)

	this.SendRes(0, "success", uid)
}
