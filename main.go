package main

import (
	"github.com/mikejeuga/go-exercises/src/adapters/cli"
)

func main()  {
	cmd := cli.NewCommand()

	if cmd.AddManualEntries() {
		return
	}
	cmd.AddWithFileInputs()
}