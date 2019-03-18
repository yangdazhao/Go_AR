package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"go_AR/models"
)

type CompanyController struct {
	beego.Controller
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

	//user1 := models.TaskInfo{Id:1}
	//_ = o.Read(&user1)

	//var qs orm.QuerySeter
	//user := new([]models.Company)
	var users []*models.Company
	//qs =
	fmt.Println(o.QueryTable(new(models.Company)).Filter("id", 1))
	_, _ = o.QueryTable(new(models.Company)).All(&users)
	//.Filter("Com_ID", "Com_ID").All(&user)
	//_, _ = qs.Filter("id", "1").All(&user)
	c.Data["Website"] = "Auto Declare"
	c.Data["buffer"] = users
	c.Data["Email"] = "yangdazhao@live.com"
	c.TplName = "task.tpl"
}
