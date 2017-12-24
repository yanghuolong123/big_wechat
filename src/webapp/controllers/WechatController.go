package controllers

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
	"webapp/models"
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
			help.Log.Info("===== 接入成功 ======")
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

		listen(msgBody)

		/*	replyBody := &wechat.MsgBody{}
			replyBody.ToUserName = msgBody.FromUserName
			replyBody.FromUserName = msgBody.ToUserName
			replyBody.CreateTime = time.Since(time.Now())
			replyBody.MsgType = "text"
			replyBody.Content = "回复:" + msgBody.Content

			this.Data["xml"] = replyBody
			this.ServeXML()
		*/
	}

	this.StopRun()
}

func listen(msgBody *wechat.MsgBody) {
	if msgBody.MsgType == "event" && msgBody.Event == "subscribe" {
		subscribe(msgBody)
		return
	}
	if msgBody.MsgType == "event" && msgBody.Event == "SCAN" {
		if strings.HasPrefix(msgBody.EventKey, "login_") {
			scanLogin(msgBody)
			return
		}
	}
}

func subscribe(msgBody *wechat.MsgBody) {
}

func scanLogin(msgBody *wechat.MsgBody) {
	cache := help.Cache
	expire := 18000

	user, err := models.GetUserByOpenid(msgBody.FromUserName)
	fmt.Println("============== user, err:", user, err)
	if err == nil {
		cache.Put(msgBody.EventKey, user, time.Duration(expire)*time.Second)
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
		e := cache.Put(msgBody.EventKey, u, time.Duration(expire)*time.Second)
		fmt.Println("=========== e:", e)
	}

}
