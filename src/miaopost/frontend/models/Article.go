package models

import (
	"github.com/astaxie/beego/orm"
	"math/rand"
	"time"
	"yhl/help"
)

const (
	Type_Nav = iota + 1
	Type_Adv
)

const (
	Nav_Footer = iota + 1
)

const (
	Adv_Side = iota + 1
	Adv_Side_1
	Adv_List_Bottom
	Adv_View_Bottom
)

func init() {
	orm.RegisterModelWithPrefix("tbl_", new(Article))
}

type Article struct {
	Id          int
	Type        int
	Group_id    int
	Title       string
	Logo        string
	Link        string
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

func GetArticleGroupMap() map[int]map[int]string {
	m := map[int]map[int]string{}
	m1 := map[int]string{
		Nav_Footer: "底部导航",
	}
	m2 := map[int]string{
		Adv_Side:        "侧边栏广告位1",
		Adv_Side_1:      "侧边栏广告位2",
		Adv_List_Bottom: "列表底部广告位",
		Adv_View_Bottom: "详情页底部广告位",
	}

	m[Type_Nav] = m1
	m[Type_Adv] = m2

	return m
}

func GetArticleByType(tp int) []Article {
	var as []Article
	_, err := orm.NewOrm().QueryTable("tbl_article").Filter("type", tp).Filter("status", 0).OrderBy("-sort").All(&as)
	help.Error(err)

	return as
}

func GetArticleByTypeAndGroup(tp, gp int) []Article {
	var as []Article
	_, err := orm.NewOrm().QueryTable("tbl_article").Filter("type", tp).Filter("group_id", gp).Filter("status", 0).OrderBy("-sort").All(&as)
	help.Error(err)

	return as
}

func GetArticleStatusMap() map[int]string {
	return map[int]string{
		0:  "启用",
		-1: "禁用",
	}
}

func RandAdv(advs []Article, size int) (adv []Article) {
	s := len(advs)
	if s == 0 {
		return
	}
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(s)
	adv = advs[r : r+size]

	return
}
