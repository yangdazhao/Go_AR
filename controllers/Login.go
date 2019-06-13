package controllers

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/session"
	"go_AR/models"
)

type LoginControler struct {
	AuthController
}

var globalSessions *session.Manager

func init() {
	scf := &session.ManagerConfig{
		CookieName:      "gosessionid",
		EnableSetCookie: true,
		Gclifetime:      3600,
		Maxlifetime:     3600,
		Secure:          false,
		CookieLifeTime:  3600,
		ProviderConfig:  "./tmp",
	}

	globalSessions, _ = session.NewManager("memory", scf)
	go globalSessions.GC()
}

func (this *LoginControler) Post() {
	param := this.Ctx.Input.Param(":TaskID")

	if param == "task" {
		var tasks []*models.TaskId
		filter := orm.NewOrm().QueryTable(new(models.TaskId))
		filter.OrderBy("TaskID").All(&tasks)

		// 预分配足够多的元素切片
		srcData := make([]int64, len(tasks))
		categories := make([]string, len(tasks))
		// 将切片赋值
		for i := 0; i < len(tasks); i++ {
			srcData[i] = tasks[i].Count
			categories[i] = tasks[i].Name
		}
		data := &JSONS{categories, srcData, "任务数量", "2019年4月份任务数量"}
		this.Data["json"] = data
	} else if param == "taxpayer" {
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
		data := &JSONS{categories, srcData, "税号数量", "2019年4月份税号数量"}
		this.Data["json"] = data
	}

	//this.SetSecureCookie()

	this.ServeJSON()
}

func (this *LoginControler) Get() {
	var _, _ = this.GetSecureCookie("asta", "dgid")
	ope := this.Ctx.Input.Param(":TaskID")
	this.Data["Website"] = "beego.me"

	this.Data["Email"] = "yangdazhao@live.com"
	this.Data["Param"] = ope
	this.TplName = "index.tpl"
}

//this.SetSecureCookie()
