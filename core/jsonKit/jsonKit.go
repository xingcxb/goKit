package jsonKit

import (
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"strings"
)

// FindContainingJSONByValue 通过json的值查找包含的json
/*
 * @param jsonStr json字符串(并不验证是否为json)
 * @param value 值
 * @param childrenName 子节点名称
 */
func FindContainingJSONByValue(jsonStr, findKey, findValue string, childrenName string) string {
	founds := "[]"
	jsonParse := gjson.Parse(jsonStr)
	found := false
	jsonParse.ForEach(func(key, value gjson.Result) bool {
		// 解析顶层的值看是否存在
		if strings.Contains(value.Get(findKey).String(), findValue) {
			founds, _ = sjson.Set(founds, "-1", value.Value())
			found = true
		}
		// 检查子节点中的值是否存在
		if children := value.Get(childrenName); children.Exists() {
			children.ForEach(func(_, child gjson.Result) bool {
				if strings.Contains(child.Get(findKey).String(), findValue) {
					founds, _ = sjson.Set(founds, "-1", value.Value())
				}
				return true
			})
		}
		return true
	})
	if found {
		return founds
	} else {
		return ""
	}
}
