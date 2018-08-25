package models

import (
	"github.com/astaxie/beego/orm"
	"yhl/help"
)

func init() {
	orm.RegisterModelWithPrefix("tbl_", new(AdvPos))
}

type AdvPos struct {
	Id     int
	Type   int
	Name   string
	Status int
}

func GetAdvPosList() (plist []AdvPos) {
	_, err := orm.NewOrm().QueryTable("tbl_adv_pos").Filter("status", 0).OrderBy("id").All(&plist)
	help.Error(err)

	return
}

func GetAdvPosByType(ptype int) (plist []AdvPos) {
	_, err := orm.NewOrm().QueryTable("tbl_adv_pos").Filter("type", ptype).Filter("status", 0).OrderBy("id").All(&plist)
	help.Error(err)

	return
}

func GetAdvPosById(id int) (ap *AdvPos) {
	ap = new(AdvPos)
	ap.Id = id
	err := orm.NewOrm().Read(ap)
	help.Error(err)

	return
}
