package scanner

// pathReceiver is responsible to hash the files in parallel
func (scanner *Scanner) pathReceiver(paths chan string, ch chan bool) {
	logger := scanner.logger.WithField("method", "pathReceiver")

	for path := range paths {
		go func(scanner *Scanner, p string) {
			ch <- true
			logger.Infof("starting hash for %s", path)
			_ = scanner.hashFile(p)
			scanner.waitGroup.Done()
			<-ch
		}(scanner, path)
	}
}
