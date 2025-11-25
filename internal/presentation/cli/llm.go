package cli

import (
	"errors"
	"fmt"
)

func (cli *CLI) commandGenerate(args []string) error {
	if len(args) != 2 {
		return errors.New("generate command expects exactly 2 arguments: <sentence> <target>")
	}

	resp, err := cli.noteGenerator.GenerateNote(args[0], args[1])
	if err != nil {
		return fmt.Errorf("failed to generate text: %w", err)
	}

	fmt.Println(resp)

	return nil
}
