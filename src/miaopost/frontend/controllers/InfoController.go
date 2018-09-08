package controllers

import (
	"fmt"
	"labix.org/v2/mgo/bson"
	"miaopost/frontend/models"
	"strings"
	"time"
	"yhl/help"
	"yhl/model"
	//	"yhl/wechat"
)

const pageSize int = 40

type InfoController struct {
	BaseController
}

// 首页
func (this *InfoController) Get() {
	infos := models.GetInfoPage(0, this.Rid, 0, pageSize)
	this.Data["infos"] = models.ConvertInfosToVo2(&infos)
	count := models.GetInfoCount(0, this.Rid)
	this.Data["hasMore"] = 0
	this.Data["page"] = 0
	this.Data["size"] = len(infos)
	if 1*pageSize < count {
		this.Data["hasMore"] = 1
	}

	adv := models.GetArticleByTypeAndGroup(this.Rid, models.Type_Adv, models.Adv_List_Bottom)
	this.Data["adv_list"] = models.RandAdv(adv, 1)

	end := time.Now()
	begin := end.AddDate(0, 0, -7)
	weekViews := models.StatPvByDomain(begin, end, this.Ctx.Input.Domain())
	this.Data["weekviews"] = weekViews

	this.Layout = "layout/main.tpl"
	this.TplName = "info/home.tpl"
}

// 列表页
func (this *InfoController) List() {
	cid, _ := this.GetInt("cid")
	page, _ := this.GetInt("page")

	this.Data["hasMore"] = 0
	this.Data["page"] = int(page)

	infos := []*models.Info{}
	if catId := int(cid); catId > 0 {
		count := models.GetInfoCount(catId, this.Rid)
		infos = models.GetInfoPage(catId, this.Rid, int(page), pageSize)
		if 1*pageSize < count {
			this.Data["hasMore"] = 1
		}
		//infos = models.GetInfoByCid(catId)
	}

	search := this.GetString("search")
	if search != "" {
		infos = models.SearchInfo(search)
	}

	this.Data["cid"] = int(cid)
	this.Data["infos"] = models.ConvertInfosToVo(infos)
	this.Data["size"] = len(infos)

	adv := models.GetArticleByTypeAndGroup(this.Rid, models.Type_Adv, models.Adv_List_Bottom)
	this.Data["adv_list"] = models.RandAdv(adv, 1)

	end := time.Now()
	begin := end.AddDate(0, 0, -7)
	weekViews := models.StatPvByDomain(begin, end, this.Ctx.Input.Domain())
	this.Data["weekviews"] = weekViews

	this.Layout = "layout/main.tpl"
	//this.TplName = "info/list.tpl"
	this.TplName = "info/home.tpl"
}

// 创建页面
func (this *InfoController) CreateGet() {
	cid, _ := this.GetInt("cid")
	cat := &models.Category{}
	if catId := int(cid); catId > 0 {
		cat = models.GetCategoryById(catId)
	}
	this.Data["cat"] = cat

	//if !this.IsWeixin() {
	//	this.Data["qr_url"] = wechat.GetTmpStrQrImg("create")
	//}

	this.Layout = "layout/main.tpl"
	this.TplName = "info/create.tpl"
}

func (this *InfoController) CreatePost() {
	cid, _ := this.GetInt("cid")
	content := this.GetString("content")
	valid_day, _ := this.GetInt("valid_day")
	email := this.GetString("email")
	photo := this.GetString("photo")
	reward_type, _ := this.GetInt("reward_type")
	reward_num, _ := this.GetInt("reward_num")
	reward_amount, _ := this.GetFloat("reward_amount")
	rtype := int(reward_type)
	rnum := int(reward_num)

	u := this.GetSession("user")

	info := new(models.Info)
	info.Rid = this.Rid
	info.Cid = int(cid)
	info.Content = content
	info.Valid_day = int(valid_day)
	info.Email = email
	info.Reward_type = rtype
	info.Reward_num = rnum
	info.Reward_amount = reward_amount
	if u != nil {
		info.Uid = u.(*models.User).Id
	}

	id := models.CreateInfo(info, this.Ctx)
	if id > 0 {
		if photo != "" {
			plist := strings.Split(photo, ",")
			for _, p := range plist {
				models.CreatePhoto(id, p)
			}
		}

		go func(id int, email string) {
			if email == "" {
				return
			}
			code := help.DesEncrypt(fmt.Sprintf("%v", id)+","+email, help.DesKey)
			linkUrl := "http://www.miaopost.com/info/edit?code=" + code
			msg := "亲，欢迎您使用秒Po，您可以通过点击链接修改您发布的信息 <a href=\"" + linkUrl + "\">进入</a>"
			help.SendMail(email, "秒Po-编辑链接", msg, "html")
		}(id, email)

		this.SendRes(0, "success", info)
	}

	this.SendRes(-1, "failed", nil)
}

// 展示页
func (this *InfoController) View() {
	id, _ := this.GetInt("id")
	models.IncInfoViews(int(id))

	info, err := models.GetInfoById(int(id))
	if err != nil {
		this.Redirect("/tips?msg="+err.Error(), 302)
		return
	}
	if info.Status < 0 {
		this.Tips("此信息已经删除!")
		return
	}
	info.Content = strings.Replace(info.Content, "\n", "<br/>", -1)
	this.Data["info"] = info

	cat := models.GetCategoryById(info.Cid)
	this.Data["cat"] = cat

	photos := models.GetPhotoByInfoid(int(id))
	//this.Data["photos"] = photos
	this.Data["photos"] = models.GetPhotoVoList(photos)

	share := WxShare
	if info.Content != "" {
		share.Desc = strings.Replace(info.Content, "<br/>", " ", -1)
	}
	if len(photos) > 0 {
		share.Img = this.Ctx.Input.Site() + photos[0].Url + "!200!200"
	}
	share.Title = cat.Name + " - 秒Po"
	share.Link = "http://www.miaopost.com" + this.Ctx.Input.URI()
	this.Data["wxshare"] = share

	adv := models.GetArticleByTypeAndGroup(this.Rid, models.Type_Adv, models.Adv_View_Bottom)
	this.Data["adv"] = models.RandAdv(adv, 1)

	ims := models.GetInfoMessageByInfoId(info.Id)
	if len(ims) > 0 {
		imvos := models.ConvertInfoMsgToVos(ims)
		this.Data["imvos"] = imvos
	}

	reward_type := 0
	len := help.Redis.Llen("list_reward_info_" + help.ToStr(info.Id))
	if len > 0 {
		reward_type = info.Reward_type
	}
	this.Data["reward_type"] = reward_type
	chance := this.GetString("chance")
	this.Data["chance"] = chance

	this.Layout = "layout/main.tpl"
	this.TplName = "info/view.tpl"
}

// 编辑
func (this *InfoController) EditGet() {
	code := this.GetString("code")
	infoId, _ := this.GetInt("id")
	var id int
	if code != "" {
		code = help.DesDecrypt(code, help.DesKey)

		s := strings.Split(code, ",")
		id = help.StrToInt(s[0])
		if id <= 0 {
			this.Redirect("/tips?msg=code不正确", 302)
			return
		}
	} else {
		id = int(infoId)
	}

	info, err := models.GetInfoById(id)
	if err != nil {
		this.Redirect("/tips?msg="+err.Error(), 302)
		return
	}
	if info.Status < 0 {
		this.Tips("此信息已被删除")
		return
	}
	this.Data["info"] = info

	cats := models.GetAllCategory()
	this.Data["cats"] = cats
	photos := models.GetPhotoByInfoid(id)
	this.Data["photos"] = photos

	this.Layout = "layout/main.tpl"
	this.TplName = "info/edit.tpl"
}

func (this *InfoController) EditPost() {
	id, _ := this.GetInt("id")
	cid, _ := this.GetInt("cid")
	content := this.GetString("content")
	valid_day, _ := this.GetInt("valid_day")
	email := this.GetString("email")
	photo := this.GetString("photo")

	info, err := models.GetInfoById(int(id))
	if err != nil {
		this.SendRes(-1, err.Error(), nil)
	}
	info.Cid = int(cid)
	info.Content = content
	info.Valid_day = int(valid_day)
	info.Email = email
	err = models.UpdateInfo(info)
	if err != nil {
		this.SendRes(-1, err.Error(), nil)
	}

	models.DelPhotoByInfoid(info.Id)
	if photo != "" {
		plist := strings.Split(photo, ",")
		for _, p := range plist {
			models.CreatePhoto(info.Id, p)
		}
	}

	this.SendRes(0, "success", info)
}

// 建议删除
func (this *InfoController) SuggestDel() {
	infoId, _ := this.GetInt("infoId")
	iid := int(infoId)
	slist := models.GetSuggestByInfoidAndGroupByIp(iid)
	if len(slist) > 3 {
		models.DelInfoById(iid)
	}

	id := models.CreateSuggest(iid, this.Ctx)
	if id > 0 {
		this.SendRes(0, "success", nil)
	}

	this.SendRes(-1, "failed", nil)
}

// 分页列表
func (this *InfoController) ListPage() {
	page, _ := this.GetInt("page")
	cid, _ := this.GetInt("cid")
	uid, _ := this.GetInt("uid")

	q := model.Query{}
	q.Table = "tbl_info"
	cm := map[string]interface{}{}
	if int(cid) > 0 {
		cm["cid"] = int(cid)
	}
	if int(uid) > 0 {
		cm["uid"] = int(uid)
		this.Data["isMy"] = true
	}
	cm["rid"] = this.Rid
	cm["status"] = 0
	q.Condition = cm
	q.OrderBy = []string{"-update_time"}
	var slice []*models.Info
	q.ReturnModelList = &slice
	p := help.GetPageList(q, int(page), pageSize)
	data := p.DataList
	infos := data.(*[]*models.Info)

	this.Data["infos"] = models.ConvertInfosToVo2(infos)
	this.TplName = "info/listPage.tpl"
	s, _ := this.RenderString()

	m := map[string]interface{}{}
	m["listData"] = s
	m["page"] = p.CurrentPage
	m["hasMore"] = p.HasMore
	m["size"] = p.CurrentSize

	this.SendRes(0, "success", m)
}

// 删除发布的信息
func (this *InfoController) Delete() {
	id, _ := this.GetInt("id")
	if models.DelInfoById(int(id)) {
		this.SendRes(0, "success", nil)
	}
	this.SendRes(-1, "failed", nil)
}

// 置顶
func (this *InfoController) Top() {
	id, _ := this.GetInt("id")
	info, err := models.GetInfoById(int(id))
	if err != nil {
		this.SendRes(-1, err.Error(), nil)
	}

	condition := bson.M{
		"infoId": info.Id,
		"uid":    info.Uid,
	}
	c := help.MongoDb.C("top_record")
	count, err := c.Find(condition).Count()
	help.Error(err)
	if count >= 3 {
		this.SendRes(-1, "亲，您的置顶机会已用完，最多可以免费置顶3次", nil)
	}

	info.Update_time = time.Now()
	err = models.UpdateInfo(info)
	if err != nil {
		this.SendRes(-1, err.Error(), nil)
	}
	m := map[string]interface{}{"uid": info.Uid, "infoId": info.Id, "time": time.Now()}
	c.Insert(m)

	this.SendRes(0, "置顶成功，您还有"+help.ToStr(2-count)+"机会", nil)
}
