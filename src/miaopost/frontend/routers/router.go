package routers

import (
	"github.com/astaxie/beego"
	"miaopost/frontend/controllers"
	"yhl/api"
)

func init() {
	beego.Router("/test", &controllers.TestController{})
	beego.Router("/delete", &controllers.TestController{}, "*:DelInfo")

	// 微信接入
	beego.Router("/wechat", &controllers.WechatController{}, "get,post:Index")
	beego.Router("/uploadfile", &api.UploadController{}, "post:Uploadfile")
	beego.Router("/webupload", &api.UploadController{}, "post:WebUpload")

	beego.Router("/tips", &api.TipsController{}, "get:Tips")
	beego.Router("/qrcode/png", &api.QrcodeController{}, "get:Png")

	// Home
	beego.Router("/home", &controllers.HomeController{})
	beego.Router("/setRegion", &controllers.HomeController{}, "*:SetRegion")

	// 登陆
	beego.Router("/login", &controllers.LoginController{}, "post:LoginPost")
	beego.Router("/login", &controllers.LoginController{}, "get:LoginGet")
	beego.Router("/loginByKey", &controllers.LoginController{}, "get:LoginByKey")
	beego.Router("/logout", &controllers.LoginController{}, "get,post:Logout")

	// 用户
	beego.Router("/user", &controllers.UserController{}, "get:Index")
	beego.Router("/user/edit", &controllers.UserController{}, "get:EditGet")
	beego.Router("/user/edit", &controllers.UserController{}, "post:EditPost")

	// 信息
	beego.Router("/", &controllers.InfoController{})
	beego.Router("/info/create", &controllers.InfoController{}, "get:CreateGet")
	beego.Router("/info/create", &controllers.InfoController{}, "post:CreatePost")
	beego.Router("/info/edit", &controllers.InfoController{}, "get:EditGet")
	beego.Router("/info/edit", &controllers.InfoController{}, "post:EditPost")
	beego.Router("/info/view", &controllers.InfoController{}, "*:View")
	beego.Router("/info/list", &controllers.InfoController{}, "*:List")
	beego.Router("/info/suggestDel", &controllers.InfoController{}, "post:SuggestDel")
	beego.Router("/info/listPage", &controllers.InfoController{}, "post:ListPage")
	beego.Router("/info/delete", &controllers.InfoController{}, "post:Delete")
	beego.Router("/info/my", &controllers.InfoController{}, "*:My")
	beego.Router("/info/top", &controllers.InfoController{}, "*:Top")

	// 文章
	beego.Router("/article/view", &controllers.ArticleController{}, "*:View")
}
