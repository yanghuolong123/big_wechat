package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"yhl/help"
)

func init() {
	orm.RegisterModelWithPrefix("tbl_", new(InfoMessage))
}

type InfoMessage struct {
	Id          int
	Uid         int
	Info_id     int
	Pid         int
	Content     string
	Status      int
	Create_time time.Time
}

type InfoMsgVo struct {
	Im   *InfoMessage
	User *User
}

func CreateInfoMessage(im *InfoMessage) int {
	im.Create_time = time.Now()
	i, err := orm.NewOrm().Insert(im)
	help.Error(err)

	return int(i)
}

func GetInfoMessageByInfoId(info_id int) (ims []*InfoMessage) {
	_, err := orm.NewOrm().QueryTable("tbl_info_message").Filter("info_id", info_id).OrderBy("-create_time").All(&ims)
	help.Error(err)

	return
}

func ConvertInfoMsgToVo(im *InfoMessage) (vo InfoMsgVo) {
	u, err := GetUserById(im.Uid)
	help.Error(err)

	vo.Im = im
	vo.User = u

	return
}

func ConvertInfoMsgToVos(ims []*InfoMessage) (vos []InfoMsgVo) {
	for _, im := range ims {
		vo := ConvertInfoMsgToVo(im)
		vos = append(vos, vo)
	}

	return
}
