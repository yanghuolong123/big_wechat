package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"yhl/help"
)

func init() {
	orm.RegisterModelWithPrefix("tbl_", new(User))
}

type User struct {
	Id         int
	Openid     string
	Username   string
	Email      string
	Password   string
	Nickname   string
	Mobile     string
	Avatar     string
	Level      int
	Status     int
	Createtime time.Time
}

func GetUserByOpenid(openid string) (user *User, err error) {
	o := orm.NewOrm()
	user = &User{Openid: openid}
	err = o.Read(user, "Openid")
	help.Error(err)

	return
}

func CreateUser(user *User) int {
	user.Createtime = time.Now()
	i, _ := orm.NewOrm().Insert(user)

	return int(i)
}

func UpdateUser(user *User) error {
	_, err := orm.NewOrm().Update(user)

	return err
}
