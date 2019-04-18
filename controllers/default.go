package controllers

type MainController struct {
	AuthController
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "yangdazhao@live.com"
	c.TplName = "echarts.tpl"
}
