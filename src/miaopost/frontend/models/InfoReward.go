package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"yhl/help"
)

func init() {
	orm.RegisterModelWithPrefix("tbl_", new(InfoReward))
}

type InfoReward struct {
	Id          int
	Info_id     int
	Amount      float64
	Status      int
	Uid         int
	Gain_time   time.Time
	Effect_time time.Time
}

func CreateInfoReward(ir *InfoReward) int {
	_, err := orm.NewOrm().Insert(ir)
	help.Error(err)

	return ir.Id
}

func GetInfoRewardByInfoId(info_id int) []InfoReward {
	var irs []InfoReward
	_, err := orm.NewOrm().QueryTable("tbl_info_reward").Filter("info_id", info_id).OrderBy("-gain_time").All(&irs)
	help.Error(err)

	return irs
}
