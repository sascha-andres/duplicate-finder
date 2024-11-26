package reporter

import (
	"errors"
	"livingit.de/code/dupfinder/scanner"
)

var (
	registeredReporters = make(map[string]func() (scanner.Reporter, error))
)

// GetReporter takes a reporter type and returns an instance
func GetReporter(reporterType string) (scanner.Reporter, error) {
	if val, ok := registeredReporters[reporterType]; ok {
		return val()
	}
	return nil, errors.New("no such reporter")
}
