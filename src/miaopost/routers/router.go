package routers

import (
	"github.com/astaxie/beego"
	"miaopost/controllers"
	"yhl/api"
)

func init() {
	beego.Router("/test", &controllers.TestController{})

	beego.Router("/", &controllers.InfoController{})
	beego.Router("/uploadfile", &api.UploadController{}, "post:Uploadfile")
}
