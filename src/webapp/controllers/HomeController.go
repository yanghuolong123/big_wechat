package controllers

import (
	"yhl/help"
)

type HomeController struct {
	help.BaseController
}

func (this *HomeController) Get() {
	this.Data["welcome"] = "hello, welcome to bigwechart! thank you!"
	this.TplName = "home/index.tpl"
}
