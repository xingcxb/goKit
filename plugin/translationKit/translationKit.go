package translationKit

type translationInterface interface {
	Translation(content, from, to string) (string, error)
}
