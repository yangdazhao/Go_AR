package routers

import (
	"github.com/astaxie/beego"
	"go_AR/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/Task/:Time/:TaxpayerId", &controllers.TaskController{})
	beego.Router("/Task/:Time/", &controllers.TaskController{})
	beego.Router("/Task/", &controllers.TaskController{})
	//beego.Router("/USB/:USBOper", &controllers.USBController{})
	//beego.Router("/task/:Mac", &controllers.TaskController{})
	//beego.Router("/task/", &controllers.TaskController{})
}
