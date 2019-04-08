package models

import "github.com/astaxie/beego/orm"

func init() {
	orm.RegisterModel(new(TaskId), new(TaskTaxpayer))
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type TaskId struct {
	Id     int64  `json:"id" 				   	orm:"column(id)"`
	Taskid string `json:"Value"            	orm:"column(TaskID)"`
	Name   string `json:"Code"            	orm:"column(Name)"`
	Count  int64  `json:"Desc"            	orm:"column(Count)"`
}

func (u *TaskId) TableName() string {
	return "statistical_taskid"
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type TaskTaxpayer struct {
	Id     int64  `json:"id" 				   	orm:"column(id)"`
	Taskid string `json:"Value"            	orm:"column(TaskID)"`
	Name   string `json:"Code"            	orm:"column(Name)"`
	Count  int64  `json:"Desc"            	orm:"column(Count)"`
}

func (u *TaskTaxpayer) TableName() string {
	return "statictical_taskid_taxpayerid"
}
