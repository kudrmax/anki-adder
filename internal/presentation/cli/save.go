package cli

import (
	"errors"
	"fmt"
	"io"

	"my/addToAnki/internal/infrastructure/source"
)

func (cli *CLI) commandSave(args []string) error {
	if len(args) > 1 {
		return errors.New("usage: save <file> or save")
	}

	isFromClipboard := len(args) == 0
	isFromNextArgument := len(args) == 1

	var line string
	switch {
	case isFromClipboard:
		readerGetter := source.NewClipboardReaderGetter()
		r, err := readerGetter.Open()
		if err != nil {
			return fmt.Errorf("open clipboard: %w", err)
		}
		defer r.Close()

		data, err := io.ReadAll(r)
		if err != nil {
			return fmt.Errorf("read clipboard: %w", err)
		}
		line = string(data)
	case isFromNextArgument:
		line = args[0]
	}
	err := cli.sentenceSaver.Save(line)
	if err != nil {
		return fmt.Errorf("save sentence: %w", err)
	}

	fmt.Println("Sentence saved successfully")
	return nil
}
