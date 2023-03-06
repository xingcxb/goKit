// Package redisKit Redis 字符串(String)
package redisKit

import (
	"context"
	"errors"
	"reflect"
	"time"
)

// GetStr 获取redis string类型的数据
// @param key 键
// @return 返回值，错误信息
func GetStr(ctx context.Context, key string) (string, error) {
	if val, err := rdb.Get(ctx, key).Result(); err != nil {
		return "", err
	} else {
		return val, nil
	}
}

// GetRange 返回 key 中字符串值的子字符
// @param key 键
// @param start 起始下标
// @param end 结束下标
func GetRange(ctx context.Context, key string, start, end int64) (string, error) {
	if str, err := rdb.GetRange(ctx, key, start, end).Result(); err != nil {
		return "", err
	} else {
		return str, nil
	}
}

// GetSet 将给定 key 的值设为 value ，并返回 key 的旧值(old value)
// @param key 键
// @param newValue 新值
// @return 当前键被替换前的值
func GetSet(ctx context.Context, key, newValue string) (string, error) {
	if str, err := rdb.GetSet(ctx, key, newValue).Result(); err != nil {
		return "", err
	} else {
		return str, nil
	}
}

// SetStr 设置redis中的值，该值永久有效
// @param key 键
// @param value 值
// @return 返回成功还是失败，错误信息
func SetStr(ctx context.Context, key, value string) (bool, error) {
	return SetStrEX(ctx, key, value, -1)
}

// SetStrEX 设置redis中的值
// @param key 键
// @param value 值
// @param seconds 数据存活时间，当值为-1或0时为永久有效
// @return 返回成功还是失败，错误信息
func SetStrEX(ctx context.Context, key, value string, seconds int) (bool, error) {
	err := errors.New("")
	if seconds == -1 || seconds == 0 {
		err = rdb.Set(ctx, key, value, 0).Err()
	} else {
		err = rdb.Set(ctx, key, value, time.Duration(seconds)*time.Second).Err()
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

// SetNX （SET if Not eXists）只有在 key 不存在时设置 key 的值
// @param key 键
// @param value 值
// @param seconds 数据存活时间，当值为-1或0时为永久有效
// @return 返回成功还是失败，错误信息
func SetNX(ctx context.Context, key, value string, seconds int) (bool, error) {
	err := errors.New("")
	if seconds == -1 || seconds == 0 {
		err = rdb.SetNX(ctx, key, value, 0).Err()
	} else {
		err = rdb.SetNX(ctx, key, value, time.Duration(seconds)*time.Second).Err()
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

// SetRange 用 value 参数覆写给定 key 所储存的字符串值，从偏移量 offset 开始
// @param key 键
// @param replaceValue 替换字符串
// @param start 替换开始的下标
// @return 返回成功还是失败，错误信息
func SetRange(ctx context.Context, key, replaceValue string, start int64) (bool, error) {
	if err := rdb.SetRange(ctx, key, start, replaceValue).Err(); err != nil {
		return false, err
	} else {
		return true, nil
	}
}

// StrLen 返回 key 所储存的字符串值的长度
// @param key 键
// @return 返回字符串的长度
func StrLen(ctx context.Context, key string) (int64, error) {
	if strLen, err := rdb.StrLen(ctx, key).Uint64(); err != nil {
		return -1, err
	} else {
		return int64(strLen), nil
	}
}

// Incr 将 key 中储存的数字值增1
// @param key 键
// @return 返回增加后的值，错误信息
func Incr(ctx context.Context, key string) (int64, error) {
	if v, err := rdb.Incr(ctx, key).Result(); err != nil {
		return -1, err
	} else {
		return v, nil
	}
}

// Incrby 将 key 中储存的数字值增加指定的数
// @param key 键
// @param value 类型仅仅支持浮点和整数
// @return 返回增加后的值，错误信息
func Incrby(ctx context.Context, key string, value interface{}) (interface{}, error) {
	// 获取value的类型
	valueType := reflect.TypeOf(value).String()
	switch valueType {
	case "int":
		return rdb.IncrBy(ctx, key, int64(value.(int))).Result()
	case "int8":
		return rdb.IncrBy(ctx, key, int64(value.(int8))).Result()
	case "int16":
		return rdb.IncrBy(ctx, key, int64(value.(int16))).Result()
	case "int32":
		return rdb.IncrBy(ctx, key, int64(value.(int32))).Result()
	case "int64":
		return rdb.IncrBy(ctx, key, value.(int64)).Result()
	case "uint":
		return rdb.IncrBy(ctx, key, int64(value.(uint))).Result()
	case "uint8":
		return rdb.IncrBy(ctx, key, int64(value.(uint8))).Result()
	case "uint16":
		return rdb.IncrBy(ctx, key, int64(value.(uint16))).Result()
	case "uint32":
		return rdb.IncrBy(ctx, key, int64(value.(uint32))).Result()
	case "uint64":
		return rdb.IncrBy(ctx, key, int64(value.(uint64))).Result()
	case "uintptr":
		return rdb.IncrBy(ctx, key, int64(value.(uintptr))).Result()
	case "float32":
		return rdb.IncrByFloat(ctx, key, float64(value.(float32))).Result()
	case "float64":
		return rdb.IncrByFloat(ctx, key, value.(float64)).Result()
	default:
		return false, errors.New("value type error")
	}
}

// Decr 将 key 中储存的数字值减一
// @param key 键
// @return 返回减少后的值，错误信息
func Decr(ctx context.Context, key string) (int64, error) {
	if v, err := rdb.Decr(ctx, key).Result(); err != nil {
		return -1, err
	} else {
		return v, nil
	}
}
