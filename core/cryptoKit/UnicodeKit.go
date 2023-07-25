package cryptoKit

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// UnicodeEncode 将字符串转为Unicode编码
/**
 * @param str 字符串
 * @return Unicode编码字符串
 */
func UnicodeEncode(str string) string {
	DD := []rune(str) //需要分割的字符串内容，将它转为字符，然后取长度。
	finallStr := ""
	for i := 0; i < len(DD); i++ {
		if unicode.Is(unicode.Scripts["Han"], DD[i]) {
			textQuoted := strconv.QuoteToASCII(string(DD[i]))
			finallStr += textQuoted[1 : len(textQuoted)-1]
		} else {
			h := fmt.Sprintf("%x", DD[i])
			finallStr += "\\u" + isFullFour(h)
		}
	}
	return finallStr
}

// UnicodeDecode 将Unicode编码转为字符串
/**
 * @param str Unicode编码字符串
 * @return 字符串
 */
func UnicodeDecode(str string) (string, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(str), `\\u`, `\u`, -1))
	if err != nil {
		return "", err
	}
	return str, nil
}

// isFullFour 补全4位
/**
 * @param str 字符串
 * @return 补全后的字符串
 */
func isFullFour(str string) string {
	if len(str) == 1 {
		str = "000" + str
	} else if len(str) == 2 {
		str = "00" + str
	} else if len(str) == 3 {
		str = "0" + str
	}
	return str
}
