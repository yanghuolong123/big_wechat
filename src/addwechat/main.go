package main

import (
	"github.com/astaxie/beego"
	//	"addwechat/chat"
	_ "addwechat/init"
	_ "addwechat/routers"
)

func main() {
	//	server := chat.NewServer("/chat")
	//	go server.Listen()

	beego.SetStaticPath("/uploads", "uploads")
	beego.Run()
}
