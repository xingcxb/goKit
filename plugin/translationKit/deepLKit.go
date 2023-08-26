package translationKit

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
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
//func checkLanguage(languageName string) bool {
//	supportLanguage := map[string]string{
//		"Auto":                     "",
//		"Simplified Chinese":       "zh",
//		"Traditional Chinese":      "zh-tw",
//		"English":                  "en",
//		"Abkhazian":                "ab",
//		"Albanian":                 "sq",
//		"Akan":                     "ak",
//		"Arabic":                   "ar",
//		"Aragonese":                "an",
//		"Amharic":                  "am",
//		"Assamese":                 "as",
//		"Azerbaijani":              "az",
//		"Asturian":                 "ast",
//		"Central Huasteca Nahuatl": "nch",
//		"Ewe":                      "ee",
//		"Aymara":                   "ay",
//		"Irish":                    "ga",
//		"Estonian":                 "et",
//		"Ojibwa":                   "oj",
//		"Occitan":                  "oc",
//		"Oriya":                    "or",
//		"Oromo":                    "om",
//		"Ossetian":                 "os",
//		"Tok Pisin":                "tpi",
//		"Bashkir":                  "ba",
//		"Basque":                   "eu",
//		"Belarusian":               "be",
//		"Berber languages":         "ber",
//		"Bambara":                  "bm",
//		"Pangasinan":               "pag",
//		"Bulgarian":                "bg",
//		"Northern Sami":            "se",
//		"Bemba (Zambia)":           "bem",
//		"Blin":                     "byn",
//		"Bislama":                  "bi",
//		"Baluchi":                  "bal",
//		"Icelandic":                "is",
//		"Polish":                   "pl",
//		"Bosnian":                  "bs",
//		"Persian":                  "fa",
//		"Bhojpuri":                 "bho",
//		"Breton":                   "br",
//		"Chamorro":                 "ch",
//		"Chavacano":                "cbk",
//		"Chuvash":                  "cv",
//		"Tsonga":                   "ts",
//		"Tatar":                    "tt",
//		"Danish":                   "da",
//		"Shan":                     "shn",
//		"Tetum":                    "tet",
//		"German":                   "de",
//		"Low German":               "nds",
//		"Scots":                    "sco",
//		"Dhivehi":                  "dv",
//		"Kam":                      "kdx",
//		"Kadazan Dusun":            "dtp",
//		"Russian":                  "ru",
//		"Faroese":                  "fo",
//		"French":                   "fr",
//		"Sanskrit":                 "sa",
//		"Filipino":                 "fil",
//		"Fijian":                   "fj",
//		"Finnish":                  "fi",
//		"Friulian":                 "fur",
//		"Fur":                      "fvr",
//		"Kongo":                    "kg",
//		"Khmer":                    "km",
//		"Guerrero Nahuatl":         "ngu",
//		"Kalaallisut":              "kl",
//		"Georgian":                 "ka",
//		"Gronings":                 "gos",
//		"Gujarati":                 "gu",
//		"Guarani":                  "gn",
//		"Kazakh":                   "kk",
//		"Haitian":                  "ht",
//		"Korean":                   "ko",
//		"Hausa":                    "ha",
//		"Dutch":                    "nl",
//		"Montenegrin":              "cnr",
//		"Hupa":                     "hup",
//		"Gilbertese":               "gil",
//		"Rundi":                    "rn",
//		"K'iche'":                  "quc",
//		"Kirghiz":                  "ky",
//		"Galician":                 "gl",
//		"Catalan":                  "ca",
//		"Czech":                    "cs",
//		"Kabyle":                   "kab",
//		"Kannada":                  "kn",
//		"Kanuri":                   "kr",
//		"Kashubian":                "csb",
//		"Khasi":                    "kha",
//		"Cornish":                  "kw",
//		"Xhosa":                    "xh",
//		"Corsican":                 "co",
//		"Creek":                    "mus",
//		"Crimean Tatar":            "crh",
//		"Klingon":                  "tlh",
//		"Serbo-Croatian":           "hbs",
//		"Quechua":                  "qu",
//		"Kashmiri":                 "ks",
//		"Kurdish":                  "ku",
//		"Latin":                    "la",
//		"Latgalian":                "ltg",
//		"Latvian":                  "lv",
//		"Lao":                      "lo",
//		"Lithuanian":               "lt",
//		"Limburgish":               "li",
//		"Lingala":                  "ln",
//		"Ganda":                    "lg",
//		"Letzeburgesch":            "lb",
//		"Rusyn":                    "rue",
//		"Kinyarwanda":              "rw",
//		"Romanian":                 "ro",
//		"Romansh":                  "rm",
//		"Romany":                   "rom",
//		"Lojban":                   "jbo",
//		"Malagasy":                 "mg",
//		"Manx":                     "gv",
//		"Maltese":                  "mt",
//		"Marathi":                  "mr",
//		"Malayalam":                "ml",
//		"Malay":                    "ms",
//		"Mari (Russia)":            "chm",
//		"Macedonian":               "mk",
//		"Marshallese":              "mh",
//		"Kekchí":                   "kek",
//		"Maithili":                 "mai",
//		"Morisyen":                 "mfe",
//		"Maori":                    "mi",
//		"Mongolian":                "mn",
//		"Bengali":                  "bn",
//		"Burmese":                  "my",
//		"Hmong":                    "hmn",
//		"Umbundu":                  "umb",
//		"Navajo":                   "nv",
//		"Afrikaans":                "af",
//		"Nepali":                   "ne",
//		"Niuean":                   "niu",
//		"Norwegian":                "no",
//		"Pam":                      "pmn",
//		"Papiamento":               "pap",
//		"Panjabi":                  "pa",
//		"Portuguese":               "pt",
//		"Pushto":                   "ps",
//		"Nyanja":                   "ny",
//		"Twi":                      "tw",
//		"Cherokee":                 "chr",
//		"Japanese":                 "ja",
//		"Swedish":                  "sv",
//		"Samoan":                   "sm",
//		"Sango":                    "sg",
//		"Sinhala":                  "si",
//		"Upper Sorbian":            "hsb",
//		"Esperanto":                "eo",
//		"Slovenian":                "sl",
//		"Swahili":                  "sw",
//		"Somali":                   "so",
//		"Slovak":                   "sk",
//		"Tagalog":                  "tl",
//		"Tajik":                    "tg",
//		"Tahitian":                 "ty",
//		"Telugu":                   "te",
//		"Tamil":                    "ta",
//		"Thai":                     "th",
//		"Tonga (Tonga Islands)":    "to",
//		"Tonga (Zambia)":           "toi",
//		"Tigrinya":                 "ti",
//		"Tuvalu":                   "tvl",
//		"Tuvinian":                 "tyv",
//		"Turkish":                  "tr",
//		"Turkmen":                  "tk",
//		"Walloon":                  "wa",
//		"Waray (Philippines)":      "war",
//		"Welsh":                    "cy",
//		"Venda":                    "ve",
//		"Volapük":                  "vo",
//		"Wolof":                    "wo",
//		"Udmurt":                   "udm",
//		"Urdu":                     "ur",
//		"Uzbek":                    "uz",
//		"Spanish":                  "es",
//		"Interlingue":              "ie",
//		"Western Frisian":          "fy",
//		"Silesian":                 "szl",
//		"Hebrew":                   "he",
//		"Hiligaynon":               "hil",
//		"Hawaiian":                 "haw",
//		"Modern Greek":             "el",
//		"Lingua Franca Nova":       "lfn",
//		"Sindhi":                   "sd",
//		"Hungarian":                "hu",
//		"Shona":                    "sn",
//		"Cebuano":                  "ceb",
//		"Syriac":                   "syr",
//		"Sundanese":                "su",
//		"Armenian":                 "hy",
//		"Achinese":                 "ace",
//		"Iban":                     "iba",
//		"Igbo":                     "ig",
//		"Ido":                      "io",
//		"Iloko":                    "ilo",
//		"Inuktitut":                "iu",
//		"Italian":                  "it",
//		"Yiddish":                  "yi",
//		"Interlingua":              "ia",
//		"Hindi":                    "hi",
//		"Indonesia":                "id",
//		"Ingush":                   "inh",
//		"Yoruba":                   "yo",
//		"Vietnamese":               "vi",
//		"Zaza":                     "zza",
//		"Javanese":                 "jv",
//		"Chinese":                  "zh",
//		"Cantonese":                "yue",
//		"Zulu":                     "zu",
//	}
//
//}

type DeepL struct {
}

// Translation DeepL翻译
/**
 * @param content 待翻译的内容
 * @param from 源语言
 * @param to 目标语言
 */
func (d *DeepL) Translation(content, from, to string) (string, error) {
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
