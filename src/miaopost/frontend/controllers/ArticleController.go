package controllers

import (
	"miaopost/frontend/models"
)

type ArticleController struct {
	BaseController
}

func (this *ArticleController) View() {
	id, _ := this.GetInt("id")
	article, err := models.GetArticleById(int(id))
	if err != nil {
		this.Tips(err.Error())
		return
	}

	if article.Type == models.Type_Adv && article.Link != "" {
		this.Redirect(article.Link, 302)
		return
	}

	this.Data["article"] = article
	this.Layout = "layout/main.tpl"
	this.TplName = "article/view.tpl"
}

func (this *ArticleController) AdvEntry() {
	arts := models.GetArticleByType(3)

	if len(arts) <= 0 {
		this.Tips("广告入口页没有创建")
		return
	}

	this.Data["article"] = arts[0]
	this.Layout = "layout/main.tpl"
	this.TplName = "article/advEntry.tpl"
}
