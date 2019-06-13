package controllers

import (
	"github.com/astaxie/beego/orm"
	"go_AR/models"
)

type StatisticalConrollerEx struct {
	//beego.Controller
	AuthController
}

type Series struct {
	Name  string                 `json:"name" `
	Type  string                 `json:"type" `
	Stack string                 `json:"stack" `
	Data  []int64                `json:"data" `
	Label map[string]interface{} `json:"label"`
	//"label": map[string]interface{}{"normal": normal},
}

func NewSeries(name string, Type string, stack string, length int, label map[string]interface{}) *Series {
	return &Series{Name: name, Type: Type, Stack: stack, Data: make([]int64, length), Label: label}
}

//noinspection ALL
//func NewSeries(name string, _type string, stack string, length int) *Series {
//	normal := map[string]interface{}{
//		"show":     true,
//		"position": "right",
//	}
//	return &Series{
//		Name: name,
//		Type: _type,
//		Stack: stack,
//		Data: make([]int64, length),
//		normal,
//	}
//}

type JSONS2 struct {
	Categories []string `json:"categories" `
	Name       string   `json:"name" `
	Legend     []string `json:"legend" `
	Title      string   `json:"title" `
	Series     []Series `json:"series" `
}

func (c *StatisticalConrollerEx) Post() {
	group := c.Ctx.Input.Query("group")
	var tasks []models.TaxSuccess
	if len(group) == 0 {
		filter := orm.NewOrm().QueryTable(new(models.TaxSuccess))
		_, _ = filter.OrderBy("TaskID").All(&tasks)
	} else {
		tasks = QueryInfoByGroup(group, "SB")
	}

	categories := make([]string, len(tasks))
	legend := []string{
		"一般纳税人增值税",
		"小规模增值税",
		"财务报表一般企业会计制度",
		"财务报表小企业会计准则",
		"财务报表一般企业会计准则",
		"印花税",
		"附加税",
		"通用申报表",
		"企业所得税A类",
	}

	normal := map[string]interface{}{
		//"show":     true,
		//"position": "top",
	}

	Series := make([]Series, 9)
	Series[0] = *NewSeries("一般纳税人增值税", "line", "", len(tasks), map[string]interface{}{"normal": normal})
	Series[1] = *NewSeries("小规模增值税", "line", "", len(tasks), map[string]interface{}{"normal": normal})
	Series[2] = *NewSeries("财务报表一般企业会计制度", "line", "", len(tasks), map[string]interface{}{"normal": normal})
	Series[3] = *NewSeries("财务报表小企业会计准则", "line", "", len(tasks), map[string]interface{}{"normal": normal})
	Series[4] = *NewSeries("财务报表一般企业会计准则", "line", "", len(tasks), map[string]interface{}{"normal": normal})
	Series[5] = *NewSeries("印花税", "line", "", len(tasks), map[string]interface{}{"normal": normal})
	Series[6] = *NewSeries("附加税", "line", "", len(tasks), map[string]interface{}{"normal": normal})
	Series[7] = *NewSeries("通用申报表", "line", "", len(tasks), map[string]interface{}{"normal": normal})
	Series[8] = *NewSeries("企业所得税A类", "line", "", len(tasks), map[string]interface{}{"normal": normal})

	// 将切片赋值
	for i := 0; i < len(tasks); i++ {
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

	data := &JSONS2{categories, "任务数量", legend, group + "当月税种成功数量", Series}
	c.Data["json"] = data
	c.ServeJSON()
}

func (c *StatisticalConrollerEx) Get() {
	ope := c.Ctx.Input.Param(":TaskID")
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "yangdazhao@live.com"
	c.Data["Param"] = ope
	c.TplName = "index.tpl"
}
