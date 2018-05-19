package models

import (
	"github.com/astaxie/beego/orm"
	"yhl/help"
)

func init() {
	orm.RegisterModelWithPrefix("tbl_", new(Category))
}

type Category struct {
	Id     int
	Rid    int
	Name   string
	Sort   int
	Status int
}

func CreateCategory(name string) bool {
	c := new(Category)
	c.Name = name

	i, err := orm.NewOrm().Insert(c)
	if err != nil {
		help.Log("error", err.Error())
	}

	return i > 0
}

func GetCategoryById(id int) Category {
	cat := Category{Id: id}
	err := orm.NewOrm().Read(&cat)
	help.Error(err)

	return cat
}

func GetAllCategory() []Category {
	var clist []Category
	_, err := orm.NewOrm().QueryTable("tbl_category").Filter("status", 0).OrderBy("-sort").All(&clist)
	if err != nil {
		help.Log("error", err.Error())
	}

	return clist
}

func GetCatsByRid(rid int) []Category {
	var clist []Category
	_, err := orm.NewOrm().QueryTable("tbl_category").Filter("status", 0).Filter("rid", rid).OrderBy("-sort").All(&clist)
	if err != nil {
		help.Log("error", err.Error())
	}

	return clist
}
