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
	Id        int
	Info_id   int
	Amount    float64
	Status    int
	Uid       int
	Gain_time time.Time
}

func CreateInfoReward(ir *InfoReward) int {
	_, err := orm.NewOrm().Insert(ir)
	help.Error(err)

	return ir.Id
}

func GetInfoRewardById(id int) (*InfoReward, error) {
	ir := new(InfoReward)
	ir.Id = id
	err := orm.NewOrm().Read(ir)

	return ir, err
}

func GetInfoRewardByInfoIdAndUid(info_id, uid int) (*InfoReward, error) {
	ir := new(InfoReward)
	ir.Info_id = info_id
	ir.Uid = uid
	err := orm.NewOrm().Read(ir, "Info_id", "Uid")
	//help.Error(err)

	return ir, err
}

func GetInfoRewardByInfoId(info_id int) []InfoReward {
	var irs []InfoReward
	_, err := orm.NewOrm().QueryTable("tbl_info_reward").Filter("info_id", info_id).OrderBy("-gain_time").All(&irs)
	help.Error(err)

	return irs
}

func GenBathInfoRewardByInfoId(info_id int) bool {
	info, err := GetInfoById(info_id)
	if err != nil {
		help.Error(err)
		return false
	}

	for i := info.Reward_num; i > 0; i-- {
		ir := new(InfoReward)
		ir.Info_id = info.Id
		ir.Amount = info.Reward_amount
		CreateInfoReward(ir)
		help.Redis.Lpush("list_reward_info_"+help.ToStr(info.Id), help.ToStr(ir.Id))
	}

	return true
}

func GainReward(id, uid int) *InfoReward {
	//	num, err := orm.NewOrm().QueryTable("tbl_info_reward").Filter("id", id).Update(orm.Params{"status": 1, "uid": uid, "gain_time": time.Now()})
	//	help.Error(err)

	ir := new(InfoReward)
	ir.Id = id
	o := orm.NewOrm()
	if o.Read(ir) == nil {
		ir.Status = 1
		ir.Uid = uid
		ir.Gain_time = time.Now()
		o.Update(ir)
		return ir
	}

	return ir
}
