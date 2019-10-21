package Tracing

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/yangdazhao/SplunkRestFul"
	"io/ioutil"
	"strings"
)

type LogQueryController struct {
	beego.Controller
	// AuthController
}

func (c *LogQueryController) Post() {
	bBody, _ := ioutil.ReadAll(c.Ctx.Request.Body)
	var p SplunkRestFul.JobParam
	_ = json.Unmarshal(bBody, &p)
	query := SplunkRestFul.NewSplunkQuery("dazhao.yang", "dazhao.yang")
	sid := query.SubmitJob(p.ToString())
	logSet := query.QueryResults(sid)
	for k, v := range logSet {
		logSet[k] = strings.Replace(strings.Replace(v, p.SerialNumber+" ", "", 1), "fw=", "", 1)
	}
	c.Data["json"] = logSet
	c.ServeJSON()
}
