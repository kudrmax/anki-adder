package source

import (
	"fmt"
	"io"
	"strings"

	"github.com/atotto/clipboard"
)

type ClipboardReaderGetter struct{}

func NewClipboardReaderGetter() *ClipboardReaderGetter {
	return &ClipboardReaderGetter{}
}

func (f *ClipboardReaderGetter) Open() (io.ReadCloser, error) {
	data, err := clipboard.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read from clipboard: %w", err)
	}

	return io.NopCloser(strings.NewReader(data)), nil
}
