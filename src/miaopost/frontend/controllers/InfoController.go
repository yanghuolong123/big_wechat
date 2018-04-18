package controllers

import (
	"fmt"
	"miaopost/frontend/models"
	"strings"
	"yhl/help"
)

const pageSize int = 15

type InfoController struct {
	BaseController
}

// 首页
func (this *InfoController) Get() {
	cats := models.GetAllCategory()
	this.Data["cats"] = cats

	infos := models.GetInfoPage(0, 0, pageSize)
	this.Data["infos"] = models.ConvertInfosToVo(&infos)
	count := models.GetInfoCount(0)
	this.Data["hasMore"] = 0
	this.Data["page"] = 0
	if 1*pageSize < count {
		this.Data["hasMore"] = 1
	}

	this.Layout = "layout/main.tpl"
	this.TplName = "info/home.tpl"
}

// 列表页
func (this *InfoController) List() {
	cid, _ := this.GetInt("cid")
	page, _ := this.GetInt("page")

	cats := models.GetAllCategory()
	this.Data["cats"] = cats

	this.Data["hasMore"] = 0
	this.Data["page"] = int(page)

	infos := []models.Info{}
	if catId := int(cid); catId > 0 {
		count := models.GetInfoCount(catId)
		infos = models.GetInfoPage(catId, int(page), pageSize)
		if 1*pageSize < count {
			this.Data["hasMore"] = 1
		}
		//infos = models.GetInfoByCid(catId)
	}

	search := this.GetString("search")
	if search != "" {
		infos = models.SearchInfo(search)
	}

	this.Data["cid"] = int(cid)
	this.Data["infos"] = models.ConvertInfosToVo(&infos)

	this.Layout = "layout/main1.tpl"
	this.TplName = "info/list.tpl"
}

// 创建页面
func (this *InfoController) CreateGet() {
	cid, _ := this.GetInt("cid")
	this.Data["cid"] = int(cid)

	cats := models.GetAllCategory()
	this.Data["cats"] = cats

	this.Layout = "layout/main.tpl"
	this.TplName = "info/create.tpl"
}

func (this *InfoController) CreatePost() {
	cid, _ := this.GetInt("cid")
	content := this.GetString("content")
	valid_day, _ := this.GetInt("valid_day")
	email := this.GetString("email")
	photo := this.GetString("photo")

	info := new(models.Info)
	info.Cid = int(cid)
	info.Content = content
	info.Valid_day = int(valid_day)
	info.Email = email

	id := models.CreateInfo(info)
	if id > 0 {
		if photo != "" {
			plist := strings.Split(photo, ",")
			for _, p := range plist {
				models.CreatePhoto(id, p)
			}
		}

		go func(id int, email string) {
			if email == "" {
				return
			}
			code := help.DesEncrypt(fmt.Sprintf("%v", id)+","+email, help.DesKey)
			linkUrl := "http://www.miaopost.com/info/edit?code=" + code
			msg := "亲，欢迎您使用秒Po，您可以通过点击链接修改您发布的信息 <a href=\"" + linkUrl + "\">进入</a>"
			help.SendMail(email, "秒Po-编辑链接", msg, "html")
		}(id, email)

		this.SendRes(0, "success", info)
	}

	this.SendRes(-1, "failed", nil)
}

// 展示页
func (this *InfoController) View() {
	cats := models.GetAllCategory()
	this.Data["cats"] = cats

	id, _ := this.GetInt("id")
	models.IncInfoViews(int(id))

	info, err := models.GetInfoById(int(id))
	if err != nil {
		this.Redirect("/tips?msg="+err.Error(), 302)
	}
	if info.Status < 0 {
		this.Tips("此信息已经删除!")
	}
	info.Content = strings.Replace(info.Content, "\n", "<br/>", -1)
	this.Data["info"] = info

	cat := models.GetCategoryById(info.Cid)
	this.Data["cat"] = cat

	photos := models.GetPhotoByInfoid(int(id))
	this.Data["photos"] = photos

	share := WxShare
	if info.Content != "" {
		share.Desc = info.Content
	}
	if len(photos) > 0 {
		share.Img = this.Ctx.Input.Site() + photos[0].Url
	}
	share.Title = cat.Name + " - 秒Po"
	share.Link = help.ClientRoute
	this.Data["wxshare"] = share

	this.Layout = "layout/main1.tpl"
	this.TplName = "info/view.tpl"
}

// 编辑
func (this *InfoController) EditGet() {
	code := this.GetString("code")
	code = help.DesDecrypt(code, help.DesKey)

	s := strings.Split(code, ",")
	id := help.StrToInt(s[0])
	if id <= 0 {
		this.Redirect("/tips?msg=code不正确", 302)
	}

	info, err := models.GetInfoById(id)
	if err != nil {
		this.Redirect("/tips?msg="+err.Error(), 302)
	}
	if info.Status < 0 {
		this.Tips("此信息已被删除")
	}
	this.Data["info"] = info

	cats := models.GetAllCategory()
	this.Data["cats"] = cats
	photos := models.GetPhotoByInfoid(id)
	this.Data["photos"] = photos

	this.Layout = "layout/main.tpl"
	this.TplName = "info/edit.tpl"
}

func (this *InfoController) EditPost() {
	id, _ := this.GetInt("id")
	cid, _ := this.GetInt("cid")
	content := this.GetString("content")
	valid_day, _ := this.GetInt("valid_day")
	email := this.GetString("email")
	photo := this.GetString("photo")

	info, err := models.GetInfoById(int(id))
	if err != nil {
		this.SendRes(-1, err.Error(), nil)
	}
	info.Cid = int(cid)
	info.Content = content
	info.Valid_day = int(valid_day)
	info.Email = email
	err = models.UpdateInfo(info)
	if err != nil {
		this.SendRes(-1, err.Error(), nil)
	}

	models.DelPhotoByInfoid(info.Id)
	if photo != "" {
		plist := strings.Split(photo, ",")
		for _, p := range plist {
			models.CreatePhoto(info.Id, p)
		}
	}

	this.SendRes(0, "success", info)
}

// 建议删除
func (this *InfoController) SuggestDel() {
	infoId, _ := this.GetInt("infoId")
	iid := int(infoId)
	id := models.CreateSuggest(iid)
	slist := models.GetSuggestByInfoidAndGroupByIp(iid)
	if len(slist) > 3 {
		models.DelInfoById(iid)
	}
	if id > 0 {
		this.SendRes(0, "success", nil)
	}

	this.SendRes(-1, "failed", nil)
}

// 分页列表
func (this *InfoController) ListPage() {
	page, _ := this.GetInt("page")
	cid, _ := this.GetInt("cid")
	infos := models.GetInfoPage(int(cid), int(page)*pageSize, pageSize)
	this.Data["infos"] = models.ConvertInfosToVo(&infos)
	count := models.GetInfoCount(int(cid))
	hasMore := 0
	if (int(page)+1)*pageSize < count {
		hasMore = 1
	}

	this.Data["hasMore"] = hasMore
	this.TplName = "info/listPage.tpl"
	s, _ := this.RenderString()

	m := map[string]interface{}{}
	m["listData"] = s
	m["page"] = int(page)
	m["hasMore"] = hasMore
	this.SendRes(0, "success", m)
}

// 删除发布的信息
func (this *InfoController) Delete() {
	id, _ := this.GetInt("id")
	if models.DelInfoById(int(id)) {
		this.SendRes(0, "success", nil)
	}
	this.SendRes(-1, "failed", nil)
}
