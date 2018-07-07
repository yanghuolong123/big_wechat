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
	if msgBody.MsgType == "event" && (msgBody.Event == "subscribe" || msgBody.Event == "SCAN") {
		if msgBody.Event == "subscribe" {
			content := `欢迎关注秒Po~ 使用秒Po发布时，我们会自动为您创建专属通道，无需注册即可在发布后进行便捷的编辑、删除和置顶操作。
专属通道除了使用您的微信昵称用于标示外，不会获取任何其他信息。如介意使用微信昵称，可以在进入秒po后修改标示。`
			wechat.SendTextMsg(msgBody.FromUserName, content)
		}

		//		if strings.Contains(msgBody.EventKey, "create") {
		//
		//			msg := "关注秒Po的同时我们即为您创建专属通道，发布后即可进行便捷的修改、删除和免费置顶操作"
		//			return replyText(msgBody, msg)
		//		}

		if strings.Contains(msgBody.EventKey, "login_") {
			msgBody.EventKey = strings.TrimLeft(msgBody.EventKey, "qrscene_")
			key := scanLogin(msgBody)
			_ = key
			//url := "http://www.miaopost.com"
			//			msg := "通道创建成功，通过电脑或公众号发布后都可以进行修改、删除和置顶。专属通道除了使用您的微信昵称用于标示外，不会获取任何其他信息。如介意使用微信昵称，可以 <a href=\"" + url + "/user/edit" + "\">修改标示</a>。"
			msg := "已登录！此时通过电脑或公众号发布后都可以进行编辑、删除和免费置顶操作。"

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
			Rid:      1,
			Nickname: v.(string),
			Avatar:   userinfo["headimgurl"].(string),
		}
		models.CreateUser(&u)
		b, _ := json.Marshal(u)
		cache.Put(msgBody.EventKey, b, time.Duration(expire)*time.Second)
	}

	return
}
