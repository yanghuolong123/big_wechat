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
	En_name    string
	Short_name string
	Region     string
	Introduce  string
	Status     int
}

func GetGroupById(id int) (group *Group) {
	group = &Group{Id: id}
	orm.NewOrm().Read(group)
	return
}

func GetGroupAll() (glist []Group) {
	orm.NewOrm().QueryTable("tbl_group").All(&glist)
	return
}

func SearchGroup(name string) (groups []Group) {
	cond := orm.NewCondition()
	cond1 := cond.Or("short_name__icontains", name).Or("en_name__icontains", name).Or("name__icontains", name)
	orm.NewOrm().QueryTable("tbl_group").SetCond(cond1).All(&groups)
	return
}
