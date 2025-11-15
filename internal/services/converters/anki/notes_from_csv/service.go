package notes_from_csv

import (
	"my/addToAnki/internal/models"
	"my/addToAnki/internal/services/files/csv"
)

// NotesFromCSV читает заметки из csv
func NotesFromCSV(path string) ([]models.Note, error) {
	header, records, err := csv.ReadCSV(path)
	if err != nil {
		return nil, err
	}

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
