package controllers

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"go_AR/models"
)

type CompanyController struct {
	AuthController
}

func (c *CompanyController) Post() {
	//var ob models.Task
	//_ = json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	//fmt.Println(ob.SerialNumber)
	c.TplName = "index.tpl"
}

func (c *CompanyController) Get() {
	orm.Debug = false
	o := orm.NewOrm()
	_ = o.Using("default")

	var users []*models.Company
	fmt.Println(o.QueryTable(new(models.Company)).Filter("id", 1))
	_, _ = o.QueryTable(new(models.Company)).All(&users)
	c.Data["Website"] = "Auto Declare"
	c.Data["buffer"] = users
	c.Data["Email"] = "yangdazhao@live.com"
	c.TplName = "task.tpl"
}
