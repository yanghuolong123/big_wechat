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
	//this.Redirect("http://utd.miaopost.com/info", 302)
	this.Redirect("/info", 302)

	setRegion := this.Ctx.GetCookie("setRegion")
	fmt.Println("=================== :", setRegion)
	if setRegion != "" {
		this.Redirect("http://"+setRegion+".miaopost.com/info", 302)
		return
	}

	this.Layout = "layout/main.tpl"
	this.TplName = "home/index.tpl"
}

func (this *HomeController) SetRegion() {
	rid, _ := this.GetInt("rid")
	region := models.GetRegionById(int(rid))
	fmt.Println("============ regionName:", region.Name)

	this.Ctx.SetCookie("setRegion", region.Name, 30*24*3600, "/", "miaopost.com")
	user := this.GetSession("user")
	if user != nil {
		u := user.(*models.User)
		u.Rid = region.Id
		this.SetSession("user", u)
		go func(u *models.User) {
			models.UpdateUser(u)
		}(u)
	}

	this.Redirect("http://"+region.Name+".miaopost.com/info", 302)

}
