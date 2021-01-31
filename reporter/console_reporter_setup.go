package reporter

import "fmt"

// Setup is called once to prepare everything a report needs
func (reporter *ConsoleReporter) Setup() error {
	fmt.Println("== potentially duplicated files ==")
	fmt.Println()
	return nil
}
