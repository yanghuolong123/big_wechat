package controllers

import (
	"miaopost/backend/models"
	//"miaopost/backend/models"
	"yhl/help"
)

type AdvController struct {
	BaseController
}

func (this *AdvController) RegionPos() {
	regions := models.GetAllRegion()
	for _, r := range regions {
		advRegions := models.GetAdvRegionByRegionId(r.Id)
		if len(advRegions) <= 0 {
			posList := models.GetAdvPosList()
			for _, p := range posList {
				ar := &models.AdvRegion{}
				ar.Region_id = r.Id
				ar.Pos_id = p.Id
				models.CreateAdvRegion(ar)
			}
		}
	}

	rs := models.GetAllAdvRegion()
	vos := models.ConvertAdvRegionToVos(rs)

	this.Data["vos"] = vos

	this.Layout = "layout/main.tpl"
	this.TplName = "adv/regionPos.tpl"
}

func (this *AdvController) UpdatePosPrice() {
	id := this.Int("id")
	price, err := this.GetFloat("price")

	if id <= 0 || price <= 0 {
		this.SendRes(-1, "failed, 参数错误", nil)
	}

	rp, err := models.GetAdvRegionById(id)
	if err != nil {
		help.Error(err)
		this.SendRes(-1, err.Error(), nil)
	}

	rp.Price = price
	err = models.UpdateAdvRegion(rp)
	if err != nil {
		help.Error(err)
		this.SendRes(-1, err.Error(), nil)
	}

	this.SendRes(0, "success", nil)
}
