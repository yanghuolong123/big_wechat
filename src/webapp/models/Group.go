package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModelWithPrefix("tbl_", new(Group))
}

type Group struct {
	Id         int
	Name       string
	Short_name string
	Introduce  string
	Status     int
}

func GetGroupById(id int) (group *Group) {
	group = &Group{Id: id}
	orm.NewOrm().Read(group)
	return
}

func getGroupAll() (glist []Group) {
	orm.NewOrm().QueryTable("tbl_group").All(&glist)
	return
}
