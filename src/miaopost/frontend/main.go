package main

import (
	"github.com/astaxie/beego"
	_ "miaopost/frontend/init"
	_ "miaopost/frontend/routers"
)

func main() {
	beego.SetStaticPath("/uploads", "uploads")
	beego.Run()
}
