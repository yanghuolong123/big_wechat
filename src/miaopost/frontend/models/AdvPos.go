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
	_, err := orm.NewOrm().QueryTable("tbl_adv_pos").Filter("status", 0).All(&plist)
	help.Error(err)

	return
}

func GetAdvPosById(id int) (ap *AdvPos) {
	ap = new(AdvPos)
	err := orm.NewOrm().Read(ap)
	help.Error(err)

	return
}
