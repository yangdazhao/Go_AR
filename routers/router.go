package routers

import (
	"github.com/astaxie/beego"
	"go_AR/controllers"
	"go_AR/controllers/Login"
)

func init() {
	beego.Router("/home/login", &Login.AuthController{}, "post:PostData;get:Login")
	beego.Router("/", &controllers.StatisticalConrollerEx{})
	beego.Router("/statisticalNx", &controllers.StatisticalNxConroller{})
	beego.Router("/index", &controllers.StatisticalConrollerEx{})
	beego.Router("/Task/:Time/:TaxpayerId", &controllers.TaskController{})
	beego.Router("/Task/:Time/", &controllers.TaskController{})
	beego.Router("/Task/:Time/", &controllers.TaskController{})
	beego.Router("/Task/", &controllers.TaskController{})
	beego.Router("/task/::Mac", &controllers.TaskController{})
	beego.Router("/currentday/", &controllers.CurrentController{})
	beego.Router("/statistical/:TaskID", &controllers.StatisticalConroller{})
	beego.Router("/statisticalEx/", &controllers.StatisticalConrollerEx{})
	beego.Router("/USB/:USBOper", &controllers.USBController{})
}
