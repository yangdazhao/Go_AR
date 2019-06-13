package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
)

type XMLCell struct {
	Id     int64
	CellId string    `xml:"id,attr,omitempty" json:"id,omitempty" orm:"column(CellId)"`
	Name   string    `xml:"name,attr,omitempty" json:"name,omitempty"`
	Index  string    `xml:"index,attr,omitempty" json:"index,omitempty"`
	Token  string    `xml:"token,attr,omitempty"`
	Value  string    `xml:",chardata" orm:"column(Value)"`
	Float  *XMLFloat `orm:"null;rel(fk);on_delete(set_null)"`
	Table  *XMLTable `orm:"rel(fk);null;on_delete(set_null)"`
}

func (u *XMLCell) TableName() string {
	return "tax.Cell"
}

type XMLFloat struct {
	Id    int64
	Name  string     `xml:"name,attr" json:"id,omitempty" orm:"column(Name)" `
	Index string     `xml:"index,attr" json:"Index,omitempty" orm:"column(Index)"`
	Input []*XMLCell `xml:"Input" orm:"reverse(many)" `
	Check []*XMLCell `xml:"Check" orm:"reverse(many)"`
	Table *XMLTable  `orm:"rel(fk)"`
}

func (task *XMLFloat) Serialize() {
	o := orm.NewOrm()
	task.Id, _ = o.Insert(task)
	for _, v := range task.Input {
		v.Float = task
		v.Table = task.Table
	}

	for _, v := range task.Check {
		v.Float = task
		v.Table = task.Table
	}

	_, _ = o.InsertMulti(len(task.Input), task.Input)
	_, _ = o.InsertMulti(len(task.Check), task.Check)
}

func (task *XMLFloat) TableName() string {
	return "tax.Float"
}

type XMLTable struct {
	Id       int64
	TableId  string       `xml:"id,attr"`
	Name     string       `xml:"name,attr"`
	Page     string       `xml:"page,attr"`
	Code     string       `xml:"Result>Code"`
	Desc     string       `xml:"Result>Desc"`
	Input    []*XMLCell   `xml:"Param>Input" orm:"reverse(many)"`
	Check    []*XMLCell   `xml:"Param>Check" orm:"reverse(many)"`
	Float    []*XMLFloat  `xml:"Param>Float" orm:"reverse(many)"`
	TableSet *XMLTableSet `orm:"rel(fk)"`
}

func (task *XMLTable) TableName() string {
	return "tax.Table"
}

func (table *XMLTable) Serialize() {
	o := orm.NewOrm()
	table.Id, _ = o.Insert(table)

	for _, v := range table.Input {
		v.Table = table
	}

	for _, v := range table.Check {
		v.Table = table
	}

	fmt.Println(o.InsertMulti(len(table.Input), table.Input))
	_, _ = o.InsertMulti(len(table.Check), table.Check)

	for _, v := range table.Float {
		v.Table = table
		v.Serialize()
	}
}

type XMLTableSet struct {
	Id      int64
	TsId    string      `xml:"id,attr" json:"id,omitempty"`
	Type    string      `xml:"type,attr" json:"type,omitempty"`
	Ssqs    string      `xml:"ssqs,attr" json:"ssqs,omitempty"`
	Ssqz    string      `xml:"ssqz,attr" json:"ssqz,omitempty"`
	Submit  string      `xml:"submit,attr" json:"submit,omitempty"`
	SsqType string      `xml:"ssqType,attr" json:"ssqType,omitempty"`
	Se      string      `xml:"da,attr" json:"Se,omitempty"`
	Code    string      `xml:"Result>Code"`
	Desc    string      `xml:"Result>Desc"`
	Table   []*XMLTable `xml:"Table" json:"Table,omitempty" orm:"reverse(many)"`
	Param   *Param      `xml:"Param" json:"Param,omitempty" orm:"-"`
}

func (u *XMLTableSet) TableName() string {
	return "Tax.TableSet"
}

func (ts *XMLTableSet) Serialize() {
	o := orm.NewOrm()
	ts.Id, _ = o.Insert(ts)

	for _, v := range ts.Table {
		v.TableSet = ts
		v.Serialize()
	}
}

func (task *XMLTaskSet) String() string {
	b, err := json.Marshal(*task)
	if err != nil {
		return fmt.Sprintf("%+v", *task)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "  ")
	if err != nil {
		return fmt.Sprintf("%+v", *task)
	}
	return out.String()
}
