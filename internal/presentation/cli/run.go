package cli

import (
	"errors"

	"my/addToAnki/config"
)

const (
	commandAdd  = "add"
	commandHelp = "help"
	commandSave = "save"
	commandGUI  = "gui"
	commandLLM  = "generate"
)

type CLI struct {
	cfg           config.Config
	ankiAdder     ankiAdder
	sentenceSaver sentenceSaver
	llmGenerator  llmGenerator
}

func NewCLI(
	cfg config.Config,
	ankiAdder ankiAdder,
	sentenceSaver sentenceSaver,
	llmGenerator llmGenerator,
) *CLI {
	return &CLI{
		cfg:           cfg,
		ankiAdder:     ankiAdder,
		sentenceSaver: sentenceSaver,
		llmGenerator:  llmGenerator,
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
	case commandLLM:
		return cli.llmGenerate(args[1:])
	default:
		return errors.New("bad command")
	}
}
