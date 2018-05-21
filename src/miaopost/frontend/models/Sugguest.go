package models

import (
	"github.com/astaxie/beego/context"
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

func CreateSuggest(infoId int, ctx *context.Context) int {
	var sug Suggest

	sug.Info_id = infoId
	sug.Ip = ctx.Input.IP()
	sug.Create_time = time.Now()

	i, err := orm.NewOrm().Insert(&sug)
	help.Error(err)

	return int(i)
}

func GetSuggestByInfoid(infoId int) (ss []Suggest) {
	_, err := orm.NewOrm().QueryTable("tbl_suggest").Filter("info_id", infoId).OrderBy("-create_time").All(&ss)
	help.Error(err)

	return
}

func GetSuggestByInfoidAndGroupByIp(infoId int) (ss []Suggest) {
	_, err := orm.NewOrm().QueryTable("tbl_suggest").Filter("info_id", infoId).GroupBy("ip").OrderBy("-create_time").All(&ss, "Info_id", "Ip")
	help.Error(err)

	return
}
