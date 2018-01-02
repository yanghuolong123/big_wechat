package controllers

import (
	"webapp/models"
	"yhl/help"
)

type SearchController struct {
	help.BaseController
}

func (this *SearchController) Group() {
	name := this.GetString("name")
	groups := models.SearchGroup(name)

	this.SendRes(0, "success", groups)
}
