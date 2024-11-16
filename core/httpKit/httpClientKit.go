package httpKit

import (
	"bufio"
	"errors"
	"github.com/valyala/fasthttp"
	"github.com/xingcxb/goKit/core/strKit"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// HttpDownload 下载文件
/**
 * @param urlString 网址
 * @param savePath 保存路径 末尾是否携带/都可以
 * @param fileName 文件名，如果不存在则自动获取
 * @param isCover 是否覆盖 true 覆盖 false 不覆盖(当文件存在的时候返回该文件已存在)
 * @return string 文件路径,error
 */
func HttpDownload(urlString, savePath, fileName string, isCover bool) (string, error) {
	if savePath == "" {
		return "", errors.New("保存路径为空")
	}
	// 发起网络请求
	// 必须要优先请求的原因是使用的开发人员可能没有指定文件名称，需要从url中获取
	res, err := http.Get(urlString)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if fileName == "" {
		//定义文件名字
		path := strings.Split(urlString, "/")
		fileName = path[len(path)-1]
	}
	// 检查保存路径是否以/结尾
	if !strings.HasSuffix(savePath, "/") {
		savePath = strKit.Splicing(savePath, "/")
	}

	filePath := strKit.Splicing(savePath, fileName)
	// 判断文件是否存在，默认不存在
	checkFile := false
	// 判断文件是否存在
	if _, err := os.Stat(filePath); err == nil {
		checkFile = true
	}
	if isCover && checkFile {
		// 允许覆盖，删除文件
		err = os.Remove(filePath)
		if err != nil {
			return "", err
		}
	} else if isCover {
		// 不允许覆盖，返回错误
		if checkFile {
			return "", errors.New("文件已存在")
		}
	}
	//创建文件
	out, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	// defer延迟调用 关闭文件，释放资源
	defer out.Close()
	//添加缓冲 bufio 是通过缓冲来提高效率。
	wt := bufio.NewWriter(out)
	_, err = io.Copy(out, res.Body)
	if err != nil {
		return "", err
	}
	//将缓存的数据写入到文件中
	if wt.Flush() != nil {
		return "", err
	}
	return filePath, nil
}

// HttpGet 发送get请求
/*
 * @param urlString 网址
 * @return string 网页内容,error
 */
func HttpGet(urlString string) (string, error) {
	stream, err := HttpBasic(urlString, http.MethodGet, nil, nil, "", -1)
	return string(stream), err
}

// HttpGetFull 发送get请求[完整版]
/**
 * @param urlString 网址
 * @param headers header信息
 * @param paramMap post表单数据
 * @param body body数据
 * @param timeout 超时时长，-1表示默认超时，单位毫秒
 * @return string 网页内容,error
 */
func HttpGetFull(urlString string, headers, paramMap map[string]string, body string, timeout int) (string, error) {
	stream, err := HttpBasic(urlString, http.MethodGet, headers, paramMap, body, timeout)
	return string(stream), err
}

// HttpPost 发送post基础请求
/**
 * @param urlString 网址
 * @param params post表单数据
 * @return string 网页内容,error
 */
func HttpPost(urlString string, params map[string]string) (string, error) {
	stream, err := HttpBasic(urlString, http.MethodPost, nil, params, "", -1)
	return string(stream), err
}

// HttpPostFull 发送post请求[完整版]
/**
 * @param urlString 网址
 * @param headers header信息
 * @param paramMap post表单数据
 * @param body body数据
 * @param timeout 超时时长，-1表示默认超时，单位毫秒
 * @return string 网页内容,error
 */
func HttpPostFull(urlString string, headers, paramMap map[string]string, body string, timeout int) (string, error) {
	stream, err := HttpBasic(urlString, http.MethodPost, headers, paramMap, body, timeout)
	return string(stream), err
}

// HttpBasic 发送http请求[基础]
/**
 * @param urlString 网址
 * @param httpMethod http请求方法 http.MethodPost http.MethodGet
 * @param headers header信息
 * @param paramMap post表单数据
 * @param body body数据
 * @param timeout 超时时长，-1表示默认超时，单位毫秒
 * @return string 网页内容,error
 */
func HttpBasic(urlString, httpMethod string, headers, paramMap map[string]string, body string, timeout int) (stream []byte, err error) {
	client := &fasthttp.Client{
		MaxConnWaitTimeout: time.Duration(timeout) * time.Millisecond,
	}
	req := &fasthttp.Request{}
	queryStr := strKit.MapParamsToUrlParams(paramMap)
	req.SetRequestURI(urlString)
	req.URI().SetQueryString(queryStr)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	req.Header.SetMethod(httpMethod)
	resp := &fasthttp.Response{}
	// 发起请求
	if err = client.Do(req, resp); err != nil {
		return nil, err
	}
	return resp.Body(), nil
}

// HttpProxyGet 发送get代理请求
/**
 * @param urlStr 网址
 * @param proxyIpPort 代理ip和端口
 * @return string 网页内容,error
 */
func HttpProxyGet(urlStr, proxyIpPort string) (map[string]string, string, error) {
	return HttpProxyGetFull(urlStr, nil, nil, "", -1,
		"http", "", "", proxyIpPort)
}

// HttpProxyGetFull 发送get代理请求[完整版]
/**
 * @param urlString 网址
 * @param headers header信息
 * @param paramMap 参数
 * @param body body数据
 * @param timeout 超时时长，-1表示默认超时，单位毫秒
 * @param proxyHttpType 代理类型 http/https
 * @param username 用户名 用户名和密码为空时不使用代理
 * @param password 密码
 * @param proxyIpPort 代理ip和端口
 * @return string 网页内容,error
 */
func HttpProxyGetFull(urlString string, headers, paramMap map[string]string, body string,
	timeout int, proxyHttpType, username, password, proxyIpPort string) (map[string]string, string, error) {
	return HttpProxyBasic(urlString, http.MethodGet, headers, paramMap, body, timeout,
		proxyHttpType, username, password, proxyIpPort)
}

// HttpProxyPost 发送post代理请求
/**
 * @param urlStr 网址
 * @param paramMap 参数
 * @param proxyIpPort 代理ip和端口
 * @return string 网页内容,error
 */
func HttpProxyPost(urlStr string, paramMap map[string]string, proxyIpPort string) (map[string]string, string, error) {
	return HttpProxyPostFull(urlStr, nil, paramMap, "", -1,
		"http", "", "", proxyIpPort)
}

// HttpProxyPostFull 发送post代理请求[完整版]
/**
 * @param urlString 网址
 * @param headers header信息
 * @param paramMap 参数
 * @param body body数据
 * @param timeout 超时时长，-1表示默认超时，单位毫秒
 * @param proxyHttpType 代理类型 http/https
 * @param username 用户名 用户名和密码为空时不使用代理
 * @param password 密码
 * @param proxyIpPort 代理ip和端口
 * @return string 网页内容,error
 */
func HttpProxyPostFull(urlString string, headers, paramMap map[string]string, body string,
	timeout int, proxyHttpType, username, password, proxyIpPort string) (map[string]string, string, error) {
	return HttpProxyBasic(urlString, http.MethodPost, headers, paramMap, body, timeout,
		proxyHttpType, username, password, proxyIpPort)
}

// HttpProxyBasic 发送http代理请求[基础]
/**
 * 注意：proxyIpPort的格式并未校验需要自行校验，原因是有些代理ip可能为ipv6
 * 感谢：感谢巨量IP(https://juliangip.com?goKit)提供测试ip
 * @param urlStr 网址
 * @param httpMethod http请求方法 http.MethodPost http.MethodGet
 * @param headers header信息
 * @param paramMap post表单数据
 * @param body body数据
 * @param timeout 超时时长，-1表示默认超时，单位毫秒
 * @param proxyHttpType 代理类型 http/https
 * @param username 用户名 用户名和密码为空时默认使用无账号密码的代理
 * @param password 密码
 * @param proxyIpPort 代理ip端口 格式：ip:port
 * @return string 网页内容,error
 */
func HttpProxyBasic(urlStr, httpMethod string, headers, paramMap map[string]string,
	body string, timeout int, proxyHttpType, username, password, proxyIpPort string) (map[string]string, string, error) {
	// 判断是否需要代理
	if proxyIpPort == "" {
		// 如果不需要代理，直接调用httpBasic
		result, err := HttpBasic(urlStr, httpMethod, headers, paramMap, body, timeout)
		return headers, string(result), err
	}
	urlParam := strKit.MapParamsToUrlParams(paramMap)
	if urlParam != "" {
		urlStr = strKit.Splicing(urlStr, "?", urlParam)
	}
	bodyReader := strings.NewReader(body)
	proxyStr := strKit.Splicing(proxyHttpType, "://", username, ":", password, "@", proxyIpPort)
	//if username == "" || password == "" {
	//	proxyStr = strKit.Splicing(proxyHttpType, "://", proxyIpPort)
	//}
	proxy, err := url.Parse(proxyStr)
	if err != nil {
		return headers, "", err
	}
	//  请求目标网页
	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxy)}}
	req, _ := http.NewRequest(httpMethod, urlStr, bodyReader)
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	if timeout != -1 {
		http.DefaultClient.Timeout = time.Duration(timeout) * time.Millisecond
	}
	res, err := client.Do(req)
	if err != nil {
		return headers, "", err
	}
	defer res.Body.Close()
	respByte, err := io.ReadAll(res.Body)
	if err != nil {
		return headers, "", err
	}
	// 获取新的headers数据
	if headers == nil {
		// 如果headers为空，初始化
		headers = make(map[string]string, 0)
	}
	for k, v := range res.Header {
		headers[k] = v[0]
	}
	result := string(respByte)
	return headers, result, nil
}
