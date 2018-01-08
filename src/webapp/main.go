package main

import (
	"github.com/astaxie/beego"
	//	"webapp/chat"
	_ "webapp/init"
	_ "webapp/routers"
)

func main() {
	//	server := chat.NewServer("/chat")
	//	go server.Listen()

	beego.SetStaticPath("/uploads", "uploads")
	beego.Run()
}
