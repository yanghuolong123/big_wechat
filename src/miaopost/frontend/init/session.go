package init

import (
	"encoding/gob"
	_ "github.com/astaxie/beego/session/redis"
	"miaopost/frontend/models"
)

func init() {
	gob.Register(&models.User{})
}
