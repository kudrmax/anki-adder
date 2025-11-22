package csv_parser

import (
	"fmt"

	"github.com/privatesquare/bkst-go-utils/utils/errors"

	modelsNew "my/addToAnki/internal/domain/models"
	"my/addToAnki/internal/usecases/parsers/csv"
)

type CSVParser struct {
	FieldNames []string
	Reader     Reader
}

func New(fieldNames []string, reader Reader) *CSVParser {
	return &CSVParser{
		FieldNames: fieldNames,
		Reader:     reader,
	}
}

func (p *CSVParser) Parse() ([]modelsNew.Fields, error) {
	r, err := p.Reader.Open()
	if err != nil {
		return nil, err
	}
	defer r.Close()

	records, err := csv.ParseCSV(r)
	if err != nil {
		return nil, fmt.Errorf("failed to parse csv: %v", err)
	}

	notes, err := p.convert(records)
	if err != nil {
		return nil, fmt.Errorf("failed to convert csv to notes: %v", err)
	}

	return notes, nil
}

func (p *CSVParser) convert(records [][]string) ([]modelsNew.Fields, error) {
	result := make([]modelsNew.Fields, 0, len(records))

	for _, row := range records {
		if len(row) != len(p.FieldNames) {
			return nil, errors.New("invalid number of fields in cvs")
		}

		fields := make(modelsNew.Fields)
		for i, val := range row {
			fields[p.FieldNames[i]] = val
		}

		result = append(result, fields)
	}

	return result, nil
}
