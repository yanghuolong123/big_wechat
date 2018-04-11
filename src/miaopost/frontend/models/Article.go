package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"yhl/help"
)

const (
	Footer_Nav = iota + 1
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

func CreateArticle(article *Article) (int, error) {
	article.Create_time = time.Now()
	i, err := orm.NewOrm().Insert(article)
	help.Error(err)

	return int(i), err
}

func UpdateArticle(article *Article) error {
	_, err := orm.NewOrm().Update(article)
	help.Error(err)

	return err
}

func GetArticleById(id int) (*Article, error) {
	article := &Article{Id: id}
	err := orm.NewOrm().Read(article)
	help.Error(err)

	return article, err
}

func DelArticleById(id int) bool {
	//i, err := orm.NewOrm().QueryTable("tbl_article").Filter("id", id).Update(orm.Params{"status": -1})
	i, err := orm.NewOrm().QueryTable("tbl_article").Filter("id", id).Delete()
	help.Error(err)

	return int(i) > 0
}

func GetArticleGroupMap() map[int]string {
	m := map[int]string{
		Footer_Nav: "底部导航",
	}

	return m
}

func GetArticleByGroupId(gid int) []Article {
	var as []Article
	_, err := orm.NewOrm().QueryTable("tbl_article").Filter("group_id", gid).Filter("status", 0).OrderBy("-sort").All(&as)
	help.Error(err)

	return as
}

func GetArticleStatusMap() map[int]string {
	return map[int]string{
		0:  "启用",
		-1: "禁用",
	}
}
