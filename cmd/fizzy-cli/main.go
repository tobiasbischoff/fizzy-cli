package main

import (
	"os"

	"fizzy-cli/internal/cli"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	os.Exit(cli.Run(version, commit, date, os.Args))
}
