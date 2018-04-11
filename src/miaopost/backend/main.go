package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	beeUtils "github.com/astaxie/beego/utils"
	_ "miaopost/backend/init"
	_ "miaopost/backend/routers"
)

func main() {
	filterRoute := []string{"/login", "/register", "/upload"}
	var filterLogin = func(ctx *context.Context) {
		user := ctx.Input.Session("user")
		if user == nil && !beeUtils.InSlice(ctx.Request.RequestURI, filterRoute) {
			ctx.Redirect(302, "/login")
		}
	}

	beego.InsertFilter("/*", beego.BeforeRouter, filterLogin)
	beego.SetStaticPath("/uploads", "uploads")

	beego.Run()
}
