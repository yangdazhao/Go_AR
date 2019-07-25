package models

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type TS700100 struct {
	Id   int64 `json:"id" pk:"auto"    orm:"column(id)"`
	Type string
}

func (u *TS700100) TableName() string {
	return "tableset.7001000"
}

type Table70010001 struct {
	Id   int64     `json:"id" pk:"auto"    orm:"column(id)"`
	TsId *TS700100 `json:"Value"           orm:"column(Ts_Id);rel(fk);null"` // OneToOne relation
}

type Table70010002 struct {
	Id   int64     `json:"id" pk:"auto"    orm:"column(id)"`
	TsId *TS700100 `json:"Value"           orm:"column(Ts_Id);rel(fk);null"` // OneToOne relation
}

type Table70010003 struct {
	Id   int64     `json:"id" pk:"auto"    orm:"column(id)"`
	TsId *TS700100 `json:"Value"           orm:"column(Ts_Id);rel(fk);null"` // OneToOne relation
}

type Table70010004 struct {
	Id   int64     `json:"id" pk:"auto"    orm:"column(id)"`
	TsId *TS700100 `json:"Value"           orm:"column(Ts_Id);rel(fk);null"` // OneToOne relation
	//Ts700100 TS700100
	/// 资产
	/// 流动资产
	Cash                 float64 // 货币资金
	ShortTermInvestments float64 // 短期投资

	NotesReceivable    float64 // 应收票据
	AccountsReceivable float64 // 应收账款
	AccountsPayment    float64 // 预付账款
	DividendReceivable float64 // 应收股利
	InterestReceivable float64 // 应收利息
	OtherReceivables   float64 // 其他应收款
	Stock              float64 // 存货
	ProductsIn         float64 // 在产品
	MerchandiseInStock float64 // 库存商品
	OtherCurrentAssets float64 // 其他流动资产
	TotalCurrentAssets float64 // 动资产合计
	// / 非流动资产
	TotalNonCurrentAssets float64 //非流动资产合计
	TotalAssets           float64 //资产总计
	/// 负债及所有者权益
	/// 负债
	TotalLiabilities float64 //负载合计
	/// 所有者权益
	TotalEquity               float64 /// 所有者权益合计
	TotalLiabilitiesAndEquity float64 /// 负债和所有者权益总计
}

func (u *Table70010004) TableName() string {
	return "table.70010004"
}
