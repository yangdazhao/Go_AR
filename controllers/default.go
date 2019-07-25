package controllers

import (
	"fmt"
	"github.com/astaxie/beego/orm"
    "go_AR/controllers/Login"
    "go_AR/models"
)

type MainController struct {
    Login.AuthController
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "yangdazhao@live.com"
	c.TplName = "echarts.tpl"
}

func QueryInfoByGroup(Group string, Type string)[]models.TaxSuccess {
	var tasks []models.TaxSuccess
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
` , Type ,Group).QueryRows(&tasks)
//` , Type ,Group).QueryRows(&tasks)
	if err == nil {
		fmt.Println("user nums: ", num)
	}
	return tasks
}
