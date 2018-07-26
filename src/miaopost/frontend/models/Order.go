package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
	"yhl/help"
)

func init() {
	orm.RegisterModelWithPrefix("tbl_", new(Order))
}

type Order struct {
	Id             int
	Type           int
	Product_id     int
	Orderno        string
	Uid            int
	Pay_type       int
	Status         int
	Amount         float64
	Transaction_id string
	Create_time    time.Time
	Pay_time       time.Time
	Remark         string
	Ip             string `orm:"-"`
}

func CreateOrder(o *Order) bool {
	o.Create_time = time.Now()

	i, err := orm.NewOrm().Insert(o)
	if err != nil {
		help.Log("error", err.Error())
	}

	return i > 0
}

func UpdateOrder(o *Order) bool {
	i, err := orm.NewOrm().Update(o)
	if err != nil {
		help.Log("error", err.Error())
	}

	return i > 0
}

func GetOrderByOrderno(orderno string) (o *Order) {
	o = new(Order)
	o.Orderno = orderno
	err := orm.NewOrm().Read(o, "orderno")
	if err != nil {
		help.Log("error", err.Error())
	}

	return
}

func GenAdmireOrder(productId, uid int, amount float64) (*Order, error) {
	order := &Order{}
	order.Type = 1
	order.Product_id = productId
	order.Orderno = time.Now().Format(help.DatetimeNumFormat) + fmt.Sprintf("%d", help.RandNum(10000, 99999))
	order.Amount = amount
	order.Pay_type = 1
	order.Uid = uid
	order.Remark = "支付赞赏"
	if CreateOrder(order) {
		return order, nil
	}

	return order, errors.New("订单创建失败")
}

func GenWithdrawOrder(uid int, amount float64) (*Order, error) {
	order := &Order{}
	order.Type = 3
	order.Orderno = help.GenOrderNo()
	order.Amount = amount
	order.Pay_type = 1
	order.Uid = uid
	order.Remark = "用户提现"
	if CreateOrder(order) {
		return order, nil
	}

	return order, errors.New("订单创建失败")
}

func GenRewardOrder(productId, uid int, amount float64) (*Order, error) {
	order := &Order{}
	order.Type = 2
	order.Product_id = productId
	order.Orderno = help.GenOrderNo()
	order.Amount = amount
	order.Pay_type = 1
	order.Uid = uid
	order.Remark = "支付发布信息红包"
	if CreateOrder(order) {
		return order, nil
	}

	return order, errors.New("订单创建失败")
}
