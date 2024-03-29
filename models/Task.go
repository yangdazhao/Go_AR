package models

import (
	"fmt"
	"github.com/antchfx/xmlquery"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"time"
)

func init() {
	/*
	   // PostgreSQL 配置
	   orm.RegisterDriver("postgres", orm.DR_Postgres) // 注册驱动
	   orm.RegisterDataBase("default", "postgres", "user=postgres password=tom dbname=test host=127.0.0.1 port=5432 sslmode=disable")
	*/

	/*** MySQL 配置
	 * 注册驱动*/
	_ = orm.RegisterDriver("mysql", orm.DRMySQL)
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")
	if dbport == "" {
		dbport = "3306"
	}

	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?loc=Local&charset=utf8"
	fmt.Println(dsn)
	//_ = orm.RegisterDataBase("default", "mysql", dsn)

	//* mysql用户：root ，root的秘密：tom ， 数据库名称：test ， 数据库别名：default
	_ = orm.RegisterDataBase("default", "mysql", "yangdazhao:7721@tcp(10.10.40.3:3306)/taskinfoex?loc=Local&charset=utf8")
	// _ = orm.RegisterDataBase("default", "mysql", "yangdazhao:7721@tcp(10.10.40.3:3306)/taskinfo?loc=Local&charset=utf8")

	orm.RegisterModel(new(Company), new(LoginInfo), new(TaskInfo), new(Currentday), new(Table70010004), new(TS700100))
	orm.RegisterModel(new(XMLCell), new(XMLFloat), new(XMLTable), new(XMLTableSet))
	// 自动建表
	_ = orm.RunSyncdb("default", false, false)
}

type Company struct {
	Id          int64  `json:"id" pk:"auto"    orm:"column(id)"`
	TaxpayerId  string `json:"TaxpayerId"      orm:"column(TaxpayerId);size(20);unique"`
	CompanyName string `json:"CompanyName"     orm:"column(CompanyName)"`
}

func NewCompanyEx(company * xmlquery.Node) *Company {
	return &Company{
		TaxpayerId:  company.SelectElement("TaxpayerId").InnerText(),
		CompanyName: company.SelectElement("CompanyName").InnerText(),
	}
}

func NewCompany(taxpayerId string, companyName string) *Company {
	return &Company{TaxpayerId: taxpayerId, CompanyName: companyName}
}

func (u *Company) TableName() string {
	return "Company"
}

type LoginInfo struct {
	Id      int64    `json:"id" pk:"auto"    orm:"column(id)"`
	Company *Company `json:"Value"           orm:"column(Com_ID);rel(fk);null"` // OneToOne relation
	Key     string   `json:"Key"             orm:"column(Key)"`
	Value   string   `json:"Value"           orm:"column(Value)"`
}

func NewLoginInfo(company *Company, input *xmlquery.Node) *LoginInfo {
	return &LoginInfo{Company: company, Key: input.SelectAttr("id"), Value: input.InnerText()}
}

// 多字段唯一键
func (u *LoginInfo) TableUnique() [][]string {
	return [][]string{
		{"Key", "Company"},
	}
}

func (u *LoginInfo) TableName() string {
	return "LoginInfo"
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type TaskInfo struct {
	Id           int64     `json:"id" pk:"auto"    orm:"column(id)"`
	Uuid         string    `json:"uuid"            orm:"column(uuid)"`
	Company      *Company  `json:"Company"         orm:"column(Com_ID);rel(fk);null"`                        // OneToOne relation
	TaskID       string    `json:"TaskID"          orm:"column(TaskID);"                 description:"税局ID"` // OneToOne relation
	LoginResult  string    `json:"Code"            orm:"column(LoginResult);null;size(20)"`
	LoginDesc    string    `json:"Desc"            orm:"column(LoginDesc)"`
	TableSetID   string    `json:"TableSetID"      orm:"column(TsId);null;size(6)"`
	SsqType      string    `json:"ssqType"         orm:"column(ssqType)"`
	Type         string    `json:"Type"            orm:"column(Type)"`
	Ssqs         string    `json:"Ssqs"            orm:"column(Ssqs)"`
	Ssqz         string    `json:"Ssqz"            orm:"column(Ssqz)"`
	Submit       string    `json:"Submit"          orm:"column(Submit)"`
	TsResult     string    `json:"Code"            orm:"column(TsResult)"`
	TsDesc       string    `json:"Desc"            orm:"column(TsDesc)"`
	SerialNumber string    `json:"SerialNumber"    orm:"column(SerialNumber)"`
	Env          string    `json:"Env"             orm:"column(Env)"`
	Mac          string    `json:"Mac"             orm:"column(ClientMac)"`
	Created      time.Time `orm:"auto_now_add;type(datetime)"`
	Updated      time.Time `orm:"auto_now;type(datetime)"`
	TaskJson     string    `json:"TaskJson"        orm:"column(TaskJson)"`
	Status       string    `json:"Status"          orm:"column(JsonResult)"`
	Message      string    `json:"Message"         orm:"column(JsonDesc)"`
	Se           string    `json:"Se"              orm:"column(Se)"`
}

func NewTaskInfo(
	uuid string,
	company *Company,
	TableSet *xmlquery.Node,
	serialNumber string,
	env string,
	mac string,
) *TaskInfo {
	return &TaskInfo{
		Uuid:         uuid,
		Company:      company,
		TaskID:       TableSet.Parent.SelectAttr("id"),
		TableSetID:   TableSet.SelectAttr("id"),
		SsqType:      TableSet.SelectAttr("ssqType"),
		Type:         TableSet.SelectAttr("type"),
		Ssqs:         TableSet.SelectAttr("ssqs"),
		Ssqz:         TableSet.SelectAttr("ssqz"),
		Submit:       TableSet.SelectAttr("submit"),
		SerialNumber: serialNumber,
		Env:          env,
		Mac:          mac,
	}
}

func (t *TaskInfo) UpdateData(
	tlResult *xmlquery.Node,
	tsResult *xmlquery.Node,
	JsonFileName string,
	Message string,
	Status string,
) {
	t.LoginResult = tlResult.SelectElement("Code").InnerText()
	t.LoginDesc = tlResult.SelectElement("Desc").InnerText()
	t.TsResult = tsResult.SelectElement("Code").InnerText()
	t.TsDesc = strings.Trim(tsResult.SelectElement("Desc").InnerText(), " \r\n")
	t.Updated = time.Now()
	t.TaskJson = JsonFileName
	t.Message = Message
	t.Status = Status
}

func (t *TaskInfo) TableName() string {
	return "TaskInfo"
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type Currentday struct {
	Id          int64     `json:"id"			  orm:"column(id)"`
	CompanyName string    `json:"Value"           orm:"column(CompanyName);"` // OneToOne relation
	TaxpayerId  string    `json:"Value"            orm:"column(TaxpayerId);"` // OneToOne relation
	LoginResult string    `json:"Code"            orm:"column(LoginResult);null;size(20)"`
	TableSetID  string    `json:"TableSetID"      orm:"column(TsId);null;size(6)"`
	Type        string    `json:"Type"            orm:"column(Type)"`
	Ssqs        string    `json:"Ssqs"            orm:"column(Ssqs)"`
	Ssqz        string    `json:"Ssqz"            orm:"column(Ssqz)"`
	TsResult    string    `json:"Code"            orm:"column(TsResult)"`
	TsDesc      string    `json:"Desc"            orm:"column(TsDesc)"`
	Mac         string    `json:"Mac"             orm:"column(ClientMac)"`
	Created     time.Time `orm:"auto_now_add;	type(datetime)"`
	Time        int64     `orm:"auto_now_add;"`
}

func (u *Currentday) TableName() string {
	return "currentday"
}
