package models

import (
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
