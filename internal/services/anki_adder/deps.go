package anki_adder

import "my/addToAnki/internal/models"

type ankiAdderClient interface {
	AddNotes(notes []models.Note) error
}

type csvNoteExtracter interface {
	GetNotesFromCSVByFilePath(path string) ([]models.Note, error)
	GetNotesFromCSVByString(data string) ([]models.Note, error)
}
