package cli

import (
	"flag"
	"fmt"
	"os"

	"my/addToAnki/internal/domain/models"
	"my/addToAnki/internal/infrastructure/source"
	"my/addToAnki/internal/usecases/anki/csv_parser"
)

func (cli *CLI) commandAdd(args []string) error {
	fs := flag.NewFlagSet("add", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)

	var (
		fromClipboard      = fs.Bool("clipboard", false, "read CSV from clipboard")
		fromClipboardShort = fs.Bool("c", false, "same as --clipboard")

		filePath      = fs.String("file", "", "path to CSV file")
		filePathShort = fs.String("f", "", "same as --file")
	)

	if err := fs.Parse(args); err != nil {
		return err
	}

	isClipboard := *fromClipboard || *fromClipboardShort
	isFile := *filePath != "" || *filePathShort != ""
	if isClipboard == isFile {
		return fmt.Errorf("you must specify exactly one of --clipboard/-c or --file/-f")
	}

	var file string
	if isFile {
		if *filePath != "" {
			file = *filePath
		} else if *filePathShort != "" {
			file = *filePathShort
		}
	}

	var readerGetter source.ReaderGetter
	switch {
	case isClipboard:
		readerGetter = source.NewClipboardReaderGetter()
	case isFile:
		readerGetter = source.NewFileReaderGetter(file)
	}
	CSVParser := csv_parser.New(cli.cfg.Fields, readerGetter)

	fields, err := CSVParser.Parse()
	if err != nil {
		return fmt.Errorf("parse csv: %w", err)
	}

	err = cli.ankiAdder.AddNotes(
		models.Deck(cli.cfg.Deck),
		models.NoteModel(cli.cfg.NoteModel),
		fields,
	)
	if err != nil {
		return fmt.Errorf("add notes: %w", err)
	}

	fmt.Println("Notes added successfully")
	return nil
}
