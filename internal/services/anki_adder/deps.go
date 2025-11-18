package anki_adder

import "my/addToAnki/internal/models"

type ankiAdderClient interface {
	AddNotes(notes []models.Note) error
}
