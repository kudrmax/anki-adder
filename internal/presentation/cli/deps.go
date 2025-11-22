package cli

import "my/addToAnki/internal/domain/models"

type AnkiAdder interface {
	AddNotes(deck models.Deck, noteModel models.NoteModel, data []models.Fields) error
}
