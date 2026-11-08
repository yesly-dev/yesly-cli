// Command yesly is the Yesly command-line interface.
package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/yesly-dev/yesly-cli/internal/cli"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	code := cli.Execute(ctx)
	stop()
	os.Exit(code)
}
