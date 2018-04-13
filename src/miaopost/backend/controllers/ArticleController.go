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
	tp, _ := this.GetInt("type")
	page, _ := this.GetInt("page")
	q := model.Query{}
	q.Table = "tbl_article"
	q.Condition = map[string]interface{}{"type": int(tp)}
	q.OrderBy = []string{"-status", "-create_time"}
	var slice []models.Article
	q.ReturnModelList = &slice
	p := help.GetPageList(q, int(page), 1000)

	this.Data["totalCount"] = p.TotalCount
	this.Data["totalPage"] = p.TotalPage
	this.Data["currentPage"] = p.CurrentPage
	this.Data["hasMore"] = p.HasMore
	this.Data["dataList"] = p.DataList.(*[]models.Article)

	tpl := "list.tpl"
	if t := int(tp); t > 0 {
		tpl = "list_type_" + help.ToStr(t) + ".tpl"
	}

	this.Data["type"] = int(tp)

	this.Layout = "layout/main.tpl"
	this.TplName = "article/" + tpl
}

func (this *ArticleController) Create() {
	tp, _ := this.GetInt("type")
	t := int(tp)
	if this.Ctx.Input.IsAjax() {
		group_id, _ := this.GetInt("group_id")
		title := this.GetString("title")
		logo := this.GetString("logo")
		link := this.GetString("link")
		content := this.GetString("content")
		sort, _ := this.GetInt("sort")
		status, _ := this.GetInt("status")

		article := models.Article{
			Type:     t,
			Group_id: int(group_id),
			Title:    title,
			Logo:     logo,
			Link:     link,
			Content:  content,
			Sort:     int(sort),
			Status:   int(status),
		}

		_, err := models.CreateArticle(&article)
		if err != nil {
			this.SendRes(-1, err.Error(), nil)
		}

		this.SendRes(0, "success", article)
	}

	tpl := "create.tpl"
	if t > 0 {
		tpl = "create_type_" + help.ToStr(t) + ".tpl"
	}

	this.Data["type"] = t

	this.Layout = "layout/main.tpl"
	this.TplName = "article/" + tpl
}

func (this *ArticleController) Edit() {
	id, _ := this.GetInt("id")
	tp, _ := this.GetInt("type")
	t := int(tp)
	article, err := models.GetArticleById(int(id))

	if this.Ctx.Input.IsAjax() {
		if err != nil {
			this.SendRes(-1, err.Error(), nil)
		}
		group_id, _ := this.GetInt("group_id")
		title := strings.TrimSpace(this.GetString("title"))
		logo := this.GetString("logo")
		link := strings.TrimSpace(this.GetString("link"))
		content := strings.TrimSpace(this.GetString("content"))
		sort, _ := this.GetInt("sort")
		status, _ := this.GetInt("status")

		article.Group_id = int(group_id)
		article.Title = title
		article.Logo = logo
		article.Content = content
		article.Sort = int(sort)
		article.Status = int(status)
		article.Link = link

		err := models.UpdateArticle(article)
		if err != nil {
			this.SendRes(-1, err.Error(), nil)
		}

		this.SendRes(0, "success", article)
	}

	if err != nil {
		this.Tips(err.Error())
	}

	this.Data["article"] = article

	tpl := "edit.tpl"
	if t > 0 {
		tpl = "edit_type_" + help.ToStr(t) + ".tpl"
	}
	this.Data["type"] = t

	this.Layout = "layout/main.tpl"
	this.TplName = "article/" + tpl
}

func (this *ArticleController) Delete() {
	id, _ := this.GetInt("id")
	tp, _ := this.GetInt("type")
	if !models.DelArticleById(int(id)) {
		this.Tips("删除失败")
	}

	this.Redirect("/article/list?type="+help.ToStr(tp), 302)
}
