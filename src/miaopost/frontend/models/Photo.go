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

type PhotoVo struct {
	P      Photo
	Height int
	Width  int
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

func DelPhotoByInfoid(infoId int) bool {
	i, err := orm.NewOrm().QueryTable("tbl_photo").Filter("info_id", infoId).Delete()
	help.Error(err)

	return i > 0
}

func ConvertPhotoToVo(p Photo) (vo PhotoVo) {
	vo.P = p
	vo.Height, vo.Width = help.GetImgHW("." + p.Url)

	return
}

func GetPhotoVoList(plist []Photo) (voList []PhotoVo) {
	for _, p := range plist {
		vo := ConvertPhotoToVo(p)
		voList = append(voList, vo)
	}

	return
}
