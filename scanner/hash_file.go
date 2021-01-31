package scanner

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"os"
)

// hashFile hashes a file and saves it to the map
func (scanner *Scanner) hashFile(path string) error {
	logger := scanner.logger.WithField("method", "hashFile")
	f, err := os.Open(path)
	if err != nil {
		errorText := fmt.Sprintf("error opening file: %s", err)
		logger.Fatal(errorText)
		return errors.New(errorText)
	}
	defer f.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, f); err != nil {
		errorText := fmt.Sprintf("error hashing file: %s", err)
		logger.Fatal(errorText)
		return errors.New(errorText)
	}
	fileHash := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	scanner.saveHash(path, fileHash)
	return nil
}

