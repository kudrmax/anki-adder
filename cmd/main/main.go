package main

import (
	"log"
	"os"

	ankiconnectExternal "github.com/atselvan/ankiconnect"

	"my/addToAnki/config"
	"my/addToAnki/internal/infrastructure/ankiconnect"
	"my/addToAnki/internal/presentation/cli"
	"my/addToAnki/internal/usecases/anki/anki_adder"
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

	ankiConnectExternalClient := ankiconnectExternal.NewClient()
	restErr := ankiConnectExternalClient.Ping()
	if restErr != nil {
		log.Fatalf("failed connecting to AnkiConnect", ankiconnect.NewClientError(restErr))
	}
	ankiConnectClient := ankiconnect.New(ankiConnectExternalClient)

	ankiUseCase := anki_adder.NewUseCase(ankiConnectClient)

	cliRunner := cli.NewCLI(cfg, ankiUseCase)
	err = cliRunner.Run(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}
