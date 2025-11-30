package anki_adder_from_clipboard

import (
	"my/addToAnki/config"
	"my/addToAnki/internal/domain/models"
	"my/addToAnki/internal/infrastructure/source"
	"my/addToAnki/internal/usecases/anki/csv_parser"
)

type Controller struct {
	cfg       config.Config
	ankiAdder ankiAdder
}

func New(
	cfg config.Config,
	ankiAdder ankiAdder,
) *Controller {
	return &Controller{
		cfg:       cfg,
		ankiAdder: ankiAdder,
	}
}

func (c *Controller) AddNotesFromClipboard(deck models.Deck, noteModel models.NoteModel) error {
	readerGetter := source.NewClipboardReaderGetter()
	CSVParser := csv_parser.New(c.cfg.Fields, readerGetter)

	data, err := CSVParser.Parse()
	if err != nil {
		return err
	}

	return c.ankiAdder.AddNotes(deck, noteModel, data)
}
