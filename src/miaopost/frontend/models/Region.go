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

func GetAllRegion() []Region {
	var rlist []Region
	_, err := orm.NewOrm().QueryTable("tbl_region").Filter("status", 0).All(&rlist)
	help.Error(err)

	return rlist
}

func GetCurrentRegion() (r Region) {
	rlist := GetAllRegion()
	for _, v := range rlist {
		cDomain := v.Name + ".miaopost.com"
		if cDomain == help.ClientDomain {
			r = v
			break
		}
	}

	return
}

func GetAllRegionMap() (m orm.Params) {
	_, err := orm.NewOrm().Raw("select id,shortname from tbl_region where status=0").RowsToMap(&m, "id", "shortname")
	help.Error(err)

	return
}
