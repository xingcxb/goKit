package plugin

import (
	"context"
	"fmt"
	"github.com/xingcxb/goKit/plugin/redisKit"
	"testing"
)

func TestRT(t *testing.T) {
	// 配置redis的连接信息
	_ = redisKit.NewRedisClient(&redisKit.RedisConfig{
		Addr:     "192.168.10.58",
		Username: "",
		Password: "123456",
	})
	fmt.Println(redisKit.Ping(context.Background()))
	fmt.Println(redisKit.GetStr(context.Background(), "gost:authers:auth-disest-v"))
	fmt.Println(redisKit.Incr(context.Background(), "aa"))
	aa := uintptr(6)
	fmt.Println(redisKit.Incrby(context.Background(), "aa", aa))
}
