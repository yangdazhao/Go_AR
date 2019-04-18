package routers

import (
	"github.com/astaxie/beego"
	"go_AR/controllers"
)

func init() {
	beego.Router("/home/login", &controllers.AuthController{}, "post:PostData;get:Login")
	beego.Router("/", 		&controllers.StatisticalConrollerEx{})
	beego.Router("/index", &controllers.StatisticalConrollerEx{})
	beego.Router("/Task/:Time/:TaxpayerId", &controllers.TaskController{})
	beego.Router("/Task/:Time/", &controllers.TaskController{})
	beego.Router("/Task/", &controllers.TaskController{})
	beego.Router("/currentday/", &controllers.CurrentController{})
	beego.Router("/statistical/:TaskID", &controllers.StatisticalConroller{})
	beego.Router("/statisticalEx/", &controllers.StatisticalConrollerEx{})
}


