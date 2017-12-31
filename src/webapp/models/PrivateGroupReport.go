package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

func init() {
	orm.RegisterModelWithPrefix("tbl_", new(PrivateGroupReport))
}

type PrivateGroupReport struct {
	Id         int
	Uid        int
	Pg_id      int
	Content    string
	Status     int
	Createtime time.Time
}

func CreatePrivateGroupReport(pgr *PrivateGroupReport) int {
	pgr.Createtime = time.Now()
	i, _ := orm.NewOrm().Insert(pgr)

	return int(i)
}
