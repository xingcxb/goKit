package urlKit

import (
	"github.com/xingcxb/goKit/core/strKit"
	"net/url"
	"strings"
)

// UrlEncode URL编码
/**
 * @param str 待编码字符串
 * @return 编码结果
 */
func UrlEncode(str string) string {
	index := strings.Index(str, "?")
	if index > 0 {
		// 如果是一个完整的连接，那么需要将参数提取出来处理
		// 参数部分
		paramStr := str[index+1:]
		// 链接部分
		urlStr := str[:index]
		// 参数部分进行编码
		mapParams := strKit.StrParamsToMapParams(paramStr)
		newMapParams := make(map[string]string, 0)
		for k, v := range mapParams {
			newMapParams[k] = url.QueryEscape(v)
		}
		newParamsStr := strKit.MapParamsToUrlParams(newMapParams)
		// 拼接
		return strKit.Splicing(urlStr, "?", newParamsStr)
	} else {
		// 如果不是一个完整的连接，那么直接进行编码
		return url.QueryEscape(str)
	}
}

// UrlDecode URL解码
/**
 * @param str 待解码字符串
 * @return 解码结果
 */
func UrlDecode(str string) (string, error) {
	index := strings.Index(str, "?")
	if index > 0 {
		// 如果是一个完整的连接，那么需要将参数提取出来处理
		// 参数部分
		paramStr := str[index+1:]
		// 链接部分
		urlStr := str[:index]
		// 参数部分进行解码
		mapParams := strKit.StrParamsToMapParams(paramStr)
		newMapParams := make(map[string]string, 0)
		for k, v := range mapParams {
			nv, err := url.QueryUnescape(v)
			if err != nil {
				return "", err
			}
			newMapParams[k] = nv
		}
		newParamsStr := strKit.MapParamsToUrlParams(newMapParams)
		// 拼接
		return strKit.Splicing(urlStr, "?", newParamsStr), nil
	} else {
		// 如果不是一个完整的连接，那么直接进行解码
		return url.QueryUnescape(str)
	}
}
