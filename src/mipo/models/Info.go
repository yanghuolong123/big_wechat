package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

func init() {
	orm.RegisterModelWithPrefix("tbl_", new(Info))
}

type Info struct {
	Id          int
	Cid         int
	Content     string
	Valid_day   int
	Email       string
	Status      int
	Views       int
	Create_time time.Time
}
