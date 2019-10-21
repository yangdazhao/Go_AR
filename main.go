package main

import (
	. "github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	_ "go_AR/routers"
)

func main() {
	BConfig.WebConfig.Session.SessionOn = true
	BConfig.WebConfig.Session.SessionName = "bgsessionID"
	BConfig.WebConfig.StaticDir["/Json"] = "Json"
	SetStaticPath("/Json", "Json")

	//  允许跨域
	InsertFilter("*", BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	Run()
}
