package main

import (
	"os"

	"github.com/g2a-com/klio/pkg/cli"
)

var Version string

func main() {
	cmd := cli.CLI{
		CommandName: "ifico",
		Version:     Version,
	}

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
