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

func CreatePrivateGroup(pg *PrivateGroup) error {
	pg.Createtime = time.Now()

	_, err := orm.NewOrm().Insert(pg)

	return err
}

func UpdatePrivateGroup(pg *PrivateGroup) error {
	_, err := orm.NewOrm().Update(pg)

	return err
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

func GetPrivateGroupByUid(uid int) (pg_slice []PrivateGroup) {
	orm.NewOrm().QueryTable("tbl_private_group").Filter("uid", uid).All(&pg_slice)
	return
}

func GetPrivateGroupByLimit(limit int) (pg_slice []PrivateGroup) {
	orm.NewOrm().QueryTable("tbl_private_group").Limit(limit).OrderBy("-createtime").All(&pg_slice)
	return
}
