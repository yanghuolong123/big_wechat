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
