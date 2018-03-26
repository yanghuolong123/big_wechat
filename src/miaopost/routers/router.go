package routers

import (
	"github.com/astaxie/beego"
	"miaopost/controllers"
	"yhl/api"
)

func init() {
	beego.Router("/test", &controllers.TestController{})

	beego.Router("/uploadfile", &api.UploadController{}, "post:Uploadfile")

	beego.Router("/tips", &controllers.TipsController{}, "get:Tips")

	beego.Router("/", &controllers.InfoController{})
	beego.Router("/info/create", &controllers.InfoController{}, "get:CreateGet")
	beego.Router("/info/create", &controllers.InfoController{}, "post:CreatePost")
	beego.Router("/info/edit", &controllers.InfoController{}, "get:EditGet")
	beego.Router("/info/edit", &controllers.InfoController{}, "post:EditPost")
	beego.Router("/info/view", &controllers.InfoController{}, "*:View")
	beego.Router("/info/list", &controllers.InfoController{}, "*:List")
	beego.Router("/info/suggestDel", &controllers.InfoController{}, "post:SuggestDel")
	beego.Router("/info/listPage", &controllers.InfoController{}, "post:ListPage")
}
