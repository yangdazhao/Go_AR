package controllers

import (
	"fmt"
	"go_AR/controllers/Login"
)

type StatisticalNxConroller struct {
	Login.AuthController
	// beego.Controller
}

func (c *StatisticalNxConroller) Post() {
	Group := c.Ctx.Input.Query("Group")
	if len(Group) != 0 {
		fmt.Print(Group)
	}

	jsonResult := make(map[string]interface{})
	jsonResult["title"] = map[string]string{"subtext": "数据来自申报网关","text":"雨量流量关系图"}
	jsonResult["tooltip"] = map[string]string{"trigger": "axis"}
	c.Data["json"] = jsonResult
	c.ServeJSON()
}