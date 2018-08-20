package controllers

import (
	"miaopost/backend/models"
	//"miaopost/backend/models"
	//"yhl/help"
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
