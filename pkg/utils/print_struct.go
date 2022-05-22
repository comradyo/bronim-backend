package utils

import (
	"encoding/json"
	"fmt"
)

func PrintStruct(i interface{}) {
	s, _ := json.MarshalIndent(i, "", "\t")
	fmt.Println(string(s))
}

func StringStruct(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
