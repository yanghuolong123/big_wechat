package controllers

import (
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
	if !models.CreateUnlockOrder(order) {
		this.SendRes(-1, "创建订单失败", nil)
	}

	orderReq := new(wxpay.UnifyOrderReq)
	orderReq.Body = bodyName
	orderReq.Out_trade_no = order.Orderno
	orderReq.Total_fee = 1
	orderReq.Notify_url = "pay.feichangjuzu.com/pay/"
	orderReq.Trade_type = "NATIVE"
	orderReq.Product_id = productId
	orderReq.Time_start = time.Now().Format(help.DatetimeNumFormat)
	orderReq.Time_expire = time.Now().Add(time.Duration(600 * time.Second)).Format(help.DatetimeNumFormat)
	orderReq.Spbill_create_ip = help.ClientIp

	wxRes := wxpay.UnifiedOrder(orderReq)
	help.Log("wxpay", wxRes)
	if wxRes.Return_code == "FAIL" {
		help.Log("wxpay", "error==================================")
	}

	if wxRes.Code_url != "" {
		fmt.Println("====== code_url:", wxRes.Code_url)
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
