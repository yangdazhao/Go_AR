package models

import "github.com/astaxie/beego/orm"

func init() {
	orm.RegisterModel(new(TaskId), new(TaskTaxpayer),new(Tax_success))
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type TaskId struct {
	Id     int64  `json:"id" 				orm:"column(id)"`
	Taskid string `json:"Value"            	orm:"column(TaskID)"`
	Name   string `json:"Code"            	orm:"column(Name)"`
	Count  int64  `json:"Desc"            	orm:"column(Count)"`
}

func (u *TaskId) TableName() string {
	return "statistical_taskid"
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type TaskTaxpayer struct {
	Id     int64  `json:"id" 				orm:"column(id)"`
	Taskid string `json:"Value"            	orm:"column(TaskID)"`
	Name   string `json:"Code"            	orm:"column(Name)"`
	Count  int64  `json:"Desc"            	orm:"column(Count)"`
}

func (u *TaskTaxpayer) TableName() string {
	return "statictical_taskid_taxpayerid"
}

type Tax_success struct {
	Id 			int64 `orm:"column(id)"`
	TaskID 		string `orm:"column(TaskID)"`
	Name 		string `orm:"column(Name)"`
	Ts010100 	int64 `orm:"column(010100)"`
	Ts010200 	int64 `orm:"column(010200)"`
	Ts010300 	int64 `orm:"column(010300)"`
	Ts040101 	int64 `orm:"column(040101)"`
	Ts040102 	int64 `orm:"column(040102)"`
	Ts070100 	int64 `orm:"column(070100)"`
	Ts700100	int64 `orm:"column(700100)"`
	Ts700200 	int64 `orm:"column(700200)"`
	Ts700300 	int64 `orm:"column(700300)"`
	Ts090100 	int64 `orm:"column(090100)"`
	Ts910300	int64 `orm:"column(910300)"`
	Ts910200	int64 `orm:"column(910200)"`
	Ts170100	int64 `orm:"column(170100)"`
	Ts180100	int64 `orm:"column(180100)"`
	Ts620100	int64 `orm:"column(620100)"`
	Ts550001	int64 `orm:"column(550001)"`
	Ts550002	int64 `orm:"column(550002)"`
	Ts550003	int64 `orm:"column(550003)"`
	Ts910100	int64 `orm:"column(910100)"`
}

func (u *Tax_success) TableName() string {
	return "statictical_success"
}


