package redisKit

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"goKit/core/reflectKit"
	"goKit/core/strKit"
	"strconv"
	"strings"
	"time"
)

type RedisConfig struct {
	Addr     string //redis链接地址
	Username string //用户名
	Password string //密码
	Port     string `default:"6379"` //端口号
	Db       int    `default:"0"`    //操作数据库
	Timeout  int    `default:"1"`    // 超时时间单位s
	PoolSize int    `default:"10"`   // 连接池大小
}

var (
	Config RedisConfig
	rdb    *redis.Client
)

// NewRedisClient *新建一个redis客户端
// @param config redis配置文件，注意：Timeout的单位是s，默认超时时间为1s
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
	rdb = redis.NewClient(&redis.Options{
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
// @return 正常返回nil，错误返回错误信息
func Ping(ctx context.Context) error {
	err := rdb.Ping(ctx).Err()
	if err != nil {
		return err
	}
	return nil
}

// ChangeDb 切换Redis数据库
// @param dbId redis数据库id
func ChangeDb(ctx context.Context, dbId int) error {
	pipe := rdb.Pipeline()
	_ = pipe.Select(ctx, dbId)
	_, err := pipe.Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

// GetDbCount 获取单个库的数量
// @return 返回该库下的数量
func GetDbCount(ctx context.Context, dbId int) int {
	ChangeDb(ctx, dbId)
	count, err := rdb.DBSize(ctx).Result()
	if err != nil {
		return 0
	}
	return int(count)
}

// GetBaseAllInfo  获取redis基础信息
func GetBaseAllInfo(ctx context.Context) map[string]string {
	_info := rdb.Info(ctx).String()
	defer rdb.Close()
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

// GetDbKeys 获取指定库中的key
func GetDbKeys(ctx context.Context, cursor uint64) ([]string, error) {
	keys := make([]string, 0, 1)
	keys, cursor, err := rdb.Scan(ctx, cursor, "*", 10000).Result()
	if err != nil {
		return keys, err
	}
	return keys, nil
}

// GetTTL 获取redis数据剩余时间，返回剩余时间的秒数；如果是永久有效，返回-1
func GetTTL(ctx context.Context, key string) string {
	val, err := rdb.TTL(ctx, key).Result()
	if err != nil {
		return ""
	}
	if val == -1 {
		return "-1"
	}
	return strconv.FormatInt(int64(val.Seconds()), 10)
}

// VObj 值信息
type VObj struct {
	Size  int    `json:"size"`  //值的大小
	Value string `json:"value"` //值的内容
	Ttl   string `json:"ttl"`   //过期时间
	Type  string `json:"type"`  //值的类型
}

// GetKeyInfo 通过key获取该键下值的所有信息
func GetKeyInfo(ctx context.Context, key string) (string, error) {
	// 获取值
	v, err := GetStr(ctx, key)
	if err != nil {
		return "", err
	}
	size := len(v)
	ttl := GetTTL(ctx, key)
	keyType := GetType(ctx, key)
	info := VObj{
		Size:  size,
		Value: v,
		Ttl:   ttl,
		Type:  keyType,
	}
	strByte, _ := json.Marshal(info)
	return string(strByte), nil
}

// GetType 获取值类型，返回类型
func GetType(ctx context.Context, key string) string {
	allTypeStr := rdb.Type(ctx, key).String()
	arr := strings.Split(allTypeStr, " ")
	if len(arr) == 3 {
		typeStr := arr[2]
		typeStr = strKit.FirstUpper(typeStr)
		return typeStr
	}
	return ""
}

// GetList 获取redis list类型的数据，返回值和大小
func GetList(ctx context.Context, key string) []string {
	val, err := rdb.LRange(ctx, key, 0, 100).Result()
	if err != nil {
		return nil
	}
	return val
}
