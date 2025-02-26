package jsonKit

import "github.com/tidwall/gjson"

// FindContainingJSONByValue 通过json的值查找包含的json
/*
 * @param jsonStr json字符串(并不验证是否为json)
 * @param value 值
 * @param childrenName 子节点名称
 */
func FindContainingJSONByValue(jsonStr, findKey, findValue string, childrenName string) string {
	jsonParse := gjson.Parse(jsonStr)
	var found gjson.Result
	jsonParse.ForEach(func(key, value gjson.Result) bool {
		// 解析顶层的值看是否存在
		if value.Get(findKey).String() == findValue {
			found = value
			return false
		}
		// 检查子节点中的值是否存在
		if children := value.Get(childrenName); children.Exists() {
			children.ForEach(func(_, child gjson.Result) bool {
				if child.Get(findKey).String() == findValue {
					found = child
					return false // 找到后停止遍历
				}
				return true
			})
		}
		return true
	})
	if found.Exists() {
		return found.String()
	} else {
		return ""
	}
}
