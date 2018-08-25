package models

import (
	"github.com/astaxie/beego/orm"
	"yhl/help"
)

func init() {
	orm.RegisterModelWithPrefix("tbl_", new(AdvRegion))
}

type AdvRegion struct {
	Id        int
	Region_id int
	Pos_id    int
	Price     float64
	Status    int
}

type AdvRegionVo struct {
	AdvRe  *AdvRegion
	Region *Region
	Pos    *AdvPos
}

func CreateAdvRegion(ar *AdvRegion) error {
	_, err := orm.NewOrm().Insert(ar)
	help.Error(err)

	return err
}

func UpdateAdvRegion(ar *AdvRegion) error {
	_, err := orm.NewOrm().Update(ar)
	help.Error(err)

	return err
}

func GetAllAdvRegion() (ars []*AdvRegion) {
	_, err := orm.NewOrm().QueryTable("tbl_adv_region").Filter("status", 0).All(&ars)
	help.Error(err)

	return
}

func GetAdvRegionById(id int) (ar *AdvRegion, err error) {
	ar = &AdvRegion{}
	ar.Id = id
	err = orm.NewOrm().Read(ar)
	help.Error(err)

	return
}

func GetAdvRegionByRegionidAndPosid(rid, pid int) (*AdvRegion, error) {
	ar := &AdvRegion{}
	ar.Region_id = rid
	ar.Pos_id = pid
	err := orm.NewOrm().Read(ar, "region_id", "pos_id")
	help.Error(err)

	return ar, err
}

func GetAdvRegionByRegionId(rid int) (ars []*AdvRegion) {
	_, err := orm.NewOrm().QueryTable("tbl_adv_region").Filter("region_id", rid).Filter("status", 0).All(&ars)
	help.Error(err)

	return
}

func BanAdvRegionById(id int) bool {
	i, err := orm.NewOrm().QueryTable("tbl_adv_region").Filter("id", id).Update(orm.Params{"status": -1})
	help.Error(err)

	return i > 0
}

func UpdateAdvRegionPriceById(id int, price float64) bool {
	i, err := orm.NewOrm().QueryTable("tbl_adv_region").Filter("id", id).Update(orm.Params{"price": price})
	help.Error(err)

	return i > 0
}

func ConvertAdvRegionToVo(advRe *AdvRegion) *AdvRegionVo {
	vo := &AdvRegionVo{}
	vo.AdvRe = advRe
	vo.Region = GetRegionById(advRe.Region_id)
	vo.Pos = GetAdvPosById(advRe.Pos_id)

	return vo
}

func ConvertAdvRegionToVos(advRes []*AdvRegion) []*AdvRegionVo {
	vos := []*AdvRegionVo{}
	for _, ar := range advRes {
		vo := ConvertAdvRegionToVo(ar)
		vos = append(vos, vo)
	}

	return vos
}
