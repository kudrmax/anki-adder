package anki_adder

import (
	"fmt"

	"my/addToAnki/internal/domain/models"
)

type UseCase struct {
	client Client
}

func NewUseCase(client Client) *UseCase {
	return &UseCase{
		client: client,
	}
}

// AddNotes add notes to anki deck.
// If deck doesn't  exist, it will be added to default deck
// Default deck name is "Default"
func (uc *UseCase) AddNotes(deck models.Deck, noteModel models.NoteModel, data []models.Fields) error {
	exists, err := uc.client.IsDeckExists(deck)
	if err != nil {
		return fmt.Errorf("error on checking deck: %w", err)
	}
	if !exists {
		deck = defaultDeck
	}

	notesRow := make([]models.NoteRow, 0, len(data))
	for _, fields := range data {
		notesRow = append(notesRow, models.NoteRow{
			Deck:      deck,
			NoteModel: noteModel,
			Fields:    fields.ConvertHTMLTags(),
		})
	}

	err = uc.client.AddBatch(notesRow)
	if err != nil {
		return fmt.Errorf("error on adding notes: %w", err)
	}
	return nil
}
