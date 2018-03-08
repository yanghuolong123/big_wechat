package controllers

import (
	"fmt"
	"miaopost/models"
	"yhl/help"
)

type TestController struct {
	help.BaseController
}

func (this *TestController) Get() {
	//fmt.Println(this.Ctx.Request.RemoteAddr)
	//go createSuggest()
	//go getPhotoByInfoid()
	//go createPhoto()
	//go getPageInfo()
	//	go getInfoByEmail()
	//go getInfoById()
	//	go createInfo()
	//go getAllCategory()
	//go createCategory()

	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplName = "test/index.tpl"
}

func createSuggest() {
	i := models.CreateSuggest(2)
	fmt.Println(i)
}

func getPhotoByInfoid() {
	photos := models.GetPhotoByInfoid(2)
	fmt.Println(photos)
}

func createPhoto() {
	models.CreatePhoto(2, "/aaa/dssd/ss")
}

func getPageInfo() {
	infos := models.GetInfoPage(0, 2)
	fmt.Println(infos)
	fmt.Println(models.GetInfoCount())
}

func getInfoByEmail() {
	infos := models.GetInfoByEmail("yhl11@11.com")
	fmt.Println(infos)
}

func getInfoById() {
	info := models.GetInfoById(2)
	fmt.Println(info)
}

func createInfo() {
	info := new(models.Info)
	info.Cid = 2
	info.Content = "dsdssd搜索"
	//info.Valid_day = 2
	//info.Email = "yhl11@11.com"
	models.CreateInfo(info)
}

func getAllCategory() {
	cList := models.GetAllCategory()
	fmt.Println(cList)
}

func createCategory() {
	models.CreateCategory("求购")
	models.CreateCategory("出售")
	models.CreateCategory("求租房")
	models.CreateCategory("出租房")
	models.CreateCategory("周边服务")
	models.CreateCategory("其他")
}
