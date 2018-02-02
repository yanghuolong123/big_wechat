package controllers

import (
	//	"time"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"webapp/models"
	"yhl/help"
	//	"yhl/wechat"
	"time"
	"yhl/search"
	"yhl/wechat/wxpay"
)

type TestController struct {
	help.BaseController
}

func importGroupToEs() {
	var glist []models.Group
	orm.NewOrm().QueryTable("tbl_group").All(&glist)
	fmt.Println("======== len:", len(glist))
	for _, g := range glist {
		m := help.StructToMap(g)
		//	fmt.Println("============= g:", m["id"])
		result := search.Put("/group/school/"+m["id"].(string), m)
		fmt.Println("============ result:", result)
	}
}

func wxPay() {
	orderReq := new(wxpay.UnifyOrderReq)
	orderReq.Body = "商品名"
	orderReq.Out_trade_no = "233233333232"
	orderReq.Total_fee = 100
	orderReq.Notify_url = "pay.feichangjuzu.com/pay/"
	orderReq.Trade_type = "NATIVE"
	orderReq.Product_id = 22
	orderReq.Time_start = time.Now().Format(help.DatetimeNumFormat)
	orderReq.Time_expire = time.Now().Add(time.Duration(600 * time.Second)).Format(help.DatetimeNumFormat)
	orderReq.Spbill_create_ip = help.ClientIp

	wxRes := wxpay.UnifiedOrder(orderReq)
	help.Log("wxpay", wxRes)
	if wxRes.Return_code == "FAIL" {
		help.Log("wxpay", help.StructToMap(wxRes))
	}

	if wxRes.Code_url != "" {
		fmt.Println("====== code_url:", wxRes.Code_url)
	}
	//fmt.Println("============= randStr:", help.RandStr(32))
	help.Log("test.log", help.RandStr(10))
	help.Log("test.log", orderReq)
}

func order() {
	o := models.GetOrderByOrderno("2018020110180392508")
	fmt.Println("================ order:", o)
	o.Status = 6
	models.UpdateOrder(o)
}

func unmarsh() {
	str := `<xml><appid><![CDATA[wx3e0b8bca5b6d6606]]></appid>
<bank_type><![CDATA[CFT]]></bank_type>
<cash_fee><![CDATA[1]]></cash_fee>
<fee_type><![CDATA[CNY]]></fee_type>
<is_subscribe><![CDATA[Y]]></is_subscribe>
<mch_id><![CDATA[1497110522]]></mch_id>
<nonce_str><![CDATA[lK3sw75HdKOw6Zv5lGABaZBGHqUp6wBv]]></nonce_str>
<openid><![CDATA[oou4Vw0zizge_p2gQhYT0UL5Kwbk]]></openid>
<out_trade_no><![CDATA[2018020211255844503]]></out_trade_no>
<result_code><![CDATA[SUCCESS]]></result_code>
<return_code><![CDATA[SUCCESS]]></return_code>
<sign><![CDATA[BAE3A2EBBD8CE52C43A0975DF7770E25]]></sign>
<time_end><![CDATA[20180202112611]]></time_end>
<total_fee>1</total_fee>
<trade_type><![CDATA[NATIVE]]></trade_type>
<transaction_id><![CDATA[4200000052201802025369351717]]></transaction_id>
</xml>`

	var notifyReq = wxpay.WXPayNotifyReq{}
	xml.Unmarshal([]byte(str), &notifyReq)
	fmt.Println("\n=================\n")
	fmt.Printf("%+v", notifyReq)
	fmt.Println("\n=================\n")
	fmt.Printf("Sign:%+v", notifyReq.Sign)
	fmt.Println("\n=================\n")
	notifyReq.Sign = ""
	m := help.StructToMap(notifyReq)
	fmt.Printf("m:%+v", m)
	fmt.Println("\n=================\n")
	signStr := wxpay.Sign(m)
	fmt.Printf("signStr: %+v", signStr)
	fmt.Println("\n=================\n")
}

func (this *TestController) Get() {

	go unmarsh()

	//go order()
	//	go wxPay()
	//	go importGroupToEs()
	/*cache := help.Cache
	token := cache.Get("access_token_")

	if token == nil {
		token = "ddsdsdsdssddsds"
		cache.Put("access_token_", token, 100*time.Second)
	} else {
	 	token = string(token.([]uint8))
	}

	accessToken := token.(string) //wechat.GetAccessToken()
	accessToken := wechat.GetAccessToken()
	m := map[string]interface{}{}
	m["touser"] = "oTbmFxG5r1WRrHdb32O5y2aSAIkc"
	m["msgtype"] = "text"
	m["text"] = map[string]string{"content": "Hello World!"}
	wechat.SendMsg(m)
	wechat.SendTextMsg("oou4Vw0zizge_p2gQhYT0UL5Kwbk", "I love eou")
	userinfo := wechat.GetWxUserinfo("oou4Vw0zizge_p2gQhYT0UL5Kwbk", "")
	fmt.Println(userinfo)
	*/
	/*	userinfo := wechat.GetWxUserinfo("oou4Vw0zizge_p2gQhYT0UL5Kwbk", "")
		fmt.Println(userinfo)
		pg := models.PrivateGroup{}
		pg.Id = 1
		//pg.Gid = 2
		//pg.Uid = 1
		//pg.Name = "aaa"
		pg.Introduction = "aaaaaaaaaaasss"
		pg.Qrcode = "aaaaaaa"
		pg.Ower_qrcode = "bbbb"
		//flag := models.CreatePrivateGroup(&pg)
		flag := models.UpdatePrivateGroup(&pg)
		fmt.Println("================= flag:", flag)
		this.Data["welcome"] = "welcome to add group, accessToken:"
		qrImgUrl := wechat.GetTmpStrQrImg("login_123")
		this.Data["qrImgUrl"] = qrImgUrl
	*/
	//	user, err := models.GetUserByOpenid("oou4Vw0zizge_p2gQhYT0UL5Kwbk")
	//user, err := models.GetUserById(3)
	//	fmt.Println(user, err)
	//	t, _ := time.Parse(help.DatetimeFormat, "2018-01-01 20:15:32")
	//	s := help.ShowTime(t)
	//	fmt.Println("==== t:", s)

	/*m := search.Get("/customer/external/1")
	fmt.Println(m)
	m1 := map[string]interface{}{"name": "yhl1"}
	d1 := search.Put("/customer/external/2", m1)
	fmt.Println(d1)
	m2 := map[string]interface{}{"name": "yhl2"}
	d2 := search.Put("/customer/external/3", m2)
	fmt.Println(d2)
	*/
	this.TplName = "test/index.tpl"
}

func (this *TestController) Post() {

	this.SendRes(0, "success", nil)
}

func (this *TestController) LoginPage() {
	this.TplName = "test/login.tpl"
	s, _ := this.RenderString()

	this.SendRes(0, "success", s)
}

func (this *TestController) login(username, password string) (m map[string]interface{}, err error) {
	m = make(map[string]interface{})
	u, err := models.Login(username, password)
	if err != nil {
		return
	}
	gids := models.GetFollowByUid(u.Id)
	group := models.GetGroupById(u.Gid)

	this.SetSession("user", *u)
	this.SetSession("follow", gids)
	this.SetSession("group", group)

	m["user"] = u
	m["follow"] = gids
	m["group"] = group

	return
}

func (this *TestController) Login() {
	username := this.GetString("username")
	password := this.GetString("password")

	m, err := this.login(username, password)
	if err != nil {
		this.SendRes(-1, err.Error(), nil)
	}

	this.SendRes(0, "success", m)
}

func (this *TestController) Logout() {
	this.DelSession("user")
	this.DelSession("follow")
	this.DelSession("group")

	this.SendRes(0, "success", nil)
}

func (this *TestController) RegisterPage() {
	this.Data["groupList"] = models.GetGroupAll()
	this.TplName = "test/register.tpl"
	s, _ := this.RenderString()

	this.SendRes(0, "success", s)
}

func (this *TestController) Register() {
	gid, _ := this.GetInt("group")
	username := this.GetString("username")
	nickname := this.GetString("nickname")
	password := this.GetString("password")
	repassword := this.GetString("repassword")
	if password != repassword {
		this.SendRes(-1, "密码输入不一致", nil)
	}
	_, err := models.GetUserByUsername(username)
	if err != nil {
		this.SendRes(-1, errors.New("账号已存在").Error(), nil)
	}
	loginPasswd := password
	password = help.Md5(password)

	u := models.User{
		Gid:      gid,
		Username: username,
		Nickname: nickname,
		Password: password,
	}

	uid := models.CreateUser(&u)
	if uid <= 0 {
		this.SendRes(-1, "注册失败", nil)
	}
	models.CreateFollow(uid, gid)

	m, _ := this.login(username, loginPasswd)

	this.SendRes(0, "success", m)
}
