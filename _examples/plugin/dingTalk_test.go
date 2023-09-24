package plugin

import (
	"fmt"
	"github.com/xingcxb/goKit/plugin/dingTalkKit"
	"testing"
)

func TestDingGroupTest(t *testing.T) {
	dingGroupConfig := dingTalkKit.DingTalkBot{
		WebHookUrl: "https://oapi.dingtalk.com/robot/send?access_token=3626528b0f0f50a59634213854f9e4sfsdf0bd4fe3b45123432e4a5e20b08f2168",
		Secret:     "SECe95a73fe9571b342edb9f50c89b87cb63a83wecsdfsdse1231zcsd152bf2",
	}
	fmt.Println(dingGroupConfig.SendTextMessage("测试", 0, ""))
	fmt.Println(dingGroupConfig.SendMarkdownMessage("# 测试123", "# 测试111 \n ## abc", 0, ""))
	fmt.Println(dingGroupConfig.SendLinkMessage("# 测试123", "# 测试111", "https://xingcxb.com", "", 0, ""))
	fmt.Println(dingGroupConfig.SendActionCardMessage("# 测试123", "# 测试111", 1, "测试", "https://xingcxb.com"))
}
