package scanner

import (
	"fmt"
	"os"
	"path/filepath"
)

// Run iterates over all files and fills the internal data
func (scanner *Scanner) Run() error {
	logger := scanner.logger.WithField("method", "Run")
	pathChannel := make(chan string)
	concurrencyChannel := make(chan bool, 4)

	go func() {
		scanner.pathReceiver(pathChannel, concurrencyChannel)
	}()

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		logger.Infof("working on %s", path)
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() || !info.Mode().IsRegular() {
			return nil
		}
		scanner.waitGroup.Add(1)
		pathChannel <- path
		return nil
	})
	if err != nil {
		return err
	}

	scanner.waitGroup.Wait()

	return nil
}
