package main

import (
	"os"

	"github.com/marco-souza/pkg/cmd/cli"
)

func main() {
	err := cli.Execute()
	if err != nil {
		cli.Fallback()
		os.Exit(1)
	}
}
