package anki

import (
	"errors"
	"fmt"
	"slices"

	"github.com/atselvan/ankiconnect"

	"my/addToAnki/internal/models"
)

type Service struct {
	client *ankiconnect.Client
}

func New(ankiConnectClient *ankiconnect.Client) *Service {
	return &Service{
		client: ankiConnectClient,
	}
}

func (s *Service) AddNotes(deck, noteModel string, notes []models.Note) error {
	if len(notes) == 0 {
		return nil
	}

	decks, respError := s.client.Decks.GetAll()
	if respError != nil {
		return fmt.Errorf("ошибка при попытке получить колоды: %s", Error(respError))
	}
	if decks == nil || len(*decks) == 0 {
		return errors.New("колоды не найдены")
	}

	if !slices.Contains(*decks, deck) {
		return fmt.Errorf("колода %s не найдена среди колод", deck)
	}

	for _, note := range notes {
		respError = s.client.Notes.Add(ankiconnect.Note{
			DeckName:  deck,
			ModelName: noteModel,
			Fields:    ankiconnect.Fields(note),
		})
		if respError != nil {
			return fmt.Errorf("ошибка при добавлении слова в колоду: %s", Error(respError))
		}
	}

	return nil
}
