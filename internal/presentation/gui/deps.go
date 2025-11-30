package gui

type Saver interface {
	Save(string) error
	Copy(n int) error
}

type Generator interface {
	GenerateNote(sentence, target string) (string, error)
}

type NextProvider interface {
	Next() string
}
