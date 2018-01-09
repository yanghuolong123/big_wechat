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
	//beego.Router("/test/login", &controllers.TestController{}, "get:LoginPage")
	//beego.Router("/test/logout", &controllers.TestController{}, "get,post:Logout")
	beego.Router("/test/register", &controllers.TestController{}, "post:Register")
	beego.Router("/test/register", &controllers.TestController{}, "get:RegisterPage")

	beego.Router("/tips", &controllers.TipsController{}, "get:Tips")
	beego.Router("/tips/pgindex", &controllers.TipsController{}, "get:PgIndex")
	beego.Router("/tips/pglist", &controllers.TipsController{}, "get:PgList")

	beego.Router("/uploadfile", &controllers.UploadController{}, "post:Uploadfile")

	beego.Router("/login", &controllers.LoginController{}, "post:LoginPost")
	beego.Router("/login", &controllers.LoginController{}, "get:LoginGet")
	beego.Router("/loginByKey", &controllers.LoginController{}, "get:LoginByKey")
	beego.Router("/logout", &controllers.LoginController{}, "get,post:Logout")

	beego.Router("/search/group", &controllers.SearchController{}, "post:Group")

	beego.Router("/user", &controllers.UserController{}, "get:Index")
	beego.Router("/user/edit", &controllers.UserController{}, "get:EditGet")
	beego.Router("/user/edit", &controllers.UserController{}, "post:EditPost")

	beego.Router("/", &controllers.PrivateGroupController{})
	beego.Router("/pg/create", &controllers.PrivateGroupController{}, "get:CreateGet")
	beego.Router("/pg/create", &controllers.PrivateGroupController{}, "post:CreatePost")
	beego.Router("/pg/edit", &controllers.PrivateGroupController{}, "get:EditGet")
	beego.Router("/pg/edit", &controllers.PrivateGroupController{}, "post:EditPost")
	beego.Router("/pg/view", &controllers.PrivateGroupController{}, "*:View")
	beego.Router("/pg/createPgMsg", &controllers.PrivateGroupController{}, "post:CreatePgMsg")
	beego.Router("/pg/createReport", &controllers.PrivateGroupController{}, "post:CreateReport")
	beego.Router("/pg/list", &controllers.PrivateGroupController{}, "*:List")
	beego.Router("/pg/unlock", &controllers.PrivateGroupController{}, "post:Unlock")

}
