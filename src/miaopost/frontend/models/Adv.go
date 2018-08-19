package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"yhl/help"
)

func init() {
	orm.RegisterModelWithPrefix("tbl_", new(Adv))
}

type Adv struct {
	Id            int
	Type          int
	Uid           int
	Merch_name    string
	Contact       string
	Tag           string
	Logo          string
	Content       string
	Target        string
	Photos        string
	Region_id     int
	Pos           int
	Display_times int
	Day_limit     int
	Amount        float64
	Total_amount  float64
	Recom_code    string
	Status        int
	Display_count int
	Create_time   time.Time
}

func CreateAdv(av *Adv) error {
	av.Create_time = time.Now()
	_, err := orm.NewOrm().Insert(av)
	help.Error(err)

	return err
}

func UpdateAdv(av *Adv) error {
	_, err := orm.NewOrm().Update(av)
	help.Error(err)

	return err
}

func GetAdvById(id int) (*Adv, error) {
	adv := &Adv{Id: id}
	err := orm.NewOrm().Read(adv)
	help.Error(err)

	return adv, err
}

func BanAdvById(id int) bool {
	i, err := orm.NewOrm().QueryTable("tbl_adv").Filter("id", id).Update(orm.Params{"status": -1})
	help.Error(err)

	return int(i) > 0
}

func EnableAdvById(id int) bool {
	i, err := orm.NewOrm().QueryTable("tbl_adv").Filter("id", id).Update(orm.Params{"status": 1})
	help.Error(err)

	return int(i) > 0
}

func GetAdvByUid(uid int) []*Adv {
	var advs []*Adv
	_, err := orm.NewOrm().QueryTable("tbl_adv").Filter("uid", uid).All(&advs)
	help.Error(err)

	return advs
}

func GetAdvByTypeAndRegion(t, r int) []*Adv {
	var advs []*Adv
	_, err := orm.NewOrm().QueryTable("tbl_adv").Filter("type", t).Filter("region_id", r).Filter("status", 1).All(&advs)
	help.Error(err)

	return advs
}
