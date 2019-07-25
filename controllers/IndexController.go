package controllers

import (
	"github.com/astaxie/beego/orm"
    "go_AR/controllers/Login"
    "go_AR/models"
)

type IndexController struct {
    Login.AuthController
}

type JSON3 struct {
	Categories []string `json:"categories" `
	Data       []int64  `json:"data" `
	Name       string   `json:"name" `
	Title      string   `json:"title" `
}

func (c *IndexController) Post() {
	param := c.Ctx.Input.Param(":TaskID")

	if param == "task" {
		var tasks []*models.TaskId
		filter := orm.NewOrm().QueryTable(new(models.TaskId))
		_, _ = filter.OrderBy("TaskID").All(&tasks)

		// 预分配足够多的元素切片
		srcData := make([]int64, len(tasks))
		categories := make([]string, len(tasks))
		// 将切片赋值
		for i := 0; i < len(tasks); i++ {
			srcData[i] = tasks[i].Count
			categories[i] = tasks[i].Name
		}
		data := &JSONS{categories, srcData, "任务数量", "2019年4月份任务数量"}
		c.Data["json"] = data
	} else if param == "taxpayer" {
		var tasks []*models.TaskTaxpayer
		filter := orm.NewOrm().QueryTable(new(models.TaskTaxpayer))
		_, _ = filter.OrderBy("TaskID").All(&tasks)

		// 预分配足够多的元素切片
		srcData := make([]int64, len(tasks))
		categories := make([]string, len(tasks))
		// 将切片赋值
		for i := 0; i < len(tasks); i++ {
			srcData[i] = tasks[i].Count
			categories[i] = tasks[i].Name
		}
		data := &JSONS{categories, srcData, "税号数量", "2019年4月份税号数量"}
		c.Data["json"] = data
	}

	c.ServeJSON()
}

func (c *IndexController) Get() {
	ope := c.Ctx.Input.Param(":TaskID")
	c.Data["Website"] = "beego.me"

	c.Data["Email"] = "yangdazhao@live.com"
	c.Data["Param"] = ope
	c.TplName = "index.tpl"
}
