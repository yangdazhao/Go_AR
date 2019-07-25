package models

import (
	"encoding/xml"
	"io"
	_ "time"
)

type Param map[string]string

type xmlMapEntry struct {
	XMLName xml.Name
	Id      string `xml:"id,attr"`
	Value   string `xml:",innerxml"`
}

func (m Param) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(m) == 0 {
		return nil
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	for k, v := range m {
		_ = e.Encode(xmlMapEntry{XMLName: xml.Name{Local: "Input"}, Id: k, Value: v})
	}

	return e.EncodeToken(start.End())
}

func (m *Param) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*m = Param{}
	for {
		var e xmlMapEntry

		err := d.Decode(&e)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		(*m)[e.Id] = e.Value
	}
	return nil
}

type TableSetEx map[string]*XMLTableSet

func (m TableSetEx) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(m) == 0 {
		return nil
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	for _, v := range m {
		_ = e.Encode(v)
	}

	return e.EncodeToken(start.End())
}

func (m *TableSetEx) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*m = TableSetEx{}
	for {
		var e XMLTableSet
		err := d.Decode(&e)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		(*m)[e.TsId] = &e
	}
	return nil
}

type XMLCompanyInfo struct {
	CompanyName string
	CreditCode  string
	TaxCode     string
	TaxpayerId  string
}

type XMLTask struct {
	Id string `xml:"id,attr" Json:"id"`
	// TableSet TableSetEx 	`xml:"TableSet"`
}

type XMLTaskLogion struct {
	Id    int64
	TsId  string `xml:"id,attr" json:"id,omitempty"`
	Code  string `xml:"Result>Code"`
	Desc  string `xml:"Result>Desc"`
	Param *Param `xml:"Param" json:"Param,omitempty" orm:"-"`
}

type TableSetNx struct {
	TaskLogin *XMLTaskLogion
	TableSet  map[string]*XMLTableSet
}

func (m *TableSetNx) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*m = TableSetNx{}
	// m.TableSet = map[string]*XMLTableSet
	// (*m).TableSet = map[string]*XMLTableSet
	
	// ///////////////////////////////////////////////////////////////////////////
	var taskLogin XMLTaskLogion
	err := d.Decode(&taskLogin)
	if err == io.EOF {
	} else if err != nil {
		return err
	}

	(*m).TaskLogin = &taskLogin
	// //////////////////////////////////////////////////////////////////////////////
	for {
		var e XMLTableSet
		err := d.Decode(&e)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		(*m).TableSet[e.TsId] = &e
	}
	return nil
}

type XMLTaskSet struct {
	CompanyName string     `xml:"CompanyInfo>CompanyName"`
	CreditCode  string     `xml:"CompanyInfo>CreditCode"`
	TaxCode     string     `xml:"CompanyInfo>TaxCode"`
	TaxpayerId  string     `xml:"CompanyInfo>TaxpayerId"`
	Task        TableSetEx `xml:"Task" json:"TableSet"`
}

type XMLRoot struct {
	XMLName xml.Name   `xml:"Root"`
	TaskSet XMLTaskSet `xml:"TaskSet"`
}

type XMLTaskSetNx struct {
	CompanyName string      `xml:"CompanyInfo>CompanyName"`
	CreditCode  string      `xml:"CompanyInfo>CreditCode"`
	TaxCode     string      `xml:"CompanyInfo>TaxCode"`
	TaxpayerId  string      `xml:"CompanyInfo>TaxpayerId"`
	Task        *TableSetNx `xml:"Task" json:"TableSet"`
}
type XMLRootNx struct {
	XMLName xml.Name      `xml:"Root"`
	TaskSet *XMLTaskSetNx `xml:"TaskSet"`
}
