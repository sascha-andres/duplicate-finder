package scanner

// Report creates the report using the passed reporter
func (scanner *Scanner) Report(reporter Reporter) error {
	logger := scanner.logger.WithField("method", "Report")
	logger.Info("start writing report")
	for _, strings := range scanner.potentialDuplicates {
		reporter.Report(strings)
	}
	return nil
}
