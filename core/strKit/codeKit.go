// Package strKit 代码工具包
package strKit

import (
	"bytes"
	"fmt"
	"github.com/alecthomas/chroma/v2/formatters"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
)

// HighlightCode 高亮代码
/*
 * @param code 代码
 * @param language 语言
 * @return string
 */
func HighlightCode(codeTxt, language string) (string, error) {
	lexer := lexers.Get(language)
	if lexer == nil {
		lexer = lexers.Fallback
	}
	iterator, err := lexer.Tokenise(nil, codeTxt)
	if err != nil {
		fmt.Println("Error tokenising code:", err)
		return codeTxt, err
	}
	// 内部支持的风格在 alecthomas/chroma/v2/styles 下查看xml文件即可
	style := styles.Get("monokai")
	if style == nil {
		style = styles.Fallback
	}
	var prettyCode bytes.Buffer
	formatter := formatters.Get("terminal256")
	if formatter == nil {
		formatter = formatters.Fallback
	}
	err = formatter.Format(&prettyCode, style, iterator)
	if err != nil {
		fmt.Println("Error formatting code:", err)
		return codeTxt, err
	}
	return prettyCode.String(), nil
}
