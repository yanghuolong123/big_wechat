package controllers

import (
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego"
	qrcode "github.com/skip2/go-qrcode"
	"miaopost/frontend/models"
	"net/url"
	"time"
	"yhl/help"
	"yhl/wechat/wxpay"
)

type PayController struct {
	help.BaseController
}

func prePayOrder(order *models.Order, trade_type, openid string) (resp wxpay.UnifyOrderResp, err error) {
	notify_domain := beego.AppConfig.String("wx.pay.notify.domain")

	orderReq := new(wxpay.UnifyOrderReq)
	orderReq.Body = order.Remark
	orderReq.Out_trade_no = order.Orderno
	orderReq.Total_fee = int(order.Amount * 100)
	orderReq.Notify_url = notify_domain + "/pay/notify"
	//orderReq.Trade_type = "NATIVE"
	orderReq.Trade_type = trade_type
	if orderReq.Trade_type == "JSAPI" {
		orderReq.Openid = openid
	}
	orderReq.Product_id = order.Product_id
	orderReq.Time_start = time.Now().Format(help.DatetimeNumFormat)
	orderReq.Time_expire = time.Now().Add(time.Duration(600 * time.Second)).Format(help.DatetimeNumFormat)
	orderReq.Spbill_create_ip = order.Ip

	wxRes := wxpay.UnifiedOrder(orderReq)
	help.Log("wxpay", wxRes)

	return wxRes, nil
}

func (this *PayController) Confirm() {
	uid := 0
	user := this.GetSession("user")
	if user != nil {
		uid = user.(*models.User).Id
	}
	productId, _ := this.GetInt("product_id")
	amount, _ := this.GetFloat("amount")
	info_id, _ := this.GetInt("info_id")

	order, err := models.GenAdmireOrder(productId, uid, amount)
	order.Ip = this.Ctx.Input.IP()
	if err != nil {
		this.Redirect(err.Error(), 302)
	}

	//url := this.Ctx.Input.Site() + this.Ctx.Input.URI()
	//openid := wxpay.GetOpenId(this.Ctx, url)
	openid := user.(*models.User).Openid
	if openid == "" {
		this.Redirect("/tips?msg=openid获取失败", 302)
		return
	}

	wxRes, err := prePayOrder(order, "JSAPI", openid)
	if err != nil {
		this.Redirect("/tips?msg=预订单生成失败", 302)
	}

	sdk := wxpay.JsPaySdk(wxRes.Prepay_id)
	help.Log("wxpay", sdk)
	this.Data["sdk"] = sdk

	this.Data["user"] = user
	this.Data["info_id"] = info_id
	this.Data["amount"] = amount

	this.Layout = "layout/main.tpl"
	this.TplName = "pay/confirm.tpl"
}

func (this *PayController) WxScan() {
	uid := 0
	user := this.GetSession("user")
	if user != nil {
		uid = user.(*models.User).Id
	}
	productId, _ := this.GetInt("product_id")
	amount, _ := this.GetFloat("amount")

	order, err := models.GenAdmireOrder(productId, uid, amount)
	order.Ip = this.Ctx.Input.IP()
	if err != nil {
		this.SendRes(-1, err.Error(), nil)
	}

	wxRes, err := prePayOrder(order, "NATIVE", "")
	if err != nil {
		this.SendRes(-1, err.Error(), nil)
	}

	if wxRes.Return_code == "FAIL" {
		help.Log("error", wxRes)
	}

	v := url.Values{}
	v.Add("url", wxRes.Code_url)

	m := map[string]string{}
	m["qrurl"] = "/pay/qrcode?" + v.Encode()
	m["prepay_id"] = wxRes.Prepay_id
	m["order_no"] = order.Orderno

	this.SendRes(0, "success", m)
}

func (this *PayController) Qrcode() {
	url := this.GetString("url")
	qr, err := qrcode.New(url, qrcode.High)
	if err != nil {
		this.SendRes(-1, "failed", nil)
	}

	b, err := qr.PNG(500)
	if err != nil {
		this.SendRes(-1, "failed", nil)
	}

	png := string(b)

	this.Ctx.Output.Header("Content-Type", "image/png")
	this.Ctx.Output.Header("Content-Length", fmt.Sprintf("%d", len(png)))
	this.Ctx.WriteString(png)
}

func (this *PayController) Notify() {
	body := this.Ctx.Input.RequestBody
	help.Log("wxpay", "========= requestBody:"+string(body))

	var notifyReq wxpay.WXPayNotifyReq
	err := xml.Unmarshal(body, &notifyReq)
	if err != nil {
		help.Log("error", err.Error())
		this.SendXml(wxpay.WXPayNotifyResp{Return_code: "FAIL", Return_msg: "failed to unmarshal xml"})
	}

	notifySign := notifyReq.Sign
	notifyReq.Sign = ""

	m := help.StructToMap(notifyReq)
	signStr := wxpay.Sign(m)

	if notifySign != signStr {
		help.Log("error", "verify sign failed || signStr:"+signStr+" reqSign:"+notifySign)
		this.SendXml(wxpay.WXPayNotifyResp{Return_code: "FAIL", Return_msg: "failed to verify sign, please retry"})
	}

	order := models.GetOrderByOrderno(notifyReq.Out_trade_no)
	if order.Status < 1 {
		order.Pay_time = time.Now()
		order.Status = 1
		order.Transaction_id = notifyReq.Transaction_id
		models.UpdateOrder(order)
		// 个人账号金额增加
		// todo .....
	}
	help.Log("wxpay", "============== weixin pay success! ===============")

	this.SendXml(wxpay.WXPayNotifyResp{Return_code: "SUCCESS", Return_msg: "OK!"})
}

func (this *PayController) Check() {
	orderNo := this.GetString("order_no")
	order := models.GetOrderByOrderno(orderNo)
	if order.Status == 1 {
		this.SendRes(0, "success", nil)
	}

	this.SendRes(-1, "no pay complete", nil)
}
