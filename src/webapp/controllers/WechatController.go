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

		replyBody := &wechat.MsgBody{}
		replyBody.ToUserName = msgBody.FromUserName
		replyBody.FromUserName = msgBody.ToUserName
		replyBody.CreateTime = time.Since(time.Now())
		replyBody.MsgType = "text"
		replyBody.Content = "回复:" + msgBody.Content

		this.Data["xml"] = replyBody
		this.ServeXML()
	}

	this.StopRun()
}
