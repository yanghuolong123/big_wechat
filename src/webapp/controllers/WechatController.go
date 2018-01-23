package controllers

import (
	"encoding/json"
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
	if msgBody.MsgType == "event" && msgBody.Event == "subscribe" {
		if strings.Contains(msgBody.EventKey, "login_") {
			msgBody.EventKey = strings.TrimLeft(msgBody.EventKey, "qrscene_")
			key := scanLogin(msgBody)
			url := "http://www.addwechat.com/loginByKey?key=" + key

			return replyText(msgBody, "登陆成功! <a href=\""+url+"\">进入AddWechat</a>")
		}

		subscribe(msgBody)
		return nil
	}
	if msgBody.MsgType == "event" && msgBody.Event == "SCAN" {
		if strings.HasPrefix(msgBody.EventKey, "login_") {
			key := scanLogin(msgBody)
			url := "http://www.addwechat.com/loginByKey?key=" + key

			return replyText(msgBody, "登陆成功! <a href=\""+url+"\">进入AddWechat</a>")
		}

		return nil
	}

	return nil
}

func subscribe(msgBody *wechat.MsgBody) {
}

func scanLogin(msgBody *wechat.MsgBody) (ekey string) {
	ekey = msgBody.EventKey
	cache := help.Cache
	expire := 18000

	user, err := models.GetUserByOpenid(msgBody.FromUserName)
	if err == nil {
		b, _ := json.Marshal(*user)
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

func replyText(msgBody *wechat.MsgBody, text string) *wechat.MsgBody {
	replyBody := &wechat.MsgBody{}
	replyBody.ToUserName = msgBody.FromUserName
	replyBody.FromUserName = msgBody.ToUserName
	replyBody.CreateTime = time.Since(time.Now())
	replyBody.MsgType = "text"
	replyBody.Content = text

	return replyBody
}
