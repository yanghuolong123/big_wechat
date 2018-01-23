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

	err := models.CreatePrivateGroup(&pg)
	if err != nil {
		this.SendRes(-1, err.Error(), nil)
	}

	this.SendRes(0, "success", pg)
}

func (this *PrivateGroupController) EditGet() {
	user := this.GetSession("user")
	if user == nil {
		this.Redirect("/", 302)
	}

	id, _ := this.GetInt("id")
	pg := models.GetPrivateGroupById(int(id))
	group := models.GetGroupById(pg.Gid)

	this.Data["user"] = user
	this.Data["pg"] = pg
	this.Data["group"] = group

	this.Layout = "layout/addwechat.tpl"
	this.TplName = "privateGroup/edit.tpl"
}

func (this *PrivateGroupController) EditPost() {
	user := this.GetSession("user")
	if user == nil {
		this.SendRes(-1, "请先登录", nil)
	}

	id, _ := this.GetInt("id")
	gid, _ := this.GetInt("gid")
	name := this.GetString("name")
	introduction := this.GetString("introduction")
	qrcode := this.GetString("qrcode")
	ower_qrcode := this.GetString("ower_qrcode")
	wechat_id := this.GetString("wechat_id")

	if name == "" || gid <= 0 || (qrcode == "" && ower_qrcode == "" && wechat_id == "") {
		this.SendRes(-1, "参数错误", nil)
	}

	pg := models.GetPrivateGroupById(int(id))
	if pg == nil {
		this.SendRes(-1, "不存在", nil)
	}
	pg.Gid = int(gid)
	pg.Uid = user.(models.User).Id
	pg.Name = name
	pg.Introduction = introduction
	pg.Qrcode = qrcode
	pg.Ower_qrcode = ower_qrcode
	pg.Wechat_id = wechat_id

	models.UpdatePrivateGroup(pg)

	this.SendRes(0, "success", pg)
}

func (this *PrivateGroupController) View() {
	user := this.GetSession("user")
	if user == nil {
		this.Redirect("/", 302)
	}

	id, _ := this.GetInt("id")
	pg := models.GetPrivateGroupById(int(id))
	group := models.GetGroupById(pg.Gid)
	pgMsgs := models.GetPrivateGroupMessageVoByPgid(pg.Id)

	this.Data["pg"] = pg
	this.Data["group"] = group
	this.Data["pgMsgs"] = pgMsgs
	this.Data["user"] = user

	this.Layout = "layout/addwechat.tpl"
	this.TplName = "privateGroup/view.tpl"
}

func (this *PrivateGroupController) CreatePgMsg() {
	user := this.GetSession("user")
	if user == nil {
		this.SendRes(-1, "请先登录", nil)
	}

	content := this.GetString("content")
	pg_id, _ := this.GetInt("pg_id")

	pgm := models.PrivateGroupMessage{}
	pgm.Uid = user.(models.User).Id
	pgm.Pg_id = int(pg_id)
	pgm.Content = content
	_, err := models.CreatePrivateGroupMessage(&pgm)
	if err != nil {
		this.SendRes(-1, err.Error(), nil)
	}

	vo := models.ConvertPrivateGroupMessageToVo(pgm)

	this.SendRes(0, "success", vo)
}

func (this *PrivateGroupController) CreateReport() {
	user := this.GetSession("user")
	if user == nil {
		this.SendRes(-1, "请先登录", nil)
	}

	content := this.GetString("content")
	pg_id, _ := this.GetInt("pg_id")

	pgr := models.PrivateGroupReport{}
	pgr.Uid = user.(models.User).Id
	pgr.Pg_id = int(pg_id)
	pgr.Content = content
	models.CreatePrivateGroupReport(&pgr)

	this.SendRes(0, "success", nil)
}

func (this *PrivateGroupController) List() {
	uid := 0
	user := this.GetSession("user")
	if user != nil {
		uid = user.(models.User).Id
	}

	gid, _ := this.GetInt("gid")
	pgs := models.GetPrivateGroupByGid(int(gid))
	group := models.GetGroupById(int(gid))
	isunlock := false
	if uid > 0 {
		isunlock = models.IsUnlock(uid, int(gid))
	}

	this.Data["pgs"] = pgs
	this.Data["group"] = group
	this.Data["isunlock"] = isunlock
	this.Data["user"] = user

	this.Layout = "layout/addwechat.tpl"
	this.TplName = "privateGroup/list.tpl"
}

func (this *PrivateGroupController) Unlock() {
	user := this.GetSession("user")
	if user == nil {
		this.SendRes(-1, "请先登录", nil)
	}

	uid := user.(models.User).Id
	ugs := models.GetUnlockGroupByUid(uid)
	if len(ugs) > 1 {
		//		this.SendRes(-1, "亲，您只能解锁两所学校，可以付费解锁更多学校", nil)
		this.SendRes(1, "需要支付", nil)
	}

	gid, _ := this.GetInt("gid")
	ug := models.UnlockGroup{}
	ug.Uid = uid
	ug.Gid = int(gid)
	models.CreateUnlockGroup(&ug)

	this.SendRes(0, "success", nil)
}
