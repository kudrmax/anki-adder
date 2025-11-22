package main

import (
	"log"
	"os"

	ankiconnectExternal "github.com/atselvan/ankiconnect"

	"my/addToAnki/config"
	"my/addToAnki/internal/domain/models"
	"my/addToAnki/internal/infrastructure/ankiconnect"
	"my/addToAnki/internal/infrastructure/source"
	"my/addToAnki/internal/usecases/anki/anki_adder"
	"my/addToAnki/internal/usecases/anki/csv_parser"
)

const (
	// TODO заменить на чтение из cli
	csvFilePath = "data.csv"
	configPath  = "config.yaml" // TODO временно -> заменить на конфиг лежащий в папке .config
)

func main() {
	f, err := os.Open(configPath)
	if err != nil {
		log.Fatal(err)
	}
	cfg, err := config.Parse(f)
	if err != nil {
		log.Fatal(err)
	}

	_ = cfg

	ankiConnectExternalClient := ankiconnectExternal.NewClient()
	restErr := ankiConnectExternalClient.Ping()
	if restErr != nil {
		log.Fatalf("failed connecting to AnkiConnect", ankiconnect.NewClientError(restErr))
	}
	ankiConnectClient := ankiconnect.New(ankiConnectExternalClient)

	ankiUseCase := anki_adder.NewUseCase(ankiConnectClient)

	fileReaderGetter := source.NewFileReaderGetter(csvFilePath)
	clipboardReaderGetter := source.NewClipboardReaderGetter()

	fileCSVParser := csv_parser.New(cfg.Fields, fileReaderGetter)
	clipboardCSVParser := csv_parser.New(cfg.Fields, clipboardReaderGetter)

	fields, err := clipboardCSVParser.Parse()
	if err != nil {
		log.Fatal(err)
	}
	err = ankiUseCase.AddNotes(models.Deck(cfg.Deck), models.NoteModel(cfg.NoteModel), fields)
	if err != nil {
		log.Fatal(err)
	}

	_ = fileCSVParser
}
