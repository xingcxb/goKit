package uaKit

import (
	"github.com/xingcxb/goKit/core/strKit"
)

// UAInfo 浏览器信息
type UAInfo struct {
	DeviceType     string `json:"device_type"`
	Os             string `json:"os"`
	OsVersion      string `json:"os_version"`
	Browser        string `json:"browser"`
	BrowserVersion string `json:"browser_version"`
}

// ParseUA 获取浏览器信息
func ParseUA(uaString string) UAInfo {
	parser := NewFromSaved()
	client := parser.Parse(uaString)
	return UAInfo{
		DeviceType: client.Device.Family,
		Os:         client.Os.Family,
		OsVersion: func(*Client) string {
			osVersion := client.Os.Major
			if client.Os.Minor != "" {
				osVersion = strKit.Splicing(osVersion, ".", client.Os.Minor)
			}
			if client.Os.Patch != "" {
				osVersion = strKit.Splicing(osVersion, ".", client.Os.Patch)
			}
			if client.Os.PatchMinor != "" {
				osVersion = strKit.Splicing(osVersion, ".", client.Os.PatchMinor)
			}
			return osVersion
		}(client),
		Browser: client.UserAgent.Family,
		BrowserVersion: func(*Client) string {
			browserVersion := client.UserAgent.Major
			if client.UserAgent.Minor != "" {
				browserVersion = strKit.Splicing(browserVersion, ".", client.UserAgent.Minor)
			}
			if client.UserAgent.Patch != "" {
				browserVersion = strKit.Splicing(browserVersion, ".", client.UserAgent.Patch)
			}
			return browserVersion
		}(client),
	}
}
