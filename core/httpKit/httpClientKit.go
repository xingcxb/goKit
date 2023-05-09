package httpKit

import (
	"bufio"
	"errors"
	"github.com/xingcxb/goKit/core/strKit"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// HttpDownload 下载文件
// @param urlString 网址
// @param savePath 保存路径
// @param fileName 文件名，如果不存在则自动获取
// @param isCover 是否覆盖 true 覆盖 false 不覆盖
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
	filePath := strKit.Splicing(savePath, "/", fileName)
	// 判断文件是否存在，默认不存在
	checkFile := false
	// 判断文件是否存在
	if _, err := os.Stat(filePath); err == nil {
		checkFile = true
	}
	if isCover {
		// 允许覆盖，删除文件
		if os.Remove(filePath) != nil {
			return "", err
		}
	} else {
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

// HttpPost 发送post基础请求
// @param urlString 网址
// @param params post表单数据
func HttpPost(urlString string, params map[string]string) (string, error) {
	return HttpBasic(urlString, http.MethodPost, nil, params, "", -1)
}

// HttpPostFull 发送post请求[完整版]
// @param urlString 网址
// @param headers header信息
// @param paramMap post表单数据
// @param body body数据
// @param timeout 超时时长，-1表示默认超时，单位毫秒
func HttpPostFull(urlString string, headers, paramMap map[string]string, body string, timeout int) (string, error) {
	return HttpBasic(urlString, http.MethodPost, headers, paramMap, body, timeout)
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
	if urlParam != "" {
		urlString = strKit.Splicing(urlString, "?", urlParam)
	}
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
