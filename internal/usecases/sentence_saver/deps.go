package sentence_saver

type SentenceRepository interface {
	Save(sentence string, target *string) error
	GetAll() ([]string, error)
	DeleteFirstNLines(n int) error
}
