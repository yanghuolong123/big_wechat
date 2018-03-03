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

func CreateUnlockGroup(uid, gid int) bool {
	l := UnlockGroup{}
	l.Uid = uid
	l.Gid = gid
	l.Createtime = time.Now()

	i, _ := orm.NewOrm().Insert(l)
	return i > 0
}

//func CreateUnlockGroup(ug *UnlockGroup) bool {
//	ug.Createtime = time.Now()
//
//	i, _ := orm.NewOrm().Insert(ug)
//
//	return i > 0
//}

func GetUnlockGroupByUid(uid int) (groups []Group) {
	var ugs []UnlockGroup
	orm.NewOrm().QueryTable("tbl_unlock_group").Filter("uid", uid).All(&ugs, "Gid")
	var gids []int
	for _, ug := range ugs {
		gids = append(gids, ug.Gid)
	}
	if len(gids) > 0 {
		orm.NewOrm().QueryTable("tbl_group").Filter("id__in", gids).All(&groups)
	}
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
