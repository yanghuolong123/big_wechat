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

	beego.Router("/tips", &api.TipsController{}, "get:Tips")
	beego.Router("/qrcode/png", &api.QrcodeController{}, "get:Png")

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

	beego.Router("/article/view", &controllers.ArticleController{}, "*:View")
}
