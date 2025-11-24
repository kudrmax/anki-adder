package cli

import (
	"context"

	"my/addToAnki/internal/domain/models"
)

type (
	ankiAdder interface {
		AddNotes(deck models.Deck, noteModel models.NoteModel, data []models.Fields) error
	}

	sentenceSaver interface {
		Save(sentence string) error
	}

	llmGenerator interface {
		Generate(ctx context.Context, prompt string) (string, error)
	}
)
