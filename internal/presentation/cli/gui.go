package cli

import (
	"fmt"
	"math/rand"

	"my/addToAnki/internal/presentation/gui"
)

type DefaultNextProvider struct{}

func (DefaultNextProvider) Next() string {
	return ProcessNext()
}

func ProcessNext() string {
	// TODO: пока заглушка
	return fmt.Sprintf("Новое предложение номер: %d", rand.Int())
}

func (cli *CLI) commandGUI(_ []string) error {
	app := gui.NewApp(
		cli.cfg,
		cli.sentenceSaver,
		cli.noteGenerator,
		DefaultNextProvider{},
		cli.ankiAdder,
		cli.ankiAdderFromClipboard,
	)

	return app.Run()
}
