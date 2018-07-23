package init

import (
	"github.com/astaxie/beego"
	//	"math"
	"strings"
	"yhl/help"
)

func init() {
	beego.AddFuncMap("showtime", help.ShowTime)
	beego.AddFuncMap("showListInfo", ShowListInfo)
	beego.AddFuncMap("cutImgSize", CutImgSize)
}

func ShowListInfo(content string) (s string) {
	s = strings.Replace(content, "\n", "&nbsp;&nbsp;&nbsp;", -1)
	s = strings.TrimPrefix(s, "描述：")
	s = strings.TrimSpace(s)
	s = strings.TrimSpace(s)

	return
}

func CutImgSize(w, h int) string {
	//	size := int(math.Max(float64(w), float64(h)))
	//if size >= 640 {
	//	size = int(size / 3)
	//	return strings.Repeat("!"+help.ToStr(size), 2)
	//}
	size := w
	if w > 480 {
		size = 480
		return strings.Repeat("!"+help.ToStr(size), 2)
	}

	return ""
}
