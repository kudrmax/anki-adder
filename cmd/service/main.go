package main

import (
	"log"

	"github.com/atselvan/ankiconnect"

	"github.com/AlekSi/pointer"

	"my/addToAnki/internal/models"
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
	deck := "Default"
	words := []models.Word{
		{
			ModelName: "Main",
			Data: models.WordData{
				Sentence: pointer.To("Слово 1"),
				//Target:   nil,
				Meaning: pointer.To("Описание"),
				//IPA:      nil,
				//Image:    nil,
				//Sound:    nil,
			},
		},
	}

	err := ankiService.AddWords(deck, words)
	if err != nil {
		panic(err)
	}
}
