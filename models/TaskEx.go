package models

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
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
		e.Encode(xmlMapEntry{XMLName: xml.Name{Local: "Input"}, Id: k, Value: v})
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

type TableSetEx map[string]XMLTableSet

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
		(*m)[e.Id] = e
	}
	return nil
}

type xmlParam struct {
	//XMLName      xml.Name 				`xml:"Param"`
	Input []XMLCell  `xml:"Input"`
	Check []XMLCell  `xml:"Check"`
	Float []XMLFloat `xml:"Float"`
}

type XMLCompanyInfo struct {
	//XMLName      xml.Name 			`xml:"CompanyInfo"`
	CompanyName string
	CreditCode  string
	TaxCode     string
	TaxpayerId  string
}

type XMLTable struct {
	//XMLName		xml.Name 				`xml:"Table"`
	Id    string   `xml:"id,attr"`
	Name  string   `xml:"name,attr"`
	Page  string   `xml:"page,attr"`
	Code  string   `xml:"Result>Code"`
	Desc  string   `xml:"Result>Desc"`
	Param xmlParam `xml:"Param"`
}

type XMLTableSet struct {
	Id      string     `xml:"id,attr" json:"id,omitempty"`
	Type    string     `xml:"type,attr" json:"type,omitempty"`
	Ssqs    string     `xml:"ssqs,attr" json:"ssqs,omitempty"`
	Ssqz    string     `xml:"ssqz,attr" json:"ssqz,omitempty"`
	Submit  string     `xml:"submit,attr" json:"submit,omitempty"`
	SsqType string     `xml:"ssqType,attr" json:"ssqType,omitempty"`
	Se      string     `xml:"da,attr" json:"Se,omitempty"`
	Code    string     `xml:"Result>Code"`
	Desc    string     `xml:"Result>Desc"`
	Table   []XMLTable `xml:"Table" json:"Table,omitempty"`
	Param   Param      `xml:"Param" json:"Param,omitempty"`
}

type XMLTask struct {
	Id string `xml:"id,attr" Json:"id"`
	//TableSet TableSetEx 	`xml:"TableSet"`
}

type XMLTaskSet struct {
	CompanyName string `xml:"CompanyInfo>CompanyName"`
	CreditCode  string `xml:"CompanyInfo>CreditCode"`
	TaxCode     string `xml:"CompanyInfo>TaxCode"`
	TaxpayerId  string `xml:"CompanyInfo>TaxpayerId"`
	//Task        []XMLTask
	Task TableSetEx `xml:"Task" json:"TableSet"`
}

type XMLRoot struct {
	XMLName xml.Name   `xml:"Root"`
	TaskSet XMLTaskSet `xml:"TaskSet"`
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

func (u *XMLTableSet) TableName() string {
	return "TableSet"
}
