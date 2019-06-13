package test

import (
	"JsonEx"
	"Task"
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego/orm"
	"go_AR/models"
	"strings"
	"testing"
	_ "time"
)

func TestTask070100(t *testing.T) {
	orm.Debug = true
	task := Task.Task{}
	var tJson = "..\\Json\\Task.3700301.91370100763687736M.010100.SB.200C000000.20190517164350.json"
	JsonEx.NewJsonStruct().LoadEx(tJson, &task)
	if strings.Index(task.Status, "200C") == 0 && task.Env == "pro" {
		root := new(models.XMLRoot)
		if err := xml.Unmarshal([]byte(task.Data), root); err != nil {
			fmt.Println(err)
			return
		}
		ts := root.TaskSet
		ts.Task["010100"].Serialize()
	}
}

func TestTask070100Nx(t *testing.T) {
	orm.Debug = true
	task := Task.Task{}
	var tJson = "..\\Json\\Task.3700301.91370100763687736M.010100.SB.200C000000.20190517164350.json"
	JsonEx.NewJsonStruct().LoadEx(tJson, &task)
	if strings.Index(task.Status, "200C") == 0 && task.Env == "pro" {
		root := new(models.XMLRootNx)
		if err := xml.Unmarshal([]byte(task.Data), root); err != nil {
			fmt.Println(err)
			return
		}

		for _, v := range root.TaskSet.Task.TableSet {
			v.Serialize()
		}
	}
}
