package models

import (
	"github.com/astaxie/beego/orm"
	"yhl/help"
)

func init() {
	orm.RegisterModelWithPrefix("tbl_", new(Region))
}

type Region struct {
	Id       int
	Name     string
	Fullname string
	Status   int
}

func GetAllRegion() []Region {
	var rlist []Region
	_, err := orm.NewOrm().QueryTable("tbl_region").Filter("status", 0).All(&rlist)
	help.Error(err)

	return rlist
}
