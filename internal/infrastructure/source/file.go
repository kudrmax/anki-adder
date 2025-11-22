package source

import (
	"io"
	"os"
)

type FileReaderGetter struct {
	Path string
}

func NewFileReaderGetter(path string) *FileReaderGetter {
	return &FileReaderGetter{Path: path}
}

func (f *FileReaderGetter) Open() (io.ReadCloser, error) {
	return os.Open(f.Path)
}
