package routers

import (
	"github.com/astaxie/beego"
	"webapp/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/login", &controllers.HomeController{}, "post:Login")
	beego.Router("/login", &controllers.HomeController{}, "get:LoginPage")
	beego.Router("/logout", &controllers.HomeController{}, "post:Logout")
	beego.Router("/register", &controllers.HomeController{}, "post:Register")
	beego.Router("/register", &controllers.HomeController{}, "get:RegisterPage")
}
