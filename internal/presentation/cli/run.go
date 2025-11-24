package cli

import (
	"fmt"
	"os"

	"my/addToAnki/config"
)

const (
	commandAdd  = "add"
	commandHelp = "help"
	commandSave = "save"
	commandLLM  = "llm"

	helpText = "Use --help for help."
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
		cli.printInvalid(args)
		return nil
	}

	switch args[0] {
	case commandHelp:
		//return cli.commandHelp(args[1:])
		return nil
	case commandAdd:
		return cli.commandAdd(args[1:])
	case commandSave:
		return cli.commandSave(args[1:])
	case commandLLM:
		return cli.llmGenerate(args[1:])
	default:
		cli.printInvalid(args)
		return nil
	}
}

func (cli *CLI) printInvalid(args []string) {
	if len(args) == 0 {
		fmt.Fprint(os.Stderr, "no agruments. ", helpText)
	}
	fmt.Fprintf(os.Stderr, "unknown agruments: %s\n\n%s", args, helpText)
}
