package controllers

import (
	"AR/usbip"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"go_AR/models"
	"golang.org/x/text/encoding/simplifiedchinese"
	"strconv"
)

type USBController struct {
	beego.Controller
}

func (c *USBController) Post() {
	ope := c.Ctx.Input.Param(":USBOper")

	fmt.Println(ope, c.Ctx.Request.URL.Path)
	var Cabinet models.UsbCabinet
	if len(ope) == 0 && c.Ctx.Request.URL.Path == "/USB/CabinetInfo" {
		if json.Unmarshal(c.Ctx.Input.RequestBody, &Cabinet) == nil {
			usb := usbip.NewUsbIP(Cabinet.IP + ":" + Cabinet.CPort)
			caStatus, _ := usb.Info()

			var JsonResult models.JsonResult
			JsonResult.Result.Code = "0"
			JsonResult.CAStats = caStatus
			c.Data["json"] = JsonResult
		}
	} else if ope == "CaInfo" {
		if json.Unmarshal(c.Ctx.Input.RequestBody, &Cabinet) == nil {
			usb := usbip.NewUsbIP(Cabinet.IP + ":" + Cabinet.CPort)
			infos := usb.CaInfo()
			var msg []models.PortInfo
			for _, ca := range infos.Info {
				if ca.HasRead != 0x00 {
					index := bytes.IndexByte(ca.Name[0:], 0)
					result, err := simplifiedchinese.GBK.NewDecoder().Bytes(ca.Name[0:index])
					if err != nil {
						panic(err)
					}
					msg = append(msg, models.PortInfo{
						CompanyName: string(result),
						TaxCode:     string(ca.TaxCode[0:bytes.IndexByte(ca.TaxCode[0:], 0)]),
						UPort:       strconv.Itoa(int(ca.Port))})
				}
			}

			var JsonResult models.JsonResult
			JsonResult.Result.Code = "0"
			JsonResult.CAInfo = msg
			c.Data["json"] = JsonResult
		}
	} else if ope == "CabinetInfo" {
		if json.Unmarshal(c.Ctx.Input.RequestBody, &Cabinet) == nil {
			usb := usbip.NewUsbIP(Cabinet.IP + ":" + Cabinet.CPort)
			caStatus, _ := usb.Info()

			var JsonResult models.JsonResult
			JsonResult.Result.Code = "0"
			JsonResult.CAStats = caStatus
			c.Data["json"] = JsonResult
		}
	} else if ope == "Reboot" {
		if json.Unmarshal(c.Ctx.Input.RequestBody, &Cabinet) == nil {
			usb := usbip.NewUsbIP(Cabinet.IP + ":" + Cabinet.CPort)
			var JsonResult models.JsonResult
			JsonResult.Result.Code = usb.Reboot()
			c.Data["json"] = JsonResult
		}
		return
	} else if ope == "Close" {
		if json.Unmarshal(c.Ctx.Input.RequestBody, &Cabinet) == nil {
			usb := usbip.NewUsbIP(Cabinet.IP + ":" + Cabinet.CPort)

			var JsonResult models.JsonResult
			JsonResult.Result.Code = string(usb.Close(1))
			c.Data["json"] = JsonResult
		}
	}
	c.ServeJSON()
}

func (c *USBController) Get() {
	orm.Debug = false
	o := orm.NewOrm()
	_ = o.Using("default")
	var tasks []*models.TaskInfo
	_, _ = o.QueryTable(new(models.TaskInfo)).OrderBy("-id").Limit(100).All(&tasks)
	c.Data["Website"] = "Auto Declare"
	c.Data["tasks"] = tasks
	c.Data["Email"] = "yangdazhao@live.com"
	c.TplName = "task.tpl"
}
