package controllers

import (
	"miaopost/backend/models"
	//"miaopost/backend/models"
	"yhl/help"
	"yhl/model"
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

func (this *AdvController) AdvList() {
	page := this.Int("page")

	q := model.Query{}
	q.Table = "tbl_adv"
	q.Condition = map[string]interface{}{"status__gte": 0}
	q.OrderBy = []string{"status", "-create_time"}
	var slice []*models.Adv
	q.ReturnModelList = &slice
	p := help.GetPageList(q, int(page), 1000)

	this.Data["totalCount"] = p.TotalCount
	this.Data["totalPage"] = p.TotalPage
	this.Data["currentPage"] = p.CurrentPage
	this.Data["hasMore"] = p.HasMore
	list := p.DataList.(*[]*models.Adv)
	this.Data["dataList"] = models.ConvertAdvToVos2(list)

	this.Layout = "layout/main.tpl"
	this.TplName = "adv/advList.tpl"
}
