// Project constants
package scanner

import (
	"log"
	"time"
)

const (
	// Flag usage descriptions
	TargetUsage  = "one or more comma-separated scanning `target`s"
	InputUsage   = "`input file` with one target address on each line"
	OutputUsage  = "store results in `output file`"
	LogUsage     = "store logs in `log file(s)`"
	VerboseUsage = "print logs to console"

	// Application constants for cli package
	AppUsage = `Scans SSH servers for vulnerabilities. Addresses should be in format of 'host:port'.
	 Input file should have one address on each line. Addresses provided to -targets should be comma-separated.
	 At a minium, either -in or -targets should be set. Both can be used but not recommended.`

	AppDescription = `Scan two servers, log the results to log1.txt and also print it to console
	  ssh-sccaner -t 127.0.0.1:22,192.168.0.1:1234 -log log1.txt -v
	 Read targets from input2.txt, store output to report2.txt and log results to log2.txt
	  ssh-sccaner -in input2.txt -out output2.txt -log log2.txt
	 -l, -log support multiple io.Writers (e.g. files or os.Stderr). Default is os.Stdout if nothing is provided.
	 Scan one server and write the output to two files, os.stdOut and os.Stderr
	  ssh-scanner -t 127.0.0.1:22 -log log1.txt,log2.txt,os.Stderr`

	// Scanning constants
	// Timeout
	Timeout = 5 * time.Second

	// Default SSH username and password
	DefaultUser     = "user"
	DefaultPassword = "password"

	// Logging
	// Log prefix - appears before every line in logs - note the space
	LogPrefix = "[*] "
	// LogFlag can be any combination of (join with |):
	// Ldate | Ltime | Lmircoseconds | Llongfile | Lshortfile | LUTC
	// See https://godoc.org/log#pkg-constants
	LogFlag = log.Ltime
)
