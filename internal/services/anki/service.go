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
	cfg    Config
}

func New(ankiConnectClient *ankiconnect.Client, config Config) (*Service, error) {
	s := &Service{
		client: ankiConnectClient,
	}

	err := s.updateConfig(config)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Service) updateConfig(cfg Config) error {
	if !cfg.IsValid() {
		return errors.New("некорректная конфигурация сервиса anki")
	}

	s.cfg = cfg
	return nil
}

func (s *Service) AddNotes(notes []models.Note) error {
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

	if !slices.Contains(*decks, s.cfg.Deck) {
		return fmt.Errorf("колода %s не найдена среди колод", s.cfg.Deck)
	}

	for _, note := range notes {
		respError = s.client.Notes.Add(ankiconnect.Note{
			DeckName:  s.cfg.Deck,
			ModelName: s.cfg.NoteModel,
			Fields:    ankiconnect.Fields(note),
		})
		if respError != nil {
			return fmt.Errorf("ошибка при добавлении слова в колоду: %s", Error(respError))
		}
	}

	return nil
}
