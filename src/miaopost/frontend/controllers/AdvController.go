package controllers

import (
	"miaopost/frontend/models"
	//	"yhl/help"
	"fmt"
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
	fmt.Println(adv)

	advRe, _ := models.GetAdvRegionByRegionidAndPosid(adv.Region_id, adv.Pos)
	adv.Amount = advRe.Price
	adv.Total_amount = adv.Amount * float64(adv.Display_times)

	if adv.Merch_name == "" || adv.Contact == "" {
		this.SendRes(-1, "参数不正确", nil)
	}

	err := models.CreateAdv(adv)
	fmt.Println(err)

	this.SendRes(0, "success", adv)
}
