package utils

import (
	"encoding/json"
	"fmt"
)

func PrettyStruct(data interface{}) string {
	val, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err.Error()
	}
	return string(val)
}

func I32ToSliceString(s []int32) []string {
	ret := []string{}
	for _, v := range s {
		ret = append(ret, fmt.Sprintf("%v", v))
	}
	return ret
}
