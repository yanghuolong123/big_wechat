package controllers

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
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
