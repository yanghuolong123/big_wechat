package models

import (
	"github.com/astaxie/beego/orm"
	"sort"
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
	info, _ := RewardPaySuccess(info_id)

	for i := info.Reward_num; i > 0; i-- {
		ir := new(InfoReward)
		ir.Info_id = info.Id
		rewardlist := RandomReward(info.Reward_amount*float64(info.Reward_num), info.Reward_num) //info.Reward_amount
		ir.Amount = rewardlist[i-1]
		CreateInfoReward(ir)
		help.Redis.Lpush("list_reward_info_"+help.ToStr(info.Id), help.ToStr(ir.Id), 0)
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

		// 个人账户变更
		go func() {
			info, err := GetInfoById(ir.Info_id)
			if err != nil {
				help.Error(err)
				return
			}

			remark := ""
			if info.Reward_type == 1 {
				remark = "获得阅读红包"
			} else if info.Reward_type == 2 {
				remark = "获得留言红包"
			}

			AccountChange(ir.Amount, uid, info.Reward_type, info.Id, remark)
		}()

		return ir
	}

	return ir
}

func RandomReward(amount float64, num int) (rewardlist []float64) {
	var list []int
	total := int(amount * 100)
	avg := total / num
	for i := 0; i < num; i++ {
		list = append(list, avg)
	}

	for i := 0; i < num; i = i + 2 {
		r := help.RandNum(1, avg)
		list[i] = list[i] + r
		list[i+1] = list[i+1] - r
	}

	sort.Ints(list)
	for _, v := range list {
		rewardlist = append(rewardlist, float64(v)/100)
	}

	return rewardlist
}
