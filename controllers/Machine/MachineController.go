package Mac

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/hashicorp/consul/api"
)

type MachineController struct {
	beego.Controller
	// AuthController
}

type Machine struct {
	IP         string `json:"IP"`
	Mac        string `json:"Mac"`
	Time       string `json:"Time"`
	Uid        string `json:"uid"`
	TaxpayerId string `json:"TaxpayerId"`
	TaskType   string `json:"TaskType"`
	TableSetId string `json:"TableSetId"`
}

func (c *MachineController) Post() {
	//fmt.Println("MachineController")
	config := api.DefaultConfig()
	config.Address = "https://cabinet.bigfintax.com"
	client, _ := api.NewClient(config)
	kv := client.KV()
	data, _, _ := kv.List("Mac", nil)

	var machines []Machine
	for _, v := range data {
		//fmt.Println(string(v.Value))
		var mac Machine
		_ = json.Unmarshal(v.Value, &mac)
		kp, _, err := kv.Get("Task/"+mac.Mac, nil)
		//fmt.Printf("kp:%v", kp)
		if err == nil && kp != nil {
			//var mac1 Machine
			_ = json.Unmarshal(kp.Value, &mac)
			fmt.Printf("%v", mac)
		} else {
			// fmt.Println(string(data.Value))
		}

		//		//var mac1 Machine
		//		//_ = json.Unmarshal(data.Value, &mac1)
		machines = append(machines, mac)
	}
	c.Data["json"] = machines
	c.ServeJSON()
}

func (c *MachineController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "yangdazhao@live.com"
	// c.Data["Param"] = ope
	c.TplName = "MachineView.tpl"
}
