package models

import (
	"github.com/astaxie/beego/orm"
	"yhl/help"
)

func init() {
	orm.RegisterModelWithPrefix("tbl_", new(Region))
}

type Region struct {
	Id        int
	Name      string
	Shortname string
	Fullname  string
	Status    int
}

func GetRegionById(id int) (region *Region) {
	region = &Region{Id: id}
	err := orm.NewOrm().Read(region)
	help.Error(err)

	return
}

func GetRegionByName(name string) (region *Region) {
	region = &Region{Name: name}
	err := orm.NewOrm().Read(region, "name")
	help.Error(err)

	return
}

func GetAllRegion() []Region {
	var rlist []Region
	_, err := orm.NewOrm().QueryTable("tbl_region").Filter("status", 0).All(&rlist)
	help.Error(err)

	return rlist
}
