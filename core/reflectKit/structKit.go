// Package reflectKit 反射工具包
package reflectKit

import (
	"encoding/json"
	"reflect"
)

// StructToMapSS 结构体转换为map[string]string
// @param request 待转换的结构体
// @return 通常在参数上使用
func StructToMapSS(body interface{}) map[string]string {
	t := reflect.TypeOf(body)
	v := reflect.ValueOf(body)
	var newMap = make(map[string]string)
	for k := 0; k < t.NumField(); k++ {
		key := t.Field(k).Name
		value := v.Field(k).String()
		newMap[key] = value
	}
	return newMap
}

// StructToMapSI 结构体转换为map[string]interface
// @param v 待转换的结构体
// @return map[string]interface
func StructToMapSI(v interface{}) map[string]interface{} {
	data, _ := json.Marshal(v)
	m := make(map[string]interface{})
	json.Unmarshal(data, &m)
	return m
}
