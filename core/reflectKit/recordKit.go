// Package reflectKit 主要用于生成一个Record
package reflectKit

import "encoding/json"

type Record struct {
	data map[string]interface{}
}

// Get 查询记录中指定字段的值
func (r *Record) Get(column string) interface{} {
	return r.data[column]
}

// Set 修改记录中指定字段的值
func (r *Record) Set(column string, value interface{}) {
	r.data[column] = value
}

// Remove 删除记录中指定字段
func (r *Record) Remove(column string) {
	delete(r.data, column)
}

// ToJSON 将记录转换为 JSON 格式的字符串
func (r *Record) ToJSON() (string, error) {
	jsonBytes, err := json.Marshal(r.data)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}
