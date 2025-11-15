package main

import (
	"log"

	"github.com/atselvan/ankiconnect"

	"my/addToAnki/internal/services/anki"
)

func main() {
	ankiConnectClient := ankiconnect.NewClient()
	restErr := ankiConnectClient.Ping()
	if restErr != nil {
		log.Fatal(restErr)
	}

	ankiService := anki.New(ankiConnectClient)

	do(ankiService)
}

func do(ankiService *anki.Service) {
	err := ankiService.AddNotes("Default", "Main", []map[string]string{{
		"Sentence": "Слово",
		"Meaning":  "Слово",
		"IPA":      "Слово",
		"Target":   "Слово",
	}})
	if err != nil {
		panic(err)
	}
}
