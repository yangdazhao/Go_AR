package controllers

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"go_AR/controllers/Login"
	"go_AR/models"
)

type StatisticalConroller struct {
	Login.AuthController
}

type JSONS struct {
	Categories []string `json:"categories" `
	Data       []int64  `json:"data" `
	Name       string   `json:"name" `
	Title      string   `json:"title" `
}

func (c *StatisticalConroller) Post() {
	param := c.Ctx.Input.Param(":TaskID")

	normal := map[string]interface{}{
		"show":     true,
		"position": "right",
		//"color":  	"#f0f",
	}

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

		title := make(map[string]interface{})
		title["subtext"] = "数据来自申报网关"
		title["text"] = "2019年4月份任务数量"

		tooltip := make(map[string]interface{})
		tooltip["trigger"] = "axis"
		tooltip["axisPointer"] = map[string]string{"type": "shadow"}

		yAxis := make(map[string]interface{})
		yAxis["type"] = "category"
		yAxis["data"] = categories

		jsonResult := make(map[string]interface{})
		jsonResult["title"] = title
		jsonResult["tooltip"] = tooltip
		jsonResult["xAxis"] = map[string]string{"type": "value"}
		jsonResult["yAxis"] = yAxis
		jsonResult["grid"] = map[string]string{
			"left":         "3%",
			"right":        "4%",
			"bottom":       "3%",
			"containLabel": "true",
		}
		jsonResult["series"] = []interface{}{
			map[string]interface{}{
				"type":  "bar",
				"name":  "任务数量",
				"data":  srcData,
				"label": map[string]interface{}{"normal": normal},
			},
		}

		c.Data["json"] = jsonResult
	} else if param == "taxpayer" {

		//////////////////////////////////////////////////
		type Temp struct {
			Total string `json:"TOTAL" orm:"column(TOTAL)"`
		}

		var temp Temp
		// tasks []models.Tax_success
		o := orm.NewOrm()
		//var maps []orm.Params
		fmt.Print(
			o.Raw(`SELECT COUNT(1) as 'TOTAL' FROM (
		SELECT
			COUNT(1)
		FROM
			taskinfo AS T 
		WHERE
			date_format( T.created, '%Y%m' ) = date_format( curdate(), '%Y%m' ) 
			AND T.TsResult LIKE '200C3%' 
			AND T.Env = 'pro'
		GROUP BY T.Com_ID
		) as w
		`).QueryRow(&temp))

		fmt.Println("TOTAL: ",temp)

		////////////////////////////////////////////////////

		var tasks []*models.TaskTaxpayer
		filter := orm.NewOrm().QueryTable(new(models.TaskTaxpayer))
		filter.OrderBy("TaskID").All(&tasks)

		// 预分配足够多的元素切片
		srcData := make([]int64, len(tasks))
		categories := make([]string, len(tasks))
		// 将切片赋值
		for i := 0; i < len(tasks); i++ {
			srcData[i] = tasks[i].Count
			categories[i] = tasks[i].Name
		}

		title := make(map[string]interface{})
		title["subtext"] = "数据来自申报网关"
		//i := maps[0]["TOTAL"](valu)
		title["text"] = "2019年4月份税号数量 " + temp.Total

		tooltip := make(map[string]interface{})
		tooltip["trigger"] = "axis"
		tooltip["axisPointer"] = map[string]string{"type": "shadow"}

		yAxis := make(map[string]interface{})
		yAxis["type"] = "category"
		yAxis["data"] = categories

		jsonResult := make(map[string]interface{})
		jsonResult["title"] = title
		jsonResult["tooltip"] = tooltip
		jsonResult["xAxis"] = map[string]string{"type": "value"}
		jsonResult["yAxis"] = yAxis
		jsonResult["grid"] = map[string]string{
			"left":         "3%",
			"right":        "4%",
			"bottom":       "3%",
			"containLabel": "true",
		}

		jsonResult["series"] = []interface{}{
			map[string]interface{}{
				"type":  "bar",
				"name":  "任务数量",
				"data":  srcData,
				"label": map[string]interface{}{"normal": normal},
			},
		}

		c.Data["json"] = jsonResult
	}
	c.ServeJSON()
}

func (c *StatisticalConroller) Get() {
	ope := c.Ctx.Input.Param(":TaskID")
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "yangdazhao@live.com"
	c.Data["Param"] = ope
	c.TplName = "echarts.tpl"
}
