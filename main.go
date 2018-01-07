package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"

	"github.com/parsiya/ssh-scanner/constants"
)

var (
	// urfave-cli flags
	flags []cli.Flag

	// List of targets
	target string
	// Input file
	input string
	// Output file
	output string
	// Name of optional log file
	logFile string
	// Verbose flag - print logs to stdout
	verbose bool
)

func init() {

	// Setup flags
	flags = []cli.Flag{
		cli.StringFlag{
			Name:        "t, target, host",
			Usage:       constants.TargetUsage,
			Destination: &target,
		},
		cli.StringFlag{
			Name:        "i, in, input",
			Usage:       constants.InputUsage,
			Destination: &input,
		},
		cli.StringFlag{
			Name:        "o, out, output",
			Usage:       constants.OutputUsage,
			Destination: &output,
		},
		cli.StringFlag{
			Name:        "l, log",
			Usage:       constants.LogUsage,
			Destination: &logFile,
		},
		cli.BoolFlag{
			Name:        "v, verbose",
			Usage:       constants.VerboseUsage,
			Destination: &verbose,
		},
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "SSH Scanner"
	app.Usage = constants.AppUsage
	app.Description = constants.AppDescription
	app.HideVersion = true
	app.Flags = flags
	app.Action = checkMinimumArgs

	app.Run(os.Args)
}

// checkArgNumber checks if either target or input is set.
func checkMinimumArgs(c *cli.Context) error {
	// One of these should be set
	if target == "" && input == "" {
		fmt.Print("set either -input or -target\n\n")
		cli.ShowAppHelp(c)
		return cli.NewExitError("", 2)
	}
	return nil
}
