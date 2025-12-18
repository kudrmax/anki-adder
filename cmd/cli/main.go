package main

import (
	"log"
	"os"
	"path/filepath"

	ankiconnectExternal "github.com/atselvan/ankiconnect"

	"my/addToAnki/config"
	"my/addToAnki/internal/controllers/anki_adder_from_clipboard"
	"my/addToAnki/internal/infrastructure/clients/ankiconnect"
	"my/addToAnki/internal/infrastructure/clients/ollama"
	sentence_saver_repository "my/addToAnki/internal/infrastructure/db/sentence_saver"
	"my/addToAnki/internal/presentation/cli"
	"my/addToAnki/internal/usecases/anki/anki_adder"
	"my/addToAnki/internal/usecases/anki/note_generator"
	"my/addToAnki/internal/usecases/sentence_saver"
)

func main() {
	cfg, err := getConfig()
	if err != nil {
		log.Fatal(err)
	}

	ankiConnectExternalClient := ankiconnectExternal.NewClient()
	ankiConnectClient := ankiconnect.New(ankiConnectExternalClient)

	ankiAdderUseCase := anki_adder.NewUseCase(ankiConnectClient)
	sentenceSaverUseCase := sentence_saver.New(sentence_saver_repository.New(cfg.SentencesFilePath))

	ollamaClient, err := ollama.NewClient("llama2", false)
	if err != nil {
		ollamaClient = nil
		log.Println("Failed to initialize ollama client. <generate> command won't work")
	}
	noteGeneratorUsecase, err := note_generator.New(ollamaClient, cfg.Fields)
	if err != nil {
		log.Fatal(err)
	}

	ankiAdderFromClipboard := anki_adder_from_clipboard.New(cfg, ankiAdderUseCase)

	cliRunner := cli.NewCLI(
		cfg,
		ankiAdderUseCase,
		sentenceSaverUseCase,
		noteGeneratorUsecase,
		ankiAdderFromClipboard,
	)
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
