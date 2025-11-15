package csv

import (
	"encoding/csv"
	"os"
)

// ReadCSV превращает csv в массив из массива строк
func ReadCSV(path string) (header []string, data [][]string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, nil, err
	}

	if len(records) == 0 {
		return nil, nil, nil
	}

	return records[0], records[1:], nil
}
