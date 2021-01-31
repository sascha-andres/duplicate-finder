package scanner

// Report creates the report using the passed reporter
func (scanner *Scanner) Report(reporter Reporter) error {
	logger := scanner.logger.WithField("method", "Report")
	logger.Info("start writing report")
	err := reporter.Setup()
	if err != nil {
		logger.Errorf("error setting up the reporter: %s", err)
	}
	for _, strings := range scanner.potentialDuplicates {
		reporter.Report(strings)
	}
	err = reporter.Finish()
	if err != nil {
		logger.Errorf("error finalizing the report: %s", err)
	}
	return nil
}
