package randomKit

import "encoding/binary"

const (
	// uint32随机数的缓冲区大小
	bufferChanSize = 10000
)

var (
	// bufferChan 缓冲区是随机字节,每项存储4个字节。
	bufferChan = make(chan []byte, bufferChanSize)
)

// Intn 返回一个介于0和max之间的int数:[0,max)。
// 注意:
// 1. 'max'只能大于0，否则直接返回'max';
// 2. 结果大于或等于0，但小于'max';
// 3. 结果数字为32位，小于math. maxin32。
func Intn(max int) int {
	if max <= 0 {
		return max
	}
	n := int(binary.LittleEndian.Uint32(<-bufferChan)) % max
	if (max > 0 && n < 0) || (max < 0 && n > 0) {
		return -n
	}
	return n
}

// RandomInt 获得指定范围内的随机数
// @param min 最小数（包含）
// @param max 最大数（不包含）
// @return 随机数。
func RandomInt(min, max int) int {
	if min >= max {
		return min
	}
	if min >= 0 {
		return Intn(max-min+1) + min
	}
	return Intn(max+(0-min)+1) - (0 - min)
}
