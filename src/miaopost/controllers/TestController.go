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
	/*id := 2
	email := "yhl27ml@126.com"
	code := help.DesEncrypt(fmt.Sprintf("%v", id)+";"+email, help.DesKey)
	fmt.Println("=========== en code:", code)
	code = help.DesDecrypt(code, help.DesKey)
	fmt.Println("=========== de code:", code)
	*/

	/*
		ptext := "aaadd打算打算打算sdsds"
		dtext := help.DesEncrypt(ptext, help.DesKey)
		fmt.Println("==== dtext:" + dtext)
		pptext := help.DesDecrypt(dtext, help.DesKey)
		fmt.Println("==== pptext:" + pptext)
	*/
	id := 2
	email := "yhl27ml@126.com"
	code := help.DesEncrypt(fmt.Sprintf("%v", id)+","+email, help.DesKey)
	linkUrl := "http://localhost:8686/info/edit?code=" + code
	fmt.Println("=== linkUrl:", linkUrl)
	msg := "亲，欢迎您使用秒Po，您可以通过点击链接修改你发布的信息 <a href=\"" + linkUrl + "\">进入</a>"
	go help.SendMail("yhl27ml@126.com", "秒Po-系统发送", msg, "html")
	fmt.Println(msg)

	//go help.SendMail("yanghuolong@zhisland.com,yhl27ml@126.com", "你好，我是 Miaopost", "非常感谢你的测试哦，我们爱你哦！", "")
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
	info, _ := models.GetInfoById(2)
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
