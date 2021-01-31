package scanner

import (
	"github.com/sirupsen/logrus"
	"sync"
)

type (
	// Scanner does all the heavy work, scanning and reporting
	Scanner struct {
		logger              *logrus.Entry
		waitGroup           sync.WaitGroup
		duplicatesLock      sync.RWMutex
		potentialDuplicates map[string][]string
		directory           string
	}

	// Reporter is used to write a report about found files
	Reporter interface {
		// Report is used to send a potentially duplicated file record
		Report(files []string)
		// Setup is called once to prepare everything a report needs
		Setup() error
		// Finish is used for cleanup and post data tasks
		Finish() error
	}
)
