package routers

import (
	"github.com/astaxie/beego"
	"webapp/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
}
