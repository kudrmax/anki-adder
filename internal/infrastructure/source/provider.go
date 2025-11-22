package source

import "io"

type ReaderGetter interface {
	Open() (io.ReadCloser, error)
}
