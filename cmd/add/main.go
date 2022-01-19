package main

import (
	"context"
	"github.com/mikejeuga/go-exercises/src/adapters/cli"
)

func main()  {
	cmd := cli.NewCommand()
	ctx := context.Background()

	if cmd.AddManualEntries(ctx) {
		return
	}
	cmd.AddWithFileInputs(ctx)
}