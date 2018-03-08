package controllers

import (
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
func (this *InfoController) Create() {
	this.TplName = "info/create.tpl"
}

// 展示页
func (this *InfoController) View() {
	this.TplName = "info/view.tpl"
}
