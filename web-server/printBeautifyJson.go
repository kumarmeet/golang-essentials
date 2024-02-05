package utils

import (
	"encoding/json"
	"fmt"
)

func PrintBeautifyJson(data interface{}) {
	jsonData, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(jsonData))
}
