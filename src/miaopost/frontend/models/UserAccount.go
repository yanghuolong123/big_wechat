package models

import (
	"github.com/astaxie/beego/orm"
	"yhl/help"
)

func init() {
	orm.RegisterModelWithPrefix("tbl_", new(UserAccount))
}

type UserAccount struct {
	Id     int
	Uid    int
	Amount float64
}

func CreateUserAccount(ua *UserAccount) int {
	i, err := orm.NewOrm().Insert(ua)
	help.Error(err)

	return int(i)
}

func GetUserAccountByUid(uid int) (ua *UserAccount, err error) {
	ua = &UserAccount{Uid: uid}
	err = orm.NewOrm().Read(ua, "Uid")
	help.Error(err)

	return
}

func IncUserAccount(uid int, amount float64) bool {
	ua, err := GetUserAccountByUid(uid)
	if err != nil {
		ua.Uid = uid
		ua.Amount = amount
		CreateUserAccount(ua)
		return true
	}

	ua.Amount += amount
	i, err := orm.NewOrm().Update(ua)
	help.Error(err)

	return i > 0
}

func AccountChange(amount float64, uid, adtype, order_id int, remark string) bool {
	uad := new(UserAccountDetail)
	uad.Uid = uid
	uad.Amount = amount
	uad.Type = adtype
	uad.Product_id = order_id
	uad.Remark = remark
	i := CreateUserAccountDetail(uad)
	if i > 0 {
		return IncUserAccount(uid, amount)
	}

	return false
}
