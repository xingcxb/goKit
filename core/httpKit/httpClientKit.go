package httpKit

import (
	"goKit/core/strKit"
	"io"
	"net/http"
	"strings"
	"time"
)

// HttpGet 发送get请求
// @param urlString 网址
func HttpGet(urlString string) (string, error) {
	return HttpGetFull(urlString, nil, nil, "", -1)
}

// HttpGetFull 发送get请求[完整版]
// @param urlString 网址
// @param headers header信息
// @param paramMap post表单数据
// @param body body数据
// @param timeout 超时时长，-1表示默认超时，单位毫秒
func HttpGetFull(urlString string, headers, paramMap map[string]string, body string, timeout int) (string, error) {
	return HttpBasic(urlString, http.MethodGet, headers, paramMap, body, timeout)
}

// HttpPostFull 发送post请求[完整版]
// @param urlString 网址
// @param headers header信息
// @param paramMap post表单数据
// @param body body数据
// @param timeout 超时时长，-1表示默认超时，单位毫秒
func HttpPostFull(urlString string, headers, paramMap map[string]string, body string, timeout int) (string, error) {
	return HttpBasic(urlString, http.MethodGet, headers, paramMap, body, timeout)
}

// HttpBasic 发送http请求[基础]
// @param urlString 网址
// @param httpMethod http请求方法
// @param headers header信息
// @param paramMap post表单数据
// @param body body数据
// @param timeout 超时时长，-1表示默认超时，单位毫秒
func HttpBasic(urlString string, httpMethod string, headers, paramMap map[string]string, body string, timeout int) (string, error) {
	urlParam := strKit.MapParamsToUrlParams(paramMap)
	urlString = strKit.Splicing(urlString, urlParam)
	bodyReader := strings.NewReader(body)
	req, _ := http.NewRequest(httpMethod, urlString, bodyReader)
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	if timeout != -1 {
		http.DefaultClient.Timeout = time.Duration(timeout) * time.Millisecond
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	respByte, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	result := string(respByte)
	return result, nil
}
