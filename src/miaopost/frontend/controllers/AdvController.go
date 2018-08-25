package controllers

import (
	"miaopost/frontend/models"
	"strings"
	"yhl/help"
)

type AdvController struct {
	BaseController
}

func (this *AdvController) CreateGet() {
	if this.User == nil {
		this.Tips("请先登陆")
	}

	type_id, _ := this.GetInt("type")
	tid := int(type_id)
	if tid != 1 && tid != 2 {
		this.Tips("类型不正确")
	}
	this.Data["tid"] = tid

	posList := models.GetAdvRegionByRegionId(this.Rid)
	this.Data["posList"] = models.ConvertAdvRegionToVos(posList)

	this.Layout = "layout/main.tpl"
	this.TplName = "adv/create.tpl"
}

func (this *AdvController) CreatePost() {
	if this.User == nil {
		this.SendRes(-1, "请先登陆", nil)
	}

	adv := new(models.Adv)
	if err := this.ParseForm(adv); err != nil {
		this.SendRes(-1, err.Error(), nil)
	}
	adv.Uid = this.User.Id
	region := models.GetRegionById(adv.Region_id)
	adv.Recom_code = strings.TrimSpace(adv.Recom_code)
	if adv.Recom_code != "" && region.Recom_code != adv.Recom_code {
		this.SendRes(-1, "推荐码不正确", nil)
	}

	advRe, _ := models.GetAdvRegionByRegionidAndPosid(adv.Region_id, adv.Pos)
	adv.Amount = advRe.Price
	adv.Total_amount = adv.Amount * float64(adv.Display_times)
	if adv.Recom_code != "" {
		adv.Operator_income = adv.Total_amount * 0.85
		adv.Head_income = adv.Total_amount - adv.Operator_income
	} else {
		adv.Operator_income = adv.Total_amount * 0.15
		adv.Head_income = adv.Total_amount - adv.Operator_income
	}

	if adv.Merch_name == "" || adv.Contact == "" {
		this.SendRes(-1, "参数不正确", nil)
	}

	err := models.CreateAdv(adv)
	if err != nil {
		help.Error(err)
		this.SendRes(-1, err.Error(), adv)
	}

	this.SendRes(0, "success", adv)
}

func (this *AdvController) ShowList() {
	if this.User == nil {
		this.SendRes(-1, "请先登陆", nil)
	}

	advs := models.ShowListAdvByRegion(this.Rid)

	this.SendRes(0, "success", advs)
}

func (this *AdvController) ShowView() {
	if this.User == nil {
		this.SendRes(-1, "请先登陆", nil)
	}

	adv := models.ShowViewAdvByRegion(this.Rid)
	if adv.Id <= 0 {
		this.SendRes(-1, "failed", nil)
	}

	vo := models.ConvertAdvToVo(adv)

	this.SendRes(0, "success", vo)
}

func (this *AdvController) View() {
	id, _ := this.GetInt("id")
	aid := int(id)

	adv, err := models.GetAdvById(aid)
	if err != nil {
		this.Tips(err.Error())
	}

	this.Data["adv"] = models.ConvertAdvToVo(adv)

	this.Layout = "layout/main.tpl"
	this.TplName = "adv/view.tpl"
}
