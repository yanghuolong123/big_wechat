package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

func init() {
	orm.RegisterModelWithPrefix("tbl_", new(UnlockGroup))
}

type UnlockGroup struct {
	Id         int
	Uid        int
	Gid        int
	Createtime time.Time
}

func CreateUnlockGroup(ug *UnlockGroup) bool {
	ug.Createtime = time.Now()

	i, _ := orm.NewOrm().Insert(ug)

	return i > 0
}

func GetUnlockGroupByUid(uid int) (ugs []UnlockGroup) {
	orm.NewOrm().QueryTable("tbl_unlock_group").Filter("uid", uid).All(&ugs)
	return
}

func IsUnlock(uid, gid int) bool {
	ug := UnlockGroup{Uid: uid, Gid: gid}
	err := orm.NewOrm().Read(&ug, "uid", "gid")
	if err != nil {
		return false
	}

	return true
}
