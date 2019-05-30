package controllers

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"go_AR/models"
	"time"
)

type TaskController struct {
	AuthController
}

func (c *TaskController) Post() {
	fmt.Println(c.Ctx.Request.URL.Path)
	ope := c.Ctx.Input.Param(":Mac")
	fmt.Println(ope, c.Ctx.Request.URL.Path)
	c.Data["json"] = ""
	c.ServeJSON()
}

func Query(c *TaskController, currentDay time.Time) {
	fmt.Println(currentDay.Format("2006-01-02 00:00:00")) //打印结果：2017-04-11 12:52:52.794351777 +0800 CST
	orm.Debug = false
	o := orm.NewOrm()
	_ = o.Using("default")
	var tasks []*models.TaskInfo
	filter := o.QueryTable(new(models.TaskInfo)).Filter("created__gt", currentDay).OrderBy("-created").RelatedSel()
	total, _ := filter.Count()
	_, _ = filter.Limit(500).All(&tasks)
	//fmt.Println(tasks)
	c.Data["Website"] = "Auto Declare"
	c.Data["total"] = total
	c.Data["tasks"] = tasks
	c.Data["Email"] = "yangdazhao@live.com"
	c.TplName = "task.tpl"
}

func QueryEx(c *TaskController, currentDay time.Time, TaxpayerId string) {
	fmt.Println(currentDay.Format("2006-01-02 00:00:00")) //打印结果：2017-04-11 12:52:52.794351777 +0800 CST
	orm.Debug = false

	var company models.Company
	var tasks []*models.TaskInfo
	_ = orm.NewOrm().QueryTable(new(models.Company)).Filter("TaxpayerId", TaxpayerId).One(&company)
	filter := orm.NewOrm().QueryTable(new(models.TaskInfo)).Filter("Company", company.Id).Filter("created__gt", currentDay).OrderBy("-created").RelatedSel()

	//.RelatedSel().OrderBy("-created")
	//filter := o.QueryTable(new(models.TaskInfo)).Filter("created__gt", currentDay).OrderBy("-created").RelatedSel()
	total, _ := filter.Count()
	_, _ = filter.Limit(500).All(&tasks)
	c.Data["Website"] = "Auto Declare"
	c.Data["total"] = total
	c.Data["tasks"] = tasks
	c.Data["Email"] = "yangdazhao@live.com"
	c.TplName = "task.tpl"
}

func (c *TaskController) Get() {

	taskid := c.Ctx.Input.Query("taskid")
	fmt.Print(taskid)
	ope := c.Ctx.Input.Param(":Time")
	TaxpayerId := c.Ctx.Input.Param(":TaxpayerId")
	t1 := time.Now().Year()  //年
	t2 := time.Now().Month() //月
	t3 := time.Now().Day()   //日
	var currentDay time.Time
	if len(taskid) != 0 {
		fmt.Println(currentDay.Format("2006-01-02 00:00:00"))	 //打印结果：2017-04-11 12:52:52.794351777 +0800 CST
		orm.Debug = false
		o := orm.NewOrm()
		_ = o.Using("default")
		var tasks []*models.TaskInfo
		filter := o.QueryTable(new(models.TaskInfo)).Filter("TaskID",taskid).OrderBy("-created").RelatedSel()
		total, _ := filter.Count()
		fmt.Print(filter.Count())
		_, _ = filter.Limit(200).All(&tasks)
		//fmt.Println(tasks)
		c.Data["Website"] = "Auto Declare"
		c.Data["total"] = total
		c.Data["tasks"] = tasks
		c.Data["Email"] = "yangdazhao@live.com"
		c.TplName = "task.tpl"
	} else {
		if len(ope) == 0 {
			currentDay = time.Date(t1, t2, t3, 0, 0, 0, 0, time.Local) //获取当前时间，返回当前时间Time
			Query(c, currentDay)
			return
		} else if ope == "day" {
			currentDay = time.Date(t1, t2, t3, 0, 0, 0, 0, time.Local) //获取当前时间，返回当前时间Time
			if len(TaxpayerId) == 0 {
				Query(c, currentDay)
			} else {
				QueryEx(c, currentDay, TaxpayerId)
			}
			return
		} else if ope == "month" {
			currentDay = time.Date(t1, t2, 0, 0, 0, 0, 0, time.Local) //获取当前时间，返回当前时间Time
			if len(TaxpayerId) == 0 {
				Query(c, currentDay)
			} else {
				QueryEx(c, currentDay, TaxpayerId)
			}
			return
		}
	}
}
