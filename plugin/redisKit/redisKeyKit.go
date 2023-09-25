package redisKit

import (
	"context"
	"encoding/json"
	"github.com/xingcxb/goKit/core/strKit"
	"strconv"
	"strings"
)

// GetDbKeys 获取指定库中的key
/*
 * @param ctx 上下文
 * @param cursor 游标
 */
func GetDbKeys(ctx context.Context, cursor uint64) ([]string, error) {
	keys := make([]string, 0, 1)
	keys, cursor, err := Rdb.Scan(ctx, cursor, "*", 10000).Result()
	if err != nil {
		return keys, err
	}
	return keys, nil
}

// GetTTL 获取redis数据剩余时间，返回剩余时间的秒数；如果是永久有效，返回-1
/*
 * @param ctx 上下文
 * @param key 键
 */
func GetTTL(ctx context.Context, key string) string {
	val, err := Rdb.TTL(ctx, key).Result()
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
/*
 * @param ctx 上下文
 * @param key 键
 * @return 返回值的所有信息,如果获取失败，返回错误信息
 */
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
/*
 * @param ctx 上下文
 * @param key 键
 * @return 返回值的类型
 */
func GetType(ctx context.Context, key string) string {
	allTypeStr := Rdb.Type(ctx, key).String()
	arr := strings.Split(allTypeStr, " ")
	if len(arr) == 3 {
		typeStr := arr[2]
		typeStr = strKit.FirstUpper(typeStr)
		return typeStr
	}
	return ""
}

// GetList 获取redis list类型的数据，返回值和大小
/*
 * @param ctx 上下文
 * @param key 键
 * @return 返回值和大小
 */
func GetList(ctx context.Context, key string) []string {
	val, err := Rdb.LRange(ctx, key, 0, 100).Result()
	if err != nil {
		return nil
	}
	return val
}
