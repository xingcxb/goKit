// Package strKit 身份证
// 代码源自于https://github.com/guanguans/id-validator
package strKit

import (
	"errors"
	"github.com/xingcxb/goKit/core/numKit"
	"github.com/xingcxb/goKit/core/strKit/data"
	"math"
	"regexp"
	"strings"
	"time"
)

// CheckIdCardValid 判断身份证号是否合法
/*
 * @param id 身份证号
 * @param strict 是否严格模式
 * @return bool
 */
func CheckIdCardValid(id string, strict bool) bool {
	code, err := generateCode(id)
	if err != nil {
		return false
	}

	// 检查顺序码、生日码、地址码
	if !checkOrderCode(code["order"]) || !checkBirthdayCode(code["birthdayCode"]) || !checkAddressCode(code["addressCode"], code["birthdayCode"], strict) {
		return false
	}
	// 15位身份证不含校验码
	if code["type"] == "15" {
		return true
	}

	// 校验码
	return code["checkBit"] == generatorCheckBit(code["body"])
}

// 生成Bit码
func generatorCheckBit(body string) string {
	// 位置加权
	var posWeight [19]float64
	for i := 2; i < 19; i++ {
		weight := int(math.Pow(2, float64(i-1))) % 11
		posWeight[i] = float64(weight)
	}

	// 累身份证号body部分与位置加权的积
	var bodySum int
	bodyArray := strings.Split(body, "")
	count := len(bodyArray)
	for i := 0; i < count; i++ {
		bodySum += Str2Int(bodyArray[i], "int").(int) * int(posWeight[18-i])
	}

	// 生成校验码
	checkBit := (12 - (bodySum % 11)) % 11
	if checkBit == 10 {
		return "x"
	}
	return numKit.Int2Str(checkBit)
}

// 检查地址码
func checkAddressCode(addressCode string, birthdayCode string, strict bool) bool {
	addressInfo := getAddressInfo(addressCode, birthdayCode, strict)
	// 用于判断是否是港澳台居民居住证（8字开头）
	// 港澳台居民居住证无市级、县级信息
	firstCharacter := SubString(addressCode, 0, 1)
	if firstCharacter == "8" && addressInfo["province"] != "" {
		return true
	}
	// 这里不判断市级信息的原因：
	// 1. 直辖市，无市级信息
	// 2. 省直辖县或县级市，无市级信息
	if addressInfo["province"] == "" || addressInfo["district"] == "" {
		return false
	}
	return true
}

// 获取地址信息
func getAddressInfo(addressCode string, birthdayCode string, strict bool) map[string]string {
	addressInfo := map[string]string{
		"province": "",
		"city":     "",
		"district": "",
	}
	// 省级信息
	addressInfo["province"] = getAddress(SubString(addressCode, 0, 2)+"0000", birthdayCode, strict)
	// 用于判断是否是港澳台居民居住证（8字开头）
	firstCharacter := SubString(addressCode, 0, 1)
	// 港澳台居民居住证无市级、县级信息
	if firstCharacter == "8" {
		return addressInfo
	}
	// 市级信息
	addressInfo["city"] = getAddress(SubString(addressCode, 0, 4)+"00", birthdayCode, strict)
	// 县级信息
	addressInfo["district"] = getAddress(addressCode, birthdayCode, strict)
	return addressInfo
}

// 获取省市区地址码
func getAddress(addressCode string, birthdayCode string, strict bool) string {
	address := ""
	timeline := data.GetAddressCodeTimeline(Str2Int(addressCode, "uint32").(uint32))
	if len(timeline) == 0 {
		// 修复 \d\d\d\d01、\d\d\d\d02、\d\d\d\d11 和 \d\d\d\d20 的历史遗留问题
		// 以上四种地址码，现实身份证真实存在，但民政部历年公布的官方地址码中可能没有查询到
		// 如：440401 450111 等
		// 所以这里需要特殊处理
		// 1980年、1982年版本中，未有制定省辖市市辖区的代码，所有带县的省辖市给予“××××20”的“市区”代码。
		// 1984年版本开始对地级市（前称省辖市）市辖区制定代码，其中“××××01”表示市辖区的汇总码，同时撤销“××××20”的“市区”代码（追溯至1983年）。
		// 1984年版本的市辖区代码分为城区和郊区两类，城区由“××××02”开始排起，郊区由“××××11”开始排起，后来版本已不再采用此方式，已制定的代码继续沿用。
		suffixes := SubString("123456", 4, 6)
		switch suffixes {
		case "20":
			address = "市区"
		case "01":
			address = "市辖区"
		case "02":
			address = "城区"
		case "11":
			address = "郊区"
		}

		return address
	}
	year := Str2Int(SubString(birthdayCode, 0, 4), "int").(int)
	startYear := "0001"
	endYear := "9999"
	for _, val := range timeline {
		if val["start_year"] != "" {
			startYear = val["start_year"]
		}
		if val["end_year"] != "" {
			endYear = val["end_year"]
		}
		if year >= Str2Int(startYear, "int").(int) && year <= Str2Int(endYear, "int").(int) {
			address = val["address"]
		}
	}
	if address == "" && !strict {
		// 由于较晚申请户口或身份证等原因，导致会出现地址码正式启用于2000年，但实际1999年出生的新生儿，由于晚了一年报户口，导致身份证上的出生年份早于地址码正式启用的年份
		// 由于某些地区的地址码已经废弃，但是实际上在之后的几年依然在使用
		// 这里就不做时间判断了
		address = timeline[0]["address"]
	}
	return address
}

// 检查出生日期码
func checkBirthdayCode(birthdayCode string) bool {
	year := Str2Int(SubString(birthdayCode, 0, 4), "int").(int)
	if year < 1800 {
		return false
	}
	if year > time.Now().Year() {
		return false
	}
	_, err := time.Parse("20060102", birthdayCode)
	return err == nil
}

// 检查顺序码
func checkOrderCode(orderCode string) bool {
	return len(orderCode) == 3
}

// 生成短数据
func generateShortCode(id string) (map[string]string, error) {
	if len(id) != 15 {
		return map[string]string{}, errors.New("invalid ID card number length")
	}

	mustCompile := regexp.MustCompile("(.{6})(.{6})(.{3})")
	subMatch := mustCompile.FindStringSubmatch(strings.ToLower(id))
	if len(subMatch) < 4 {
		return nil, errors.New("error extract submatch(shortCode)")
	}

	return map[string]string{
		"body":         subMatch[0],
		"addressCode":  subMatch[1],
		"birthdayCode": "19" + subMatch[2],
		"order":        subMatch[3],
		"checkBit":     "",
		"type":         "15",
	}, nil
}

// 生成长数据
func generateLongCode(id string) (map[string]string, error) {
	if len(id) != 18 {
		return map[string]string{}, errors.New("invalid ID card number length")
	}
	mustCompile := regexp.MustCompile("((.{6})(.{8})(.{3}))(.)")
	subMatch := mustCompile.FindStringSubmatch(strings.ToLower(id))
	if len(subMatch) < 6 {
		return nil, errors.New("error extract submatch(longCode)")
	}

	return map[string]string{
		"body":         subMatch[1],
		"addressCode":  subMatch[2],
		"birthdayCode": subMatch[3],
		"order":        subMatch[4],
		"checkBit":     subMatch[5],
		"type":         "18",
	}, nil
}

// 生成数据
func generateCode(id string) (map[string]string, error) {
	length := len(id)
	if length == 15 {
		return generateShortCode(id)
	}

	if length == 18 {
		return generateLongCode(id)
	}

	return map[string]string{}, errors.New("invalid ID card number length")
}
