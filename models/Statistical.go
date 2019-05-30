package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(TaskId), new(TaskTaxpayer), new(TaxSuccess))
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type TaskId struct {
	Id     int64  `json:"id" orm:"column(id)"`
	Taskid string `json:"Value" orm:"column(TaskID)"`
	Name   string `json:"Code"  orm:"column(Name)"`
	Count  int64  `json:"Desc"  orm:"column(Count)"`
}

func (u *TaskId) TableName() string {
	return "statistical_taskid"
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type TaskTaxpayer struct {
	Id     int64  `json:"id" orm:"column(id)"`
	Taskid string `json:"Value" orm:"column(TaskID)"`
	Name   string `json:"Code"  orm:"column(Name)"`
	Count  int64  `json:"Desc"  orm:"column(Count)"`
}

func (u *TaskTaxpayer) TableName() string {
	return "statictical_taskid_taxpayerid"
}

type TaxSuccess struct {
	Id       int64  `orm:"column(id)"`
	TaskID   string `orm:"column(TaskID)"`
	Name     string `orm:"column(Name)"`
	Ts010100 int64  `orm:"column(010100)"`
	Ts010200 int64  `orm:"column(010200)"`
	Ts010300 int64  `orm:"column(010300)"`
	Ts040101 int64  `orm:"column(040101)"`
	Ts040102 int64  `orm:"column(040102)"`
	Ts070100 int64  `orm:"column(070100)"`
	Ts700100 int64  `orm:"column(700100)"`
	Ts700200 int64  `orm:"column(700200)"`
	Ts700300 int64  `orm:"column(700300)"`
	Ts090100 int64  `orm:"column(090100)"`
	Ts910300 int64  `orm:"column(910300)"`
	Ts910200 int64  `orm:"column(910200)"`
	Ts170100 int64  `orm:"column(170100)"`
	Ts180100 int64  `orm:"column(180100)"`
	Ts620100 int64  `orm:"column(620100)"`
	Ts550001 int64  `orm:"column(550001)"`
	Ts550002 int64  `orm:"column(550002)"`
	Ts550003 int64  `orm:"column(550003)"`
	Ts910100 int64  `orm:"column(910100)"`
}

func (u *TaxSuccess) TableName() string {
	return "statictical_success"
}

func QueryInfoByGroup(Group string, Type string) []TaxSuccess {
	var tasks [] TaxSuccess
	// tasks []models.Tax_success
	o := orm.NewOrm()
	num, err := o.Raw(`SELECT
	t3.id AS id,
	t3.taxid AS TaskID,
	t3.name AS Name,
	sum( t2.010100 ) AS "010100",
	sum( t2.010200 ) AS "010200",
	sum( t2.010300 ) AS "010300",
	sum( t2.040101 ) AS "040101",
	sum( t2.040102 ) AS "040102",
	sum( t2.070100 ) AS "070100",
	sum( t2.700100 ) AS "700100",
	sum( t2.700200 ) AS "700200",
	sum( t2.700300 ) AS "700300",
	sum( t2.090100 ) AS "090100",
	sum( t2.910300 ) AS "910300",
	sum( t2.910200 ) AS "910200",
	sum( t2.170100 ) AS "170100",
	sum( t2.180100 ) AS "180100",
	sum( t2.620100 ) AS "620100",
	sum( t2.550001 ) AS "550001",
	sum( t2.550002 ) AS "550002",
	sum( t2.550003 ) AS "550003",
	sum( t2.910100 ) AS "910100" 
FROM
	(
		(
		SELECT
			t1.TaskID AS TaskID,
			( CASE WHEN ( t1.TsId = '010100' ) THEN t1.num END ) AS '010100',
			( CASE WHEN ( t1.TsId = '010200' ) THEN t1.num END ) AS '010200',
			( CASE WHEN ( t1.TsId = '010300' ) THEN t1.num END ) AS '010300',
			( CASE WHEN ( t1.TsId = '040101' ) THEN t1.num END ) AS '040101',
			( CASE WHEN ( t1.TsId = '040102' ) THEN t1.num END ) AS '040102',
			( CASE WHEN ( t1.TsId = '070100' ) THEN t1.num END ) AS '070100',
			( CASE WHEN ( t1.TsId = '700100' ) THEN t1.num END ) AS '700100',
			( CASE WHEN ( t1.TsId = '700200' ) THEN t1.num END ) AS '700200',
			( CASE WHEN ( t1.TsId = '700300' ) THEN t1.num END ) AS '700300',
			( CASE WHEN ( t1.TsId = '090100' ) THEN t1.num END ) AS '090100',
			( CASE WHEN ( t1.TsId = '910300' ) THEN t1.num END ) AS '910300',
			( CASE WHEN ( t1.TsId = '910200' ) THEN t1.num END ) AS '910200',
			( CASE WHEN ( t1.TsId = '170100' ) THEN t1.num END ) AS '170100',
			( CASE WHEN ( t1.TsId = '180100' ) THEN t1.num END ) AS '180100',
			( CASE WHEN ( t1.TsId = '620100' ) THEN t1.num END ) AS '620100',
			( CASE WHEN ( t1.TsId = '550001' ) THEN t1.num END ) AS '550001',
			( CASE WHEN ( t1.TsId = '550002' ) THEN t1.num END ) AS '550002',
			( CASE WHEN ( t1.TsId = '550003' ) THEN t1.num END ) AS '550003',
			( CASE WHEN ( t1.TsId = '910100' ) THEN t1.num END ) AS '910100' 
		FROM
			(
			SELECT
				t.TaskID AS TaskID,
				t.TsId AS TsId,
				count( 1 ) AS num 
			FROM
				(
				SELECT
					t1.TaskID AS TaskID,
					t1.TaxpayerId AS TaxpayerId,
					t1.TsId AS TsId,
					t1.Type AS type,
					t1.Ssqs AS Ssqs 
				FROM
					currentmonth t1 JOIN logininfoex J
				WHERE
					(
						( t1.TsResult LIKE '200C%' ) 
						AND ( t1.Submit = '1' ) 
						AND ( t1.Env = 'pro' ) 
						AND ( t1.Type = ? ) 
						and T1.TaxpayerId = j.LoginNsrsbh
						and j.GroupName = ?
					) 
				GROUP BY
					t1.TaskID,
					t1.TaxpayerId,
					t1.TsId,
					t1.Type,
					t1.Ssqs 
				ORDER BY
					t1.TaskID,
					t1.TaxpayerId,
					t1.TsId,
					t1.Type,
					t1.Ssqs 
				) t 
			GROUP BY
				t.TaskID,
				t.TsId 
			) t1 
		) t2
		JOIN taxinfo t3 
	) 
WHERE
	( LEFT ( t2.TaskID, 4 ) = t3.taxid ) 
GROUP BY
	LEFT ( t2.TaskID, 4 ) 
ORDER BY
	t3.taxid
`, Type, Group).QueryRows(&tasks)
	if err == nil {
		fmt.Println("user nums: ", num)
	}
	return tasks
}
