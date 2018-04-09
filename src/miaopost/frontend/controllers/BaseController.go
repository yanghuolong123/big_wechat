package controllers

import (
	"miaopost/frontend/models"
	"yhl/help"
)

type BaseController struct {
	help.BaseController
}

func (this *BaseController) Prepare() {
	this.BaseController.Prepare()
	footer_nav := models.GetArticleByGroupId(models.Footer_Nav)
	this.Data["footer_nav"] = footer_nav
	this.Data["last_footer_nav_index"] = len(footer_nav) - 1
}
