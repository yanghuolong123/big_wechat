package controllers

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"miaopost/frontend/models"
	"strings"
	"time"
	"yhl/help"
	"yhl/wechat"
)

type WechatController struct {
	help.BaseController
}

func (this *WechatController) Index() {
	timestamp, nonce, signatureIn := this.GetString("timestamp"), this.GetString("nonce"), this.GetString("signature")
	echostr := strings.TrimSpace(this.GetString("echostr"))

	if wechat.Check(timestamp, nonce, signatureIn) {
		if len(echostr) > 0 {
			this.Ctx.WriteString(echostr)
			help.Log("wx.log", "===== 接入成功 ======")
			return
		}

		body, err := ioutil.ReadAll(this.Ctx.Request.Body)
		if err != nil {
			return
		}

		fmt.Println(string(body))

		msgBody := &wechat.MsgBody{}
		xml.Unmarshal(body, msgBody)
		fmt.Println(msgBody)

		reply := listen(msgBody)
		if reply != nil {
			this.Data["xml"] = reply
			this.ServeXML()
		}

	}

	this.StopRun()
}

func listen(msgBody *wechat.MsgBody) *wechat.MsgBody {
	url := "http://www.miaopost.com/info/create"
	if msgBody.MsgType == "event" && (msgBody.Event == "subscribe" || msgBody.Event == "SCAN") {
		if strings.Contains(msgBody.EventKey, "create") {

			return replyText(msgBody, "关联成功，发布信息完成后，您可以通过公众号的 “我的发布” 菜单进行便捷操作! <a href=\""+url+"\">发布我的信息</a>")
		}

		if strings.Contains(msgBody.EventKey, "login_") {
			msgBody.EventKey = strings.TrimLeft(msgBody.EventKey, "qrscene_")
			key := scanLogin(msgBody)
			_ = key
			url := "http://www.miaopost.com"
			msg := "关联成功，发布后即可随时通过“我的发布”菜单进行编辑、删除和置顶操作，<a href=\"" + url + "/info/create" + "\">去发布</a>。此过程中我们只会使用您的微信昵称用于标示，不会获取您的任何其他信息。如介意使用微信昵称，可以<a href=\"" + url + "/user/edit" + "\">修改标示</a>。"

			return replyText(msgBody, msg)
		}

		return nil
	}

	return nil
}

func replyText(msgBody *wechat.MsgBody, text string) *wechat.MsgBody {
	replyBody := &wechat.MsgBody{}
	replyBody.ToUserName = msgBody.FromUserName
	replyBody.FromUserName = msgBody.ToUserName
	replyBody.CreateTime = time.Since(time.Now())
	replyBody.MsgType = "text"
	replyBody.Content = text

	return replyBody
}

func scanLogin(msgBody *wechat.MsgBody) (ekey string) {
	ekey = msgBody.EventKey
	cache := help.Cache
	expire := 1800

	user, err := models.GetUserByOpenid(msgBody.FromUserName)
	if err == nil {
		b, _ := json.Marshal(user)
		cache.Put(msgBody.EventKey, b, time.Duration(expire)*time.Second)
		return
	}

	userinfo := wechat.GetWxUserinfo(msgBody.FromUserName, "")
	if v, ok := userinfo["nickname"]; ok {
		u := models.User{
			Openid:   msgBody.FromUserName,
			Nickname: v.(string),
			Avatar:   userinfo["headimgurl"].(string),
		}
		models.CreateUser(&u)
		b, _ := json.Marshal(u)
		cache.Put(msgBody.EventKey, b, time.Duration(expire)*time.Second)
	}

	return
}
