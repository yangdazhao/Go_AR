package Tracing

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"go_AR/models"
	"io/ioutil"
)

type TaskQueryController struct {
	beego.Controller
	// AuthController
}

func init() {
	orm.Debug = false
}

type QueryParam struct {
	CompanyName  string `json:"companyname"`
	TaxpayerId   string `json:"taxpayerid"`
	SerialNumber string `json:"serialnumber"`
	TableSetId   string `json:"TableSetId"`
}

func QueryByCompanyName(param QueryParam) []*models.TaskInfo {
	var company models.Company
	var tasks []*models.TaskInfo
	var table orm.QuerySeter
	table = orm.NewOrm().QueryTable(new(models.TaskInfo))
	table = table.OrderBy("-created")

	if len(param.CompanyName) > 0 || len(param.TaxpayerId) > 0 {
		fmt.Println(param.CompanyName)
		var _Company orm.QuerySeter
		_Company = orm.NewOrm().QueryTable(new(models.Company))
		if len(param.CompanyName) > 0 {
			_Company = _Company.Filter("CompanyName", param.CompanyName)
		}

		if len(param.TaxpayerId) > 0 {
			_Company = _Company.Filter("TaxpayerId", param.TaxpayerId)
		}

		_ = _Company.One(&company)
		table = table.Filter("Company", company.Id)
	}

	if len(param.SerialNumber) > 0 {
		fmt.Println(param.SerialNumber)
		table = table.Filter("SerialNumber", param.SerialNumber)
	}

	_, err := table.RelatedSel().Limit(100).All(&tasks)
	if err != nil {

	}
	return tasks
}

func (c *TaskQueryController) Post() {
	bBody, _ := ioutil.ReadAll(c.Ctx.Request.Body)
	var p QueryParam
	err := json.Unmarshal(bBody, &p)
	if err != nil {

	}
	c.Data["json"] = QueryByCompanyName(p)
	c.ServeJSON()
}

func (c *TaskQueryController) Get() {
	ope := c.Ctx.Input.Param(":TaskID")
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "yangdazhao@live.com"
	c.Data["Param"] = ope
	c.TplName = "QueryTask.tpl"
}
