package controllers

import (
	"fmt"
	"miaopost/frontend/models"
	"yhl/help"
	"yhl/wechat"
)

var WxShare = wechat.Share{
	Title: "秒Po-中国学生极简信息发布平台",
	Desc:  "“秒Po”是中国学生的极简信息发布平台。平台以各校区独立运营、免注册极简发布等模式，让学生间的信息交互更加简洁、高效和相对安全，有效解决聊天群信息分数、大众平台繁琐杂乱等问题。解决学生间的买卖二手，房屋租赁，求租，出租单间，卖车，买车，办卡，开电灯问题。常见二手物品包括床垫，床架，书桌，台灯，洗衣机，烘干机等。我们将会在主要学校开通秒Po，包括UTD,UT, UIUC, UW,SMU, USC, PU, NEU, Columbia, OSU, UCLA, Indiana University, Berkeley, NYU, PSU, ASU, UMAA, Boston, IIT, Rutgers.",
	Link:  "http://www.miaopost.com",
	Img:   "http://www.miaopost.com/static/img/logo.png",
}

type BaseController struct {
	help.BaseController
	Rid int // region id
}

func (this *BaseController) Prepare() {

	isWx := this.IsWeixin()
	this.Data["isWeixin"] = isWx
	if isWx {
		if !this.IsLogin() {
			openid := wechat.GetOpenId(this.Ctx, help.ClientRoute)
			if openid == "" {
				goto loginEnd
			}

			user, err := models.GetUserByOpenid(openid)
			if err == nil {
				this.SetSession("user", user)
				goto loginEnd
			}

			userinfo := wechat.GetWxUserinfo(openid, "")
			if v, ok := userinfo["nickname"]; ok {
				u := models.User{
					Openid:   openid,
					Nickname: v.(string),
					Avatar:   userinfo["headimgurl"].(string),
				}
				if models.CreateUser(&u) > 0 {
					this.SetSession("user", &u)
				}
			}
		}

	loginEnd:

		signPackage := wechat.GetSignPackage()
		this.Data["signPackage"] = signPackage
		this.Data["wxshare"] = WxShare
	}

	user := this.GetSession("user")
	this.Data["user"] = user

	if this.Ctx.Input.Site() == "http://www.miaopost.com" || this.Ctx.Input.Site() == "http://home.miaopost.com" {
		subDomain := ""
		setRegion := this.Ctx.GetCookie("setRegion")
		fmt.Println("============ setRegion:", setRegion)

		if setRegion != "" {
			subDomain = setRegion
		} else if user != nil {
			region := models.GetRegionById(user.(*models.User).Rid)
			if region.Name != "" {
				this.Ctx.SetCookie("setRegion", region.Name, 30*24*3600, "/", "miaopost.com")
				subDomain = region.Name
			}
		}

		if subDomain != "" {
			fmt.Println("============ subDomain:", subDomain)
			uri := this.Ctx.Input.URI()
			if uri == "/" || uri == "" {
				uri = "/info"
			}
			this.Redirect("http://"+subDomain+".miaopost.com"+uri, 302)
		}
	}

	region := this.GetCurrentRegion()
	this.Rid = region.Id

	isMobile := this.IsMobile()
	this.Data["isMobile"] = isMobile
	if !isMobile {
		side_adv := models.GetArticleByTypeAndGroup(this.Rid, models.Type_Adv, models.Adv_Side)
		this.Data["side_adv"] = models.RandAdv(side_adv, 1)
		side_adv_1 := models.GetArticleByTypeAndGroup(this.Rid, models.Type_Adv, models.Adv_Side_1)
		this.Data["side_adv_1"] = models.RandAdv(side_adv_1, 1)
	}

	this.Data["version"] = help.Version

	cats := models.GetCatsByRid(this.Rid)
	this.Data["cats"] = cats

	footer_nav := models.GetArticleByType(models.Type_Nav)
	this.Data["footer_nav"] = footer_nav
	this.Data["last_footer_nav_index"] = len(footer_nav) - 1

	this.BaseController.Prepare()
}

func (this *BaseController) IsLogin() bool {
	user := this.GetSession("user")
	if user != nil {
		return true
	}

	return false
}

func (this *BaseController) GetCurrentRegion() (region models.Region) {
	regions := models.GetAllRegion()
	this.Data["regions"] = regions
	for _, r := range regions {
		if this.Ctx.Input.Domain() == r.Name+".miaopost.com" {
			region = r
			break
		}
	}
	this.Data["region"] = region

	return
}
