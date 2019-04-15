package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"go_AR/models"
)

type StatisticalConrollerEx struct {
	beego.Controller
}

type Series struct {
	Name  string  `json:"name" `
	Type  string  `json:"type" `
	Stack string  `json:"stack" `
	Data  []int64 `json:"data"`
}

type JSONS2 struct {
	Categories []string `json:"categories" `
	//Data       []int64  `json:"data" `
	Name       string   `json:"name" `
	Legend     []string   `json:"legend" `
	Title      string   `json:"title" `
	Series     []Series `json:"series" `
}
func (c *StatisticalConrollerEx) Post() {
	var tasks []*models.Tax_success
	filter := orm.NewOrm().QueryTable(new(models.Tax_success))
	_, _ = filter.OrderBy("TaskID").All(&tasks)
	fmt.Print(filter.Count())
	// 预分配足够多的元素切片
	//srcData := make([]int64, len(tasks))
	categories := make([]string, len(tasks))
	legend := make([]string, 9)
	legend[0] = "一般纳税人增值税"
	legend[1] = "小规模增值税"
	legend[2] = "财务报表一般企业会计制度"
	legend[3] = "财务报表小企业会计准则"
	legend[4] = "财务报表一般企业会计准则"
	legend[5] = "印花税"
	legend[6] = "附加税"
	legend[7] = "通用申报表"
	legend[8] = "企业所得税A类"

	Series := make([]Series, 9)

	Series[0].Name = "一般纳税人增值税"
	Series[0].Type = "line"
	Series[0].Data = make([]int64, len(tasks))

	Series[1].Name = "小规模增值税"
	Series[1].Type = "line"
	Series[1].Data = make([]int64, len(tasks))

	Series[2].Name = "财务报表一般企业会计制度"
	Series[2].Type = "line"
	Series[2].Data = make([]int64, len(tasks))

	Series[3].Name = "财务报表小企业会计准则"
	Series[3].Type = "line"
	Series[3].Data = make([]int64, len(tasks))

	Series[4].Name = "财务报表一般企业会计准则"
	Series[4].Type = "line"
	Series[4].Data = make([]int64, len(tasks))

	Series[5].Name = "印花税"
	Series[5].Type = "line"
	Series[5].Data = make([]int64, len(tasks))

	Series[6].Name = "附加税"
	Series[6].Type = "line"
	Series[6].Data = make([]int64, len(tasks))

	Series[7].Name = "通用申报表"
	Series[7].Type = "line"
	Series[7].Data = make([]int64, len(tasks))

	Series[8].Name = "企业所得税A类"
	Series[8].Type = "line"
	Series[8].Data = make([]int64, len(tasks))

	// 将切片赋值
	for i := 0; i < len(tasks); i++ {
		//srcData[i] = tasks[i]
		categories[i] = tasks[i].Name
		Series[0].Data[i] = tasks[i].Ts010100
		Series[1].Data[i] = tasks[i].Ts010200
		Series[2].Data[i] = tasks[i].Ts700100
		Series[3].Data[i] = tasks[i].Ts700200
		Series[4].Data[i] = tasks[i].Ts700300
		Series[5].Data[i] = tasks[i].Ts090100
		Series[6].Data[i] = tasks[i].Ts070100
		Series[7].Data[i] = tasks[i].Ts910100
		Series[8].Data[i] = tasks[i].Ts040101
	}

	data := &JSONS2{categories,  "任务数量", legend,"当月税种成功数量",Series}
	c.Data["json"] = data
	fmt.Print("123")
	c.ServeJSON()
}

func (c *StatisticalConrollerEx) Get() {
	ope := c.Ctx.Input.Param(":TaskID")
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "yangdazhao@live.com"
	c.Data["Param"] = ope
	c.TplName = "echarts2.tpl"
}
