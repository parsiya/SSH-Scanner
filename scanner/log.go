// Logging
package scanner

import (
	"io"
	"log"
	"os"
	"strings"
)

// SetupLogging will return a log.Logger object that can be used for logging.
// It will write the logs to all io.Writers designated in logs.
// If there are any errors, it will return nil and the appropriate error.
// Usage:
// SetupLogging("log1.txt") = only log to one file
// SetupLogging("log1.txt", "os.Stdout") = log to both os.Stdout and file
// SetupLogging("log1.txt", "log2.txt") = log to two files
// SetupLogging("os.Stdout") = only log to stdout
// SetupLogging("os.Stderr", "os.Stdout") = log to both os.Stdout and os.Stderr
func SetupLogging(logwriters ...string) (*log.Logger, error) {

	var writers []io.Writer

	// Add all logwriters
	for _, log := range logwriters {

		// Add os.Stdout and os.Stderr if provided
		if strings.ToLower(log) == "os.stdout" {
			writers = append(writers, os.Stdout)
			continue
		}

		if strings.ToLower(log) == "os.stderr" {
			writers = append(writers, os.Stderr)
			continue
		}

		// Otherwise attempt to create/open a file with that name/path
		logFile, err := os.Create(log)
		if err != nil {
			return nil, err
		}
		writers = append(writers, logFile)
	}

	// If nothing is provided, use os.Stdout
	if len(logwriters) == 0 {
		writers = append(writers, os.Stdout)
	}

	// Create a multiwriter from all writers
	multiW := io.MultiWriter(writers...)
	// Create logger and return
	logger := log.New(multiW, LogPrefix, LogFlag)
	return logger, nil
}
