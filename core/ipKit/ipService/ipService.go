// Package ipService 读取第三方的ip地址信息
package ipService

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
	"github.com/xingcxb/goKit/core/httpKit"
	"github.com/xingcxb/goKit/core/randomKit"
	"github.com/xingcxb/goKit/core/strKit"
	"strings"
)

var (
	// 纯真地址库
	requestUrlCZ = "https://www.cz88.net/api/cz88/ip/base?ip="
	// cip库
	requestUrlCip = "http://www.cip.cc/"
)

type IpInfo struct {
	Ip       string `json:"ip"`       // ip
	Country  string `json:"country"`  // 国家
	Province string `json:"province"` // 省份
	City     string `json:"city"`     // 城市
	Isp      string `json:"isp"`      // 运营商
	Score    string `json:"score"`    // 真人概率
}

// RandomQuery 随机查询ip的地址信息
/**
 * @param ip 要查询的ip
 * @return IpInfo,error
 */
func RandomQuery(ip string) (IpInfo, error) {
	if randomKit.RandomInt(0, 1) == 0 {
		return GetIpInfoCZ(ip)
	}
	return GetIpInfoCip(ip)
}

// GetIpInfoCZ 获取ip信息，数据来自纯真
/**
 * 数据来源网址 https://www.cz88.net/
 * @param ip 要查询的ip
 * @return IpInfo,error
 */
func GetIpInfoCZ(ip string) (IpInfo, error) {
	ipInfo := IpInfo{Ip: ip}
	requestUrl := strKit.Splicing(requestUrlCZ, ip)
	result, _ := httpKit.HttpGet(requestUrl)
	code := gjson.Get(result, "code").Int()
	if code != 200 {
		// 目前属于试探性测试为200属于正常
		return ipInfo, errors.New("纯真数据获取失败，请在项目中提交Issues")
	}
	ipInfo.Country = gjson.Get(result, "data.country").String()
	ipInfo.Province = gjson.Get(result, "data.province").String()
	ipInfo.City = gjson.Get(result, "data.city").String()
	ipInfo.Isp = gjson.Get(result, "data.isp").String()
	ipInfo.Score = gjson.Get(result, "data.score").String()
	return ipInfo, nil
}

// GetIpInfoCip 获取ip信息，数据来自ip138
/**
 * 数据来源网址 http://www.cip.cc/
 * @param ip 要查询的ip
 * @return IpInfo,error
 */
func GetIpInfoCip(ip string) (IpInfo, error) {
	ipInfo := IpInfo{Ip: ip}
	requestUrl := strKit.Splicing("https://www.cip.cc/", ip)
	resultHtml, _ := httpKit.HttpGet(requestUrl)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resultHtml))
	if err != nil {
		return ipInfo, err
	}
	selectText := ""
	doc.Find("pre").Each(func(i int, selection *goquery.Selection) {
		selectText = selection.Text()
	})
	ipInfos := strings.Split(selectText, "\n")
	for _, info := range ipInfos {
		if info == "" {
			continue
		}
		datas := strings.Split(info, ":")
		if strings.Contains(datas[0], "运营商") {
			ipInfo.Isp = datas[1]
		} else if strings.Contains(datas[0], "地址") {
			addressInfo := strings.Split(datas[1], "  ")
			ipInfo.Country = strings.TrimSpace(addressInfo[0])
			if len(addressInfo) > 1 {
				ipInfo.Province = addressInfo[1]
			}
			if len(addressInfo) > 2 {
				ipInfo.City = addressInfo[2]
			}
		}
	}
	return ipInfo, err
}
