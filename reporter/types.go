package reporter

type (
	ConsoleReporter struct{}

	GroupedConsoleReporter struct {
		directoryFilesMap map[string]map[string]int
		fileList          map[int][]string
	}
)
