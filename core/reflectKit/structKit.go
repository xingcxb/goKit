// Package reflectKit 反射工具包
package reflectKit

import "reflect"

// StructToMap 结构体转换为map[string]string
// @param request 待转换的结构体
// @return map[string]string，由于value中类型的可能性太多了，为了防止出口后太麻烦统一为string
func StructToMap(body interface{}) map[string]string {
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
