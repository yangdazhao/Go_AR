package Tracing

import (
	"SplunkRestFul"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"strings"
)

type LogQueryController struct {
	beego.Controller
	// AuthController
}

func (c *LogQueryController) Post() {
	bBody, _ := ioutil.ReadAll(c.Ctx.Request.Body)
	var qParam SplunkRestFul.JobParam
	_ = json.Unmarshal(bBody, &qParam)
	query := SplunkRestFul.NewSplunkQuery("dazhao.yang", "dazhao.yang")
	sid := query.SubmitJob(qParam.ToString())
	logSet := query.QueryResults(sid)
	for k, v := range logSet {
		logSet[k] = strings.Replace(strings.Replace(v, qParam.SerialNumber+" ", "", 1), "fw=", "", 1)
	}
	c.Data["json"] = logSet
	c.ServeJSON()
}
