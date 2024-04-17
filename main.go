// Package main is the main package for bb CLI.
package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/bytebase/cli/cmd"
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	return cmd.Execute(ctx)
}
