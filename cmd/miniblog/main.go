package main

import (
	"miniblog/internal/miniblog"
	"os"
)

func main() {
	command := miniblog.NewMiniBlogCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
