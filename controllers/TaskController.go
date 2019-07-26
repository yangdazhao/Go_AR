package controllers

import (
	"JsonEx"
	"Task"
	"encoding/json"
	"fmt"
	"github.com/antchfx/xmlquery"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"go_AR/models"
	"os"
	"path/filepath"
	"strings"

	//"net/http"
	//"ORM/models"
	"time"
)

type TaskController struct {
	beego.Controller
	// AuthController
}

func TaskHandler(task Task.Task, Mac string) {
	orm.Debug = false
	o := orm.NewOrm()
	_ = o.Using("default")

	node, _ := xmlquery.Parse(strings.NewReader(task.Data))
	Company := xmlquery.FindOne(node, "//TaskSet/CompanyInfo")
	Company.SelectElement("CompanyName")
	fmt.Printf("%v", Company.SelectElement("CompanyName").InnerText())

	comEx := new(models.Company)
	com := new(models.Company)
	qc := o.QueryTable(comEx)
	if qc.Filter("TaxpayerId", Company.SelectElement("TaxpayerId").InnerText()).Exist() {
		_ = qc.Filter("TaxpayerId", Company.SelectElement("TaxpayerId").InnerText()).One(com)
	} else {
		com.CompanyName = Company.SelectElement("CompanyName").InnerText()
		com.TaxpayerId = Company.SelectElement("TaxpayerId").InnerText()
		fmt.Println(o.Insert(com))
	}

	InputSet := xmlquery.Find(node, "//TableSet[@id='TaskLogin']/Param/Input")
	for _, value := range InputSet {
		if len(value.InnerText()) > 0 {
			info := new(models.LoginInfo)
			info.Company = com
			info.Key = value.SelectAttr("id")
			info.Value = value.InnerText()
			fmt.Println(o.Insert(info))
		}
	}
	
	// tlResult := xmlquery.FindOne(node, "//TableSet[@id='TaskLogin']/Result")
	TableSet := xmlquery.FindOne(node, "//TableSet[@id!='TaskLogin']")

	taskInfo := new(models.TaskInfo)
	taskInfo.Company = com
	taskInfo.TaskID = TableSet.Parent.SelectAttr("id")
	taskInfo.TableSetID = TableSet.SelectAttr("id")
	taskInfo.SsqType = TableSet.SelectAttr("ssqType")
	taskInfo.Ssqs = TableSet.SelectAttr("ssqs")
	taskInfo.Ssqz = TableSet.SelectAttr("ssqz")
	taskInfo.Type = TableSet.SelectAttr("type")
	taskInfo.Submit = TableSet.SelectAttr("submit")
	taskInfo.SerialNumber = task.SerialNumber
	taskInfo.Env = task.Env
	taskInfo.Uuid = task.Uuid
	taskInfo.Mac = Mac
	o.Insert(taskInfo)
}

func TaskHandlerUpdate(task Task.Task, Mac string, JsonFileName string) {
	orm.Debug = false
	o := orm.NewOrm()
	_ = o.Using("default")

	node, _ := xmlquery.Parse(strings.NewReader(task.Data))
	Company := xmlquery.FindOne(node, "//TaskSet/CompanyInfo")
	Company.SelectElement("CompanyName")
	fmt.Printf("%v", Company.SelectElement("CompanyName").InnerText())

	comEx := new(models.Company)
	com := new(models.Company)
	qc := o.QueryTable(comEx)
	if qc.Filter("TaxpayerId", Company.SelectElement("TaxpayerId").InnerText()).Exist() {
		_ = qc.Filter("TaxpayerId", Company.SelectElement("TaxpayerId").InnerText()).One(com)
	} else {
		com.CompanyName = Company.SelectElement("CompanyName").InnerText()
		com.TaxpayerId = Company.SelectElement("TaxpayerId").InnerText()
		fmt.Println(o.Insert(com))
	}

	InputSet := xmlquery.Find(node, "//TableSet[@id='TaskLogin']/Param/Input")
	for _, value := range InputSet {
		if len(value.InnerText()) > 0 {
			info := new(models.LoginInfo)
			info.Company = com
			info.Key = value.SelectAttr("id")
			info.Value = value.InnerText()
			fmt.Println(o.Insert(info))
		}
	}

	tlResult := xmlquery.FindOne(node, "//TableSet[@id='TaskLogin']/Result")
	TableSet := xmlquery.FindOne(node, "//TableSet[@id!='TaskLogin']")

	taskInfo := new(models.TaskInfo)
	o.QueryTable(new(models.TaskInfo)).Filter("uuid", task.Uuid).One(taskInfo)
	taskInfo.Company = com
	taskInfo.LoginResult = tlResult.SelectElement("Code").InnerText()
	taskInfo.LoginDesc = tlResult.SelectElement("Desc").InnerText()
	taskInfo.TableSetID = TableSet.SelectAttr("id")
	taskInfo.SsqType = TableSet.SelectAttr("ssqType")
	taskInfo.Ssqs = TableSet.SelectAttr("ssqs")
	taskInfo.Ssqz = TableSet.SelectAttr("ssqz")
	taskInfo.Type = TableSet.SelectAttr("type")
	taskInfo.Submit = TableSet.SelectAttr("submit")
	tsResult := TableSet.SelectElement("Result")
	taskInfo.TsResult = tsResult.SelectElement("Code").InnerText()
	taskInfo.TsDesc = strings.Trim(tsResult.SelectElement("Desc").InnerText(), " \r\n")
	taskInfo.Mac = Mac
	taskInfo.Updated = time.Now()
	taskInfo.TaskJson = JsonFileName
	taskInfo.Env = task.Env
	taskInfo.Message = task.Message
	taskInfo.Status = task.Status
	o.Update(taskInfo)
}

func CreateDateDir(basePath string, subPath string) string {
	folderName := subPath
	if len(subPath) == 0 {
		folderName = time.Now().Format("200601")
	}
	folderPath := filepath.Join(basePath, folderName)
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// 必须分成两步
		_ = os.Mkdir(folderPath, 0777) // 先创建文件夹
		_ = os.Chmod(folderPath, 0777) // 再修改权限
	}
	return folderPath
}

func SaveEx(Dir string, task Task.Task) string {
	JsonParse := JsonEx.NewJsonStruct()
	taskDir := Dir + "\\" + task.CA.TaxCode + "_" + task.Status + "_" + time.Now().Format("20060102T150405") + ".json"
	fmt.Println(taskDir)

	node, _ := xmlquery.Parse(strings.NewReader(task.Data))
	TaskId := xmlquery.FindOne(node, "//TaskSet/Task/@id")
	fmt.Println(TaskId.Attr)
	JsonParse.Save(taskDir, task)
	return taskDir
}

func (c *TaskController) Post() {
	ope := c.Ctx.Input.Param(":Mac")
	var task Task.Task
	if json.Unmarshal(c.Ctx.Input.RequestBody, &task) == nil {
		TaskHandler(task, ope)
	}
	c.Data["json"] = ""
	c.ServeJSON()
}

func (c *TaskController) Put() {
	ope := c.Ctx.Input.Param(":Mac")
	var task Task.Task
	if json.Unmarshal(c.Ctx.Input.RequestBody, &task) == nil {
		CreateDateDir(".\\", "Json")
		CreateDateDir(".\\Json\\", ope)
		TaskJson := SaveEx(CreateDateDir(".\\Json\\"+ope, ""), task)
		TaskHandlerUpdate(task, ope, TaskJson)
	}
	c.Data["json"] = ""
	c.ServeJSON()
}

func Query(c *TaskController, currentDay time.Time) {
	fmt.Println(currentDay.Format("2006-01-02 00:00:00")) // 打印结果：2017-04-11 12:52:52.794351777 +0800 CST
	orm.Debug = false
	o := orm.NewOrm()
	_ = o.Using("default")
	var tasks []*models.TaskInfo
	filter := o.QueryTable(new(models.TaskInfo)).Filter("created__gt", currentDay).OrderBy("-created").RelatedSel()
	total, _ := filter.Count()
	_, _ = filter.Limit(500).All(&tasks)
	// fmt.Println(tasks)
	c.Data["Website"] = "Auto Declare"
	c.Data["total"] = total
	c.Data["tasks"] = tasks
	c.Data["Email"] = "yangdazhao@live.com"
	c.TplName = "task.tpl"
}

func QueryEx(c *TaskController, currentDay time.Time, TaxpayerId string) {
	fmt.Println(currentDay.Format("2006-01-02 00:00:00")) // 打印结果：2017-04-11 12:52:52.794351777 +0800 CST
	orm.Debug = false

	var company models.Company
	var tasks []*models.TaskInfo
	_ = orm.NewOrm().QueryTable(new(models.Company)).Filter("TaxpayerId", TaxpayerId).One(&company)
	filter := orm.NewOrm().QueryTable(new(models.TaskInfo)).Filter("Company", company.Id).Filter("created__gt", currentDay).OrderBy("-created").RelatedSel()
	
	// .RelatedSel().OrderBy("-created")
	// filter := o.QueryTable(new(models.TaskInfo)).Filter("created__gt", currentDay).OrderBy("-created").RelatedSel()
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
	t1 := time.Now().Year()  // 年
	t2 := time.Now().Month() // 月
	t3 := time.Now().Day()   // 日
	var currentDay time.Time
	if len(taskid) != 0 {
		fmt.Println(currentDay.Format("2006-01-02 00:00:00")) // 打印结果：2017-04-11 12:52:52.794351777 +0800 CST
		orm.Debug = false
		o := orm.NewOrm()
		_ = o.Using("default")
		var tasks []*models.TaskInfo
		filter := o.QueryTable(new(models.TaskInfo)).Filter("TaskID", taskid).OrderBy("-created").RelatedSel()
		total, _ := filter.Count()
		fmt.Print(filter.Count())
		_, _ = filter.Limit(200).All(&tasks)
		// fmt.Println(tasks)
		c.Data["Website"] = "Auto Declare"
		c.Data["total"] = total
		c.Data["tasks"] = tasks
		c.Data["Email"] = "yangdazhao@live.com"
		c.TplName = "task.tpl"
	} else {
		if len(ope) == 0 {
			currentDay = time.Date(t1, t2, t3, 0, 0, 0, 0, time.Local) // 获取当前时间，返回当前时间Time
			Query(c, currentDay)
			return
		} else if ope == "day" {
			currentDay = time.Date(t1, t2, t3, 0, 0, 0, 0, time.Local) // 获取当前时间，返回当前时间Time
			if len(TaxpayerId) == 0 {
				Query(c, currentDay)
			} else {
				QueryEx(c, currentDay, TaxpayerId)
			}
			return
		} else if ope == "month" {
			currentDay = time.Date(t1, t2, 0, 0, 0, 0, 0, time.Local) // 获取当前时间，返回当前时间Time
			if len(TaxpayerId) == 0 {
				Query(c, currentDay)
			} else {
				QueryEx(c, currentDay, TaxpayerId)
			}
			return
		}
	}
}
