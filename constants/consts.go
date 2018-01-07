// Project constants
package constants

const (
	// Flag usage descriptions
	TargetUsage  = "one or more comma-separated scanning `target`s"
	InputUsage   = "`input file` with one target address on each line"
	OutputUsage  = "store results in `output file`"
	LogUsage     = "store logs in `log file`"
	VerboseUsage = "print logs to console"

	// Application constants for cli package
	AppUsage = `Scans SSH servers for vulnerabilities. Addresses should be in format of 'host:port'.
	 Input file should have one address on each line and addresses provided to -targets should be separated by commas.
	 At a minium, either -in or -targets (mutually exclusive) should be set.`

	AppDescription = `Scan two servers, log the results to log1.txt and also print it to console
	  ssh-sccaner -t 127.0.0.1:22,192.168.0.1:1234 -log log1.txt -v
	 Read targets from input2.txt, store output to report2.txt and log results to log2.txt
	  ssh-sccaner -in input2.txt -out output2.txt -log log2.txt`
)
