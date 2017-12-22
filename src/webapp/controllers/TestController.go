package controllers

import (
	//	"time"
	"errors"
	"fmt"
	"webapp/models"
	"yhl/help"
	"yhl/wechat"
)

type TestController struct {
	help.BaseController
}

func (this *TestController) Get() {
	/*cache := help.Cache
	token := cache.Get("access_token_")

	if token == nil {
		token = "ddsdsdsdssddsds"
		cache.Put("access_token_", token, 100*time.Second)
	} else {
		token = string(token.([]uint8))
	}

	accessToken := token.(string) //wechat.GetAccessToken()
	accessToken := wechat.GetAccessToken()
	m := map[string]interface{}{}
	m["touser"] = "oTbmFxG5r1WRrHdb32O5y2aSAIkc"
	m["msgtype"] = "text"
	m["text"] = map[string]string{"content": "Hello World!"}
	wechat.SendMsg(m)
	wechat.SendTextMsg("oou4Vw0zizge_p2gQhYT0UL5Kwbk", "I love you")
	userinfo := wechat.GetWxUserinfo("oTbmFxG5r1WRrHdb32O5y2aSAIkc", "")
	fmt.Println(userinfo)
	*/
	pg := models.PrivateGroup{}
	pg.Id = 1
	//pg.Gid = 2
	//pg.Uid = 1
	//pg.Name = "aaa"
	pg.Introduction = "aaaaaaaaaaasss"
	pg.Qrcode = "aaaaaaa"
	pg.Ower_qrcode = "bbbb"
	//flag := models.CreatePrivateGroup(&pg)
	flag := models.UpdatePrivateGroup(&pg)
	fmt.Println("================= flag:", flag)
	this.Data["welcome"] = "welcome to add group, accessToken:"
	qrImgUrl := wechat.GetTmpStrQrImg("login_123")
	this.Data["qrImgUrl"] = qrImgUrl
	this.TplName = "test/index.tpl"
}

func (this *TestController) Post() {

	this.SendRes(0, "success", nil)
}

func (this *TestController) LoginPage() {
	this.TplName = "test/login.tpl"
	s, _ := this.RenderString()

	this.SendRes(0, "success", s)
}

func (this *TestController) login(username, password string) (m map[string]interface{}, err error) {
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

func (this *TestController) Login() {
	username := this.GetString("username")
	password := this.GetString("password")

	m, err := this.login(username, password)
	if err != nil {
		this.SendRes(-1, err.Error(), nil)
	}

	this.SendRes(0, "success", m)
}

func (this *TestController) Logout() {
	this.DelSession("user")
	this.DelSession("follow")
	this.DelSession("group")

	this.SendRes(0, "success", nil)
}

func (this *TestController) RegisterPage() {
	this.Data["groupList"] = models.GetGroupAll()
	this.TplName = "test/register.tpl"
	s, _ := this.RenderString()

	this.SendRes(0, "success", s)
}

func (this *TestController) Register() {
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
