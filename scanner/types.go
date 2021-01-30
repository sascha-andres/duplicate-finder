package scanner

import (
	"github.com/sirupsen/logrus"
)

type (
	// Scanner does all the heavy work, scanning and reporting
	Scanner struct {
		logger *logrus.Entry
		potentialDuplicates map[string][]string
		directory           string
	}

	Reporter interface {
		Report(files []string)
	}
)

