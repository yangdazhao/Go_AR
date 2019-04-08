package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"go_AR/models"
)

type StatisticalConroller struct {
	beego.Controller
}

type JSONS struct {
	Categories []string `json:"categories" `
	Data       []int64  `json:"data" `
	Name       string   `json:"name" `
	Title      string   `json:"title" `
}

func (c *StatisticalConroller) Post() {
	var tasks []*models.TaskTaxpayer
	filter := orm.NewOrm().QueryTable(new(models.TaskTaxpayer))
	fmt.Println(filter.Count())
	fmt.Println(filter.OrderBy("TaskID").All(&tasks))

	// 预分配足够多的元素切片
	srcData := make([]int64, len(tasks))
	categories := make([]string, len(tasks))
	// 将切片赋值
	for i := 0; i < len(tasks); i++ {
		srcData[i] = tasks[i].Count
		categories[i] = tasks[i].Name
	}
	data := &JSONS{categories, srcData, "税号数量", "2019年4月份税号务量"}
	c.Data["json"] = data
	c.ServeJSON()
}

func (c *StatisticalConroller) Get() {
	ope := c.Ctx.Input.Param(":TaskID")
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "yangdazhao@live.com"
	c.Data["TaskID"] = ope
	c.TplName = "echarts.tpl"
}
