package routers

import (
	"github.com/astaxie/beego"
	"miaopost/controllers"
	"yhl/api"
)

func init() {
	beego.Router("/test", &controllers.TestController{})

	beego.Router("/uploadfile", &api.UploadController{}, "post:Uploadfile")

	beego.Router("/", &controllers.InfoController{})
	beego.Router("/info/create", &controllers.InfoController{}, "get:CreateGet")
	beego.Router("/info/create", &controllers.InfoController{}, "post:CreatePost")
	beego.Router("/info/view", &controllers.InfoController{}, "*:View")
}
