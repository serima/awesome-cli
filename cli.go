package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

const (
	ExitCodeOK             = iota
	ExitCodeParseFlagError // = iota
)

type CLI struct {
	// both are io.Writer
	outStream, errStream io.Writer
}

// CLI structure as pointer receiver
func (c *CLI) Run(args []string) int {
	var version bool
	flags := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	flags.SetOutput(c.errStream)
	flags.BoolVar(&version, "version", false, "Print version information and quit")

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagError
	}

	if version {
		fmt.Fprintf(c.errStream, "awesome-cli version %s\n", Version)
		return ExitCodeOK
	}

	fmt.Fprint(c.outStream, "Do awesome working\n")

	return ExitCodeOK

}
