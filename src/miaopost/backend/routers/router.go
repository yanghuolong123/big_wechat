package routers

import (
	"github.com/astaxie/beego"
	"miaopost/backend/controllers"
	"yhl/api"
)

func init() {
	beego.Router("/uploadfile", &api.UploadController{}, "post:Uploadfile")
	beego.Router("/ckuploadfile", &api.UploadController{}, "post:CkUploadfile")
	beego.Router("/tips", &api.TipsController{}, "get:Tips")

	beego.Router("/", &controllers.SiteController{})
	beego.Router("/login", &controllers.SiteController{}, "get:LoginGet")
	beego.Router("/login", &controllers.SiteController{}, "post:LoginPost")
	beego.Router("/logout", &controllers.SiteController{}, "*:Logout")

	beego.Router("/article/list", &controllers.ArticleController{}, "*:List")
	beego.Router("/article/create", &controllers.ArticleController{}, "*:Create")
	beego.Router("/article/edit", &controllers.ArticleController{}, "*:Edit")
	beego.Router("/article/delete", &controllers.ArticleController{}, "*:Delete")

	beego.Router("/adv/regionPos", &controllers.AdvController{}, "*:RegionPos")
	beego.Router("/adv/updatePosPrice", &controllers.AdvController{}, "post:UpdatePosPrice")
	beego.Router("/adv/advList", &controllers.AdvController{}, "*:AdvList")
	beego.Router("/adv/updateStatus", &controllers.AdvController{}, "post:UpdateStatus")

}
