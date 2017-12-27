package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

func init() {
	orm.RegisterModelWithPrefix("tbl_", new(PrivateGroupMessage))
}

type PrivateGroupMessage struct {
	Id         int
	Uid        int
	Pg_id      int
	Content    string
	Status     int
	Createtime time.Time
}

func GetPrivateGroupMessageById(id int) (pgm *PrivateGroupMessage) {
	pgm = &PrivateGroupMessage{Id: id}
	orm.NewOrm().Read(pgm)
	return
}

func CreatePrivateGroupMessage(pgm *PrivateGroupMessage) int {
	pgm.Createtime = time.Now()
	i, _ := orm.NewOrm().Insert(pgm)

	return int(i)
}

func GetPrivateGroupMessageByPgid(pg_id int) (pgms []PrivateGroupMessage) {
	orm.NewOrm().QueryTable("tbl_private_group_message").Filter("pg_id", pg_id).All(&pgms)
	return
}
