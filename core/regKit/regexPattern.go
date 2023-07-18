package regKit

const (
	GeneralPattern  = "^\\w+$"                                                                                                                                                                                                                        // GeneralPattern 英文字母 、数字和下划线
	NumbersPattern  = "\\d+"                                                                                                                                                                                                                          // NumbersPattern 数字
	WordPattern     = "[a-zA-Z]+"                                                                                                                                                                                                                     // WordPattern 英文字母
	ChinesePattern  = "[\\u4E00-\\u9FFF]"                                                                                                                                                                                                             // ChinesePattern 单个中文汉字
	ChinesesPattern = ChinesePattern + "+"                                                                                                                                                                                                            // ChinesesPattern 多个中文汉字
	FilePattern     = "\\.(doc|docx|log|odt|pages|rtf|txt|wps|csv|dat|key|pps|ppt|pptx|tar|m3u|m4a|mid|mp3|mpa|wav|wma|3gp|avi|flv|m4v|mov|mp4|mpg|rm|srt|swf|vob|wmv|psd|ai|pdf|xls|xlsx|sql|apk|so|app|bat|cgi|exe|jar|7z|gz|zip|rar|dmg|iso|msi)$" // FilePattern 文件后缀名
)
