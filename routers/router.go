package routers

import (
	"github.com/astaxie/beego"
	"go_AR/controllers"
	"go_AR/controllers/Login"
	"go_AR/controllers/Machine"
	"go_AR/controllers/TaskHand"
	"go_AR/controllers/Tracing"
	"go_AR/controllers/USB"
)

func init() {
	beego.Router("/home/login", &Login.AuthController{}, "post:PostData;get:Login")
	beego.Router("/", &controllers.StatisticalConrollerEx{})
	beego.Router("/statisticalNx", &controllers.StatisticalNxConroller{})
	beego.Router("/index", &controllers.StatisticalConrollerEx{})
	beego.Router("/Task/:Time/:TaxpayerId", &TaskHand.TaskController{})
	beego.Router("/Task/:Time/", &TaskHand.TaskController{})
	beego.Router("/Task/:Time/", &TaskHand.TaskController{})
	beego.Router("/Task/", &TaskHand.TaskController{})
	beego.Router("/task/:Mac", &TaskHand.TaskController{})
	beego.Router("/currentday/", &controllers.CurrentController{})
	beego.Router("/statistical/:TaskID", &controllers.StatisticalConroller{})
	beego.Router("/statisticalEx/", &controllers.StatisticalConrollerEx{})
	beego.Router("/USB/:USBOper", &USB.USBController{})
	beego.Router("/TaskQuery", &Tracing.TaskQueryController{})
	beego.Router("/LogQuery", &Tracing.LogQueryController{})
	beego.Router("/Machine", &Mac.MachineController{})
}
