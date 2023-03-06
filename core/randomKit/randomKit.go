package randomKit

import (
	"crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	"goKit/core/strKit"
	"strings"
)

const (
	// uint32随机数的缓冲区大小
	bufferChanSize = 10000
)

var (
	// bufferChan 缓冲区是随机字节,每项存储4个字节。
	bufferChan = make(chan []byte, bufferChanSize)
	// BaseNumber 用于随机选的数字
	BaseNumber = "0123456789"
	// BaseChar 用于随机选的字符
	BaseChar = "abcdefghijklmnopqrstuvwxyz"
	// BaseCharNumber 用于随机选的字符和数字
	BaseCharNumber = fmt.Sprintf("%v%v", BaseChar, BaseNumber)
)

func init() {
	go asyncProducingRandomBufferBytesLoop()
}

// asyncProducingRandomBufferBytes is a named goroutine, which uses a asynchronous goroutine
// to produce the random bytes, and a buffer chan to store the random bytes.
// So it has high performance to generate random numbers.
func asyncProducingRandomBufferBytesLoop() {
	var step int
	for {
		buffer := make([]byte, 1024)
		if n, err := rand.Read(buffer); err != nil {
			panic(errors.New(fmt.Sprintf(`%v error reading random buffer from system`, err)))
		} else {
			// The random buffer from system is very expensive,
			// so fully reuse the random buffer by changing
			// the step with a different number can
			// improve the performance a lot.
			// for _, step = range []int{4, 5, 6, 7} {
			for _, step = range []int{4} {
				for i := 0; i <= n-4; i += step {
					bufferChan <- buffer[i : i+4]
				}
			}
		}
	}
}

// RandomLong 返回一个介于0和max之间的int数:[0,max)。
// 注意:
// 1. max只能大于0，否则直接返回max;
// 2. 结果大于或等于0，但小于max;
// 3. 结果数字为32位，小于math. max32。
func RandomLong(max int) int {
	if max <= 0 {
		return max
	}
	n := int(binary.LittleEndian.Uint32(<-bufferChan)) % max
	if (max > 0 && n < 0) || (max < 0 && n > 0) {
		return -n
	}
	return n
}

// RandomInt 获得指定范围内的随机数[min, max)
// @param min 最小数（包含）
// @param max 最大数（不包含）
// @return 随机数
func RandomInt(min, max int) int {
	if min >= max {
		return min
	}
	if min >= 0 {
		return RandomLong(max-min+1) + min
	}
	return RandomLong(max+(0-min)+1) - (0 - min)
}

// RandomBool 获得随机Boolean值
// @return true or false
func RandomBool() bool {
	return 0 == RandomLong(2)
}

// RandomStringWithoutStr 获得一个随机的字符串（只包含数字和字符） 并排除指定字符串
// @param length – 字符串的长度
// @param elemData – 要排除的字符串,如：去重容易混淆的字符串，oO0、lL1、q9Q、pP
// @return 随机字符串
func RandomStringWithoutStr(length int, elemData string) string {
	str := BaseCharNumber
	str = strKit.RemoveAll(str, strings.Split(elemData, "")...)
	return RandomStrBasic(str, length)
}

// RandomNumbers 获得一个只包含数字的字符串
// @param length – 字符串的长度
// @return 随机字符串
func RandomNumbers(length int) string {
	return RandomStrBasic(BaseNumber, length)
}

// RandomStrBasic 获得一个随机的字符串
// @param baseString 随机字符选取的样本
// @param length 字符串的长度
func RandomStrBasic(baseString string, length int) string {
	if baseString == "" {
		return strKit.EMPTY
	}
	sb := strings.Builder{}
	if length < 1 {
		length = 1
	}
	baseLength := len(baseString)
	for i := 0; i < length; i++ {
		number := RandomLong(baseLength)
		ArrayCharNumber := strings.Split(BaseCharNumber, "")
		sb.WriteString(ArrayCharNumber[number])
	}
	return sb.String()
}
