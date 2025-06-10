package plugin

import (
	"context"
	"fmt"
	"github.com/xingcxb/goKit/plugin/oss/ali"
	"os"
	"testing"
	"time"
)

func TestOss(t *testing.T) {
	ali.BuildOssInfo("accessKeyID", "accessKeySecret")
	oss := ali.AliOss{
		Region:     "cn-hangzhou",
		BucketName: "bucket",
		ObjectName: "temp/picture.svg", // 指定路径，去掉桶路径
	}
	fileInfo, err := os.ReadFile("/Users/symbol/Downloads/fogfox_430.svg")
	if err != nil {
		fmt.Println("读取文件失败", err)
	}
	fileSizeKB := float64(len(fileInfo)) / 1024
	fmt.Println(fileSizeKB)
	beginTime := time.Now().UnixMilli()
	oss.NewAliOssUpLoad(context.Background(), "http://username:password@127.0.0.1:12345", fileInfo)
	endTime := time.Now().UnixMilli()
	useTime := endTime - beginTime
	fmt.Println("耗时:", useTime, "ms;速度:", fileSizeKB/float64(useTime)*1000, "KB/s")
}
