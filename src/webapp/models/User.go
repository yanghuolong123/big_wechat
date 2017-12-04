package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"yhl/help"
)

func init() {
	orm.RegisterModelWithPrefix("tbl_", new(User))
}

type User struct {
	Id       int
	Username string
	Email    string
	Password string
}

func GetById(id int) (user *User, err error) {
	o := orm.NewOrm()
	user = &User{Id: id}
	err = o.Read(user)

	return
}

func Login(username, password string) (*User, error) {
	password = help.Md5(password)

	o := orm.NewOrm()
	user := new(User)
	user.Username = username
	user.Password = password
	err := o.Read(user, "username", "password")
	if err != nil {
		return nil, errors.New("帐号或密码有误！")
	}

	return user, nil
}
