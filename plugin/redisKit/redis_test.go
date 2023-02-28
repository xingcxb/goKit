package redisKit

import (
	"context"
	"fmt"
	"testing"
)

func TestRT(t *testing.T) {
	// 配置redis的连接信息
	_ = NewRedisClient(&RedisConfig{
		Addr:     "127.0.0.1",
		Username: "",
		Password: "123456",
	})
	fmt.Println(Ping(context.Background()))
	fmt.Println(GetStr(context.Background(), "gost:authers:auth-disest-v"))
}
