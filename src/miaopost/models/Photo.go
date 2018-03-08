package models

import (
	"github.com/astaxie/beego/orm"
	"yhl/help"
)

func init() {
	orm.RegisterModelWithPrefix("tbl_", new(Photo))
}

type Photo struct {
	Id      int
	Info_id int
	Url     string
}

func CreatePhoto(info_id int, urlstr string) bool {
	p := new(Photo)
	p.Info_id = info_id
	p.Url = urlstr

	i, err := orm.NewOrm().Insert(p)
	help.Error(err)

	return i > 0
}

func GetPhotoByInfoid(infoId int) (photos []Photo) {
	_, err := orm.NewOrm().QueryTable("tbl_photo").Filter("info_id", infoId).All(&photos)
	help.Error(err)

	return
}
