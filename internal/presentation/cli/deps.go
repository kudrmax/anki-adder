package cli

import (
	"my/addToAnki/internal/domain/models"
)

type (
	ankiAdder interface {
		AddNotes(deck models.Deck, noteModel models.NoteModel, data []models.Fields) error
	}

	sentenceSaver interface {
		SaveSentence(sentence string, target *string) error
		CopyNFirstSentencesToClipboard(n int) error
		DeleteNFirstSentences(n int) error
	}

	noteGenerator interface {
		GenerateNote(sentence, target string) (string, error)
	}

	ankiAdderFromClipboard interface {
		AddNotesFromClipboard(deck models.Deck, noteModel models.NoteModel) error
	}
)
