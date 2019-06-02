package test

import (
	"encoding/xml"
	"fmt"
	"go_AR/models"
	"testing"
)

func TestName(t *testing.T) {
	_Param := make(map[string]string)
	_Param["LoginName"] = "912101000987214892"
	_Param["id"] = "1"

	buf, _ := xml.Marshal(models.Param(_Param))
	fmt.Println(string(buf))

	stringMap := make(map[string]string)
	err := xml.Unmarshal(buf, (*models.Param)(&stringMap))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(stringMap)
}
