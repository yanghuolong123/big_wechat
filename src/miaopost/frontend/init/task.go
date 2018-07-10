package init

import (
	"fmt"
	"github.com/astaxie/beego/toolbox"
	"labix.org/v2/mgo/bson"
	"miaopost/frontend/models"
	"time"
	"yhl/help"
	"yhl/wechat"
)

func init() {
	autoDelExpireInfo := toolbox.NewTask("mytask", "0 0/10 * * * *", func() error {
		f := models.DelExpireInfo()
		help.Log("task", fmt.Sprintf("%v:%v", "自动删除过期的发布信息：", f))
		return nil
	})
	toolbox.AddTask("autoDelExpireInfo", autoDelExpireInfo)

	statEmail := toolbox.NewTask("statEmail", "0 10 0 * * *", statEmail)
	toolbox.AddTask("statEmail", statEmail)

	clearWxCache := toolbox.NewTask("clearWxCache", "0 */30 * * * *", clearWxCache)
	toolbox.AddTask("clearWxCache", clearWxCache)

	recoverRewardList := toolbox.NewTask("recoverRewardList", "0 */2 * * * *", recoverRewardList)
	toolbox.AddTask("recoverRewardList", recoverRewardList)

	toolbox.StartTask()
}

func statEmail() error {
	help.Log("task", "每日统计执行...")
	go func() {
		t := time.Now().AddDate(0, 0, -1)
		date_begin := help.GetDateBegin(t)
		date_end := help.GetDateEnd(t)

		date_pv := models.StatPv(date_begin, date_end)
		date_uv := models.StatUv(date_begin, date_end)
		date_num := models.StatCountInfo(date_begin, date_end)

		mon_begin := help.GetMonthBegin(t)
		mon_end := help.GetMonthEnd(t)
		mon_pv := models.StatPv(mon_begin, mon_end)
		mon_uv := models.StatUv(mon_begin, mon_end)
		mon_num := models.StatCountInfo(mon_begin, mon_end)

		email := "yhl27ml@126.com,wenbo1zhao@126.com"
		msg := "<br/><h3>本日数据统计(" + t.Format("2006年01月02日") + ")：</h3><br/>"
		msg += "&nbsp;&nbsp;&nbsp;&nbsp; PV: " + help.ToStr(date_pv) + "<br/>"
		msg += "&nbsp;&nbsp;&nbsp;&nbsp; UV: " + help.ToStr(date_uv) + "<br/>"
		msg += "&nbsp;&nbsp;&nbsp;&nbsp; 信息发布量: " + help.ToStr(date_num) + "<br/>"
		msg += "<br><h3>当月数据统计(" + t.Format("2006年01月") + ")：</h3><br/>"
		msg += "&nbsp;&nbsp;&nbsp;&nbsp; PV: " + help.ToStr(mon_pv) + "<br/>"
		msg += "&nbsp;&nbsp;&nbsp;&nbsp; UV: " + help.ToStr(mon_uv) + "<br/>"
		msg += "&nbsp;&nbsp;&nbsp;&nbsp; 信息发布量: " + help.ToStr(mon_num) + "<br/>"
		msg += "<p style='color:red;'>PV: 点击量，用户每点击一次计数 <br/> UV: 独立IP用户，每个独立ip算一个用户</p"

		help.SendMail(email, "秒Po每日数据统计", msg, "html")

	}()

	return nil
}

func clearWxCache() error {
	help.Log("task", "更新微信token.....")
	cache := help.Cache
	cache.Delete("access_token_" + wechat.Appid)
	cache.Delete("jsapi_ticket_" + wechat.Appid)
	//	cache.Delete("qr_img_create")

	accessToken := wechat.GetAccessToken()
	jsapiTickey := wechat.GetJsApiTickey()
	//	createImg := wechat.GetTmpStrQrImg("create")
	help.Log("task", "=============== new accessToken:"+accessToken)
	help.Log("task", "=============== new jsapiTickey:"+jsapiTickey)
	//	help.Log("task", "=============== new createImg:"+createImg)

	return nil
}

func recoverRewardList() error {
	help.Log("task", "恢复红包队列...")

	c := help.MongoDb.C("pre_msg_reward")
	var irs []*models.InfoReward
	c.Find(nil).All(&irs)
	for _, ir := range irs {
		if !time.Now().Before(ir.Gain_time.Add(time.Minute * 45)) {
			c.Remove(bson.M{"id": ir.Id})
			help.Redis.Lpush("list_reward_info_"+help.ToStr(ir.Info_id), help.ToStr(ir.Id), 0)
			help.Log("task", ir)
		}
	}

	return nil
}
