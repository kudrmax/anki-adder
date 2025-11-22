package sentence_saver

type SentenceRepository interface {
	Save(sentence string) error
}
