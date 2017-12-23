package routers

import (
	"github.com/astaxie/beego"
	"webapp/controllers"
)

func init() {
	beego.Router("/test", &controllers.TestController{})
	// 微信接入
	beego.Router("/wechat", &controllers.WechatController{}, "get,post:Index")

	beego.Router("/home", &controllers.HomeController{})

	// test
	beego.Router("/login", &controllers.TestController{}, "post:Login")
	beego.Router("/login", &controllers.TestController{}, "get:LoginPage")
	beego.Router("/logout", &controllers.TestController{}, "post:Logout")
	beego.Router("/register", &controllers.TestController{}, "post:Register")
	beego.Router("/register", &controllers.TestController{}, "get:RegisterPage")

	beego.Router("/", &controllers.PrivateGroupController{})

}
