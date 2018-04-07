package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"yhl/help"
)

func init() {
	orm.RegisterModelWithPrefix("tbl_", new(Article))
}

type Article struct {
	Id          int
	Group_id    int
	Title       string
	Content     string
	Sort        int
	Status      int
	Create_time time.Time
}

func CreateArticle(article *Article) int {
	article.Create_time = time.Now()
	i, err := orm.NewOrm().Insert(article)
	help.Error(err)

	return int(i)
}

func UpdateArticle(article *Article) error {
	_, err := orm.NewOrm().Update(article)
	help.Error(err)

	return err
}

func GetArticleById(id int) *Article {
	article := &Article{Id: id}
	err := orm.NewOrm().Read(article)
	help.Error(err)

	return article
}

func DelArticleById(id int) bool {
	i, err := orm.NewOrm().QueryTable("tbl_article").Filter("id", id).Update(orm.Params{"status": -1})
	help.Error(err)

	return int(i) > 0
}
