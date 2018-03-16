package main

import (
	"github.com/astaxie/beego"
	_ "miaopost/init"
	_ "miaopost/routers"
)

func main() {
	beego.SetStaticPath("/uploads", "uploads")
	beego.Run()
}
