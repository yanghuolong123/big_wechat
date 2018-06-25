package controllers

import (
	"labix.org/v2/mgo/bson"
	"miaopost/frontend/models"
	"time"
	"yhl/help"
)

type InfoRewardController struct {
	BaseController
}

func (this *InfoRewardController) Chance() {
	info_id, _ := this.GetInt("info_id")

	u := this.GetSession("user")
	if u == nil {
		this.SendRes(-1, "need login", nil)
	}
	user := u.(*models.User)
	_, err := models.GetInfoRewardByInfoIdAndUid(int(info_id), user.Id)

	if err == nil {
		this.SendRes(0, "has change", nil)
	}

	info, _ := models.GetInfoById(int(info_id))

	if info.Reward_type == 1 {
		val := help.Redis.Rpop("list_reward_info_" + help.ToStr(info_id))
		if val == "" {
			this.SendRes(0, "no change", nil)
		}

		reward_id := help.StrToInt(val)
		ir := models.GainReward(reward_id, user.Id)
		this.SendRes(1, "success", ir)
	}

	if info.Reward_type == 2 {
		c := help.MongoDb.C("pre_msg_reward")

		var ir1 models.InfoReward
		err := c.Find(bson.M{"info_id": int(info_id), "uid": user.Id}).One(&ir1)
		if err == nil {
			this.SendRes(0, "has pre change", ir1)
		}

		val := help.Redis.Rpop("list_reward_info_" + help.ToStr(info_id))
		if val == "" {
			this.SendRes(0, "no change", nil)
		}
		reward_id := help.StrToInt(val)

		ir, err := models.GetInfoRewardById(reward_id)
		ir.Uid = user.Id
		ir.Gain_time = time.Now()
		c.Insert(ir)

		if err == nil {
			this.SendRes(2, "success", ir)
		}
	}

	this.SendRes(-1, "unkown", nil)

}
