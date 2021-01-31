package reporter

import "livingit.de/code/dupfinder/scanner"

// NewGroupedConsoleReporter creates a new console reporter with directory grouping
func NewGroupedConsoleReporter() (*GroupedConsoleReporter, error) {
	return &GroupedConsoleReporter{
		directoryFilesMap: make(map[string]map[string]int),
		fileList:          make(map[int][]string),
	}, nil
}

func consoleGroupReporterFactory() (scanner.Reporter, error) {
	reporter, err := NewGroupedConsoleReporter()
	if err != nil {
		return nil, err
	}
	return reporter, nil
}

func init() {
	registeredReporters["grouped-console"] = consoleGroupReporterFactory
}
