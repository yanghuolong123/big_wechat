package controllers

import (
	"miaopost/frontend/models"
	"yhl/help"
	"yhl/wechat"
)

var WxShare = wechat.Share{
	Title: "出租房 - 秒Po",
	Desc:  "“秒Po”是中国学生的极简信息发布平台。平台以各校区独立运营、免注册极简发布等模式，让学生间的信息交互更加简洁、高效和相对安全，有效解决聊天群信息分数、大众平台繁琐杂乱等问题。解决学生间的买卖二手，房屋租赁，求租，出租单间，卖车，买车，办卡，开电灯问题。常见二手物品包括床垫，床架，书桌，台灯，洗衣机，烘干机等。我们将会在主要学校开通秒Po，包括UTD,UT, UIUC, UW,SMU, USC, PU, NEU, Columbia, OSU, UCLA, Indiana University, Berkeley, NYU, PSU, ASU, UMAA, Boston, IIT, Rutgers.",
	Link:  "http://www.miaopost.com",
	Img:   "http://www.miaopost.com/static/img/logo.png",
}

type BaseController struct {
	help.BaseController
}

func (this *BaseController) Prepare() {
	this.BaseController.Prepare()

	footer_nav := models.GetArticleByType(models.Type_Nav)
	this.Data["footer_nav"] = footer_nav
	this.Data["last_footer_nav_index"] = len(footer_nav) - 1

	side_adv := models.GetArticleByTypeAndGroup(models.Type_Adv, models.Adv_Side)
	this.Data["side_adv"] = side_adv

	this.Data["version"] = help.Version

	signPackage := wechat.GetSignPackage()
	this.Data["signPackage"] = signPackage
	this.Data["wxshare"] = WxShare
}
