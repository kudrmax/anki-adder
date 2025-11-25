package gui

type Saver interface {
	Save(string) error
}

type Generator interface {
	Generate(text string) string
}

type NextProvider interface {
	Next() string
}
