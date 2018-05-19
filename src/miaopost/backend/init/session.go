package init

import (
	"encoding/gob"
	_ "github.com/astaxie/beego/session/redis"
	"miaopost/backend/models"
)

func init() {
	gob.Register(&models.Administrator{})
	gob.Register(&models.User{})
}
