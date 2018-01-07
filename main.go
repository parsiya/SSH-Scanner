package main

import (
	"fmt"

	"github.com/urfave/cli"

	"scanner/consts"
)

var (
	// urfave-cli flags
	flags []cli.Flag

	// List of targets
	targets []string
	// Input file
	input string
	// Output file
	output string
	// Name of optional log file
	logFile string
	// Verbose flag - if true logs are printed to stdout
	verbose bool
)

func init() {

	// Setup flags
	flags = []cli.Flag{
		cli.StringFlag{
			Name:  "t, target, host",
			Usage: targetUsage,
		},
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "SSH Scanner"
	app.Usage = appUsage
	app.HideVersion = true
	app.Flags = flags
	app.Action = noArgs

	app.Run(os.Args)
}

func noArgs(c *cli.Context) error {
	cli.ShowAppHelp(c)
	return cli.NewExitError("no commands provided", 2)
}
