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
func (scanner *Scanner) hashFile(info os.FileInfo, path string) error {
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

	if val, ok := scanner.potentialDuplicates[fileHash]; ok {
		val = append(val, path)
		scanner.potentialDuplicates[fileHash] = val
	} else {
		val = make([]string, 1)
		val[0] = path
		scanner.potentialDuplicates[fileHash] = val
	}
	return nil
}
