package controllers

import (
	"miaopost/models"
	"yhl/help"
)

type InfoController struct {
	help.BaseController
}

// 首页
func (this *InfoController) Get() {
	cats := models.GetAllCategory()
	infos := models.GetInfoPage(0, 10)

	this.Data["cats"] = cats
	this.Data["infos"] = infos

	this.Layout = "layout/main.tpl"
	this.TplName = "info/index.tpl"
}

// 列表页
func (this *InfoController) List() {
	cid, _ := this.GetInt("cid")

	infos := models.GetInfoByCid(int(cid))
	cats := models.GetAllCategory()

	this.Data["cats"] = cats
	this.Data["infos"] = infos

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
		_ = photo
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
	info := models.GetInfoById(int(id))
	this.Data["info"] = info

	photos := models.GetPhotoByInfoid(int(id))
	this.Data["photos"] = photos

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
