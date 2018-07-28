package controllers

import (
	"miaopost/frontend/models"
	"yhl/help"
	"yhl/model"
)

type UserController struct {
	BaseController
}

func (this *UserController) Prepare() {
	user := this.GetSession("user")
	if user == nil {
		this.Redirect("/", 302)
		return
	}

	this.BaseController.Prepare()
	this.Data["uri"] = this.Ctx.Input.URI()
}

func (this *UserController) Index() {
	this.Redirect("/user/my", 302)

	this.Layout = "layout/main.tpl"
	this.TplName = "user/index.tpl"
}

// 编辑用户信息 昵称等
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
		return
	}

	this.SetSession("user", u)

	this.Redirect("/user", 302)

}

// 我发布的信息
func (this *UserController) My() {
	u := this.GetSession("user")
	user := u.(*models.User)
	q := model.Query{}
	q.Table = "tbl_info"
	cm := map[string]interface{}{}
	cm["uid"] = user.Id
	cm["status"] = 0
	q.Condition = cm
	q.OrderBy = []string{"-update_time"}
	var slice []*models.Info
	q.ReturnModelList = &slice
	p := help.GetPageList(q, 0, 15)
	data := p.DataList

	infos := data.(*[]*models.Info)
	this.Data["infos"] = models.ConvertInfosToVo2(infos)
	this.Data["page"] = p.CurrentPage
	this.Data["hasMore"] = p.HasMore
	this.Data["uid"] = user.Id
	this.Data["isMy"] = true

	this.Layout = "layout/main.tpl"
	this.TplName = "user/my.tpl"
}

// 钱包，账户金额/明细
func (this *UserController) Account() {
	u := this.GetSession("user")
	user := u.(*models.User)

	ua, _ := models.GetUserAccountByUid(user.Id)
	this.Data["ua"] = ua

	uad := models.GetUserAccountDetailByUid(user.Id)
	this.Data["uad"] = uad

	this.Layout = "layout/main.tpl"
	this.TplName = "user/account.tpl"
}
