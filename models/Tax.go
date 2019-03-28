package models

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type TS700100 struct {
	Id   int64 `json:"id" pk:"auto"    orm:"column(id)"`
	Type string
}

type TS70030003 struct {
	Ts700100 TS700100
}

type TS70030002 struct {
	Ts700100 TS700100
}

type TS70030001 struct {
	Ts700100 TS700100
}
