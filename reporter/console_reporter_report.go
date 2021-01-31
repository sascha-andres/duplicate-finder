package reporter

import "fmt"

// Report writes the paths to the console if there are more than 1 entries
// in the passed argument
func (reporter *ConsoleReporter) Report(files []string) {
	if len(files) <= 1 {
		return
	}

	fmt.Println()
	fmt.Println("Potentially duplicated files:")
	for _, file := range files {
		fmt.Println(fmt.Sprintf("  %s", file))
	}
}
