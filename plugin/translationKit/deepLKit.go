package translationKit

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/xingcxb/goKit/core/arrayKit"
	"github.com/xingcxb/goKit/core/cryptoKit"
	"github.com/xingcxb/goKit/core/httpKit"
	"github.com/xingcxb/goKit/core/randomKit"
	"strings"
	"time"
)

const (
	deepLUrl = "https://www2.deepl.com/jsonrpc"
)

// 检查语言是否支持
func checkLanguage(languageName string) bool {
	supportLanguage := []string{"zh", "zh-tw", "en", "ab", "sq", "ak", "ar", "an", "am", "as", "az", "ast", "nch", "ee", "ay", "ga", "et", "oj", "oc", "or", "om", "os", "tpi", "ba", "eu", "be", "ber", "bm", "pag", "bg", "se", "bem", "byn", "bi", "bal", "is", "pl", "bs", "fa", "bho", "br", "ch", "cbk", "cv", "ts", "tt", "da", "shn", "tet", "de", "nds", "sco", "dv", "kdx", "dtp", "ru", "fo", "fr", "sa", "fil", "fj", "fi", "fur", "fvr", "kg", "km", "ngu", "kl", "ka", "gos", "gu", "gn", "kk", "ht", "ko", "ha", "nl", "cnr", "hup", "gil", "rn", "quc", "ky", "gl", "ca", "cs", "kab", "kn", "kr", "csb", "kha", "kw", "xh", "co", "mus", "crh", "tlh", "hbs", "qu", "ks", "ku", "la", "ltg", "lv", "lo", "lt", "li", "ln", "lg", "lb", "rue", "rw", "ro", "rm", "rom", "jbo", "mg", "gv", "mt", "mr", "ml", "ms", "chm", "mk", "mh", "kek", "mai", "mfe", "mi", "mn", "bn", "my", "hmn", "umb", "nv", "af", "ne", "niu", "no", "pmn", "pap", "pa", "pt", "ps", "ny", "tw", "chr", "ja", "sv", "sm", "sg", "si", "hsb", "eo", "sl", "sw", "so", "sk", "tl", "tg", "ty", "te", "ta", "th", "to", "toi", "ti", "tvl", "tyv", "tr", "tk", "wa", "war", "cy", "ve", "vo", "wo", "udm", "ur", "uz", "es", "ie", "fy", "szl", "he", "hil", "haw", "el", "lfn", "sd", "hu", "sn", "ceb", "syr", "su", "hy", "ace", "iba", "ig", "io", "ilo", "iu", "it", "yi", "ia", "hi", "id", "inh", "yo", "vi", "zza", "jv", "zh", "yue", "zu"}
	return arrayKit.Contains(supportLanguage, languageName)
}

type DeepL struct {
}

// Translation DeepL翻译
/**
 * @param content 待翻译的内容
 * @param from 源语言
 * @param to 目标语言
 */
func (d *DeepL) Translation(content, from, to string) (string, error) {
	// 检查语言是否支持
	if !checkLanguage(from) || !checkLanguage(to) {
		return "", fmt.Errorf("暂不支持该语言翻译")
	}
	// 头部信息
	headers := map[string]string{
		"accept":          "*/*",
		"content-type":    "application/json",
		"user-agent":      "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_0_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36",
		"accept-language": "zh-Hans-CN;q=1, en-CN;q=0.9",
	}
	// id随机数
	idRandom := randomKit.RandomLong(99999) + 8300000*1e3
	// 请求参数
	dataMap := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "LMT_handle_texts",
		"params": map[string]interface{}{
			"splitting": "newlines",
			"lang": map[string]interface{}{
				"source_lang_user_selected": strings.ToUpper(from),
				"target_lang":               strings.ToUpper(to),
			},
			"texts": []map[string]interface{}{
				{
					"text":                content,
					"requestAlternatives": 3,
				},
			},
			"timestamp": int(time.Now().UnixNano() / int64(time.Millisecond)),
		},
		"id": idRandom,
	}
	dataBytes, err := json.Marshal(dataMap)
	if err != nil {
		return "", err
	}
	data := string(dataBytes)
	if (idRandom+5)%29 == 0 || (idRandom+3)%13 == 0 {
		data = strings.ReplaceAll(data, "\"method\":\"", "\"method\" : \"")
	} else {
		data = strings.ReplaceAll(data, "\"method\":\"", "\"method\": \"")
	}
	response, err := httpKit.HttpPostFull(deepLUrl, headers, nil, data, -1)
	if err != nil {
		return "", err
	}
	code := gjson.Get(response, "error.code").Int()
	if code != 0 {
		return "", fmt.Errorf("翻译失败，错误码：%d", code)
	}
	bestResult := gjson.Get(response, "result.texts").String()
	// 解码unicode
	bestResult, _ = cryptoKit.UnicodeDecode(bestResult)
	return bestResult, nil
}
