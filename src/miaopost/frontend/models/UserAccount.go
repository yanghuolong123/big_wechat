package models

import (
	"github.com/astaxie/beego/orm"
	"yhl/help"
)

func init() {
	orm.RegisterModelWithPrefix("tbl", new(UserAccount))
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
	//	ua = &UserAccount{Uid: uid}
	ua.Uid = uid
	err = orm.NewOrm().Read(ua)
	help.Error(err)

	return
}
