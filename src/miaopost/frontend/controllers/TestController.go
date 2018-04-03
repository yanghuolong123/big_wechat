package controllers

import (
	"fmt"
	"miaopost/frontend/models"
	"time"
	"yhl/help"
	m "yhl/model"
)

type TestController struct {
	help.BaseController
}

func (this *TestController) Get() {
	t := time.Now()
	q := m.Query{}
	q.Model = models.Info{}
	p := help.GetPageList(q, 1, 10)
	fmt.Println(p.String())

	/*	t, _ := time.Parse(help.DatetimeFormat, "2018-04-03 00:10:00")
		//local, _ := time.LoadLocation("Local")
		//t = t.In(local)
		fmt.Println("================= t:", t)
		t = t.AddDate(0, 0, -1)
		fmt.Println("================= t:", t)
		date_begin := help.GetDateBegin(t)
		date_end := help.GetDateEnd(t)
		fmt.Println("========= date_begin:", date_begin)
		fmt.Println("========= date_end:", date_end)
		date_pv := models.StatPv(date_begin, date_end)
		date_uv := models.StatUv(date_begin, date_end)
		date_num := models.StatCountInfo(date_begin, date_end)

		fmt.Println("========= date_pv:", date_pv)
		fmt.Println("========= date_uv:", date_uv)
		fmt.Println("========= date_num:", date_num)
		mon_begin := help.GetMonthBegin(t)
		mon_end := help.GetMonthEnd(t)
		fmt.Println("========= mon_begin:", mon_begin)
		fmt.Println("========= mon_end:", mon_end)
	*/

	//	q := m.Query{}
	//	q.Model = models.Info{}
	//
	//	p := help.GetPageList(q, 0, 10)
	//	fmt.Println(p.String())
	//	fmt.Println("========= :", t.AddDate(0, 0, -1))
	//	begin := help.GetDateBegin(t)
	//	end := help.GetDateEnd(t)
	//begin := help.GetMonthBegin(t)
	//end := help.GetMonthEnd(t)

	//	pv := models.StatPv(begin, end)
	//	fmt.Println("============ pv:", pv)
	//	uv := models.StatUv(begin, end)
	//	fmt.Println("============ uv:", uv)

	//count := models.StatCountInfo(begin, end)
	//fmt.Println("=============== count:", count)

	_ = t

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
	/*
		for i := 100; i < 130; i++ {
			id := i
			//email := "yanghuolong@zhisland.com"
			code := help.DesEncrypt(fmt.Sprintf("%v", id), help.DesKey)
			linkUrl := "http://localhost:8686/info/edit?code=" + code
			fmt.Println("=== linkUrl:", linkUrl)
			msg := "亲，欢迎您使用秒Po，您可以通过点击链接修改你发布的信息 <a href=\"" + linkUrl + "\">进入</a>"
			//	go help.SendMail("yhl27ml@126.com", "秒Po-系统发送", msg, "html")
			fmt.Println(msg)
		}
	*/

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
	infos := models.GetInfoPage(0, 0, 2)
	fmt.Println(infos)
	fmt.Println(models.GetInfoCount(0))
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

func (this *TestController) DelInfo() {
	id, _ := this.GetInt("id")
	if models.DelInfoById(int(id)) {
		this.Tips("删除发布信息成功！ id:" + help.ToStr(id))
	} else {
		this.Tips("删除操作失败!")
	}
}
