package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModelWithPrefix("tbl_", new(Photo))
}

type Photo struct {
	Id      int
	Info_id int
	Url     string
}
