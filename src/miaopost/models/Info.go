package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"yhl/help"
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

func CreateInfo(info *Info) bool {
	info.Create_time = time.Now()
	i, err := orm.NewOrm().Insert(info)
	if err != nil {
		help.Log("error", err.Error())
	}

	return i > 0
}

func GetInfoById(id int) *Info {
	info := &Info{Id: id}

	err := orm.NewOrm().Read(info)
	if err != nil {
		help.Log("error", err.Error())
	}

	return info
}

func GetInfoByCid(cid int) []Info {
	var infos []Info
	_, err := orm.NewOrm().QueryTable("tbl_info").Filter("cid", cid).All(&infos)
	help.Error(err)

	return infos
}

func GetInfoByEmail(email string) []Info {
	var infos []Info
	_, err := orm.NewOrm().QueryTable("tbl_info").Filter("email", email).All(&infos)
	if err != nil {
		help.Log("error", err.Error())
	}

	return infos
}

func GetInfoPage(offset, size int) (infos []Info) {
	_, err := orm.NewOrm().QueryTable("tbl_info").Filter("status", 0).OrderBy("-create_time").Limit(size, offset).All(&infos)
	help.Error(err)

	return
}

func GetInfoCount() int {
	count, err := orm.NewOrm().QueryTable("tbl_info").Filter("status", 0).Count()
	help.Error(err)

	return int(count)
}
