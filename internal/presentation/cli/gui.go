package cli

import (
	"fmt"
	"math/rand"

	"my/addToAnki/internal/presentation/gui"
)

type DefaultSaver struct{}

func (DefaultSaver) Save(text string) {
	ProcessSave(text)
}

type DefaultGenerator struct{}

func (DefaultGenerator) Generate(text string) string {
	return ProcessGenerate(text)
}

type DefaultNextProvider struct{}

func (DefaultNextProvider) Next() string {
	return ProcessNext()
}

func ProcessGenerate(text string) string {
	// TODO: пока заглушка
	return fmt.Sprintf("Сгенерировано: %s", text)
}

func ProcessSave(text string) {
	// TODO: пока заглушка
}

func ProcessNext() string {
	// TODO: пока заглушка
	return fmt.Sprintf("Новое предложение номер: %d", rand.Int())
}

func (cli *CLI) commandGUI(_ []string) error {
	app := gui.NewApp(
		cli.sentenceSaver,
		DefaultGenerator{},
		DefaultNextProvider{},
	)

	return app.Run()
}
