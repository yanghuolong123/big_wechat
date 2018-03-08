package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"yhl/help"
)

func init() {
	orm.RegisterModelWithPrefix("tbl_", new(Suggest))
}

type Suggest struct {
	Id          int
	Info_id     int
	Ip          string
	Content     string
	Create_time time.Time
}

func CreateSuggest(infoId int) int {
	var sug Suggest

	sug.Info_id = infoId
	sug.Ip = help.ClientIp
	sug.Create_time = time.Now()

	i, err := orm.NewOrm().Insert(&sug)
	help.Error(err)

	return int(i)
}
