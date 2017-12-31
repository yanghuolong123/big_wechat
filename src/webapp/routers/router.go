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
	beego.Router("/test/login", &controllers.TestController{}, "get,post:Login")
	//beego.Router("/login", &controllers.TestController{}, "get:LoginPage")
	//beego.Router("/logout", &controllers.TestController{}, "post:Logout")
	beego.Router("/register", &controllers.TestController{}, "post:Register")
	beego.Router("/register", &controllers.TestController{}, "get:RegisterPage")

	beego.Router("/", &controllers.PrivateGroupController{})
	beego.Router("/uploadfile", &controllers.UploadController{}, "post:Uploadfile")
	beego.Router("/login", &controllers.LoginController{}, "post:LoginPost")
	beego.Router("/login", &controllers.LoginController{}, "get:LoginGet")
	beego.Router("/pg/create", &controllers.PrivateGroupController{}, "get:CreateGet")
	beego.Router("/pg/create", &controllers.PrivateGroupController{}, "post:CreatePost")
	beego.Router("/pg/user", &controllers.PrivateGroupController{}, "*:User")
	beego.Router("/pg/view", &controllers.PrivateGroupController{}, "*:View")
	beego.Router("/pg/createPgMsg", &controllers.PrivateGroupController{}, "post:CreatePgMsg")
	beego.Router("/pg/createReport", &controllers.PrivateGroupController{}, "post:CreateReport")

}
