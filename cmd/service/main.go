package main

import (
	"github.com/atselvan/ankiconnect"

	"my/addToAnki/internal/models"
	"my/addToAnki/internal/services/anki"
	"my/addToAnki/internal/services/converters/anki/notes_from_csv"
)

func main() {
	ankiConnectClient := ankiconnect.NewClient()
	restErr := ankiConnectClient.Ping()
	if restErr != nil {
		panic(restErr)
	}

	ankiService := anki.New(ankiConnectClient)
	_ = ankiService
}

func doFromCSV(ankiService *anki.Service) {
	path := "data.csv"
	notes, err := notes_from_csv.NotesFromCSV(path)
	if err != nil {
		panic(err)
	}

	err = ankiService.AddNotes("Default", "Main", notes)
	if err != nil {
		panic(err)
	}
}

func doTestData(ankiService *anki.Service) {
	err := ankiService.AddNotes("Default", "Main", []models.Note{{
		"Sentence": "Слово",
		"Meaning":  "Слово",
		"IPA":      "Слово",
		"Target":   "Слово",
	}})
	if err != nil {
		panic(err)
	}
}
