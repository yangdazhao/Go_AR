package models

import (
	"github.com/astaxie/beego/orm"
)

type XMLCell struct {
	Id     int64
	CellId string    `xml:"id,attr,omitempty" json:"id,omitempty" orm:"column(CellId)"`
	Name   string    `xml:"name,attr,omitempty" json:"name,omitempty"`
	Index  string    `xml:"index,attr,omitempty" json:"index,omitempty"`
	Token  string    `xml:"token,attr,omitempty"`
	Value  string    `xml:",chardata" orm:"column(Value)"`
	Float  *XMLFloat `orm:"rel(fk)"`
}

func (u *XMLCell) TableName() string {
	return "XMLCell"
}

type XMLFloat struct {
	Id    int64
	Name  string     `xml:"name,attr" json:"id,omitempty" orm:"column(Name)" `
	Index string     `xml:"index,attr" json:"Index,omitempty" orm:"column(Index)"`
	Input []*XMLCell `xml:"Input" orm:"reverse(many)" `
	Check []*XMLCell `xml:"Check" orm:"reverse(many)"`
}

func (task *XMLFloat) Serialize() {
	o := orm.NewOrm()
	task.Id, _ = o.Insert(task)
	for _, v := range task.Input {
		v.Float = task
	}

	for _, v := range task.Check {
		v.Float = task
	}

	_, _ = o.InsertMulti(len(task.Input), task.Input)
	_, _ = o.InsertMulti(len(task.Check), task.Check)
}

func (task *XMLFloat) TableName() string {
	return "XMLFloat"
}
