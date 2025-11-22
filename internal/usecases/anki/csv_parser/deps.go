package csv_parser

import "io"

type Reader interface {
	Open() (io.ReadCloser, error)
}
