package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

func init() {
	orm.RegisterModelWithPrefix("tbl_", new(PrivateGroup))
}

type PrivateGroup struct {
	Id           int
	Gid          int
	Uid          int
	Name         string
	Introduction string
	Qrcode       string
	Ower_qrcode  string
	Wechat_id    string
	Status       int
	Createtime   time.Time
}

func CreatePrivateGroup(pg *PrivateGroup) bool {
	pg.Createtime = time.Now()

	i, _ := orm.NewOrm().Insert(pg)
	return i > 0
}

func UpdatePrivateGroup(pg *PrivateGroup) bool {
	if pg.Id <= 0 {
		return false
	}

	i, _ := orm.NewOrm().Update(pg)

	return i > 0
}

func GetPrivateGroupById(id int) (pg *PrivateGroup) {
	pg = &PrivateGroup{Id: id}
	orm.NewOrm().Read(pg)

	return
}

func GetPrivateGroupByGid(gid int) (pg_slice []PrivateGroup) {
	orm.NewOrm().QueryTable("tbl_private_group").Filter("gid", gid).All(&pg_slice)
	return
}

func GetPrivateGroupByLimit(limit int) (pg_slice []PrivateGroup) {
	orm.NewOrm().QueryTable("tbl_private_group").Limit(limit).OrderBy("-createtime").All(&pg_slice)
	return
}
