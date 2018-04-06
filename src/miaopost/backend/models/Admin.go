package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"time"
	"yhl/help"
)

func init() {
	orm.RegisterModelWithPrefix("tbl_", new(Administrator))
}

type Administrator struct {
	Id             int
	Username       string
	Password       string
	Last_logintime time.Time
}

func GetAdminById(id int) (*Administrator, error) {
	admin := &Administrator{Id: id}
	err := orm.NewOrm().Read(admin)
	help.Error(err)

	return admin, err
}

func Login(username, password string) (*Administrator, error) {
	admin := &Administrator{}
	admin.Username = username
	admin.Password = help.Md5(password)
	err := orm.NewOrm().Read(admin, "username", "password")
	if err != nil {
		return nil, errors.New("帐号或密码错误")
	}
	admin.Last_logintime = time.Now()
	orm.NewOrm().Update(admin, "last_logintime")

	return admin, nil
}
