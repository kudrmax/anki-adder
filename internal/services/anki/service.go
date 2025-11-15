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

func (s *Service) AddWords(deck string, words []models.Word) error {
	if len(words) == 0 {
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

	for _, word := range words {
		wordMap := word.Map()

		note := ankiconnect.Note{
			DeckName:  deck,
			ModelName: word.ModelName,
			Fields:    ankiconnect.Fields(wordMap),
		}

		respError = s.client.Notes.Add(note)
		if respError != nil {
			return fmt.Errorf("ошибка при добавлении слова %s в колоду: %s", word.Sentence(), Error(respError))
		}
	}

	return nil
}
