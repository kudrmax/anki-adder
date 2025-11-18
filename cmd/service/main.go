package main

import (
	"flag"
	"os"
	"path/filepath"
	"runtime"

	"github.com/atselvan/ankiconnect"
	"github.com/privatesquare/bkst-go-utils/utils/errors"

	"my/addToAnki/internal/services/anki"
	"my/addToAnki/internal/services/anki_adder"
	"my/addToAnki/internal/services/config"
)

func main() {
	ankiConnectClient := ankiconnect.NewClient()
	restErr := ankiConnectClient.Ping()
	if restErr != nil {
		panic(errors.New("failed connecting to AnkiConnect"))
	}

	configPath := "config.yaml" // TODO временно, заменить на конфиг лежащий где-то вне
	cfg, err := config.Parse(configPath)
	if err != nil {
		panic(err)
	}

	ankiService, err := anki.New(ankiConnectClient, anki.Config{
		Deck:      cfg.Deck,
		NoteModel: cfg.NoteModel,
	})
	if err != nil {
		panic(err)
	}

	err = run(ankiService)
	if err != nil {
		panic(err)
	}
}

func run(ankiService *anki.Service) error {
	flagFrom := flag.String("from", "", "path to file for import")
	flagFromClipboard := flag.Bool("from_clipboard", false, "import from clipboard")
	flag.Parse()

	ankiAdder, err := getAnkiAdderService(ankiService, flagFrom, flagFromClipboard)
	if err != nil {
		return err
	}

	return ankiAdder.AddNotes()
}

func getAnkiAdderService(ankiService *anki.Service, filePath *string, fromClipboard *bool) (anki_adder.IAnkiAdder, error) {
	switch {
	case filePath != nil && *filePath != "":
		return anki_adder.NewAnkiAdderFromCSV(ankiService, *filePath), nil
	case fromClipboard != nil && *fromClipboard:
		return anki_adder.NewAnkiAdderFromClipboard(ankiService), nil
	default:
		return nil, errors.New("invalid flags")
	}
}

func getConfigPath() (string, error) {
	var baseConfigDir string
	var err error

	switch {
	case runtime.GOOS == "windows":
		baseConfigDir, err = os.UserConfigDir()
		if err != nil {
			return "", err
		}
	default:
		if xdg := os.Getenv("XDG_CONFIG_HOME"); xdg != "" {
			baseConfigDir = xdg
		} else {
			home, err := os.UserHomeDir()
			if err != nil {
				return "", err
			}
			baseConfigDir = filepath.Join(home, ".config")
		}
	}

	appDir := filepath.Join(baseConfigDir, "anki-adder")
	return filepath.Join(appDir, "config.yaml"), nil
}
