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

// 预订单生成
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

// 公众号微信支付
func (this *PayController) Confirm() {
	uid := 0
	user := this.GetSession("user")
	if user != nil {
		uid = user.(*models.User).Id
	}
	productId, _ := this.GetInt("product_id")
	amount, _ := this.GetFloat("amount")
	info_id, _ := this.GetInt("info_id")
	otype, _ := this.GetInt("type")
	msg := this.GetString("msg")

	order_type := int(otype)
	order := new(models.Order)
	var err error
	if order_type == 1 {
		// 赞赏支付订单
		order, err = models.GenAdmireOrder(productId, uid, amount)
	} else if order_type == 2 {
		// 红包信息支付发布
		order, err = models.GenRewardOrder(productId, uid, amount)
	}

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
		return
	}

	sdk := wxpay.JsPaySdk(wxRes.Prepay_id)
	help.Log("wxpay", sdk)
	this.Data["sdk"] = sdk

	this.Data["user"] = user
	this.Data["info_id"] = info_id
	this.Data["amount"] = amount
	this.Data["msg"] = msg

	this.Layout = "layout/main.tpl"
	this.TplName = "pay/confirm.tpl"
}

// pc微信扫码支付
func (this *PayController) WxScan() {
	uid := 0
	user := this.GetSession("user")
	if user != nil {
		uid = user.(*models.User).Id
	}
	productId, _ := this.GetInt("product_id")
	amount, _ := this.GetFloat("amount")
	otype, _ := this.GetInt("type")

	order_type := int(otype)
	order := new(models.Order)
	var err error
	if order_type == 1 {
		// 赞赏支付订单
		order, err = models.GenAdmireOrder(productId, uid, amount)
	} else if order_type == 2 {
		// 红包信息支付发布
		order, err = models.GenRewardOrder(productId, uid, amount)
	}

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

// 生成二维码
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

// 支付后监听微信后台返回
func (this *PayController) Notify() {
	body := this.Ctx.Input.RequestBody
	help.Log("wxpay", "========= requestBody:"+string(body))

	var notifyReq wxpay.WXPayNotifyReq
	err := xml.Unmarshal(body, &notifyReq)
	if err != nil {
		help.Log("wxpay", err.Error())
		this.SendXml(wxpay.WXPayNotifyResp{Return_code: "FAIL", Return_msg: "failed to unmarshal xml"})
	}

	if notifyReq.Return_code == "SUCCESS" && notifyReq.Result_code == "SUCCESS" {
		order := models.GetOrderByOrderno(notifyReq.Out_trade_no)
		if order.Status < 1 {
			order.Pay_time = time.Now()
			order.Status = 1
			order.Transaction_id = notifyReq.Transaction_id
			models.UpdateOrder(order)

			ua, _ := models.GetUserAccountByUid(order.Uid)
			if order.Type == 1 {
				// 赞赏支付,赞赏对象个人的账号金额增加, 自己的账号余额清0

				msg, _ := models.GetInfoMessageById(order.Product_id)
				amount := int(order.Amount*100) + int(ua.Amount*100)
				models.AccountChange(float64(amount)/100, msg.Uid, order.Type, order.Product_id, "获得赞赏")

				models.AccountChange(-ua.Amount, order.Uid, order.Type, order.Product_id, "支付赞赏")

				go models.AdmireWxTip(order.Product_id, order.Amount, this.Ctx)
			} else if order.Type == 2 {
				models.AccountChange(-ua.Amount, order.Uid, order.Type, order.Product_id, "支付发布红包信息")
				// 红包发布信息
				go models.GenBathInfoRewardByInfoId(order.Product_id)
			}
		}

		help.Log("wxpay", "============== weixin pay success! ===============")
		this.SendXml(wxpay.WXPayNotifyResp{Return_code: "SUCCESS", Return_msg: "OK!"})
	}

	/*
		notifySign := notifyReq.Sign
		notifyReq.Sign = ""

		m := help.StructToMap(notifyReq)
		signStr := wxpay.Sign(m)

		if notifySign != signStr {
			help.Log("wxpay", "verify sign failed || signStr:"+signStr+" reqSign:"+notifySign)
			this.SendXml(wxpay.WXPayNotifyResp{Return_code: "FAIL", Return_msg: "failed to verify sign, please retry"})
		}
	*/

	help.Log("wxpay", "未知错误")
	this.SendXml(wxpay.WXPayNotifyResp{Return_code: "FAIL", Return_msg: "unknow error"})
}

// pc扫码支付后监听支付结果
func (this *PayController) Check() {
	orderNo := this.GetString("order_no")
	order := models.GetOrderByOrderno(orderNo)
	if order.Status == 1 {
		this.SendRes(0, "success", order)
	}

	this.SendRes(-1, "no pay complete", nil)
}

// 用户申请提现
func (this *PayController) Withdraw() {
	amount, _ := this.GetFloat("amount")
	if amount < 0 {
		this.SendRes(-1, "错误金额", nil)
	}

	u := this.GetSession("user")
	user := u.(*models.User)

	order, _ := models.GenWithdrawOrder(user.Id, amount)
	//partnerTradeNo := help.GenOrderNo()
	certDir := beego.AppConfig.String("wx.pay.cert.dir")
	certPath := help.GetAPPRootPath() + certDir
	//remark := "用户提现"
	help.Log("wxpay", "certPath:"+certPath)

	res := wxpay.PayToUser(amount, user.Openid, order.Orderno, order.Remark, this.Ctx.Input.IP(), certPath)
	help.Log("wxpay", res)

	if res.ReturnCode == "SUCCESS" && res.ResultCode == "SUCCESS" {
		help.Log("wxpay", user.Nickname+" 提现"+fmt.Sprintf("%v", amount)+"元成功!")
		order = models.GetOrderByOrderno(res.PartnerTradeNo)
		if order.Status < 1 {
			order.Pay_time = time.Now()
			order.Status = 1
			order.Transaction_id = res.PaymentNo
			models.UpdateOrder(order)

			models.AccountChange(-amount, user.Id, order.Type, 0, order.Remark)
		}
		this.SendRes(0, "success", nil)
	}

	help.Log("wxpay", "code:"+res.ErrCodeDes+" msg:"+res.ReturnMsg)
	this.SendRes(-1, "code:"+res.ErrCodeDes+" msg:"+res.ReturnMsg, nil)
}

// 获取用户余额
func (this *PayController) Balance() {
	pay_way, _ := this.GetInt("type")
	amount, _ := this.GetFloat("amount")
	toUid, _ := this.GetInt("toUid")
	product_id, _ := this.GetInt("product_id")

	payToUid := int(toUid)
	productId := int(product_id)

	u := this.GetSession("user")
	user := u.(*models.User)

	ua, err := models.GetUserAccountByUid(user.Id)
	if err != nil {
		this.SendRes(-1, "obtain balance failed", nil)
	}

	if amount <= ua.Amount {
		adtype := int(pay_way)
		remark := ""
		if adtype == 1 {
			remark = "支付赞赏"
		} else if adtype == 2 {
			remark = "支付发布信息红包"
		}
		if models.AccountChange(-amount, user.Id, adtype, productId, remark) {
			if adtype == 1 {
				remark = "获得赞赏"
				models.AccountChange(amount, payToUid, adtype, productId, remark)
			} else if adtype == 2 {
				go models.GenBathInfoRewardByInfoId(productId)
			}

		}

		this.SendRes(0, "success", 0)
	}

	this.SendRes(1, "success", ua)
}
