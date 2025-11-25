package ankiconnect

import (
	"fmt"
	"slices"

	"github.com/atselvan/ankiconnect"

	modelsNew "my/addToAnki/internal/domain/models" // TODO rename
)

type Client struct {
	client *ankiconnect.Client
}

func New(ankiConnectClient *ankiconnect.Client) *Client {
	return &Client{
		client: ankiConnectClient,
	}
}

// GetAllDecks returns a list of all decks
func (c *Client) GetAllDecks() ([]modelsNew.Deck, error) {
	decks, respError := c.client.Decks.GetAll()
	if respError != nil {
		return nil, fmt.Errorf("error on getting all decks: %s", NewClientError(respError))
	}
	if decks == nil || len(*decks) == 0 {
		return nil, nil
	}

	result := make([]modelsNew.Deck, 0, len(*decks))
	for _, deck := range *decks {
		result = append(result, modelsNew.Deck(deck))
	}
	return result, nil
}

// IsDeckExists checks if a deck exists
func (c *Client) IsDeckExists(deck modelsNew.Deck) (bool, error) {
	decks, err := c.GetAllDecks()
	if err != nil {
		return false, fmt.Errorf("error on checking if deck exists: %w", err)
	}
	return slices.Contains(decks, deck), nil
}

// Add adds a note to the deck
func (c *Client) Add(noteRow modelsNew.NoteRow) error {
	if err := noteRow.Deck.IsValid(); err != nil {
		return fmt.Errorf("invalid deck: %w", err)
	}
	if err := noteRow.NoteModel.IsValid(); err != nil {
		return fmt.Errorf("invalid note model: %w", err)
	}

	if noteRow.Fields == nil || len(noteRow.Fields) == 0 {
		return nil
	}

	respError := c.client.Notes.Add(ankiconnect.Note{
		DeckName:  noteRow.Deck.String(),
		ModelName: noteRow.NoteModel.String(),
		Fields:    noteRow.Fields.Map(),
	})
	if respError != nil {
		return fmt.Errorf("error on adding note: %s", NewClientError(respError))
	}

	return nil
}

// AddBatch adds a notes to the deck
func (c *Client) AddBatch(notesRow []modelsNew.NoteRow) error {
	for _, noteRow := range notesRow {
		err := c.Add(noteRow)
		if err != nil {
			return fmt.Errorf("error on adding batch on note: %w", err)
		}
	}

	return nil
}
