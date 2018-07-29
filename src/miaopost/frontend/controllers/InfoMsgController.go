package controllers

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"miaopost/frontend/models"
	"time"
	"yhl/help"
)

type InfoMsgController struct {
	BaseController
}

// 留言/回复
func (this *InfoMsgController) CreateMsg() {
	u := this.GetSession("user")
	if u == nil {
		this.SendRes(-1, "请先登录", nil)
	}
	user := u.(*models.User)

	content := this.GetString("content")
	info_id, _ := this.GetInt("info_id")
	pid, _ := this.GetInt("pid")
	im := models.InfoMessage{
		Uid:     user.Id,
		Info_id: int(info_id),
		Pid:     int(pid),
		Content: content,
	}
	i := models.CreateInfoMessage(&im)
	if i > 0 {
		vo := models.ConvertInfoMsgToVo(&im)

		// 留言红包处理
		c := help.MongoDb.C("pre_msg_reward")
		var ir *models.InfoReward
		err := c.Find(bson.M{"info_id": int(info_id), "uid": user.Id}).One(&ir)
		if err == nil {
			ir := models.GainReward(ir.Id, user.Id)
			vo.Ireward = ir

			c.Remove(bson.M{"id": ir.Id})
		}

		if int(pid) > 0 {
			// 微信提醒回复人
			go models.ReplyWxTip(im.Id, this.Ctx)
		} else {
			// 微信提醒信息发布人
			go models.MessageWxTip(im.Id, this.Ctx)
		}

		this.SendRes(0, "success", vo)
	}

	this.SendRes(-1, "failed", nil)
}

// 建删
func (this *InfoMsgController) SuggestDel() {
	id, _ := this.GetInt("id")

	c := help.MongoDb.C("info_msg_del_sug")
	condition := bson.M{
		"msg_id": int(id),
	}
	job := &mgo.MapReduce{
		Map:    "function(){ emit(this.ip, 1) }",
		Reduce: "function(key, values) { return Array.sum(values) }",
	}
	type record struct {
		Ip    string "_id"
		Count int    "value"
	}
	var result []record
	_, err := c.Find(condition).MapReduce(job, &result)
	if err != nil {
		this.SendRes(-1, err.Error(), nil)
	}
	if len(result) > 3 {
		models.DelInfoMsgById(int(id))
		this.SendRes(0, "success", nil)
	}

	m := map[string]interface{}{
		"msg_id": int(id),
		"time":   time.Now(),
		"ip":     this.Ctx.Input.IP(),
	}
	user := this.GetSession("user")
	if user != nil {
		m["uid"] = user.(*models.User).Id
	}
	c.Insert(m)

	this.SendRes(0, "success", nil)
}

// 赞赏
func (this *InfoMsgController) Admire() {
	mid, _ := this.GetInt("mid")
	// 生成支付订单
	_ = mid
	this.SendRes(0, "success", nil)
	// 支付后个人账号变更
}

// 点赞
func (this *InfoMsgController) Support() {
	id, _ := this.GetInt("id")
	models.Support(int(id))

	this.SendRes(0, "success", nil)
}

func (this *InfoMsgController) View() {
	id, _ := this.GetInt("id")
	infoMsg, err := models.GetInfoMessageById(int(id))
	if err != nil {
		this.Tips(err.Error())
	}
	infoMsgVo := models.ConvertInfoMsgToVo(infoMsg)

	this.Data["infoMsgVo"] = infoMsgVo
	this.Layout = "layout/main.tpl"
	this.TplName = "msg/view.tpl"
}
