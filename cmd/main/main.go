package main

import (
	"log"
	"os"
	"path/filepath"

	ankiconnectExternal "github.com/atselvan/ankiconnect"

	"my/addToAnki/config"
	"my/addToAnki/internal/infrastructure/ankiconnect"
	santence_saver_repository "my/addToAnki/internal/infrastructure/santence_saver"
	"my/addToAnki/internal/presentation/cli"
	"my/addToAnki/internal/usecases/anki/anki_adder"
	"my/addToAnki/internal/usecases/anki/sentence_saver"
)

func main() {
	f, err := os.Open(getConfigPath())
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
		log.Fatalf("failed connecting to AnkiConnect: %s", ankiconnect.NewClientError(restErr))
	}
	ankiConnectClient := ankiconnect.New(ankiConnectExternalClient)

	ankiUseCase := anki_adder.NewUseCase(ankiConnectClient)
	sentenceSaverUseCase := sentence_saver.New(santence_saver_repository.New(cfg.DBFile))

	cliRunner := cli.NewCLI(cfg, ankiUseCase, sentenceSaverUseCase)
	err = cliRunner.Run(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}

func getConfigPath() string {
	dir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	return filepath.Join(dir, ".config", "anki-adder", "config.yaml")
}
