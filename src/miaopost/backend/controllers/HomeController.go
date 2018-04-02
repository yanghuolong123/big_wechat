package controllers

import (
	"fmt"
	"yhl/help"
)

type HomeController struct {
	help.BaseController
}

func (this *HomeController) Get() {
	fmt.Println("========")
}
