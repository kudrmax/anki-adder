package cli

import (
	"fmt"
	"log"

	"my/addToAnki/internal/presentation/gui"
)

type Saver struct{}

func (Saver) Save(s string) {
	log.Printf("saved")
}

func (cli *CLI) commandGUI(_ []string) error {
	saver := Saver{}

	app, err := gui.New(saver)
	if err != nil {
		return fmt.Errorf("error on creating app: %w", err)
	}

	return app.Run()
}
