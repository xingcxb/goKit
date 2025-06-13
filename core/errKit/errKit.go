package errKit

import "sync"

// 错误码定义
const (
	StatusOK            = 200 // 成功
	StatusParamsLack    = 201 // 参数缺失
	StatusDataLoss      = 202 // 数据不存在
	StatusDataParse     = 203 // 数据解析失败
	StatusLoadFile      = 300 // 加载文件失败
	StatusReadFile      = 301 // 读取文件失败
	StatusWriteFile     = 302 // 写入文件失败
	StatusPipe          = 303 // 管道错误
	StatusDbQuery       = 304 // 数据库查询错误
	StatusDbSave        = 305 // 数据库保存错误
	StatusDbUpdate      = 306 // 数据库更新错误
	StatusParseJson     = 400 // 解析json失败
	StatusCreatePty     = 401 // 创建pty失败
	StatusSSHCut        = 402 // SSH断开
	StatusSSHConnect    = 403 // SSH连接失败
	StatusSSHSession    = 404 // SSH会话失败
	StatusSSHPTY        = 405 // SSH PTY失败
	StatusSSHStartShell = 406 // SSH shell启动失败
)

// 错误码描述映射表（支持动态扩展）
var (
	mu         sync.RWMutex
	statusText = map[int]string{
		StatusOK:            "成功",
		StatusParamsLack:    "参数缺失",
		StatusDataLoss:      "数据不存在",
		StatusDataParse:     "数据解析失败",
		StatusLoadFile:      "加载文件失败",
		StatusReadFile:      "读取文件失败",
		StatusWriteFile:     "写入文件失败",
		StatusPipe:          "管道错误",
		StatusDbQuery:       "数据库查询错误",
		StatusDbSave:        "数据库保存错误",
		StatusDbUpdate:      "数据库更新错误",
		StatusParseJson:     "解析json失败",
		StatusCreatePty:     "创建pty失败",
		StatusSSHCut:        "SSH断开",
		StatusSSHConnect:    "SSH连接失败",
		StatusSSHSession:    "SSH会话失败",
		StatusSSHPTY:        "SSH PTY失败",
		StatusSSHStartShell: "SSH shell启动失败",
	}
)

// StatusText 返回错误码对应的错误描述
func StatusText(code int) string {
	mu.RLock()
	defer mu.RUnlock()
	if text, ok := statusText[code]; ok {
		return text
	}
	return "未知错误"
}

// RegisterStatus 允许主项目注册自定义错误码描述
func RegisterStatus(code int, text string) {
	mu.Lock()
	defer mu.Unlock()
	statusText[code] = text
}
