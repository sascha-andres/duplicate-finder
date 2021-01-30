package reporter

// NewConsoleReporter creates a new console reporter
func NewConsoleReporter() (*ConsoleReporter, error) {
	return &ConsoleReporter{}, nil
}

