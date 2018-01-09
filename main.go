package main

import (
	"fmt"
	"log"
	"os"
	"strings"

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
	// Name of optional logwriters
	logwriters string
	// Verbose flag - print logs to stdout
	verbose bool

	// Logger
	logSSH *log.Logger
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
			Destination: &logwriters,
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
		return cli.NewExitError(err.Error(), 2)
	}

	var writers []string

	// Parse logwriters
	if logwriters != "" {
		// Split by comma to get all logwriters
		writers = strings.Split(logwriters, ",")
	}

	// Setup logging
	logSSH, err := scanner.SetupLogging(writers...)
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}

	// fmt.Printf("%+v", c)

	// Parse target list
	var addresses []string

	// Read file if input is provided
	if input != "" {
		addresses, err = scanner.ReadTargetFile(input)
		if err != nil {
			logSSH.Println("error reading from file")
		}
	}

	// Parse addresses supplied by -t
	if target != "" {
		// Split by "," and add to addresses
		targets := strings.Split(target, ",")
		addresses = append(addresses, targets...)
	}

	// logSSH.Println(addresses)

	// Remove duplicate addresses
	uniques := scanner.RemoveDuplicates(addresses)

	// fmt.Println(uniques)

	// Process addresses
	var servers scanner.SSHServers
	servers.Initialize(uniques, logSSH)

	// Check if input had any correct addresses
	if len(servers) == 0 {
		logSSH.Println("no correct addresses were provided - terminating")
		return cli.NewExitError("", 2)
	}

	// fmt.Println(len(servers))

	// fmt.Println(servers.String())

	servers.Process(logSSH)

	return nil
}
