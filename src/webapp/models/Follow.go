package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

func init() {
	orm.RegisterModelWithPrefix("tbl_", new(Follow))
}

type Follow struct {
	Id         int
	Uid        int
	Gid        int
	Createtime time.Time
}

func CreateFollow(uid, gid int) bool {
	o := orm.NewOrm()
	f := new(Follow)
	f.Uid = uid
	f.Gid = gid
	f.Createtime = time.Now()

	i, _ := o.Insert(f)
	return i > 0
}

func GetFollowByUid(uid int) (gids []int) {
	var flist []Follow
	orm.NewOrm().QueryTable("tbl_follow").Filter("uid", uid).All(&flist)

	for _, f := range flist {
		gids = append(gids, f.Gid)
	}
	return
}
