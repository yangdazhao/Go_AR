package controllers

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"go_AR/controllers/Login"
	"go_AR/models"
	"time"
)

type CurrentController struct {
	Login.AuthController
}

func QueryC(c *CurrentController, currentDay time.Time) {
	fmt.Println(currentDay.Format("2006-01-02 00:00:00")) // 打印结果：2017-04-11 12:52:52.794351777 +0800 CST
	orm.Debug = false
	o := orm.NewOrm()
	_ = o.Using("default")
	var tasks []*models.Currentday
	filter := o.QueryTable(new(models.Currentday)).OrderBy("-created").RelatedSel()
	total, _ := filter.Count()
	_, _ = filter.All(&tasks)
	fmt.Println(tasks)
	c.Data["Website"] = "Auto Declare"
	c.Data["total"] = total
	c.Data["tasks"] = tasks
	c.Data["Email"] = "yangdazhao@live.com"
	c.TplName = "view.tpl"
}

func (c *CurrentController) Get() {
	t1 := time.Now().Year()  // 年
	t2 := time.Now().Month() // 月
	t3 := time.Now().Day()   // 日
	var currentDay time.Time

	currentDay = time.Date(t1, t2, t3, 0, 0, 0, 0, time.Local) // 获取当前时间，返回当前时间Time
	QueryC(c, currentDay)
}
