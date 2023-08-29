package translationKit

// Translation 翻译
type translationInterface interface {
	Translation(content, from, to string) (string, error)
}
