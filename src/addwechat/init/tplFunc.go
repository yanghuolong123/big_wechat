package init

import (
	"github.com/astaxie/beego"
	"yhl/help"
)

func init() {
	beego.AddFuncMap("showtime", help.ShowTime)
}
