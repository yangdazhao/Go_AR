package controllers

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"go_AR/controllers/Login"
	"go_AR/models"
)

type CompanyController struct {
	Login.AuthController
}

func (c *CompanyController) Post() {
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
	c.Data["Email"] = ZuoZheEmail
	c.TplName = "task.tpl"
}
