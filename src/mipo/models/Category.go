package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModelWithPrefix("tbl_", new(Category))
}

type Category struct {
	Id     int
	Name   string
	status int
}
