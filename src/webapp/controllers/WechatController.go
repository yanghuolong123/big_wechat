package controllers

import (
	"crypto/sha1"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strings"
	"time"
	"yhl/help"
)

const Token = "feichangjuzu123456"

type MsgBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	Content      string
	MsgId        int
}

type WechatController struct {
	help.BaseController
}

func (this *WechatController) Index() {
	timestamp, nonce, signatureIn := this.GetString("timestamp"), this.GetString("nonce"), this.GetString("signature")
	echostr := strings.TrimSpace(this.GetString("echostr"))

	if check(timestamp, nonce, signatureIn) {
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

		msgBody := &MsgBody{}
		xml.Unmarshal(body, msgBody)
		fmt.Println(msgBody)

		replyBody := &MsgBody{}
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

func check(timestamp, nonce, signatureIn string) bool {
	sl := []string{Token, timestamp, nonce}
	sort.Strings(sl)
	s := sha1.New()
	io.WriteString(s, strings.Join(sl, ""))

	signatureGen := fmt.Sprintf("%x", s.Sum(nil))

	return signatureGen == signatureIn
}
