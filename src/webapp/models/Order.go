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

func CreateUnlockOrder(o *Order) bool {
	o.Create_time = time.Now()

	i, err := orm.NewOrm().Insert(o)
	if err != nil {
		help.Log("error", err.Error())
	}

	return i > 0
}
