// Package regKit 常用正则表达式字符串集合
package regKit

const (
	// GeneralPattern 英文字母 、数字和下划线
	GeneralPattern = "^\\w+$"
	// GeneralWithChinesePattern 中文字、英文字母、数字和下划线
	GeneralWithChinesePattern = "^[\u4E00-\u9FFF\\w]+$"
	// NumbersPattern 数字
	NumbersPattern = "\\d+"
	// WordPattern 英文字母
	WordPattern = "[a-zA-Z]+"
	// ChinesePattern 单个中文汉字
	ChinesePattern = "[\\u2E80-\\u2EFF\\u2F00-\\u2FDF\\u31C0-\\u31EF\\u3400-\\u4DBF\\u4E00-\\u9FFF\\uF900-\\uFAFF\\uD840\\uDC00-\\uD869\\uDEDF\\uD869\\uDF00-\\uD86D\\uDF3F\\uD86D\\uDF40-\\uD86E\\uDC1F\\uD86E\\uDC20-\\uD873\\uDEAF\\uD87E\\uDC00-\\uD87E\\uDE1F]"
	// ChinesesPattern 多个中文汉字
	ChinesesPattern = ChinesePattern + "+"
	// GroupVarPattern 分组
	GroupVarPattern = "\\$(\\d+)"
	// IpV4Pattern IpV4
	IpV4Pattern = "^(25[0-5]|2[0-4]\\d|[0-1]?\\d?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]?\\d?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]?\\d?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]?\\d?\\d)$"
	// IpV6Pattern IpV6
	IpV6Pattern = "(([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]+|::(ffff(:0{1,4})?:)?((25[0-5]|(2[0-4]|1?[0-9])?[0-9])\\.){3}(25[0-5]|(2[0-4]|1?[0-9])?[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1?[0-9])?[0-9])\\.){3}(25[0-5]|(2[0-4]|1?[0-9])?[0-9]))"
	// MoneyPattern 货币
	MoneyPattern = "^(\\d+(?:\\.\\d+)?)$"
	// EmailPattern 邮箱地址 符合RFC 5322规范，正则来自：http://emailregex.com/
	EmailPattern = "(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21\\x23-\\x5b\\x5d-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21-\\x5a\\x53-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])+)])"
	// EmailWithChinesePattern 规则同EMAIL，添加了对中文的支持
	EmailWithChinesePattern = "(?:[a-z0-9\\u4e00-\\u9fa5!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9\\u4e00-\\u9fa5!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21\\x23-\\x5b\\x5d-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])*\")@(?:(?:[a-z0-9\\u4e00-\\u9fa5](?:[a-z0-9\\u4e00-\\u9fa5-]*[a-z0-9\\u4e00-\\u9fa5])?\\.)+[a-z0-9\\u4e00-\\u9fa5](?:[a-z0-9\\u4e00-\\u9fa5-]*[a-z0-9\\u4e00-\\u9fa5])?|\\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9\\u4e00-\\u9fa5-]*[a-z0-9\\u4e00-\\u9fa5]:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21-\\x5a\\x53-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])+)])"
	// MobileCNPattern 移动电话-中国大陆 +86 2位区域码标示+11位数字
	MobileCNPattern = "(?:0|86|\\+86)?1[3-9]\\d{9}"
	// MobileHKPattern 移动电话-中国香港 +052 三位区域码+10位数字, 中国香港手机号码8位数
	MobileHKPattern = "(?:0|852|\\+852)?\\d{8}"
	// MobileMOPattern 移动电话-中国澳门 中国澳门 +853 Macao 国际域名缩写：MO 三位区域码 +号码以数字6开头 + 7位数字, 中国澳门手机号码8位数
	MobileMOPattern = "(?:0|853|\\+853)?(?:|-)6\\d{7}"
	// MobileTWPattern 移动电话-中国台湾 +886 三位区域码+号码以数字09开头 + 8位数字, 中国台湾手机号码10位数
	MobileTWPattern = "(?:0|886|\\+886)?(?:|-)09\\d{8}"
	// CitizenIdPattern 18位身份证号码
	CitizenIdPattern = "[1-9]\\d{5}[1-2]\\d{3}((0\\d)|(1[0-2]))(([012]\\d)|3[0-1])\\d{3}(\\d|X|x)"
	// ZipCodePattern 邮编，兼容港澳台
	ZipCodePattern = "^(0[1-7]|1[0-356]|2[0-7]|3[0-6]|4[0-7]|5[0-7]|6[0-7]|7[0-5]|8[0-9]|9[0-8])\\d{4}|99907[78]$"
	// BirthdayPattern 生日
	BirthdayPattern = "^(\\d{2,4})([/\\-.年]?)(\\d{1,2})([/\\-.月]?)(\\d{1,2})日?$"
	// UriPattern uri 定义见：https://www.ietf.org/rfc/rfc3986.html#appendix-B
	UriPattern = "^(([^:/?#]+):)?(//([^/?#]*))?([^?#]*)(\\?([^#]*))?(#(.*))?"
	// UrlPattern url
	UrlPattern = "[a-zA-Z]+://[\\w-+&@#/%?=~_|!:,.;]*[\\w-+&@#/%=~_|]"
	// UrlHttpPattern Http URL（来自：http://urlregex.com/）此正则同时支持FTP、File等协议的URL
	UrlHttpPattern = "(https?|ftp|file)://[\\w-+&@#/%?=~_|!:,.;]*[\\w-+&@#/%=~_|]"
	// UUIDPattern uuid
	UUIDPattern = "^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$"
	// UUIDSimplePattern 不带横线的UUID
	UUIDSimplePattern = "^[0-9a-fA-F]{32}$"
	// MacAddressPattern MAC地址正则
	MacAddressPattern = "((?:[a-fA-F0-9]{1,2}[:-]){5}[a-fA-F0-9]{1,2})|0x(\\d{12}).+ETHER"
	// HEXPattern 16进制字符串
	HEXPattern = "^[a-fA-F0-9]+$"
	// TimePattern 时间正则
	TimePattern = "\\d{1,2}:\\d{1,2}(:\\d{1,2})?"
	// PlateNumberPattern 中国车牌号码（兼容新能源车牌）
	PlateNumberPattern = "^(([京津沪渝冀豫云辽黑湘皖鲁新苏浙赣鄂桂甘晋蒙陕吉闽贵粤青藏川宁琼使领][A-Z](([0-9]{5}[ABCDEFGHJK])|([ABCDEFGHJK]([A-HJ-NP-Z0-9])[0-9]{4})))|([京津沪渝冀豫云辽黑湘皖鲁新苏浙赣鄂桂甘晋蒙陕吉闽贵粤青藏川宁琼使领]\\d{3}\\d{1,3}[领])|([京津沪渝冀豫云辽黑湘皖鲁新苏浙赣鄂桂甘晋蒙陕吉闽贵粤青藏川宁琼使领][A-Z][A-HJ-NP-Z0-9]{4}[A-HJ-NP-Z0-9挂学警港澳使领]))$"
	// CreditCodePattern 统一社会信用代码
	/**
	 * 第一部分：登记管理部门代码1位 (数字或大写英文字母)
	 * 第二部分：机构类别代码1位 (数字或大写英文字母)
	 * 第三部分：登记管理机关行政区划码6位 (数字)
	 * 第四部分：主体标识码（组织机构代码）9位 (数字或大写英文字母)
	 * 第五部分：校验码1位 (数字或大写英文字母)
	 */
	CreditCodePattern = "^[0-9A-HJ-NPQRTUWXY]{2}\\d{6}[0-9A-HJ-NPQRTUWXY]{10}$"
	// CarVinPattern 车架号（车辆识别代号由世界制造厂识别代号(WMI、车辆说明部分(VDS)车辆指示部分(VIS)三部分组成，共 17 位字码。
	/**
	 * 别名：车辆识别代号、车辆识别码、车架号、十七位码<br>
	 * 标准号：GB 16735-2019<br>
	 * 标准官方地址：https://openstd.samr.gov.cn/bzgk/gb/newGbInfo?hcno=E2EBF667F8C032B1EDFD6DF9C1114E02
	 * 对年产量大于或等于1 000 辆的完整车辆和/或非完整车辆制造厂：
	 * <pre>
	 *   第一部分为世界制造厂识别代号(WMI)，3位
	 *   第二部分为车辆说明部分(VDS)，     6位
	 *   第三部分为车辆指示部分(VIS)，     8位
	 * </pre>
	 *
	 * 对年产量小于 1 000 辆的完整车辆和/或非完整车辆制造厂：
	 * <pre>
	 *   第一部分为世界制造广识别代号(WMI),3位;
	 *   第二部分为车辆说明部分(VDS)，6位;
	 *   第三部分的三、四、五位与第一部分的三位字码起构成世界制造厂识别代号(WMI),其余五位为车辆指示部分(VIS)，8位。
	 * </pre>
	 *
	 * <pre>
	 *   eg:LDC613P23A1305189
	 *   eg:LSJA24U62JG269225
	 *   eg:LBV5S3102ESJ25655
	 * </pre>
	 */
	CarVinPattern = "^[A-HJ-NPR-Z0-9]{8}[X0-9]([A-HJ-NPR-Z0-9]{3}\\d{5}|[A-HJ-NPR-Z0-9]{5}\\d{3})$"
	// CarDrivingLicencePattern  驾驶证  别名：驾驶证档案编号、行驶证编号 仅限：中国驾驶证档案编号
	CarDrivingLicencePattern = "^[0-9]{12}$"
	// ChineseNamePattern 中文姓名
	/**
	 * 维吾尔族姓名里面的点是 · 输入法中文状态下，键盘左上角数字1前面的那个符号；<br>
	 * 错误字符：{@code ．.。．.}<br>
	 * 正确维吾尔族姓名：
	 * <pre>
	 * 霍加阿卜杜拉·麦提喀斯木
	 * 玛合萨提别克·哈斯木别克
	 * 阿布都热依木江·艾斯卡尔
	 * 阿卜杜尼亚孜·毛力尼亚孜
	 * </pre>
	 * <pre>
	 * ----------
	 * 错误示例：孟  伟                reason: 有空格
	 * 错误示例：连逍遥0               reason: 数字
	 * 错误示例：依帕古丽-艾则孜        reason: 特殊符号
	 * 错误示例：牙力空.买提萨力        reason: 新疆人的点不对
	 * 错误示例：王建鹏2002-3-2        reason: 有数字、特殊符号
	 * 错误示例：雷金默(雷皓添）        reason: 有括号
	 * 错误示例：翟冬:亮               reason: 有特殊符号
	 * 错误示例：李                   reason: 少于2位
	 * ----------
	 * </pre>
	 * 总结中文姓名：2-60位，只能是中文和维吾尔族的点·
	 * 放宽汉字范围：如生僻姓名 刘欣䶮yǎn
	 */
	ChineseNamePattern = "^[\u2E80-\u9FFF·]{2,60}$"
	// FilePattern 文件后缀名
	FilePattern = "\\.(doc|docx|log|odt|pages|rtf|txt|wps|csv|dat|key|pps|ppt|pptx|tar|m3u|m4a|mid|mp3|mpa|wav|wma|3gp|avi|flv|m4v|mov|mp4|mpg|rm|srt|swf|vob|wmv|psd|ai|pdf|xls|xlsx|sql|apk|so|app|bat|cgi|exe|jar|7z|gz|zip|rar|dmg|iso|msi)$"
)
