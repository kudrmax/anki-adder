package cli

import (
	"errors"

	"my/addToAnki/config"
)

const (
	commandAdd      = "add"
	commandHelp     = "help"
	commandSave     = "save"
	commandGUI      = "gui"
	commandGenerate = "generate"
)

type CLI struct {
	cfg           config.Config
	ankiAdder     ankiAdder
	sentenceSaver sentenceSaver
	noteGenerator noteGenerator
}

func NewCLI(
	cfg config.Config,
	ankiAdder ankiAdder,
	sentenceSaver sentenceSaver,
	noteGenerator noteGenerator,
) *CLI {
	return &CLI{
		cfg:           cfg,
		ankiAdder:     ankiAdder,
		sentenceSaver: sentenceSaver,
		noteGenerator: noteGenerator,
	}
}

// Run executes the CLI
// args are the command line arguments without the executable name
func (cli *CLI) Run(args []string) error {
	if len(args) == 0 {
		return errors.New("no args")
	}

	switch args[0] {
	case commandHelp:
		return nil
	case commandAdd:
		return cli.commandAdd(args[1:])
	case commandSave:
		return cli.commandSave(args[1:])
	case commandGUI:
		return cli.commandGUI(args[1:])
	case commandGenerate:
		return cli.commandGenerate(args[1:])
	default:
		return errors.New("bad command")
	}
}
