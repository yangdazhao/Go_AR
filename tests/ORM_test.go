package test

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"go_AR/models"
	"testing"
)

func Test_ORM(t *testing.T) {
	var qParam QueryParam
	_ = json.Unmarshal([]byte(`{
	"CompanyName":"玩具反斗城（中国）商贸有限公司威海经区九龙城店",
	"TaxpayerId":"91371000MA3N6WYW1J",
	"serialnumber":"252d340cb084421b8cb4c9af23f29309"
}`), &qParam)
	QueryByCompanyName(qParam)
}
