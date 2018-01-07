// Project constants
package scanner

const (
	// Flag usage descriptions
	targetUsage  = "comma-separated target addresses"
	inputUsage   = "`input file` with one target address on each line"
	outputUsage  = "store results in `output file`"
	logUsage     = "store logs in `log file`"
	verboseUsage = "print logs to console"

	// Application constants for cli package
	appUsage = "SSH Scanner scans SSH servers for vulnerabilities.\n" +
		"Addresses should be in format of 'host:port'.\n" +
		"Input file should have one address on each line " +
		"and addresses provided to -targets should be separated by commas.\n" +
		"-in and -targets are mutually exclusive, use one.\n"

	appDescription = "Scan two servers, log the results to log1.txt and also print it to console.\n" +
		"ssh-sccaner -t 127.0.0.1:22,192.168.0.1:1234 -log log1.txt -v\n" +
		"Read targets from input2.txt, store output to report2.txt and log results to log2.txt\n" +
		"ssh-sccaner -in input2.txt -out output2.txt -log log2.txt\n"
)
