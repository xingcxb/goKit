// Package redisKit Redis 字符串(Hash)
package redisKit

import "context"

// HMSet 同时将多个 field-value (域-值)对设置到哈希表 key 中
// @param key 键
// @param value... 值
// @return 返回的值
func HMSet(ctx context.Context, key string, value ...interface{}) (bool, error) {
	return rdb.HMSet(ctx, key, value...).Result()
}

// HMGet 同时将多个 field-value (域-值)对设置到哈希表 key 中
// @param key 键
// @param value... 值
// @return 返回的值
func HMGet(ctx context.Context, key string, fields ...string) (interface{}, error) {
	return rdb.HMGet(ctx, key, fields...).Result()
}
