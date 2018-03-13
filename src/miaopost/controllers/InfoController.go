package controllers

import (
	"miaopost/models"
	"strings"
	"yhl/help"
)

type InfoController struct {
	help.BaseController
}

// 首页
func (this *InfoController) Get() {
	cats := models.GetAllCategory()
	this.Data["cats"] = cats

	infos := models.GetInfoPage(0, 10)
	this.Data["infos"] = models.ConvertInfosToVo(infos)

	this.Layout = "layout/main.tpl"
	this.TplName = "info/home.tpl"
}

// 列表页
func (this *InfoController) List() {
	cid, _ := this.GetInt("cid")

	cats := models.GetAllCategory()
	this.Data["cats"] = cats

	infos := []models.Info{}
	if catId := int(cid); catId > 0 {
		infos = models.GetInfoByCid(catId)
	}

	search := this.GetString("search")
	if search != "" {
		infos = models.SearchInfo(search)
	}

	this.Data["infos"] = models.ConvertInfosToVo(infos)

	this.Layout = "layout/main1.tpl"
	this.TplName = "info/list.tpl"
}

// 创建页面
func (this *InfoController) CreateGet() {
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
			// todo .... 发邮件
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
		this.Redirect("/", 302)
	}
	this.Data["info"] = info

	photos := models.GetPhotoByInfoid(int(id))
	this.Data["photos"] = photos

	this.Layout = "layout/main1.tpl"
	this.TplName = "info/view.tpl"
}

// 编辑
func (this *InfoController) EditGet() {
	this.TplName = "info/edit.tpl"
}

func (this *InfoController) EditPost() {
	this.SendRes(0, "", nil)
}

// 建议删除
func (this *InfoController) SuggestDel() {
	infoId, _ := this.GetInt("infoId")
	id := models.CreateSuggest(int(infoId))
	if id > 0 {
		this.SendRes(0, "success", nil)
	}

	this.SendRes(-1, "failed", nil)
}