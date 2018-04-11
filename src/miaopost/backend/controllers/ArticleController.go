package controllers

import (
	"miaopost/backend/models"
	"strings"
	"yhl/help"
	"yhl/model"
)

type ArticleController struct {
	BaseController
}

func (this *ArticleController) Prepare() {
	this.BaseController.Prepare()

	this.Data["groupMap"] = models.GetArticleGroupMap()
	this.Data["statusMap"] = models.GetArticleStatusMap()
}

func (this *ArticleController) List() {
	page, _ := this.GetInt("page")
	q := model.Query{}
	q.Table = "tbl_article"
	q.Condition = map[string]interface{}{"group_id": 1}
	q.OrderBy = []string{"-status", "-create_time"}
	var slice []models.Article
	q.ReturnModelList = &slice
	p := help.GetPageList(q, int(page), 1000)

	this.Data["totalCount"] = p.TotalCount
	this.Data["totalPage"] = p.TotalPage
	this.Data["currentPage"] = p.CurrentPage
	this.Data["hasMore"] = p.HasMore
	this.Data["dataList"] = p.DataList.(*[]models.Article)

	this.Layout = "layout/main.tpl"
	this.TplName = "article/list.tpl"
}

func (this *ArticleController) Create() {
	if this.Ctx.Input.IsAjax() {
		group_id, _ := this.GetInt("group_id")
		title := this.GetString("title")
		content := this.GetString("content")
		sort, _ := this.GetInt("sort")
		status, _ := this.GetInt("status")

		article := models.Article{
			Group_id: int(group_id),
			Title:    title,
			Content:  content,
			Sort:     int(sort),
			Status:   int(status),
		}

		_, err := models.CreateArticle(&article)
		if err != nil {
			this.SendRes(-1, err.Error(), nil)
		}

		this.SendRes(0, "success", nil)
	}

	this.Layout = "layout/main.tpl"
	this.TplName = "article/create.tpl"
}

func (this *ArticleController) Edit() {
	id, _ := this.GetInt("id")
	article, err := models.GetArticleById(int(id))

	if this.Ctx.Input.IsAjax() {
		if err != nil {
			this.SendRes(-1, err.Error(), nil)
		}
		title := strings.TrimSpace(this.GetString("title"))
		content := strings.TrimSpace(this.GetString("content"))
		sort, _ := this.GetInt("sort")
		status, _ := this.GetInt("status")

		if title != "" {
			article.Title = title
		}
		if content != "" {
			article.Content = content
		}
		if s := int(sort); s > 0 {
			article.Sort = s
		}
		article.Status = int(status)

		err := models.UpdateArticle(article)
		if err != nil {
			this.Redirect("/article/list", 302)
		}

		this.SendRes(0, "", nil)
	}

	if err != nil {
		this.Tips(err.Error())
	}

	this.Data["article"] = article

	this.Layout = "layout/main.tpl"
	this.TplName = "article/edit.tpl"
}

func (this *ArticleController) Delete() {
	id, _ := this.GetInt("id")
	if !models.DelArticleById(int(id)) {
		this.Tips("删除失败")
	}

	this.Redirect("/article/list", 302)
}
