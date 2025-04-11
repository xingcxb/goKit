package retKit

import (
	"encoding/json"
	"fmt"
)

type Ret struct {
	data  map[string]interface{}
	state bool
}

const (
	StateOk   = true
	StateFail = false
)

func (r *Ret) init() {
	r.data = make(map[string]interface{})
}

// Set 设置返回的数据
func (r *Ret) Set(key string, value interface{}) *Ret {
	if r.data == nil {
		r.init()
	}
	r.data[key] = value
	return r
}

// Get 获取返回的数据
func (r *Ret) Get(key string) interface{} {
	return r.data[key]
}

// Delete 删除返回的数据
func (r *Ret) Delete(key string) *Ret {
	delete(r.data, key)
	return r
}

// Fail 设置返回的数据为失败状态
func (r *Ret) Fail() *Ret {
	r.state = StateFail
	return r
}

// Ok 设置返回的数据为成功状态
func (r *Ret) Ok() *Ret {
	r.state = StateOk
	return r
}

// IsOk 获取返回的数据是否正确
func (r *Ret) IsOk() bool {
	return r.state
}

// By 设置返回的数据为 key-value 对
func (r *Ret) By(key string, value interface{}) *Ret {
	return r.Set(key, value)
}

// As 转换为指定类型
/*
 * @param data interface{} 数据
 */
func As[T any](data interface{}) (T, error) {
	var result T
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return result, fmt.Errorf("failed to marshal data: %w", err)
	}
	err = json.Unmarshal(dataBytes, &result)
	if err != nil {
		return result, fmt.Errorf("failed to unmarshal data into target type: %w", err)
	}
	return result, nil
}

// ToJSON 将 Ret 结构体转换为 JSON 格式的字符串
func (r *Ret) ToJSON() (string, error) {
	jsonBytes, err := json.Marshal(r.data)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}
