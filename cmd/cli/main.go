package main

import (
	"log"
	"os"
	"path/filepath"

	ankiconnectExternal "github.com/atselvan/ankiconnect"

	"my/addToAnki/config"
	"my/addToAnki/internal/infrastructure/clients/ankiconnect"
	"my/addToAnki/internal/infrastructure/clients/ollama"
	santence_saver_repository "my/addToAnki/internal/infrastructure/db/santence_saver"
	"my/addToAnki/internal/presentation/cli"
	"my/addToAnki/internal/usecases/anki/anki_adder"
	"my/addToAnki/internal/usecases/anki/sentence_saver"
)

func main() {
	cfg, err := getConfig()
	if err != nil {
		log.Fatal(err)
	}

	ankiConnectExternalClient := ankiconnectExternal.NewClient()
	ankiConnectClient := ankiconnect.New(ankiConnectExternalClient)

	ankiUseCase := anki_adder.NewUseCase(ankiConnectClient)
	sentenceSaverUseCase := sentence_saver.New(santence_saver_repository.New(cfg.SentencesFilePath))

	ollamaClient, err := ollama.NewClient("llama2", false)
	if err != nil {
		log.Fatal(err)
	}

	cliRunner := cli.NewCLI(cfg, ankiUseCase, sentenceSaverUseCase, ollamaClient)
	err = cliRunner.Run(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}

func getConfig() (config.Config, error) {
	f, err := os.Open(getConfigPath())
	if err != nil {
		return config.Config{}, err
	}

	cfg, err := config.Parse(f)
	if err != nil {
		return config.Config{}, err
	}

	return cfg, nil
}

func getConfigPath() string {
	dir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	return filepath.Join(dir, ".config", "anki-adder", "config.yaml")
}
