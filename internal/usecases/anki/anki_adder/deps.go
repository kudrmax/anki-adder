package anki_adder

import "my/addToAnki/internal/domain/models"

type Client interface {
	IsDeckExists(deck models.Deck) (bool, error)
	AddBatch(notesRow []models.NoteRow) error
}
