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

	Reporter interface {
		Report(files []string)
	}
)
