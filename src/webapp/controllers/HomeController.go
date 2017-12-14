package controllers

import (
	"errors"
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

	this.Data["user"] = u
	this.Data["follow"] = follow
	this.Data["group"] = group

	this.TplName = "home/index.tpl"
}

func (this *HomeController) LoginPage() {
	this.TplName = "home/login.tpl"
	s, _ := this.RenderString()

	this.SendRes(0, "success", s)
}

func (this *HomeController) login(username, password string) (m map[string]interface{}, err error) {
	m = make(map[string]interface{})
	u, err := models.Login(username, password)
	if err != nil {
		return
	}
	gids := models.GetFollowByUid(u.Id)
	group := models.GetGroupById(u.Gid)

	this.SetSession("user", u)
	this.SetSession("follow", gids)
	this.SetSession("group", group)

	m["user"] = u
	m["follow"] = gids
	m["group"] = group

	return
}

func (this *HomeController) Login() {
	username := this.GetString("username")
	password := this.GetString("password")

	m, err := this.login(username, password)
	if err != nil {
		this.SendRes(-1, err.Error(), nil)
	}

	this.SendRes(0, "success", m)
}

func (this *HomeController) Logout() {
	this.DelSession("user")
	this.DelSession("follow")
	this.DelSession("group")

	this.SendRes(0, "success", nil)
}

func (this *HomeController) RegisterPage() {
	this.Data["groupList"] = models.GetGroupAll()
	this.TplName = "home/register.tpl"
	s, _ := this.RenderString()

	this.SendRes(0, "success", s)
}

func (this *HomeController) Register() {
	gid, _ := this.GetInt("group")
	username := this.GetString("username")
	nickname := this.GetString("nickname")
	password := this.GetString("password")
	repassword := this.GetString("repassword")
	if password != repassword {
		this.SendRes(-1, "密码输入不一致", nil)
	}
	_, err := models.GetUserByUsername(username)
	if err != nil {
		this.SendRes(-1, errors.New("账号已存在").Error(), nil)
	}
	loginPasswd := password
	password = help.Md5(password)

	u := models.User{
		Gid:      gid,
		Username: username,
		Nickname: nickname,
		Password: password,
	}

	uid := models.CreateUser(&u)
	if uid <= 0 {
		this.SendRes(-1, "注册失败", nil)
	}
	models.CreateFollow(uid, gid)

	m, _ := this.login(username, loginPasswd)

	this.SendRes(0, "success", m)
}
