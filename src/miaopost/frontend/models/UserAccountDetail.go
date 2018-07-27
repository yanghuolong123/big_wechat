package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"yhl/help"
)

func init() {
	orm.RegisterModelWithPrefix("tbl_", new(UserAccountDetail))
}

type UserAccountDetail struct {
	Id          int
	Uid         int
	Type        int
	Product_id  int
	Amount      float64
	Remark      string
	Create_time time.Time
}

func CreateUserAccountDetail(ucd *UserAccountDetail) int {
	ucd.Create_time = time.Now()
	i, err := orm.NewOrm().Insert(ucd)
	help.Error(err)

	return int(i)
}

func GetUserAccountDetailByUid(uid int) (ucds []UserAccountDetail) {
	_, err := orm.NewOrm().QueryTable("tbl_user_account_detail").Filter("uid", uid).OrderBy("-create_time").All(&ucds)
	help.Error(err)

	return
}
