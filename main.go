package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"

	"github.com/parsiya/ssh-scanner/scanner"
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
			Usage:       scanner.TargetUsage,
			Destination: &target,
		},
		cli.StringFlag{
			Name:        "i, in, input",
			Usage:       scanner.InputUsage,
			Destination: &input,
		},
		cli.StringFlag{
			Name:        "o, out, output",
			Usage:       scanner.OutputUsage,
			Destination: &output,
		},
		cli.StringFlag{
			Name:        "l, log",
			Usage:       scanner.LogUsage,
			Destination: &logFile,
		},
		cli.BoolFlag{
			Name:        "v, verbose",
			Usage:       scanner.VerboseUsage,
			Destination: &verbose,
		},
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "SSH Scanner"
	app.Usage = scanner.AppUsage
	app.Description = scanner.AppDescription
	app.HideVersion = true
	app.Flags = flags
	app.Action = action

	app.Run(os.Args)
}

// checkMinimumFlags checks if either target or input is set.
func checkMinimumFlags(c *cli.Context) error {
	// One of these should be set
	if target == "" && input == "" {
		return cli.NewExitError("", 2)
	}
	return nil
}

// Main action
func action(c *cli.Context) error {
	// Check if there are enough flags to run the application.
	if err := checkMinimumFlags(c); err != nil {
		fmt.Print("set either -input or -target\n\n")
		cli.ShowAppHelp(c)
		return err
	}

	// Parse target list
	var servers scanner.SSHServers
	if target != nil {

	}

	return nil
}
