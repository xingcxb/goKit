package ali

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
	"io"
	"net/http"
	"net/url"
)

var providerInfo credentials.CredentialsProvider

// BuildOssInfo 构建oss信息
func BuildOssInfo(accessKeyID, accessKeySecret string) {
	providerInfo = credentials.NewStaticCredentialsProvider(accessKeyID, accessKeySecret, "")
}

type AliOss struct {
	Region     string // 存储区域
	BucketName string // 存储空间名称
	ObjectName string // 对象名称
}

// NewAliOssUpLoad 新建阿里云OSS上传
/*
 * @param data {[]byte} 待上传数据
 */
func (a *AliOss) NewAliOssUpLoad(ctx context.Context, proxyInfo string, data []byte) (string, error) {
	client := a.createClient(proxyInfo)
	// 创建上传对象的请求
	request := &oss.PutObjectRequest{
		Bucket: oss.Ptr(a.BucketName), // 存储空间名称
		Key:    oss.Ptr(a.ObjectName), // 对象名称
		Body:   bytes.NewReader(data), // 要上传的内容
	}
	// 执行上传对象的请求
	result, err := client.PutObject(ctx, request)
	if err != nil {
		return "", err
	}
	// 打印上传对象的结果
	fmt.Println("put object result:", result)
	return result.Status, nil
}

// NewAliOssDownload 新建阿里云OSS下载
func (a *AliOss) NewAliOssDownload(ctx context.Context, proxyInfo string) ([]byte, error) {
	client := a.createClient(proxyInfo)
	// 创建获取对象的请求
	request := &oss.GetObjectRequest{
		Bucket: oss.Ptr(a.BucketName), // 存储空间名称
		Key:    oss.Ptr(a.ObjectName), // 对象名称
	}
	// 执行获取对象的操作并处理结果
	result, err := client.GetObject(ctx, request)
	defer result.Body.Close()
	if err != nil {
		return nil, err
	}
	data, _ := io.ReadAll(result.Body)
	return data, nil
}

// NewAliOssDel 新建阿里云OSS删除
func (a *AliOss) NewAliOssDel(ctx context.Context) error {
	client := a.createClient("")
	// 创建删除对象的请求
	request := &oss.DeleteObjectRequest{
		Bucket: oss.Ptr(a.BucketName), // 存储空间名称
		Key:    oss.Ptr(a.ObjectName), // 对象名称
	}
	// 执行删除对象的操作并处理结果
	_, err := client.DeleteObject(ctx, request)
	if err != nil {
		fmt.Println("failed to delete object", err)
		return err
	}
	return nil
}

// 创建OSS客户端
func (a *AliOss) createClient(proxyInfo string) *oss.Client {
	// 加载默认配置并设置凭证提供者和区域
	cfg := oss.LoadDefaultConfig().
		WithCredentialsProvider(providerInfo).
		WithRegion(a.Region)
	if proxyInfo != "" {
		// 如果存在代理信息，走代理
		cfg.WithHttpClient(createAliOssProxy(proxyInfo))
	}
	// 创建OSS客户端
	return oss.NewClient(cfg)
}

// 创建代理信息
func createAliOssProxy(proxyInfo string) *http.Client {
	proxyURL, _ := url.Parse(proxyInfo)
	// 创建带代理的 Transport
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}
	// 创建带代理的 HTTP Client
	return &http.Client{
		Transport: transport,
	}
}
