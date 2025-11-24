package cli

import (
	"context"
	"errors"
	"fmt"
)

func (cli *CLI) llmGenerate(args []string) error {
	if len(args) != 1 {
		return errors.New("Usage: llm <prompt>")
	}

	resp, err := cli.llmGenerator.Generate(context.Background(), args[0])
	if err != nil {
		return fmt.Errorf("failed to generate text: %w", err)
	}

	fmt.Println(resp)

	return nil
}
