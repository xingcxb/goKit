// Package redisKit Redis 字符串(String)
package redisKit

import (
	"context"
	"errors"
	"time"
)

// GetStr 获取redis string类型的数据，返回值和大小
func GetStr(ctx context.Context, key string) string {
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return ""
	}
	return val
}

// GetRange 返回 key 中字符串值的子字符
// @param key 键
// @param start 起始下标
// @param end 结束下标
func GetRange(ctx context.Context, key string, start, end int64) (string, error) {
	str, err := rdb.GetRange(ctx, key, start, end).Result()
	if err != nil {
		return "", err
	}
	return str, nil
}

// GetSet 将给定 key 的值设为 value ，并返回 key 的旧值(old value)
// @param key 键
// @param newValue 新值
// @return 当前键被替换前的值
func GetSet(ctx context.Context, key, newValue string) (string, error) {
	str, err := rdb.GetSet(ctx, key, newValue).Result()
	if err != nil {
		return "", err
	}
	return str, nil
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
	err := rdb.SetRange(ctx, key, start, replaceValue).Err()
	if err != nil {
		return false, err
	}
	return true, nil
}

// StrLen 返回 key 所储存的字符串值的长度
// @param key 键
// @return 返回字符串的长度
func StrLen(ctx context.Context, key string) (int64, error) {
	strLen, err := rdb.StrLen(ctx, key).Uint64()
	if err != nil {
		return -1, err
	}
	return int64(strLen), nil
}

// Incr 将 key 中储存的数字值增1
// @param key 键
// @return 返回增加后的值，错误信息
func Incr(ctx context.Context, key string) (int64, error) {
	v, err := rdb.Incr(ctx, key).Result()
	if err != nil {
		return v, err
	}
	return v, nil
}

// Incrby 将 key 中储存的数字值增1
// @param key 键
// @return 返回增加后的值，错误信息
//func Incrby(ctx context.Context, key string) (bool, error) {
//}
