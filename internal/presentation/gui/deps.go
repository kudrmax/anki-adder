package gui

import "my/addToAnki/internal/domain/models"

type saver interface {
	SaveSentence(sentence string, target *string) error
	CopyNFirstSentencesToClipboard(n int) error
	DeleteNFirstSentences(n int) error
}

type generator interface {
	GenerateNote(sentence, target string) (string, error)
}

type nextProvider interface {
	Next() string
}

type ankiAdder interface {
	AddNotes(deck models.Deck, noteModel models.NoteModel, data []models.Fields) error
}

type ankiAdderFromClipboard interface {
	AddNotesFromClipboard(deck models.Deck, noteModel models.NoteModel) error
}
