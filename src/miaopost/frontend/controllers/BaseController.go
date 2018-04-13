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

	footer_nav := models.GetArticleByType(models.Type_Nav)
	this.Data["footer_nav"] = footer_nav
	this.Data["last_footer_nav_index"] = len(footer_nav) - 1

	side_adv := models.GetArticleByTypeAndGroup(models.Type_Adv, models.Adv_Side)
	this.Data["side_adv"] = side_adv

	this.Data["version"] = help.Version

}
