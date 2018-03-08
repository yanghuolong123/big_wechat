package routers

import (
	"github.com/astaxie/beego"
	"mipo/controllers"
)

func init() {
	beego.Router("/test", &controllers.TestController{})

	beego.Router("/", &controllers.InfoController{})
}
