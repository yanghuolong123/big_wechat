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

type PrivateGroupMessageVo struct {
	Pgm  PrivateGroupMessage
	User User
}

func GetPrivateGroupMessageById(id int) (pgm *PrivateGroupMessage) {
	pgm = &PrivateGroupMessage{Id: id}
	orm.NewOrm().Read(pgm)
	return
}

func CreatePrivateGroupMessage(pgm *PrivateGroupMessage) (int, error) {
	pgm.Createtime = time.Now()
	i, err := orm.NewOrm().Insert(pgm)

	return int(i), err
}

func GetPrivateGroupMessageByPgid(pg_id int) (pgms []PrivateGroupMessage) {
	orm.NewOrm().QueryTable("tbl_private_group_message").Filter("pg_id", pg_id).OrderBy("-createtime").All(&pgms)
	return
}

func GetPrivateGroupMessageVoByPgid(pg_id int) (vos []PrivateGroupMessageVo) {
	pgms := GetPrivateGroupMessageByPgid(pg_id)

	for _, m := range pgms {
		vo := ConvertPrivateGroupMessageToVo(m)
		vos = append(vos, vo)
	}

	return
}

func ConvertPrivateGroupMessageToVo(pgm PrivateGroupMessage) (vo PrivateGroupMessageVo) {
	vo = PrivateGroupMessageVo{}
	u, _ := GetUserById(pgm.Uid)
	vo.User = *u
	vo.Pgm = pgm

	return
}
