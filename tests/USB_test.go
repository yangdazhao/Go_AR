package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

type UsbCabinet struct {
	IP    string `form:"IP"     json:"IP"       binding:"required"`
	CPort string `form:"CPort"  json:"CPort"    binding:"required"`
	DPort string `form:"DPort"  json:"DPort"`
	UPort string `form:"UPort"  json:"UPort"`
}

// //https://cabinet.bigfintax.com/USB/CabinetInfo
func TestCabiInfo(t *testing.T){
	 usb := UsbCabinet{"182.18.75.50","10002","3242","1"}
	bytesData, err := json.Marshal(usb)
	if err != nil {
		fmt.Println(err.Error() )
		return
	}
	reader := bytes.NewReader(bytesData)

	//生成要访问的url
	url := "https://cabinet.bigfintax.com/USB/CabinetInfo"
	//url := "http://127.0.0.1:8088/USB/CabinetInfo"
	resp, err := http.Post(url, "application/x-www-form-urlencoded", reader)

	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	}
	fmt.Println(string(body))
}

func TestCaInfo(t *testing.T){
	usb := UsbCabinet{"182.18.75.50","10002","3242","1"}
	bytesData, err := json.Marshal(usb)
	if err != nil {
		fmt.Println(err.Error() )
		return
	}
	reader := bytes.NewReader(bytesData)
	
	// 生成要访问的url
	url := "https://cabinet.bigfintax.com/USB/CaInfo"
	// url := "http://127.0.0.1:8081/USB/CabinetInfo"
	resp, err := http.Post(url, "application/x-www-form-urlencoded", reader)

	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	}
	fmt.Println(string(body))
}

func TestReboot(t *testing.T){
	usb := UsbCabinet{"182.18.75.50","10002","3242","1"}
	bytesData, err := json.Marshal(usb)
	if err != nil {
		fmt.Println(err.Error() )
		return
	}
	reader := bytes.NewReader(bytesData)
	
	// 生成要访问的url
	url := "https://cabinet.bigfintax.com/USB/Reboot"
	// url := "http://127.0.0.1:8081/USB/CabinetInfo"
	resp, err := http.Post(url, "application/x-www-form-urlencoded", reader)

	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	}
	fmt.Println(string(body))
}

func TestClose(t *testing.T){
	usb := UsbCabinet{"182.18.75.50","10002","3242","1"}
	bytesData, err := json.Marshal(usb)
	if err != nil {
		fmt.Println(err.Error() )
		return
	}
	reader := bytes.NewReader(bytesData)
	
	// 生成要访问的url
	url := "https://cabinet.bigfintax.com/USB/Close"
	// url := "http://127.0.0.1:8081/USB/CabinetInfo"
	resp, err := http.Post(url, "application/x-www-form-urlencoded", reader)

	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	}
	fmt.Println(string(body))
}