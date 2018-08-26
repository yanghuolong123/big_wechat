package models

import (
	"github.com/astaxie/beego/orm"
	"math/rand"
	"strings"
	"time"
	"yhl/help"
)

func init() {
	orm.RegisterModelWithPrefix("tbl_", new(Adv))
}

type Adv struct {
	Id              int
	Type            int
	Uid             int
	Merch_name      string
	Contact         string
	Tag             string
	Logo            string
	Content         string
	Target          string
	Photos          string
	Region_id       int
	Pos             int
	Display_times   int
	Day_limit       int
	Amount          float64
	Total_amount    float64
	Recom_code      string
	Head_income     float64
	Operator_income float64
	Status          int
	Display_count   int
	Views           int
	Valid_time      time.Time
	Create_time     time.Time
}

type AdvVo struct {
	A               *Adv
	Photos          []string
	Logo            string
	ARvo            *AdvRegionVo
	StatusLabel     string
	Head_income     float64
	Operator_income float64
	Desc            string
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
	i, err := orm.NewOrm().QueryTable("tbl_adv").Filter("id", id).Update(orm.Params{"status": 1, "valid_time": time.Now()})
	help.Error(err)

	return int(i) > 0
}

func GetAdvByUid(uid int) []*Adv {
	var advs []*Adv
	_, err := orm.NewOrm().QueryTable("tbl_adv").Filter("uid", uid).OrderBy("-id").All(&advs)
	help.Error(err)

	return advs
}

func GetAdvAll() []*Adv {
	var advs []*Adv
	_, err := orm.NewOrm().QueryTable("tbl_adv").Filter("status__gte", 0).OrderBy("-id").All(&advs)
	help.Error(err)

	return advs
}

func GetAdvByTypeAndRegionAndPos(t, r, p int) []*Adv {
	var advs []*Adv
	_, err := orm.NewOrm().QueryTable("tbl_adv").Filter("type", t).Filter("region_id", r).Filter("pos", p).Filter("status", 1).All(&advs)
	help.Error(err)

	return advs
}

func UpdateViews(id int) {
	orm.NewOrm().QueryTable("tbl_adv").Filter("id", id).Update(orm.Params{"views": orm.ColValue(orm.ColAdd, 1)})
}

func ShowListAdvByRegion(r, size int) []*Adv {
	var advs []*Adv
	aps := GetAdvPosByType(1)
	for _, p := range aps {
		plist := GetAdvByTypeAndRegionAndPos(1, r, p.Id)
		l := len(plist)
		if l > 0 {
			randnum := rand.Intn(l)
			advs = append(advs, plist[randnum])
		}
	}

	new_advs := []*Adv{}
	if size < 5 {
		new_advs = advs[0:0]
	} else if size < 15 {
		for _, a := range advs {
			if a.Pos <= 1 {
				new_advs = append(new_advs, a)
			}
		}
	} else if size < 25 {
		for _, a := range advs {
			if a.Pos <= 2 {
				new_advs = append(new_advs, a)
			}
		}
	} else if size < 35 {
		for _, a := range advs {
			if a.Pos <= 3 {
				new_advs = append(new_advs, a)
			}
		}
	} else if size <= 40 {
		new_advs = advs
	}

	go func(advs []*Adv) {
		for _, adv := range advs {
			orm.NewOrm().QueryTable("tbl_adv").Filter("id", adv.Id).Update(orm.Params{"display_count": orm.ColValue(orm.ColAdd, 1)})
		}
	}(new_advs)

	return new_advs
}

func ShowViewAdvByRegion(r int) *Adv {
	adv := &Adv{}
	plist := GetAdvByTypeAndRegionAndPos(2, r, 5)
	l := len(plist)
	if l > 0 {
		randnum := rand.Intn(l)
		adv = plist[randnum]
	}

	go func(adv *Adv) {
		orm.NewOrm().QueryTable("tbl_adv").Filter("id", adv.Id).Update(orm.Params{"display_count": orm.ColValue(orm.ColAdd, 1)})

	}(adv)

	return adv
}

func ConvertAdvToVo(adv *Adv) *AdvVo {
	vo := &AdvVo{}
	vo.A = adv
	var photos []string
	photos = strings.Split(strings.Trim(adv.Photos, ","), ",")
	vo.Photos = photos
	logos := strings.Split(strings.Trim(adv.Logo, ","), ",")
	logo_size := len(logos)
	if logo_size > 0 {
		randnum := rand.Intn(logo_size)
		vo.Logo = logos[randnum]
	}
	vo.StatusLabel = AdvStatusArr()[adv.Status]
	vo.Desc = help.SubStr(help.HtmlToStr(adv.Content), 0, 300)

	ar, _ := GetAdvRegionByRegionidAndPosid(adv.Region_id, adv.Pos)
	vo.ARvo = ConvertAdvRegionToVo(ar)

	return vo
}

func ConvertAdvToVos(advs []*Adv) []*AdvVo {
	vos := []*AdvVo{}
	for _, a := range advs {
		vo := ConvertAdvToVo(a)
		vos = append(vos, vo)
	}

	return vos
}

func ConvertAdvToVos2(advs *[]*Adv) []*AdvVo {
	vos := []*AdvVo{}
	for _, a := range *advs {
		vo := ConvertAdvToVo(a)
		vos = append(vos, vo)
	}

	return vos
}

func AdvStatusArr() map[int]string {
	return map[int]string{
		0: "待支付",
		1: "已支付",
		2: "暂停投放",
	}
}
