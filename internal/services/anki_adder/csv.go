package anki_adder

import (
	"my/addToAnki/internal/services/converters/anki/notes_from_csv"
)

type AnkiAdderFromCSV struct {
	ankiAdderClient ankiAdderClient
	filePath        string
}

func NewAnkiAdderFromCSV(ankiAdderClient ankiAdderClient, filePath string) *AnkiAdderFromCSV {
	return &AnkiAdderFromCSV{
		ankiAdderClient: ankiAdderClient,
		filePath:        filePath,
	}
}

func (s *AnkiAdderFromCSV) AddNotes() error {
	notes, err := notes_from_csv.NotesFromCSV(s.filePath)
	if err != nil {
		return err
	}

	return s.ankiAdderClient.AddNotes(notes)
}
