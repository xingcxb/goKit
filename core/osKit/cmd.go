package osKit

const (
	// CmdOsSyncLocalTime 同步本地时间命令
	CmdOsSyncLocalTime = "nohup ntpdate -u ntp1.aliyun.com >/dev/null 2>&1 &"
)
