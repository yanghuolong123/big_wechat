package controllers

import (
	"fmt"
	"time"
	"webapp/models"
	"yhl/help"
)

type PayController struct {
	help.BaseController
}

func (this *PayController) WxScan() {
	productId, _ := this.GetInt("product_id")

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

	this.SendRes(0, "success", nil)
}
