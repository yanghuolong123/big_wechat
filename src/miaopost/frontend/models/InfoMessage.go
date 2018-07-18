package models

import (
	"github.com/astaxie/beego/orm"
	"regexp"
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
	Support     int
	Status      int
	Create_time time.Time
}

type InfoMsgVo struct {
	Im      *InfoMessage
	User    *User
	Parent  *InfoMsgVo
	Ireward *InfoReward
	IsImg   bool
}

func CreateInfoMessage(im *InfoMessage) int {
	im.Create_time = time.Now()
	i, err := orm.NewOrm().Insert(im)
	help.Error(err)

	return int(i)
}

func GetInfoMessageById(id int) (im *InfoMessage, err error) {
	im = &InfoMessage{Id: id}
	err = orm.NewOrm().Read(im)
	help.Error(err)

	return
}

func GetInfoMessageByInfoId(info_id int) (ims []*InfoMessage) {
	_, err := orm.NewOrm().QueryTable("tbl_info_message").Filter("info_id", info_id).Filter("status", 0).OrderBy("-create_time").All(&ims)
	help.Error(err)

	return
}

func ConvertInfoMsgToVo(im *InfoMessage) (vo InfoMsgVo) {
	u, err := GetUserById(im.Uid)
	help.Error(err)

	vo.Im = im
	vo.User = u
	if im.Pid > 0 {
		p, err := GetInfoMessageById(im.Pid)
		if err == nil {
			pvo := ConvertInfoMsgToVo(p)
			vo.Parent = &pvo
		}
	}
	re, _ := regexp.Compile(`<img.*?(?:>|\/>)`)
	if re.MatchString(im.Content) {
		vo.IsImg = true
	} else {
		vo.IsImg = false
	}

	return
}

func ConvertInfoMsgToVos(ims []*InfoMessage) (vos []InfoMsgVo) {
	for _, im := range ims {
		vo := ConvertInfoMsgToVo(im)
		vos = append(vos, vo)
	}

	return
}

func DelInfoMsgById(id int) bool {
	i, err := orm.NewOrm().QueryTable("tbl_info_message").Filter("id", id).Update(orm.Params{"status": -1})
	help.Error(err)

	return i > 0
}

func Support(id int) bool {
	i, err := orm.NewOrm().QueryTable("tbl_info_message").Filter("id", id).Update(orm.Params{"support": orm.ColValue(orm.ColAdd, 1)})
	help.Error(err)

	return i > 0
}
