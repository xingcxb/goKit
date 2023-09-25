// Package redisKit redis连接信息
package redisKit

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/xingcxb/goKit/core/reflectKit"
	"strings"
	"time"
)

var (
	Rdb *redis.Client
)

// RedisConfig redis连接信息
type RedisConfig struct {
	Addr     string //redis链接地址
	Username string //用户名
	Password string //密码
	Port     string `default:"6379"` //端口号
	Db       int    `default:"0"`    //操作数据库
	Timeout  int    `default:"1"`    // 超时时间单位s
	PoolSize int    `default:"10"`   // 连接池大小
}

// NewRedisClient *新建一个redis客户端
/*
 * 注意：Timeout的单位是s，默认超时时间为1s
 *  @param config redis配置文件
 */
func NewRedisClient(config *RedisConfig) error {
	// 读取结构体tag中为default的值
	err := reflectKit.StructDefault(config)
	if err != nil {
		return err
	}
	// 构建redis信息
	var url strings.Builder
	url.WriteString(config.Addr)
	url.WriteString(":")
	url.WriteString(config.Port)

	// 创建redis连接
	Rdb = redis.NewClient(&redis.Options{
		Addr:        url.String(),                                     //redis链接地址
		Username:    config.Username,                                  //设置用户名
		Password:    config.Password,                                  //设置密码
		DB:          config.Db,                                        //设置默认的数据库
		DialTimeout: time.Duration(config.Timeout) * time.Millisecond, //设置超时时间为1s，避免等待时间过长
		PoolSize:    config.PoolSize,                                  //设置连接池大小
	})
	return nil
}

// Ping redis测试是否联通
/*
 * @param ctx 上下文
 * @return 正常返回nil，错误返回错误信息
 */
func Ping(ctx context.Context) error {
	err := Rdb.Ping(ctx).Err()
	if err != nil {
		return err
	}
	return nil
}

// ChangeDb 切换Redis数据库
/*
 * @param ctx 上下文
 * @param dbId redis数据库id
 */
func ChangeDb(ctx context.Context, dbId int) error {
	pipe := Rdb.Pipeline()
	_ = pipe.Select(ctx, dbId)
	_, err := pipe.Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

// GetDbCount 获取单个库的数量
/*
 * @param ctx 上下文
 * @param dbId redis数据库id
 * @return 返回该库下的数量
 */
func GetDbCount(ctx context.Context, dbId int) int {
	ChangeDb(ctx, dbId)
	count, err := Rdb.DBSize(ctx).Result()
	if err != nil {
		return 0
	}
	return int(count)
}

// GetBaseAllInfo  获取redis基础信息
/*
 * @param ctx 上下文
 * @return 返回redis的基础信息
 */
func GetBaseAllInfo(ctx context.Context) map[string]string {
	_info := Rdb.Info(ctx).String()
	defer Rdb.Close()
	_vs := strings.Split(_info, "\r\n")
	infoMap := make(map[string]string)
	for _, _str := range _vs {
		_strs := strings.Split(_str, ":")
		if len(_strs) != 2 || _strs[0] == "info" {
			continue
		}
		infoMap[_strs[0]] = _strs[1]
	}
	return infoMap
}
