package anki_adder_from_clipboard

import "my/addToAnki/internal/domain/models"

type ankiAdder interface {
	AddNotes(deck models.Deck, noteModel models.NoteModel, data []models.Fields) error
}
