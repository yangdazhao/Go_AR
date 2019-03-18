package models

type UsbCabinet struct {
	IP    string `form:"IP"     json:"IP"       binding:"required"`
	CPort string `form:"CPort"  json:"CPort"    binding:"required"`
	DPort string `form:"DPort"  json:"DPort"`
	UPort string `form:"UPort"  json:"UPort"`
}

type PortInfo struct {
	CompanyName string
	TaxCode     string
	UPort       string
}

type ResultEx struct {
	Code string
}

type JsonResult struct {
	CAInfo  []PortInfo
	Result  ResultEx
	CAStats map[string]string
}
