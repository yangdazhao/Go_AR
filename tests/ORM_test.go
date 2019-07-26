package test

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"go_AR/models"
	"testing"
)

type QueryParam struct{
	CompanyName 	string	`json:"CompanyName"`
	TaxpayerId		string	`json:"TaxpayerId"`
	SerialNumber	string	`json:"serialnumber"`
	TableSetId		string	`json:"TableSetId"`
}

func QueryByCompanyName(param QueryParam) {
	fmt.Println(param)
	orm.Debug = false
	var company models.Company
	var tasks []*models.TaskInfo
	var table orm.QuerySeter
	table = orm.NewOrm().QueryTable(new(models.TaskInfo))

	if len(param.CompanyName) > 0 || len(param.TaxpayerId) > 0 {
		fmt.Println(param.CompanyName)
		var _Company orm.QuerySeter
		_Company = orm.NewOrm().QueryTable(new(models.Company))
		if len(param.CompanyName) > 0{
			_Company = _Company.Filter("CompanyName", param.CompanyName)
		}

		if len(param.TaxpayerId) > 0{
			_Company = _Company.Filter("TaxpayerId", param.TaxpayerId)
		}

		_ = _Company.One(&company)
		table = table.Filter("Company", company.Id)
	}

	if len(param.SerialNumber) > 0 {
		fmt.Println(param.SerialNumber)
		table = table.Filter("SerialNumber", param.SerialNumber)
	}

	_, _ = table.RelatedSel().Limit(100).All(&tasks)

	for k, v := range tasks {
		fmt.Println(k)
		fmt.Println(v)
	}
}

func Test_ORM(t *testing.T) {
	var qParam QueryParam
	_ = json.Unmarshal([] byte (`{
	"CompanyName":"玩具反斗城（中国）商贸有限公司威海经区九龙城店",
	"TaxpayerId":"91371000MA3N6WYW1J",
	"serialnumber":"252d340cb084421b8cb4c9af23f29309"
}`), &qParam)
	QueryByCompanyName(qParam)
}