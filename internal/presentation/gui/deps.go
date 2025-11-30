package gui

import "my/addToAnki/internal/domain/models"

type Saver interface {
	Save(string) error
	Copy(n int) error
	DeleteFirstNLines(n int) error
}

type Generator interface {
	GenerateNote(sentence, target string) (string, error)
}

type NextProvider interface {
	Next() string
}

type ankiAdder interface {
	AddNotes(deck models.Deck, noteModel models.NoteModel, data []models.Fields) error
}
