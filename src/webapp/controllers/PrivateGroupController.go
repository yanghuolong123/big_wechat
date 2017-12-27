package controllers

import (
	"webapp/models"
	"yhl/help"
)

type PrivateGroupController struct {
	help.BaseController
}

func (this *PrivateGroupController) Get() {
	pgroups := models.GetPrivateGroupByLimit(16)
	user := this.GetSession("user")

	this.Data["pgroups"] = pgroups
	this.Data["user"] = user
	this.TplName = "privateGroup/index.tpl"
}

func (this *PrivateGroupController) CreateGet() {
	user := this.GetSession("user")
	this.Data["user"] = user
	this.Layout = "layout/addwechat.tpl"
	this.TplName = "privateGroup/create.tpl"
}

func (this *PrivateGroupController) CreatePost() {
	user := this.GetSession("user")
	if user == nil {
		this.SendRes(-1, "请先登录", nil)
	}
	gid, _ := this.GetInt("gid")
	name := this.GetString("name")
	introduction := this.GetString("introduction")
	qrcode := this.GetString("qrcode")
	ower_qrcode := this.GetString("ower_qrcode")
	wechat_id := this.GetString("wechat_id")

	if name == "" || gid <= 0 || (qrcode == "" && ower_qrcode == "" && wechat_id == "") {
		this.SendRes(-1, "参数错误", nil)
	}

	pg := models.PrivateGroup{}
	pg.Gid = int(gid)
	pg.Uid = user.(models.User).Id
	pg.Name = name
	pg.Introduction = introduction
	pg.Qrcode = qrcode
	pg.Ower_qrcode = ower_qrcode
	pg.Wechat_id = wechat_id

	models.CreatePrivateGroup(&pg)

	this.SendRes(0, "success", nil)
}

func (this *PrivateGroupController) User() {

	this.Layout = "layout/addwechat.tpl"
	this.TplName = "privateGroup/user.tpl"
}

func (this *PrivateGroupController) View() {
	id, _ := this.GetInt("id")
	pg := models.GetPrivateGroupById(int(id))
	group := models.GetGroupById(pg.Gid)

	this.Data["pg"] = pg
	this.Data["group"] = group

	this.Layout = "layout/addwechat.tpl"
	this.TplName = "privateGroup/view.tpl"
}
