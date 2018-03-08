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
	this.TplName = "info/index.tpl"
}

// 列表页
func (this *InfoController) List() {
	this.TplName = "info/list.tpl"
}

// 创建页面
func (this *InfoController) CreateGet() {

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

	is := models.CreateInfo(info)
	if is {
		_ = photo
		this.SendRes(0, "success", nil)
	}

	this.SendRes(-1, "failed", nil)
}

// 展示页
func (this *InfoController) View() {
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
	this.SendRes(0, "", nil)
}
