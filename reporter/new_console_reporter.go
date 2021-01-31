package reporter

import "livingit.de/code/dupfinder/scanner"

// NewConsoleReporter creates a new console reporter
func NewConsoleReporter() (*ConsoleReporter, error) {
	return &ConsoleReporter{}, nil
}

func consoleReporterFactory() (scanner.Reporter, error) {
	reporter, err := NewConsoleReporter()
	if err != nil {
		return nil, err
	}
	return reporter, nil
}

func init() {
	registeredReporters["console"] = consoleReporterFactory
}
