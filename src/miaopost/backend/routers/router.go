package routers

import (
	"github.com/astaxie/beego"
	"miaopost/backend/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
