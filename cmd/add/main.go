package main

import (
	"context"
	"github.com/mikejeuga/go-exercises/src/adapters/cli"
	add "github.com/mikejeuga/go-exercises/src/domain"
)

func main()  {
	cmd := cli.NewCommand(add.Default)
	ctx := context.Background()

	if cmd.AddManualEntries(ctx) {
		return
	}
	cmd.AddWithFileInputs(ctx)
}