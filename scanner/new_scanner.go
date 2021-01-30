package scanner

import (
	"github.com/sirupsen/logrus"
	"os"
)

// create a new scanner to iterate over the provided directory
func NewScanner(directory string, logger *logrus.Entry) (*Scanner, error) {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		return nil, err
	}
	return &Scanner{
		logger:              logger.WithField("package", "livingit.de/code/dupfinder/scanner"),
		directory:           directory,
		potentialDuplicates: make(map[string][]string),
	}, nil
}
