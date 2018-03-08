package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

func init() {
	orm.RegisterModelWithPrefix("tbl_", new(Sugguest))
}

type Suggest struct {
	Id          int
	Info_id     int
	Ip          string
	Content     string
	Create_time time.Time
}
