package csv

import (
	"encoding/csv"
	"io"
)

func ParseCSV(r io.ReadCloser) (records [][]string, err error) {
	return csv.NewReader(r).ReadAll()
}
