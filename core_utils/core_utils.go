package core_utils

import (
	"encoding/json"
	"fmt"
)

func ConvertMapToStruct(source map[string]interface{} , destination interface{}) {
	jsonStr , err := json.Marshal(source)
	if err != nil {
		return
	}
	if err = json.Unmarshal(jsonStr , &destination); err != nil {
		fmt.Println(err.Error())
		return
	}
}