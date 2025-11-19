package csv_reader

import (
	"encoding/csv"
	"os"
	"strings"

	"my/addToAnki/internal/models"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) GetNotesFromCSVByFilePath(path string) ([]models.Note, error) {
	data, err := s.readCSV(path)
	if err != nil {
		return nil, err
	}

	records, err := s.readString(data)
	if err != nil {
		return nil, err
	}
	if len(records) == 0 || len(records) == 1 {
		return nil, nil
	}

	notes, err := s.notesFromCSV(records[0], records[1:])
	if err != nil {
		return nil, err
	}

	return notes, nil
}

func (s *Service) GetNotesFromCSVByString(data string) ([]models.Note, error) {
	records, err := s.readString(data)
	if err != nil {
		return nil, err
	}
	if len(records) == 0 || len(records) == 1 {
		return nil, nil
	}

	notes, err := s.notesFromCSV(records[0], records[1:])
	if err != nil {
		return nil, err
	}

	return notes, nil
}

func (s *Service) readCSV(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (s *Service) readString(data string) ([][]string, error) {
	reader := csv.NewReader(strings.NewReader(data))
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func (s *Service) notesFromCSV(header []string, records [][]string) ([]models.Note, error) {
	result := make([]models.Note, 0, len(records))
	for _, record := range records {
		noteMap := make(map[string]string)
		for i := range len(header) {
			fieldName, value := header[i], record[i]
			noteMap[fieldName] = value
		}
		result = append(result, noteMap)
	}

	return result, nil
}
