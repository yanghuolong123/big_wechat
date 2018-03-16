package task

import (
	"fmt"
	"github.com/astaxie/beego/toolbox"
	"miaopost/models"
	"yhl/help"
)

func init() {
	autoDelExpireInfo := toolbox.NewTask("mytask", "0 0/10 * * * *", func() error {
		f := models.DelExpireInfo()
		help.Log("task", fmt.Sprintf("%v:%v", "自动删除过期的发布信息：", f))
		return nil
	})

	toolbox.AddTask("autoDelExpireInfo", autoDelExpireInfo)

	toolbox.StartTask()
}
