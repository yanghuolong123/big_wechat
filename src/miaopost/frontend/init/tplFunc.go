package init

import (
	"github.com/astaxie/beego"
	"strings"
	"yhl/help"
)

func init() {
	beego.AddFuncMap("showtime", help.ShowTime)
	beego.AddFuncMap("showListInfo", ShowListInfo)
}

func ShowListInfo(content string) (s string) {
	s = strings.Replace(content, "\n", "&nbsp;&nbsp;&nbsp;", -1)
	s = strings.TrimPrefix(s, "描述：")
	s = strings.TrimLeft(s, "&nbsp;")
	s = strings.TrimSpace(s)

	return
}
