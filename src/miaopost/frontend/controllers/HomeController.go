package controllers

import (
	"fmt"
	"miaopost/frontend/models"
	//	"yhl/help"
)

type HomeController struct {
	BaseController
}

func (this *HomeController) Get() {
	//regions := models.GetAllRegion()

	//this.Data["version"] = help.Version
	//this.Data["regions"] = regions

	setRegion := this.Ctx.GetCookie("setRegion")
	fmt.Println("=================== :", setRegion)
	if setRegion != "" {
		this.Redirect("http://"+setRegion+".miaopost.com", 302)
	}

	this.Layout = "layout/main.tpl"
	this.TplName = "home/index.tpl"
}

func (this *HomeController) SetRegion() {
	rid, _ := this.GetInt("rid")
	region := models.GetRegionById(int(rid))

	this.Ctx.SetCookie("setRegion", region.Name, 30*24*3600, "/")

	this.Redirect("http://"+region.Name+".miaopost.com", 302)

}
