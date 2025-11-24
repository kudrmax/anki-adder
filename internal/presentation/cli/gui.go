package cli

import "my/addToAnki/internal/presentation/gui"

func (cli *CLI) commandGUI(_ []string) error {
	gui.Run()

	return nil
}
