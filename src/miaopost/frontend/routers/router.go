package routers

import (
	"github.com/astaxie/beego"
	"miaopost/frontend/controllers"
	"yhl/api"
)

func init() {
	beego.Router("/test", &controllers.TestController{})
	beego.Router("/delete", &controllers.TestController{}, "*:DelInfo")

	// 微信接入
	beego.Router("/wechat", &controllers.WechatController{}, "get,post:Index")

	// 文件上传
	beego.Router("/uploadfile", &api.UploadController{}, "post:Uploadfile")
	beego.Router("/kuploadfile", &api.UploadController{}, "post:KUploadfile")
	beego.Router("/webupload", &api.UploadController{}, "post:WebUpload")

	beego.Router("/tips", &api.TipsController{}, "get:Tips")
	beego.Router("/qrcode/png", &api.QrcodeController{}, "get:Png")

	// Home
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/setRegion", &controllers.HomeController{}, "*:SetRegion")

	// 登陆
	beego.Router("/login", &controllers.LoginController{}, "post:LoginPost")
	beego.Router("/login", &controllers.LoginController{}, "get:LoginGet")
	beego.Router("/loginByKey", &controllers.LoginController{}, "get:LoginByKey")
	beego.Router("/logout", &controllers.LoginController{}, "get,post:Logout")

	// 用户
	beego.Router("/user", &controllers.UserController{}, "get:Index")
	beego.Router("/user/edit", &controllers.UserController{}, "get:EditGet")
	beego.Router("/user/edit", &controllers.UserController{}, "post:EditPost")
	beego.Router("/user/my", &controllers.UserController{}, "*:My")
	beego.Router("/user/account", &controllers.UserController{}, "*:Account")

	// 信息
	beego.Router("/info", &controllers.InfoController{})
	beego.Router("/info/create", &controllers.InfoController{}, "get:CreateGet")
	beego.Router("/info/create", &controllers.InfoController{}, "post:CreatePost")
	beego.Router("/info/edit", &controllers.InfoController{}, "get:EditGet")
	beego.Router("/info/edit", &controllers.InfoController{}, "post:EditPost")
	beego.Router("/info/view", &controllers.InfoController{}, "*:View")
	beego.Router("/info/list", &controllers.InfoController{}, "*:List")
	beego.Router("/info/suggestDel", &controllers.InfoController{}, "post:SuggestDel")
	beego.Router("/info/listPage", &controllers.InfoController{}, "post:ListPage")
	beego.Router("/info/delete", &controllers.InfoController{}, "post:Delete")
	beego.Router("/info/top", &controllers.InfoController{}, "*:Top")

	// 留言
	beego.Router("/msg/create", &controllers.InfoMsgController{}, "post:CreateMsg")
	beego.Router("/msg/support", &controllers.InfoMsgController{}, "post:Support")
	beego.Router("/msg/suggestDel", &controllers.InfoMsgController{}, "post:SuggestDel")
	beego.Router("/msg/admire", &controllers.InfoMsgController{}, "post:Admire")

	// 红包
	beego.Router("/reward/chance", &controllers.InfoRewardController{}, "post:Chance")

	// 文章
	beego.Router("/article/view", &controllers.ArticleController{}, "*:View")

	// 支付
	beego.Router("/pay/confirm", &controllers.PayController{}, "get,post:Confirm")
	beego.Router("/pay/wxscan", &controllers.PayController{}, "get,post:WxScan")
	beego.Router("/pay/qrcode", &controllers.PayController{}, "get,post:Qrcode")
	beego.Router("/pay/notify", &controllers.PayController{}, "*:Notify")
	beego.Router("/pay/check", &controllers.PayController{}, "post:Check")
	beego.Router("/pay/withdraw", &controllers.PayController{}, "post:Withdraw")
}
