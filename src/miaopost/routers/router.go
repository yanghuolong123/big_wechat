package routers

import (
	"github.com/astaxie/beego"
	"miaopost/controllers"
)

func init() {
	beego.Router("/test", &controllers.TestController{})

	beego.Router("/", &controllers.InfoController{})
}
