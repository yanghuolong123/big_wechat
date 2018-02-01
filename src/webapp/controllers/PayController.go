package controllers

import (
	"encoding/xml"
	"fmt"
	qrcode "github.com/skip2/go-qrcode"
	"net/url"
	"time"
	"webapp/models"
	"yhl/help"
	"yhl/wechat/wxpay"
)

type PayController struct {
	help.BaseController
}

func (this *PayController) WxScan() {
	productId, _ := this.GetInt("product_id")

	group := models.GetGroupById(productId)
	bodyName := "解锁 " + group.Name
	if group.Name == "" {
		bodyName += group.En_name
	}

	order := &models.Order{}
	order.Type = 1
	order.Product_id = productId
	order.Orderno = time.Now().Format(help.DatetimeNumFormat) + fmt.Sprintf("%d", help.RandNum(10000, 99999))
	order.Amount = 0.01
	order.Pay_type = 1
	order.Uid = 2
	if !models.CreateOrder(order) {
		this.SendRes(-1, "创建订单失败", nil)
	}

	orderReq := new(wxpay.UnifyOrderReq)
	orderReq.Body = bodyName
	orderReq.Out_trade_no = order.Orderno
	orderReq.Total_fee = int(order.Amount * 100)
	orderReq.Notify_url = "addwechat.feichangjuzu.com/pay/notify"
	orderReq.Trade_type = "NATIVE"
	orderReq.Product_id = productId
	orderReq.Time_start = time.Now().Format(help.DatetimeNumFormat)
	orderReq.Time_expire = time.Now().Add(time.Duration(600 * time.Second)).Format(help.DatetimeNumFormat)
	orderReq.Spbill_create_ip = help.ClientIp

	wxRes := wxpay.UnifiedOrder(orderReq)
	help.Log("wxpay", wxRes)
	if wxRes.Return_code == "FAIL" {
		help.Log("error", wxRes)
	}

	v := url.Values{}
	v.Add("url", wxRes.Code_url)

	m := map[string]string{}
	m["orderno"] = order.Orderno
	m["qrurl"] = "/pay/qrcode?" + v.Encode()

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
		strResp := wxpay.ToWxRespXmlStr("FAIL", "failed to unmarshal xml")
		this.Ctx.Output.XML(strResp, true)
		return
	}

	notifySign := notifyReq.Sign
	notifyReq.Sign = ""

	m := help.StructToMap(notifyReq)
	signStr := wxpay.Sign(m)
	help.Log("wxpay", "signStr:"+signStr+" reqSign:"+notifySign)

	if notifySign != signStr {
		help.Log("error", "sign error:: signStr:"+signStr+" reqSign:"+notifySign)
		strResp := wxpay.ToWxRespXmlStr("FAIL", "failed to verify sign, please retry!")
		this.Ctx.Output.XML(strResp, true)
		return
	}

	order := models.GetOrderByOrderno(notifyReq.Out_trade_no)
	order.Pay_time = time.Now()
	order.Status = 1
	order.Transaction_id = notifyReq.Transaction_id
	models.UpdateOrder(order)

	strResp := wxpay.ToWxRespXmlStr("SUCCESS", "OK")
	this.Ctx.Output.XML(strResp, true)
}
