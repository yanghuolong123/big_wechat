package models

import (
	"github.com/astaxie/beego/context"
	"yhl/help"
	"yhl/wechat"
)

func ReplyWxTip(id int, ctx *context.Context) {
	m, err := GetInfoMessageById(id)
	if err != nil {
		return
	}
	p, err := GetInfoMessageById(m.Pid)
	if err != nil {
		return
	}

	user, _ := GetUserById(p.Uid)
	u, _ := GetUserById(m.Uid)
	viewUrl := ctx.Input.Site() + "/info/view?id=" + help.ToStr(p.Info_id)
	msg := u.Nickname + " 回复了您的留言:  " + help.HtmlToStr(m.Content) + "\n" + viewUrl
	wechat.SendTextMsg(user.Openid, msg)
}

func MessageWxTip(id int, ctx *context.Context) {
	m, err := GetInfoMessageById(id)
	if err != nil {
		return
	}

	info, _ := GetInfoById(m.Info_id)

	user, _ := GetUserById(info.Uid)
	u, _ := GetUserById(m.Uid)
	viewUrl := ctx.Input.Site() + "/info/view?id=" + help.ToStr(m.Info_id)
	msg := u.Nickname + " 给您留言:  " + help.HtmlToStr(m.Content) + "\n" + viewUrl
	wechat.SendTextMsg(user.Openid, msg)
}

func AdmireWxTip(mid int, amount float64, ctx *context.Context) {
	m, err := GetInfoMessageById(mid)
	if err != nil {
		return
	}

	user, _ := GetUserById(m.Uid)
	viewUrl := ctx.Input.Site() + "/user/account"
	msg := " 您收到赞赏:  " + help.ToStr(amount) + "元 " + viewUrl
	wechat.SendTextMsg(user.Openid, msg)
}
