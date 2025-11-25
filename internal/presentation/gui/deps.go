package gui

type Saver interface {
	Save(text string)
}

type Generator interface {
	Generate(text string) string
}

type NextProvider interface {
	Next() string
}
