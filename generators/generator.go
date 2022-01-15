package generators

type Generator interface {
	Generate(data interface{}) (string, error)
}
