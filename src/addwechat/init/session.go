package init

import (
	"addwechat/models"
	"encoding/gob"
	_ "github.com/astaxie/beego/session/redis"
)

func init() {
	gob.Register(models.User{})
	gob.Register(models.Group{})
}
